package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/module/core/model"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type IElasticStorage interface {
	IndexUsers(ctx context.Context, users []model.UserModel) error
	SearchUsers(ctx context.Context, name string) ([]model.UserModel, error)
}

type elasticStorage struct {
	es *elasticsearch.Client
}

func NewElasticStorage(es *elasticsearch.Client) IElasticStorage {
	return elasticStorage{
		es: es,
	}
}

func (s elasticStorage) IndexUsers(ctx context.Context, users []model.UserModel) error {
	for _, user := range users {
		userJSON, err := json.Marshal(user)
		if err != nil {
			return err
		}

		indexReq := esapi.IndexRequest{
			Index:      "users",
			DocumentID: user.ID,
			Body:       strings.NewReader(string(userJSON)),
			Refresh:    "true", // Refresh the index after indexing the document
		}
		indexRes, err := indexReq.Do(ctx, s.es)
		if err != nil {
			return err
		}
		defer indexRes.Body.Close()
		if indexRes.IsError() {
			return fmt.Errorf("Elasticsearch indexing error: %s", indexRes.String())
		}
	}
	return nil
}

type Test struct {
}

func (s elasticStorage) SearchUsers(ctx context.Context, search string) ([]model.UserModel, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"wildcard": map[string]interface{}{
							"phone": fmt.Sprintf("*%s*", search),
						},
					},
					{
						"wildcard": map[string]interface{}{
							"email": fmt.Sprintf("*%s*", search),
						},
					},
				},
			},
		},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	searchReq := esapi.SearchRequest{
		Index: []string{"users"},
		Body:  strings.NewReader(string(queryJSON)),
	}

	searchRes, err := searchReq.Do(ctx, s.es)
	if err != nil {
		return nil, err
	}
	defer searchRes.Body.Close()

	if searchRes.IsError() {
		return nil, fmt.Errorf("Elasticsearch search error: %s", searchRes.String())
	}

	var result struct {
		Hits struct {
			Hits []struct {
				Source model.UserModel `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(searchRes.Body).Decode(&result); err != nil {
		return nil, err
	}

	var users []model.UserModel
	for _, hit := range result.Hits.Hits {
		users = append(users, hit.Source)
	}
	return users, nil
}
