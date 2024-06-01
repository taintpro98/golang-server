package elastic

import (
	"context"
	"fmt"
	"golang-server/config"
	"golang-server/pkg/logger"

	"github.com/elastic/go-elasticsearch/v7"
)

func New(ctx context.Context, config *config.ElasticConfig) (*elasticsearch.Client, error) {
	escfg := elasticsearch.Config{
		Addresses: []string{config.Addresses},
		Username:  config.Username,
		Password:  config.Password,
	}

	es, err := elasticsearch.NewClient(escfg)
	if err != nil {
		logger.Error(ctx, err, "elastic search connection err")
		return nil, err
	}
	// Ping Elasticsearch
	res, err := es.Ping()
	if err != nil {
		logger.Error(ctx, err, "Error pinging Elasticsearch")
		return nil, err
	}
	// Check the response status
	if res.IsError() {
		logger.Error(ctx, err, fmt.Sprintf("Elasticsearch ping error: %s", res.String()))
	}
	return es, nil
}
