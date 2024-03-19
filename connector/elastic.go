package connector

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/shing1211/go-lib/config"
	log "github.com/sirupsen/logrus"
)

// Elastic search global variable  for access
var EsClient *elasticsearch.Client

func GetESClient() *elasticsearch.Client {
	if EsClient == nil {
		if err := InitESConnection; err != nil {
			log.Warn(err)
			return nil
		}
	}
	return EsClient
}

func InitESConnection(config config.ElasticConfig) *error {
	esURL := "http://" + config.ElasticHost + ":" + config.ElasticPort
	esUser := config.ElasticUser
	esPwd := config.ElasticPwd

	log.Info("Connecting to elasticsearch at: " + esURL)

	esConfig := elasticsearch.Config{
		Addresses: []string{esURL},
		Username:  esUser,
		Password:  esPwd,
	}

	if EsClient, err = elasticsearch.NewClient(esConfig); err != nil {
		log.Warn(err)
		return &err
	}

	if EsClient != nil {
		log.Info("Connected to elasticsearch at: " + esURL)

		res, err := EsClient.Info()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()
		// Check response status
		if res.IsError() {
			log.Fatalf("Error: %s", res.String())
		}
		// Deserialize the response into a map.
		//if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		//	log.Fatalf("Error parsing the response body: %s", err)
		//}
		// Print client and server version numbers.
		log.Printf("Client: %s", elasticsearch.Version)
		//log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
		log.Println(strings.Repeat("~", 37))
	}
	return nil
}
