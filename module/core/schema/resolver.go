package schema

import (
	"golang-server/module/core/dto"
	"golang-server/module/core/storage"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type IResolver interface {
	ResolveUser(p graphql.ResolveParams) (interface{}, error)
	ResolveUsers(p graphql.ResolveParams) (interface{}, error)
}

type resolver struct {
	userStorage storage.IUserStorage
}

func NewResolver(
	userStorage storage.IUserStorage,
) IResolver {
	return resolver{
		userStorage: userStorage,
	}
}

func (r resolver) ResolveUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if ok {
		ctx := p.Context.Value("GinContext").(*gin.Context)
		userDB, err := r.userStorage.FindOne(ctx, dto.FilterUser{
			ID: id,
		})
		return userDB, err
	}
	return nil, nil
}

func (r resolver) ResolveUsers(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context.Value("GinContext").(*gin.Context)
	usersDB, err := r.userStorage.List(ctx, dto.FilterUser{
		CommonFilter: dto.CommonFilter{
			Limit: 10,
		},
	})
	return usersDB, err
}

// func createUser(p graphql.ResolveParams) (interface{}, error) {
// 	user := map[string]interface{}{
// 		"id":    "2", // Thông thường, ID sẽ được tạo tự động
// 		"name":  p.Args["name"].(string),
// 		"email": p.Args["email"].(string),
// 	}
// 	users = append(users, user)
// 	return user, nil
// }
