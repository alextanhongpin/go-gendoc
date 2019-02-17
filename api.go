// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/policypalnet/go-test/service/namesvc"
)

const (
	SRC  = "API.tpl"
	DEST = "API.md"
)

type Template map[string]interface{}

func RenderMarkdown(writer io.Writer, t Template) {
	for k, v := range t {
		b, _ := json.MarshalIndent(v, "", "  ")
		t[k] = template.HTML(string(b))
	}
	tmpl, err := template.ParseFiles(SRC)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(writer, t)
	fmt.Println("generated documentation")
}

func main() {
	t := Template{}
	t["GetNameRequest"] = namesvc.GetNameRequest{"john doe"}
	f, err := os.Create(DEST)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	RenderMarkdown(f, t)
}
