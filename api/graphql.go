package api

import (
	"net/http"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func graphQLHandler(debug bool) http.Handler {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootQuery",
				Fields: graphql.Fields{
					"hello": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return "world", nil
						},
					},
				},
			},
		),
	})

	h := gqlhandler.New(&gqlhandler.Config{
		Schema:   &schema,
		Pretty:   debug,
		GraphiQL: debug,
	})
	return h
}
