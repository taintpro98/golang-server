package http_client_api

import (
	"context"
	http_integration_dto "golang-server/module/http_integration/dto"
	"golang-server/pkg/httpclient"
	"golang-server/pkg/logger"
	"golang-server/pkg/utils"
	"net/http"
	"time"
)

type IHttpClientApi interface {
	GetPrices(ctx context.Context, data http_integration_dto.GetPricesRequest) (http_integration_dto.GetPricesResponse, error)
}

type httpClientApi struct {
	httpclient.HttpClient
}

func NewHttpClientApi() IHttpClientApi {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	return httpClientApi{
		HttpClient: httpclient.HttpClient{
			Client: client,
		},
	}
}

// GetPrices implements IHttpClientApi.
func (c httpClientApi) GetPrices(ctx context.Context, data http_integration_dto.GetPricesRequest) (http_integration_dto.GetPricesResponse, error) {
	var result http_integration_dto.GetPricesResponse
	paramsReader, err := utils.ConvertToReader(data)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "", paramsReader)
	if err != nil {
		logger.Error(ctx, err, "")
		return result, err
	}
	if err = c.DoRequest(ctx, httpclient.DoRequestParam{
		Request: req,
	}, &result, nil); err != nil {
		logger.Error(ctx, err, "Error sending request to clevertap push notification")
		return result, err
	}
	logger.Info(ctx, "send to clevertap success")
	return result, nil
}
