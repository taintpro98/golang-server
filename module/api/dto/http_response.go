package dto

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func isEmptyArray(i interface{}) bool {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		return val.Len() == 0
	}
	return false
}

type PageSizeResponse struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

type SuccessResponse struct {
	Data       interface{}       `json:"data"`
	Pagination *PageSizeResponse `json:"pagination,omitempty"`
}

func HandleSuccess(ctx *gin.Context, data interface{}, metadata ...interface{}) {
	if isEmptyArray(data) {
		ctx.JSON(http.StatusOK, SuccessResponse{Data: []int{}})
		return
	}
	result := SuccessResponse{
		Data: data,
	}

	for _, field := range metadata {
		paginate, ok := field.(PageSizeResponse)
		if ok {
			result.Pagination = &paginate
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func HandleResponse(ctx *gin.Context, data interface{}, err error, metadata ...interface{}) {
	if err == nil {
		HandleSuccess(ctx, data)
	}
}
