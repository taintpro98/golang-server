package schema

import (
	"golang-server/module/core/storage"

	"github.com/graphql-go/graphql"
)

func NewGraphSchema(
	userStorage storage.IUserStorage,
) (graphql.Schema, error) {
	resolver := NewResolver(userStorage)
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

	// Định nghĩa Root Query với resolve function
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
				Resolve: resolver.ResolveUser,
			},
			"users": &graphql.Field{
				Type:    graphql.NewList(userType),
				Resolve: resolver.ResolveUsers,
			},
		},
	})

	// var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	// 	Name: "RootMutation",
	// 	Fields: graphql.Fields{
	// 		"createUser": &graphql.Field{
	// 			Type: userType,
	// 			Args: graphql.FieldConfigArgument{
	// 				"name": &graphql.ArgumentConfig{
	// 					Type: graphql.String,
	// 				},
	// 				"email": &graphql.ArgumentConfig{
	// 					Type: graphql.String,
	// 				},
	// 			},
	// 			Resolve: createUser,
	// 		},
	// 	},
	// })
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation,
	})
}
