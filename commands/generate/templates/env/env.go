package env

import (
	"bytes"
	_ "embed"
	"github.com/gogen/domain"
	"github.com/gogen/utils"
	"strings"
	"text/template"
)

var (
	//go:embed env.tmpl
	envTemplateString string
)

func GetEnv(basepath string, config domain.Application) (map[string]string,error) {
	all := make(map[string]string)

	tmpl, tmplError := template.New("env").Funcs(template.FuncMap{
		"ToCap": utils.ToCap,
		"ToUpperCase": strings.ToUpper,
		"Basepath": func(args ...string) string { return basepath },
	}).Parse(envTemplateString)
	if tmplError != nil {
		return nil, tmplError
	}

	// TODO: Handle this case
	output := bytes.NewBufferString("")
	tmpl.Execute(output, config)

	all[".env.sample"] = output.String()
	return all, nil
}
