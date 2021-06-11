package start

import (
	"bytes"
	_ "embed"
	"github.com/gogen/domain"
	"github.com/gogen/utils"
	"html/template"
)

var (
	//go:embed main.tmpl
	mainTemplateString string
)

func GenerateMain(basepath string, config domain.Application) (map[string]string, error) {
	all := make(map[string]string)

	tmpl, tmplError := template.New("mainfile").Funcs(template.FuncMap{
		"ToCap": utils.ToCap,
		"Basepath": func(args ...string) string { return basepath },
	}).Parse(mainTemplateString)
	if tmplError != nil {
		return nil, tmplError
	}

	// TODO: Handle this case
	output := bytes.NewBufferString("")
	tmpl.Execute(output, config)

	all["main.go"] = output.String()
	return all, nil
}
