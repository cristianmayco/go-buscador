package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go-buscador/config"
	"go-buscador/models"
	"log"
	"strings"
)

func SearchDocuments(searchParams models.GlobalSearchElastic) ([]models.DocumentModel, error) {
	client := config.ElasticsearchClient

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"extractedText": searchParams.Filter,
						},
					},
				},
				"filter": []interface{}{
					map[string]interface{}{
						"terms": map[string]interface{}{
							"masterContextName": searchParams.Modules,
						},
					},
				},
			},
		},
	}

	var buf strings.Builder
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	reader := strings.NewReader(buf.String())

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("document"),
		client.Search.WithBody(reader),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithFrom((searchParams.Page-1)*searchParams.Limit),
		client.Search.WithSize(searchParams.Limit),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	var documents []models.DocumentModel
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var doc models.DocumentModel
		source := hit.(map[string]interface{})["_source"]
		sourceBytes, err := json.Marshal(source)
		if err != nil {
			log.Printf("Failed to marshal source: %v", err)
			continue
		}
		if err := json.Unmarshal(sourceBytes, &doc); err != nil {
			log.Printf("Failed to unmarshal document: %v", err)
			continue
		}
		documents = append(documents, doc)
	}
	return documents, nil
}

func SearchByWord(word string) ([]models.DocumentModel, error) {
	client := config.ElasticsearchClient

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"extractedText": word,
			},
		},
	}

	var buf strings.Builder
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	reader := strings.NewReader(buf.String())

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("document"),
		client.Search.WithBody(reader),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithSize(100), // Limite de 100 documentos
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	var documents []models.DocumentModel
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var doc models.DocumentModel
		source := hit.(map[string]interface{})["_source"]
		sourceBytes, err := json.Marshal(source)
		if err != nil {
			log.Printf("Failed to marshal source: %v", err)
			continue
		}
		if err := json.Unmarshal(sourceBytes, &doc); err != nil {
			log.Printf("Failed to unmarshal document: %v", err)
			continue
		}
		documents = append(documents, doc)
	}
	return documents, nil
}

func GetAllDocuments() ([]models.DocumentModel, error) {
	client := config.ElasticsearchClient

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size": 1000, // Ajuste o tamanho conforme necessário
	}

	var buf strings.Builder
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	reader := strings.NewReader(buf.String())

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("document"),
		client.Search.WithBody(reader),
		client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	var documents []models.DocumentModel
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var doc models.DocumentModel
		source := hit.(map[string]interface{})["_source"]
		sourceBytes, err := json.Marshal(source)
		if err != nil {
			log.Printf("Failed to marshal source: %v", err)
			continue
		}
		if err := json.Unmarshal(sourceBytes, &doc); err != nil {
			log.Printf("Failed to unmarshal document: %v", err)
			continue
		}
		documents = append(documents, doc)
	}
	return documents, nil
}

func SearchByWords(words []string) ([]models.DocumentModel, error) {
	client := config.ElasticsearchClient

	// Construindo a consulta corretamente
	shouldQueries := make([]interface{}, len(words))
	for i, word := range words {
		shouldQueries[i] = map[string]interface{}{
			"match": map[string]interface{}{
				"extractedText": word,
			},
		}
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should":               shouldQueries,
				"minimum_should_match": 1,
			},
		},
	}

	var buf strings.Builder
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	reader := strings.NewReader(buf.String())

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("document"),
		client.Search.WithBody(reader),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithSize(100), // Limite de 100 documentos
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	var documents []models.DocumentModel
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var doc models.DocumentModel
		source := hit.(map[string]interface{})["_source"]
		sourceBytes, err := json.Marshal(source)
		if err != nil {
			log.Printf("Failed to marshal source: %v", err)
			continue
		}
		if err := json.Unmarshal(sourceBytes, &doc); err != nil {
			log.Printf("Failed to unmarshal document: %v", err)
			continue
		}
		documents = append(documents, doc)
	}
	return documents, nil
}
