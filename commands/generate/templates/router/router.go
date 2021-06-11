package router

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
	//go:embed crud.tmpl
	crudTemplateString string
	//go:embed router.tmpl
	routerTemplateString string
)

func GenerateRouter(basepath string, config domain.Application) (string, string, error) {
	tmpl, parseErr := template.New("router").Funcs(template.FuncMap{
		"ToCap": utils.ToCap,
		"Basepath": func(args ...string) string { return basepath },
	}).Parse(routerTemplateString)
	if parseErr != nil {
		return "", "", parseErr
	}

	bs := bytes.NewBufferString("")
	err := tmpl.Execute(bs, config)
	if err != nil {
		return "", "", err
	}
	return "server/router.go", bs.String(), nil
}

func GenerateCruds(basepath string, config domain.Crud) (string, string, error) {
	tmpl, parseErr := template.New("crud").Funcs(template.FuncMap{
		"ToCap": utils.ToCap,
		"Basepath": func(args ...string) string { return basepath },
	}).Parse(crudTemplateString)
	if parseErr != nil {
		return "", "", parseErr
	}

	bs := bytes.NewBufferString("")
	err := tmpl.Execute(bs, config)
	if err != nil {
		return "", "", err
	}
	fileName := fmt.Sprintf("server/%s.go", strings.ToLower(config.Name))
	return fileName, bs.String(), nil
}

func GetAll(basepath string, config domain.Application) (map[string]string, error) {
	all := make(map[string]string)
	routerContextConfig := config.Router
	configPath, configContent, err := GenerateRouter(basepath, config)
	if err != nil {
		return nil, err
	}
	all[configPath] = configContent

	subConfigs := routerContextConfig.Cruds
	for _, subconfig := range subConfigs {
		sConfigPath, sConfigContent, sErr := GenerateCruds(basepath, subconfig)
		if sErr != nil {
			return nil, sErr
		}
		all[sConfigPath] = sConfigContent
	}

	return all, nil
}
