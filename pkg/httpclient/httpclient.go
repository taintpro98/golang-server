package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang-server/pkg/constants"
	"golang-server/pkg/e"
	"golang-server/pkg/logger"

	"io"
	"net"
	"net/http"
	"time"
)

type DoRequestParam struct {
	Request *http.Request
	Headers map[string]string
	Output  interface{}
}

type HttpClient struct {
	Client *http.Client
}

func (c HttpClient) DoRequest(ctx context.Context, param DoRequestParam) error {
	param.Request.Header.Add("Content-Type", "application/json")
	for key, value := range param.Headers {
		param.Request.Header.Add(key, value)
	}
	requestID := ctx.Value(constants.TraceID)
	if requestID != nil {
		param.Request.Header.Add(constants.KeyRequestID, fmt.Sprintf("%s", requestID))
	}

	start := time.Now()
	res, err := c.Client.Do(param.Request)
	end := time.Since(start)

	if err != nil {
		logger.LogInfoRequest(ctx, end, *param.Request, http.Response{}, nil, nil)
		var netErr net.Error
		ok := errors.As(err, &netErr)
		if ok && netErr.Timeout() {
			return e.ErrTimeout
		}
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(ctx, err, "close body error")
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	logger.LogInfoRequest(ctx, end, *param.Request, *res, body, err)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		var errorBody e.CustomErr
		if err := json.Unmarshal(body, &errorBody); err != nil {
			return err
		}
		errorBody.HttpStatusCode = res.StatusCode
		return errorBody
	}
	if param.Output != nil {
		if err := json.Unmarshal(body, param.Output); err != nil {
			logger.Error(ctx, err, "")
			return err
		}
	}
	return nil
}
