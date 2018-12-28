package swagger2gql

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"

	"github.com/EGT-Ukraine/go2gql/generator/plugins/graphql"
	"github.com/EGT-Ukraine/go2gql/generator/plugins/graphql/lib/importer"
	"github.com/EGT-Ukraine/go2gql/generator/plugins/graphql/lib/names"
	"github.com/EGT-Ukraine/go2gql/generator/plugins/swagger2gql/parser"
)

var ErrMultipleSuccessResponses = errors.New("method  contains multiple success responses")

func (p *Plugin) graphqlMethod(methodCfg MethodConfig, file *parsedFile, tag parser.Tag, method parser.Method) (*graphql.Method, error) {
	name := method.OperationID
	if methodCfg.Alias != "" {
		name = methodCfg.Alias
	}
	var successResponse parser.MethodResponse
	var successResponseFound bool
	for _, resp := range method.Responses {
		if resp.StatusCode/100 == 2 {
			if successResponseFound {
				return nil, ErrMultipleSuccessResponses
			}

			successResponse = resp
			successResponseFound = true
		}
	}
	responseType, err := p.TypeOutputTypeResolver(file, successResponse.ResultType, false)
	if err != nil {
		return nil, errors.Wrap(err, "can't get response output type resolver")
	}
	gqlInputObjName := p.methodParamsInputObjectGQLName(file, method)
	var args []graphql.MethodArgument
	for _, param := range method.Parameters {
		gqlName := names.FilterNotSupportedFieldNameCharacters(param.Name)

		paramCfg, err := file.Config.FieldConfig(gqlInputObjName, param.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve property %s config", param.Name)
		}

		if paramCfg.ContextKey != "" {
			continue
		}
		paramType, err := p.TypeInputTypeResolver(file, param.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve parameter '%s' type resolver", param.Name)
		}

		goType, err := p.goTypeByParserType(file, param.Type, true)

		if err != nil {
			return nil, err
		}

		args = append(args, graphql.MethodArgument{
			Name:          gqlName,
			Type:          paramType,
			GoType:        goType,
			QuotedComment: strconv.Quote(param.Description),
		})
	}
	reqType := graphql.GoType{
		Kind: reflect.Ptr,
		ElemType: &graphql.GoType{
			Kind: reflect.Interface,
			Pkg:  file.Config.Tags[tag.Name].ClientGoPackage,
			Name: pascalize(method.OperationID) + "Params",
		},
	}

	return &graphql.Method{
		Name:                   name,
		QuotedComment:          strconv.Quote(method.Description),
		GraphQLOutputType:      responseType,
		Arguments:              args,
		RequestResolver:        graphql.ResolverCall(file.OutputPkg, "Resolve"+pascalize(method.OperationID)+"Params"),
		RequestResolverWithErr: true,
		ClientMethodCaller: func(client, req string, ctx graphql.BodyContext) string {
			var res string
			var err error
			if successResponse.ResultType.Kind() == parser.KindNull {
				res, err = p.renderNullMethodCaller(reqType.String(ctx.Importer), req, client, pascalize(method.OperationID))
			} else {
				respType, err := p.goTypeByParserType(file, successResponse.ResultType, true)
				if err != nil {
					panic(errors.Wrap(err, "failed to resolve result go type"))
				}
				res, err = p.renderMethodCaller(respType.String(ctx.Importer), reqType.String(ctx.Importer), req, client, pascalize(method.OperationID))
			}
			if err != nil {
				panic(errors.Wrap(err, "failed to render method caller"))
			}
			return res
		},
		RequestType:          reqType,
		PayloadErrorChecker:  nil,
		PayloadErrorAccessor: nil,
		GraphQLOutputDataLoaderType: func(ctx graphql.BodyContext) string {
			resType, ok := successResponse.ResultType.(*parser.Array)

			if !ok {
				panic("Response type must be array")
			}

			subResType, ok := successResponse.ResultType.(*parser.Array)

			var outType parser.Type

			outType = resType

			if ok {
				outType = subResType.ElemType
			}

			dataLoaderOutType, err := p.TypeOutputTypeResolver(file, outType, false)

			if err != nil {
				panic(err)
			}

			return dataLoaderOutType(ctx)
		},
		DataLoaderResponseType: func() (graphql.GoType, error) {
			resType, ok := successResponse.ResultType.(*parser.Array)

			if !ok {
				return graphql.GoType{}, errors.New("Response type must be array")
			}

			goType, err := p.goTypeByParserType(file, resType.ElemType, true)

			if err != nil {
				return graphql.GoType{}, err
			}

			return goType, nil
		},
		DataLoaderFetch: func(importer *importer.Importer) string {
			paramsType := reqType.ElemType.String(importer)
			argName := ucFirst(method.Parameters[0].Name)
			methodName := ucFirst(method.OperationID)

			return `
			params := &` + paramsType + `{
				` + argName + `: keys,
				Context: ctx,
			}

			response, err := client.` + methodName + `(params)

			if err != nil {
				return nil, []error{err}
			}

			return response.Payload, nil
			`
		},
	}, nil
}

func ucFirst(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func (p *Plugin) tagQueriesMethods(tagCfg TagConfig, file *parsedFile, tag parser.Tag) ([]graphql.Method, error) {
	var res []graphql.Method
	for _, method := range tag.Methods {
		var methodCfg MethodConfig
		if tagCfg.Methods[method.Path] != nil {
			methodCfg = tagCfg.Methods[method.Path][strings.ToLower(method.HTTPMethod)]
		}
		if methodCfg.RequestType == "" {
			if method.HTTPMethod != "GET" {
				continue
			}
		} else if methodCfg.RequestType != "QUERY" {
			continue
		}

		meth, err := p.graphqlMethod(methodCfg, file, tag, method)
		if err != nil {
			if err == ErrMultipleSuccessResponses {
				fmt.Println("Warning: Method: ", method.Path, "have multiple successful responses. I'll skip it")
				continue
			}
			return nil, errors.Wrap(err, "failed to resolve graphql method")
		}
		res = append(res, *meth)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Name > res[j].Name
	})
	return res, nil
}
func (p *Plugin) tagMutationsMethods(tagCfg TagConfig, file *parsedFile, tag parser.Tag) ([]graphql.Method, error) {
	var res []graphql.Method
	for _, method := range tag.Methods {
		var methodCfg MethodConfig
		if tagCfg.Methods[method.Path] != nil {
			methodCfg = tagCfg.Methods[method.Path][strings.ToLower(method.HTTPMethod)]
		}
		if methodCfg.RequestType == "" {
			if method.HTTPMethod == "GET" {
				continue
			}
		} else if methodCfg.RequestType != "MUTATION" {
			continue
		}
		meth, err := p.graphqlMethod(methodCfg, file, tag, method)
		if err != nil {
			if err == ErrMultipleSuccessResponses {
				fmt.Println("Warning: Method: ", method.Path, "have multiple successful responses. I'll skip it")
				continue
			}
			return nil, errors.Wrap(err, "failed to resolve graphql method")
		}
		res = append(res, *meth)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Name > res[j].Name
	})
	return res, nil
}
func (p *Plugin) fileServices(file *parsedFile) ([]graphql.Service, error) {
	var res []graphql.Service
	for _, tag := range file.File.Tags {
		tagCfg, ok := file.Config.Tags[tag.Name]

		if !ok {
			fmt.Println("Warning: Skip tag:", tag.Name, "from", file.Config.Path)
			continue
		}

		if tagCfg.ClientGoPackage == "" {
			return nil, errors.Errorf("file: `%s`. Need to specify tag %s `client_go_package` option", file.Config.Name, tag.Name)
		}
		queriesMethods, err := p.tagQueriesMethods(*tagCfg, file, tag)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get tag queries methods")
		}

		mutationsMethods, err := p.tagMutationsMethods(*tagCfg, file, tag)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get tag mutations methods")
		}

		name := pascalize(tag.Name)

		if tagCfg.ServiceName != "" {
			name = tagCfg.ServiceName
		}
		res = append(res, graphql.Service{
			Name:            name,
			QuotedComment:   strconv.Quote(tag.Description),
			QueryMethods:    queriesMethods,
			MutationMethods: mutationsMethods,
			CallInterface: graphql.GoType{
				Kind: reflect.Interface,
				Pkg:  tagCfg.ClientGoPackage,
				Name: "IClient",
			},
		})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Name > res[j].Name
	})
	return res, nil
}
