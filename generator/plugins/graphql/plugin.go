package graphql

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"

	"github.com/EGT-Ukraine/go2gql/generator"
	"github.com/EGT-Ukraine/go2gql/generator/plugins/graphql/lib/importer"
)

const (
	PluginName            = "graphql"
	SchemasConfigsKey     = "graphql_schemas"
	DataLoadersConfigsKey = "data_loaders"
)

type Plugin struct {
	files             map[string]*TypesFile
	dataLoaderConfigs *DataLoadersConfig
	schemaConfigs     []SchemaConfig
	generateCfg       *generator.GenerateConfig
	dataLoader        *DataLoader
}

type SchemaObjects struct {
	SchemaName     string
	GoPkg          string
	Services       []SchemaService
	QueryObject    string
	MutationObject string
	Objects        []*gqlObject
}

func (p *Plugin) Prepare() error {
	return nil
}

func (p *Plugin) Init(config *generator.GenerateConfig, plugins []generator.Plugin) error {
	var cfgs []SchemaConfig
	p.files = make(map[string]*TypesFile)
	err := mapstructure.Decode(config.PluginsConfigs[SchemasConfigsKey], &cfgs)
	if err != nil {
		return errors.Wrap(err, "failed to decode config")
	}
	p.schemaConfigs = cfgs

	var dataLoadersConfig DataLoadersConfig

	if config.PluginsConfigs[DataLoadersConfigsKey] != nil {
		if err := mapstructure.Decode(config.PluginsConfigs[DataLoadersConfigsKey], &dataLoadersConfig); err != nil {
			return errors.Wrap(err, "failed to decode dataloaders config")
		}

		dataLoadersConfig.OutputPath, err = filepath.Abs(dataLoadersConfig.OutputPath)
		if err != nil {
			return errors.Wrapf(err, "Failed to normalize path")
		}

		p.dataLoaderConfigs = &dataLoadersConfig
	}

	p.generateCfg = config

	if err = p.parseImports(); err != nil {
		return errors.Wrap(err, "failed to decode imports")
	}

	return nil
}

func (p *Plugin) parseImports() error {
	for _, pluginsConfigsImports := range p.generateCfg.PluginsConfigsImports {
		cfg := new([]*SchemaConfig)

		if err := mapstructure.Decode(pluginsConfigsImports.PluginsConfigs[SchemasConfigsKey], cfg); err != nil {
			return errors.Wrap(err, "failed to decode config")
		}

		for _, schema := range *cfg {
			if err := p.parseImportedSchema(schema); err != nil {
				return errors.Wrap(err, "Failed to parse imported schema")
			}
		}
	}

	return nil
}

func (p *Plugin) parseImportedSchema(cfg *SchemaConfig) error {
	if cfg == nil {
		return nil
	}

	var importedSchemaName = cfg.Name

	schema := p.findSchemaByName(importedSchemaName)

	if schema == nil {
		return errors.New("Schema " + importedSchemaName + " not defined")
	}

	if cfg.Queries != nil && schema.Queries.Type == SchemaNodeTypeService {
		return errors.New("Cannot merge object with service in query")
	}

	if cfg.Mutations != nil && schema.Mutations.Type == SchemaNodeTypeService {
		return errors.New("Cannot merge object with service in mutations")
	}

	if cfg.Queries != nil {
		schema.Queries.Fields = p.mergeFields(schema.Queries.Fields, cfg.Queries.Fields)
	}

	if cfg.Mutations != nil {
		schema.Mutations.Fields = p.mergeFields(schema.Mutations.Fields, cfg.Mutations.Fields)
	}

	return nil
}

func (p *Plugin) mergeFields(schemaFields []SchemaNodeConfig, importFields []SchemaNodeConfig) []SchemaNodeConfig {
	for _, importField := range importFields {
		nodeIdx := p.findNodeIndex(importField.Field, schemaFields)

		if nodeIdx != -1 {
			schemaFields[nodeIdx].Fields = p.mergeFields(schemaFields[nodeIdx].Fields, importField.Fields)
		} else {
			schemaFields = append(schemaFields, importField)
		}
	}

	return schemaFields
}

func (p *Plugin) findNodeIndex(name string, schemaFields []SchemaNodeConfig) int {
	for schemaFieldIdx, schemaField := range schemaFields {
		if schemaField.Field == name {
			return schemaFieldIdx
		}
	}

	return -1
}

func (p *Plugin) findSchemaByName(name string) *SchemaConfig {
	for _, schema := range p.schemaConfigs {
		if schema.Name == name {
			return &schema
		}
	}

	return nil
}

// Types returns info about all parsed types
func (p *Plugin) Types() map[string]*TypesFile {
	return p.files
}

func (p *Plugin) AddTypesFile(outputPath string, file *TypesFile) {
	p.files[outputPath] = file
}

func (p Plugin) Name() string {
	return PluginName
}

func (p *Plugin) PrintInfo(infos generator.Infos) {
	if infos.Contains("gql-services") {
		for path, file := range p.files {
			if len(file.Services) > 0 {
				fmt.Println(path)
				for _, service := range file.Services {
					fmt.Println("\t Service " + service.Name)
				}
			}
		}
	}
}

func (p *Plugin) Infos() map[string]string {
	return map[string]string{
		"gql-services": "Info about all parsed GraphQL services",
	}
}

func (p *Plugin) validateOutputObjects(outputObjects []OutputObject) error {
	for _, outputObject := range outputObjects {
		for _, dataLoaderField := range outputObject.DataLoaderFields {
			dataLoader, ok := p.dataLoader.Loaders[dataLoaderField.DataLoaderName]

			if !ok {
				return errors.Errorf(
					"Failed to found dataloader with name %s in object %s",
					dataLoaderField.DataLoaderName,
					outputObject.GraphQLName,
				)
			}

			outputArgument := outputObject.FindFieldByName(dataLoaderField.ParentKeyFieldName)

			if outputArgument == nil {
				return errors.Errorf(
					"Field `%s` not found in `%s`",
					dataLoaderField.ParentKeyFieldName,
					outputObject.GraphQLName,
				)
			}

			if !outputArgument.GoType.Scalar {
				return errors.Errorf(
					"Field `%s` in `%s` must be scalar",
					dataLoaderField.ParentKeyFieldName,
					outputObject.GraphQLName,
				)
			}

			if dataLoader.InputGoType.ElemType.Kind != outputArgument.GoType.Kind {
				// TODO: use type casting if possible.
				return errors.New("Input argument must be same type as output")
			}
		}
	}

	return nil
}

func (p *Plugin) generateTypes() error {
	for outputPath, file := range p.files {
		if err := p.validateOutputObjects(file.OutputObjects); err != nil {
			return errors.Wrapf(err, "failed to validate output objects in file %s", outputPath)
		}

		err := os.MkdirAll(filepath.Dir(outputPath), 0777)
		if err != nil {
			return errors.Wrapf(err, "failed to create directories for output types file %s", outputPath)
		}
		out, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			return errors.Wrapf(err, "failed to open file %s for write", outputPath)
		}
		err = typesGenerator{
			File:          file,
			tracerEnabled: p.generateCfg.GenerateTraces,
			imports: &importer.Importer{
				CurrentPackage: file.Package,
			},
			dataLoader: p.dataLoader,
		}.generate(out)
		if err != nil {
			if cerr := out.Close(); cerr != nil {
				err = errors.Wrap(err, cerr.Error())
			}

			return errors.Wrapf(err, "failed to generate types file %s", outputPath)
		}
		if err = out.Close(); err != nil {
			return errors.Wrapf(err, "failed to close generated types file %s", outputPath)
		}
	}

	return nil
}

// SchemasObjects returns info about schemas objects. Useful in external plugins.
func (p *Plugin) SchemasObjects() ([]SchemaObjects, error) {
	var schemasObjects []SchemaObjects

	for _, schema := range p.schemaConfigs {
		pkg, err := GoPackageByPath(filepath.Dir(schema.OutputPath), p.generateCfg.VendorPath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve schema %s output go package", schema.Name)
		}
		parser := NewSchemaParser(schema, p.files)

		schemaContext, err := parser.SchemaObjects()

		if err != nil {
			return nil, errors.Wrapf(err, "failed to get schema %s template context", schema.Name)
		}

		schemaObjects := SchemaObjects{
			SchemaName:     schema.Name,
			GoPkg:          pkg,
			Services:       schemaContext.Services,
			QueryObject:    schemaContext.QueryObject,
			MutationObject: schemaContext.MutationObject,
			Objects:        schemaContext.Objects,
		}

		schemasObjects = append(schemasObjects, schemaObjects)
	}

	return schemasObjects, nil
}

func (p *Plugin) Generate() error {
	dataLoader, err := CreateDataLoader(p.dataLoaderConfigs, p.generateCfg.VendorPath, p.files)

	if err != nil {
		return errors.Wrap(err, "failed to process dataloader config")
	}

	loaderGen := NewLoaderGenerator(dataLoader)

	if err := loaderGen.GenerateDataLoaders(); err != nil {
		return errors.Wrap(err, "failed to generate data loader files")
	}

	p.dataLoader = dataLoader

	if err := p.generateTypes(); err != nil {
		return errors.Wrap(err, "failed to generate types files")
	}

	if err = p.generateSchemas(); err != nil {
		return errors.Wrap(err, "failed to generate schema files")
	}

	return nil
}
