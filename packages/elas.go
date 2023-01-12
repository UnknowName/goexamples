package packages

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"strings"
)

func Search(address string, query string) error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	_query := `{
	    "query": {
	        "bool": {
	            "filter": {"range": {"timestamp": {"gte":"2023-01-04 03:20:20.631","lte": "2023-01-04 04:25:20.631"}}},
				"must":   {"match": {"source": "sixunmall-jobhost1-79584ddd7b-pr6m5"}}
	        }
	    },
		"_source": ["source", "message"],
		"from": %d,
		"size": %d
	}`
	start := 0
	q := fmt.Sprintf(_query, start, 1)
	total := GetTotal(es, q)
	log.Println("Total", total)
	size := 20
	for total >= 0 {
		if SearchWrite(es, fmt.Sprintf(_query, start, size)) != nil {
			log.Println("Get", start, err)
			break
		}
		total -= size
		start += size
	}
	return nil
}

func GetTotal(es *elasticsearch.Client, search string) int {
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithBody(strings.NewReader(search)),
		es.Search.WithPretty(),
		es.Search.WithSort("timestamp"),
	)
	if err != nil {
		return GetTotal(es, search)
	}
	defer res.Body.Close()
	record := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
		return GetTotal(es, search)
	}
	total := record["hits"].(map[string]interface{})["total"].(float64)
	return int(total)
}

func SearchWrite(es *elasticsearch.Client, search string) error {
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithBody(strings.NewReader(search)),
		es.Search.WithPretty(),
		es.Search.WithSort("timestamp"),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	record := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
		return err
	}
	hits := record["hits"].(map[string]interface{})["hits"]
	if hits == nil {
		return nil
	}
	for _, hit := range hits.([]interface{}) {
		item := hit.(map[string]interface{})["_source"].(map[string]interface{})
		log.Println(item)
	}
	return nil
}
