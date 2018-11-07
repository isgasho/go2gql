package graphql

import (
	"strings"

	"github.com/pkg/errors"
)

type SchemaParserObjects struct {
	Services       []SchemaService
	QueryObject    string
	MutationObject string
	Objects        []*gqlObject
}

type schemaParser struct {
	schemaCfg SchemaConfig
	types     map[string]*TypesFile
}

func NewSchemaParser(schemaCfg SchemaConfig, types map[string]*TypesFile) *schemaParser {
	return &schemaParser{schemaCfg, types}
}

func (g *schemaParser) resolveObjectFields(nodeCfg SchemaNodeConfig, object *gqlObject) (services []SchemaService, newObjectx []*gqlObject, err error) {
	var newObjs []*gqlObject
	var newServices []SchemaService
	switch nodeCfg.Type {
	case SchemaNodeTypeObject:
		for _, fld := range nodeCfg.Fields {
			fldObj := &gqlObject{
				QueryObject: object.QueryObject,
				Name:        strings.Replace(fld.ObjectName, " ", "_", -1),
			}
			services, subObjs, err := g.resolveObjectFields(fld, fldObj)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "can't resolve field %s object fields", fld.Field)
			}
			if len(fldObj.Fields) > 0 {
				object.Fields = append(object.Fields, fieldConfig{
					Name:   fld.Field,
					Object: fldObj,
				})
				newServices = append(newServices, services...)
				newObjs = append(newObjs, subObjs...)
				newObjs = append(newObjs, fldObj)
			}
		}
		return newServices, newObjs, nil
	case SchemaNodeTypeService:
		for _, typesFile := range g.types {
			for _, service := range typesFile.Services {
				if service.Name != nodeCfg.Service {
					continue
				}

				var serviceMethods []Method

				if object.QueryObject == true {
					serviceMethods = service.QueryMethods
				} else {
					serviceMethods = service.MutationMethods
				}

				meths := make([]string, len(serviceMethods))

				for i, meth := range serviceMethods {
					meths[i] = meth.Name
				}
				fields := g.filterMethods(meths, nodeCfg.FilterMethods, nodeCfg.ExcludeMethods)
				srv := SchemaService{
					Name:         service.Name,
					ClientGoType: service.CallInterface,
					Pkg:          typesFile.Package,
				}
				newServices = append(newServices, srv)
				for _, fld := range fields {
					object.Fields = append(object.Fields, fieldConfig{
						Name:    fld,
						Service: &srv,
					})

				}

				return newServices, nil, nil
			}
		}
		return nil, nil, errors.Errorf("service '%s' not found", nodeCfg.Service)

	default:
		return nil, nil, errors.Errorf("unknown type %s", nodeCfg.Type)
	}
}

func (g *schemaParser) filterMethods(methods []string, filter, exclude []string) []string {
	var res []string
	var filteredMethods = make(map[string]struct{})
	for _, f := range filter {
		filteredMethods[f] = struct{}{}
	}
	var excludedMethods = make(map[string]interface{})
	for _, f := range exclude {
		excludedMethods[f] = struct{}{}
	}
	for _, m := range methods {
		if len(excludedMethods) > 0 {
			if _, ok := excludedMethods[m]; ok {
				continue
			}
		}
		if len(filteredMethods) > 0 {
			if _, ok := filteredMethods[m]; !ok {
				continue
			}
		}
		res = append(res, m)
	}

	return res
}

func (g *schemaParser) resolveSchemaObjects() ([]SchemaService, []*gqlObject, string, string, error) {
	var objects []*gqlObject
	var services []SchemaService
	var queryObject, mutationObject string
	if g.schemaCfg.Queries != nil {
		var queryObj = &gqlObject{
			QueryObject: true,
			Name:        "Query",
		}
		newServices, newObjs, err := g.resolveObjectFields(*g.schemaCfg.Queries, queryObj)
		if err != nil {
			return nil, nil, "", "", errors.Wrap(err, "failed to resolve queries fields")
		}
		objects = append(objects, newObjs...)
		objects = append(objects, queryObj)
		services = append(services, newServices...)
		queryObject = queryObj.Name
	}
	if g.schemaCfg.Mutations != nil {
		var mutationObj = &gqlObject{
			QueryObject: false,
			Name:        "Mutation",
		}
		newServices, newObjs, err := g.resolveObjectFields(*g.schemaCfg.Mutations, mutationObj)

		if err != nil {
			return nil, nil, "", "", errors.Wrap(err, "failed to resolve mutations fields")
		}
		if len(mutationObj.Fields) > 0 {
			objects = append(objects, newObjs...)
			objects = append(objects, mutationObj)
			services = append(services, newServices...)
			mutationObject = mutationObj.Name
		}
	}
	var uniqueServices []SchemaService
	var handledService = map[string]struct{}{}
	for _, service := range services {
		if _, ok := handledService[service.Name]; !ok {
			uniqueServices = append(uniqueServices, service)
			handledService[service.Name] = struct{}{}
		}
	}

	return uniqueServices, objects, queryObject, mutationObject, nil
}

func (g schemaParser) SchemaObjects() (*SchemaParserObjects, error) {
	services, objects, queryObj, mutationsObj, err := g.resolveSchemaObjects()
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve objects to generate")
	}

	return &SchemaParserObjects{
		QueryObject:    queryObj,
		MutationObject: mutationsObj,
		Objects:        objects,
		Services:       services,
	}, nil
}