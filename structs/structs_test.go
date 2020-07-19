package structs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// JSON数据序列化成struct数据
func TestElastic_Echo(t *testing.T) {
	var es Elastic
	resp, err := http.Get("http://128.0.255.10:9200")
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP GET ERROR ", err)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&es)
	if err != nil {
		fmt.Println("JSON decode error ", err)
		return
	}
	fmt.Printf("%v", es)
	jsonByte := es.ToJson()
	fmt.Println(string(jsonByte))
}

func TestElastic_ToJson(t *testing.T) {
	// 如果结构体在初始化时不赋值，则使用此方法
	var es Elastic
	es.Name = "test"
	es.Version.Number = "6.6.1"
	es.Version.BuildHash = "sxxxx"
	fmt.Println(es)
	jsonByte := es.ToJson()
	fmt.Println(string(jsonByte))
}

func TestAnother_Echo(t *testing.T) {
	// 初始化方法一
	a := &Another{Version: Version{Number: "6.6.1", BuildHash: "SSNNCSXX"}}
	a.GetVersion()
	a.Echo()
	// 初始化方法二
	var b Another
	b.Number = "7.9.8"
	b.GetVersion()
	b.Echo()
}
