package graphql

import (
	"net/http"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func InitHandler(debug bool) (http.Handler, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		return nil, err
	}

	h := gqlhandler.New(&gqlhandler.Config{
		Schema:   &schema,
		Pretty:   debug,
		GraphiQL: debug,
	})
	return h, nil
}
