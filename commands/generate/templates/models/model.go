package models

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/gogen/domain"
	"github.com/gogen/utils"
	"strings"
	"text/template"
)

var (
	//go:embed model.tmpl
	modelTemplateString string
)


func GenerateModels(basepath string, config domain.Application) (map[string]string, error) {

	all := make(map[string]string)
	modelsList := config.Models

	tmpl, parseErr := template.New("config").Funcs(template.FuncMap{
		"ToCap": utils.ToCap,
		"ToLowerCase": strings.ToLower,
		"Basepath": func(arg0 string, args ...string) string { return basepath },
	}).Parse(modelTemplateString)
	if parseErr != nil {
		return nil, parseErr
	}
	for _, model := range modelsList {
		fmt.Println(model)
		bs := bytes.NewBufferString("")
		m := model
		err := tmpl.Execute(bs, m)
		if err != nil {
			return nil, err
		}
		// TODO: Should an error block the whole generation?
		filePath := fmt.Sprintf("models/%s.go", strings.ToLower(m.Name))
		all[filePath] = bs.String()
	}

	return all, nil
}
