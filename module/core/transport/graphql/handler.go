package graphql_transport

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// User type
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

var users = []map[string]interface{}{
	{"id": "1", "name": "John Doe", "email": "john@example.com"},
}

// Root Query
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
					for _, user := range users {
						if user["id"] == id {
							return user, nil
						}
					}
				}
				return nil, nil
			},
		},
		"users": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return users, nil
			},
		},
	},
})

// Root Mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := map[string]interface{}{
					"id":    "2", // Thông thường, ID sẽ được tạo tự động
					"name":  p.Args["name"].(string),
					"email": p.Args["email"].(string),
				}
				users = append(users, user)
				return user, nil
			},
		},
	},
})

// Tạo schema
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

type GraphqlTransport struct {
	h *handler.Handler
}

func NewGraphqlTransport() GraphqlTransport {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // Bật giao diện GraphiQL để dễ dàng kiểm thử
	})
	return GraphqlTransport{
		h: h,
	}
}

func (t GraphqlTransport) GraphQLHandler(ctx *gin.Context) {
	t.h.ServeHTTP(ctx.Writer, ctx.Request)
}
