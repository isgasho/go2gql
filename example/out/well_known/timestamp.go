// This file was generated by github.com/EGT-Ukraine/go2gql. DO NOT EDIT IT
package well_known

import (
	context "context"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	graphql "github.com/graphql-go/graphql"
	errors "github.com/pkg/errors"

	scalars "github.com/EGT-Ukraine/go2gql/api/scalars"
)

// Enums
// Input object
var TimestampInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:   "TimestampInput",
	Fields: graphql.InputObjectConfigFieldMap{},
})

func init() {
	TimestampInput.Fields()["seconds"] = &graphql.InputObjectField{PrivateName: "seconds", Type: scalars.GraphQLInt64Scalar}
	TimestampInput.Fields()["nanos"] = &graphql.InputObjectField{PrivateName: "nanos", Type: scalars.GraphQLInt32Scalar}
}

// Input objects resolvers
func ResolveTimestampInput(ctx context.Context, i interface{}) (_ *timestamp.Timestamp, rerr error) {
	if i == nil {
		return nil, nil
	}
	args := i.(map[string]interface{})
	_ = args
	var result = new(timestamp.Timestamp)
	if args["seconds"] != nil {
		result.Seconds = args["seconds"].(int64)
	}
	if args["nanos"] != nil {
		result.Nanos = args["nanos"].(int32)
	}

	return result, nil
}

// Output objects
var Timestamp = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Timestamp",
	Fields: graphql.Fields{},
})

func init() {
	Timestamp.AddFieldConfig("seconds", &graphql.Field{
		Name: "seconds",
		Type: scalars.GraphQLInt64Scalar,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			switch src := p.Source.(type) {
			case *timestamp.Timestamp:
				if src == nil {
					return nil, nil
				}
				s := *src
				return s.Seconds, nil
			case timestamp.Timestamp:
				return src.Seconds, nil
			}

			return nil, errors.New("source of unknown type")
		},
	})
	Timestamp.AddFieldConfig("nanos", &graphql.Field{
		Name: "nanos",
		Type: scalars.GraphQLInt32Scalar,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			switch src := p.Source.(type) {
			case *timestamp.Timestamp:
				if src == nil {
					return nil, nil
				}
				s := *src
				return s.Nanos, nil
			case timestamp.Timestamp:
				return src.Nanos, nil
			}

			return nil, errors.New("source of unknown type")
		},
	})
}

// Maps input objects
// Maps input objects resolvers
// Maps output objects
