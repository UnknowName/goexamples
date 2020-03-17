package template

import (
	"fmt"
	"html/template"
	"os"
)

type TemplData struct {
	Name string
	Other string
}

func EchoTemplate(){
	file, err := os.Create("index.html")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := TemplData{Name:"Name", Other: "Other"}
	const templ = `<html><h5>{{.Name}}<h5><h6>{{.Other}}</h6></html>`
	fmt.Printf("Origin templ string is: '%s'\n", templ)
	t := template.Must(template.New("temp").Parse(templ))
	err = t.Execute(file, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}