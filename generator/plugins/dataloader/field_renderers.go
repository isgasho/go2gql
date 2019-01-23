package dataloader

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"

	"github.com/EGT-Ukraine/go2gql/generator/plugins/graphql"
)

type fieldsRenderer struct {
	dataLoader *DataLoader
}

func (r *fieldsRenderer) RenderFields(o graphql.OutputObject, ctx graphql.BodyContext) (string, error) {
	templateFuncs := map[string]interface{}{
		"goType": func(typ graphql.GoType) string {
			return typ.String(ctx.Importer)
		},
		"gqlPkg": func() string {
			return ctx.Importer.New(graphql.GraphqlPkgPath)
		},
		"loadersPkg": func() string {
			return ctx.Importer.New(r.dataLoader.Pkg)
		},
		"graphqlOutputLoaderTypeName": func(ctx graphql.BodyContext, dataLoaderName string) string {
			dataLoaderConfig := r.dataLoader.Loaders[dataLoaderName]

			return dataLoaderConfig.OutputGraphqlType(ctx)
		},
	}

	buf := new(bytes.Buffer)
	tmpl, err := templatesOutput_object_fieldsGohtmlBytes()
	if err != nil {
		return "", errors.Wrap(err, "failed to get output fields template")
	}
	bodyTpl, err := template.New("head").Funcs(templateFuncs).Parse(string(tmpl))
	if err != nil {
		return "", errors.Wrap(err, "failed to parse template")
	}
	err = bodyTpl.Execute(buf, graphql.RenderFieldsContext{
		OutputObject:  o,
		ObjectContext: ctx,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to execute template")
	}

	return buf.String(), nil
}