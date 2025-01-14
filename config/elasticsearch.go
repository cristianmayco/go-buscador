package config

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var ElasticsearchClient *elasticsearch.Client

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.30.156:9200",
		},
		Username: "elastic", // substitua "elastic" pelo nome de usuário correto
		Password: "elastic", // substitua "senha" pela senha correta
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	ElasticsearchClient = client
}
