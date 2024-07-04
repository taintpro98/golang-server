package graphql

import (
	"github.com/graphql-go/graphql"
)

// Định nghĩa User type
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// Định nghĩa Root Query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(string)
				if ok {
					// Giả lập dữ liệu
					return map[string]interface{}{
						"id":    id,
						"name":  "John Doe",
						"email": "john@example.com",
					}, nil
				}
				return nil, nil
			},
		},
	},
})

// Tạo schema
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
