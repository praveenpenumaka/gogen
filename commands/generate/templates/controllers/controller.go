package controllers

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
	//go:embed controller.tmpl
	controllerTemplateString string

	//go:embed auth.tmpl
	authControllerTemplateString string
)

func GetAll(basepath string, config domain.Application) (map[string]string, error) {
	all := make(map[string]string)
	controllerList := config.Controllers

	authtmpl, parseErr := template.New("authcontroller").Funcs(template.FuncMap{
		"ToCap":       utils.ToCap,
		"ToLowerCase": strings.ToLower,
		"Basepath":    func(args ...string) string { return basepath },
	}).Parse(authControllerTemplateString)
	if parseErr != nil {
		return nil, parseErr
	}

	tmpl, parseErr := template.New("controller").Funcs(template.FuncMap{
		"ToCap":       utils.ToCap,
		"ToLowerCase": strings.ToLower,
		"Basepath":    func(args ...string) string { return basepath },
	}).Parse(controllerTemplateString)
	if parseErr != nil {
		return nil, parseErr
	}

	for _, model := range controllerList {
		if !model.DoNotOverwrite {
			bs := bytes.NewBufferString("")
			m := model
			if model.Ctype == "auth" {
				err := authtmpl.Execute(bs, "")
				if err == nil {
					all["controllers/auth.go"] = bs.String()
				}
			} else {
				err := tmpl.Execute(bs, m)
				if err == nil {
					filePath := fmt.Sprintf("controllers/%s.go", strings.ToLower(m.Name))
					all[filePath] = bs.String()
				}
			}
		}
	}

	return all, nil
}
