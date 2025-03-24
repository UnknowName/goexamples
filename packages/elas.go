package packages

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"os"
	"strings"
)

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
	log.Println(record)
	return int(total)
}

func SearchWrite(es *elasticsearch.Client, search, saveFile string) error {
	file, err := os.OpenFile(saveFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("open file error", err)
	}
	defer file.Close()
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
	total := record["hits"].(map[string]interface{})["total"].(float64)
	log.Println("当前取到条数为", total)
	cnt := 0
	for _, hit := range hits.([]interface{}) {
		item := hit.(map[string]interface{})["_source"].(map[string]interface{})
		msg := item["full_message"]
		msgStr := fmt.Sprintf("%s", msg)
		cnt++
		if ! strings.HasPrefix(msgStr,"接口") {
			continue
		}
		infos := strings.Split(fmt.Sprintf("%s", msg),"请求报文：")
		info := strings.Split(fmt.Sprintf("%s", infos[1]), "错误信息")
		writeStr := fmt.Sprintf("%s\n", info[0])
		line := strings.TrimSuffix(writeStr, "\n")
		if line == "\n" {
			continue
		}
		_, err = file.WriteString(fmt.Sprintf("%s\n", line))
		if err != nil {
			log.Fatalln(err)
		}
	}
	if cnt != int(total) {
		log.Fatalln("写入数据不对")
	}
	log.Println("总写入", cnt, "条数")
	return nil
}