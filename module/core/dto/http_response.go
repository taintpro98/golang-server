package dto

import (
	"errors"
	"golang-server/pkg/e"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func isEmptyArray(i interface{}) bool {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		return val.Len() == 0
	}
	return false
}

type SuccessResponse struct {
	Data       interface{}       `json:"data"`
	Pagination *PaginateResponse `json:"pagination,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
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
		paginate, ok := field.(PaginateResponse)
		if ok {
			result.Pagination = &paginate
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func getMessageCodeError(err error) (int, int, string) {
	// validation error
	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if ok {
		return http.StatusBadRequest, http.StatusBadRequest, err.Error()
	}

	// handle customer error
	var _err e.CustomErr
	ok = errors.As(err, &_err)
	if ok {
		return _err.HttpStatusCode, _err.Code, _err.Error()
	}
	return http.StatusInternalServerError, http.StatusInternalServerError, err.Error()
}

func HandleFailed(ctx *gin.Context, err error) {
	statusCode, msgCode, msg := getMessageCodeError(err)
	ctx.JSON(
		statusCode, ErrorResponse{
			Code:    msgCode,
			Message: msg,
		},
	)
}

func HandleResponse(ctx *gin.Context, data interface{}, err error, metadata ...interface{}) {
	if err == nil {
		HandleSuccess(ctx, data, metadata...)
	} else {
		HandleFailed(ctx, err)
	}
}
