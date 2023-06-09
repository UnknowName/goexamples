package packages

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Response []ChannelStatistics

type ChannelStatistics struct {
	ConnectType string `json:"log_type"`
	ChannelName string `json:"pay_type"`
	ECSCnt      int    `json:"num_ecs"`
	PayCnt      int    `json:"num_pay"`
}

func getStatistics() *Response {
	envURL := os.Getenv("BASE_URL")
	url := fmt.Sprintf("%s%d", envURL, time.Now().UnixMilli())
	_resp, err := http.Get(url)
	if err != nil {
		log.Println("get ", url, "failed: ", err)
		return nil
	}
	var resp Response
	bytes, err := io.ReadAll(_resp.Body)
	if err != nil {
		log.Println("read body", err)
		return nil
	}
	defer _resp.Body.Close()
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		log.Println("json unmarshal", err)
		return nil
	}
	return &resp
}

type countCollector struct {
	metrics []metric
}

type metric struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
}

func (cc *countCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- cc.metrics[0].desc
}

func (cc *countCollector) Collect(ch chan<- prometheus.Metric) {
	resp := getStatistics()
	for _, stat := range *resp {
		ch <- prometheus.MustNewConstMetric(cc.metrics[0].desc, prometheus.GaugeValue, float64(stat.PayCnt), stat.ChannelName)
		ch <- prometheus.MustNewConstMetric(cc.metrics[1].desc, prometheus.GaugeValue, float64(stat.ECSCnt), stat.ChannelName)
	}
}

func NewCollector() prometheus.Collector {
	metrics := []metric{
		{prometheus.NewDesc("pay_cnt", "pay cnt", []string{"pay_channel"}, nil), prometheus.GaugeValue},
		{prometheus.NewDesc("ecs_cnt", "ecs cnt", []string{"pay_channel"}, nil), prometheus.GaugeValue},
	}
	return &countCollector{metrics: metrics}
}

func main() {
	prometheus.MustRegister(NewCollector())
	http.Handle("/metrics", promhttp.Handler())
	log.Print("expose /metrics use port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
