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
	return es, nil
}

// create index if not exist
func CreateSupplierSettingIndex(es *elasticsearch.Client, index string) error {
	// check exist index
	res, err := es.Indices.Exists([]string{index})
	if err != nil {
		return fmt.Errorf("cannot check index existence: %w", err)
	}
	if res.StatusCode == 200 {
		return nil
	}
	if res.StatusCode != 404 {
		return fmt.Errorf("error in index existence response: %s", res.String())
	}

	// TODO can be define mapping
	// in this case: do not need mapping because use user_id is _id in ES
	/*mapping := `
	    {
			"mappings": {
				"properties": {
					"user_id": {
						"type": "text"
					}
				}
			}
	    }`*/

	// create if not exist
	res, err = es.Indices.Create(
		index,
		//es.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return fmt.Errorf("cannot create index: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in index creation response: %s", res.String())
	}

	// update Alias
	res, err = es.Indices.PutAlias([]string{index}, index+"_alias")
	if err != nil {
		return fmt.Errorf("cannot create index alias: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in index alias creation response: %s", res.String())
	}

	return nil
}
