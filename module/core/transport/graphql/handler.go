package graphql_transport

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type GraphqlTransport struct {
	h *handler.Handler
}

func NewGraphqlTransport(
	graphSchema graphql.Schema,
) GraphqlTransport {
	h := handler.New(&handler.Config{
		Schema:   &graphSchema,
		Pretty:   true,
		GraphiQL: true, // Bật giao diện GraphiQL để dễ dàng kiểm thử
	})
	return GraphqlTransport{
		h: h,
	}
}

func (t GraphqlTransport) GraphQLHandler(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "GinContext", c)
	c.Request = c.Request.WithContext(ctx)
	t.h.ServeHTTP(c.Writer, c.Request)
}
