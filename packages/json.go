package packages

import (
	"encoding/json"
	"fmt"
)

func Map2Json() {
	dict := make(map[string]string)
	dict["name"] = "unknowname"
	dict["country"] = "china"
	jsonByte, err := json.Marshal(dict)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonByte))
}

func Json2Map() {
	jsonStr := `{"name":"unknowname", "country": "china"}`
	dict := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &dict)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dict)
}
