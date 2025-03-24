package packages

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	address := "http://localhost:9200"
	query := `{start:%s,end:%s}`
	timeTemplate1 := "2006-01-02 15:04:05.000"
	_startTime := "2023-12-14 10:24:00.000"
	_finalTime := "2023-12-15 09:30:00.000"
	initTime, _ := time.ParseInLocation(timeTemplate1, _startTime, time.Local)
	finalTime, _ := time.ParseInLocation(timeTemplate1, _finalTime, time.Local)
	startTime := initTime.Add(time.Minute).UTC()
	endTime := startTime.Add(time.Second * 10 - time.Millisecond).UTC()
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	cnt := 1
	for finalTime.After(startTime) {
		fileNumber := cnt / 60
		fileName := fmt.Sprintf("t2_file_%d.txt", fileNumber)
		searchQuery := fmt.Sprintf(query, startTime.Format(timeTemplate1), endTime.Format(timeTemplate1))
		log.Println(startTime, endTime, fileName)
		if SearchWrite(es, searchQuery, fileName) != nil {
			log.Fatalln(startTime, "to", endTime.String(), "get failed")
		}
		startTime = endTime.Add(time.Millisecond)
		endTime = startTime.Add(time.Second*10 - time.Millisecond)
		cnt++
	}
}
