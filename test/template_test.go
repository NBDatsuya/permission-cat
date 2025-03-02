package test

import (
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	templateText := `
		Output0: {{title .Name1}}
		Output1: {{title .Name2}}
		Output2: {{.Name3|title}}
	`

	funcMap := template.FuncMap{
		"title": strings.ToTitle,
	}
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)

	data := map[string]string{
		"Name1": "go",
		"Name2": "Programming",
		"Name3": "Cat",
	}

	_ = tpl.Execute(os.Stdout, data)
}
