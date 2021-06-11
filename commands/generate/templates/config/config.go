package config

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
	//go:embed config.tmpl
	configTemplateString string
	//go:embed subconfig.tmpl
	subConfigTemplateString string
)

func GenerateConfig(basepath string, config domain.Application) (string, string, error) {

	tmpl, parseErr := template.New("config").Funcs(template.FuncMap{
		"ToCap":    utils.ToCap,
		"Basepath": func(arg0 string, args ...string) string { return basepath },
	}).Parse(configTemplateString)
	if parseErr != nil {
		return "", "", parseErr
	}

	bs := bytes.NewBufferString("")
	err := tmpl.Execute(bs, config)
	if err != nil {
		return "", "", err
	}
	return "config/config.go", bs.String(), nil
}

func GenerateSubconfig(basepath string, config domain.Config) (string, string, error) {
	tmpl, parseErr := template.New("subconfig").Funcs(template.FuncMap{
		"ToCap":       utils.ToCap,
		"ToUpperCase": strings.ToUpper,
		"Basepath":    func(arg0 string, args ...string) string { return basepath },
	}).Parse(subConfigTemplateString)
	if parseErr != nil {
		return "", "", parseErr
	}

	bs := bytes.NewBufferString("")
	err := tmpl.Execute(bs, config)
	if err != nil {
		return "", "", err
	}
	fileName := fmt.Sprintf("config/%s.go", config.Name)
	return fileName, bs.String(), nil
}

func GetAll(basepath string, config domain.Application) (map[string]string, error) {
	all := make(map[string]string)

	configPath, configContent, err := GenerateConfig(basepath, config)
	if err != nil {
		return nil, err
	}
	all[configPath] = configContent

	subConfigs := config.Configs
	for _, subconfig := range subConfigs {
		sConfigPath, sConfigContent, sErr := GenerateSubconfig(basepath, subconfig)
		if sErr != nil {
			return nil, sErr
		}
		all[sConfigPath] = sConfigContent
	}

	return all, nil
}
