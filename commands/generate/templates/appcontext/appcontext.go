package appcontext

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/gogen/domain"
	"github.com/gogen/utils"
	"strings"
	"text/template"
)

var (
	//go:embed appcontext.tmpl
	appContextTemplateString string
	//go:embed config.tmpl
	configTemplateString string
	//go:embed db.tmpl
	dbTemplateString string
	//go:embed redis.tmpl
	redisTemplateString string
)

func getTemplate(basepath string) (*template.Template, error) {
	return template.New("appcontext").Funcs(template.FuncMap{
		"ToCap":    utils.ToCap,
		"Basepath": func(args ...string) string { return basepath },
	}).Parse(appContextTemplateString)
}

// GetAppContext accepts whole app configuration
func GetAppContext(basepath string, config domain.Application) (string, string, error) {
	tmpl, parseErr := getTemplate(basepath)
	if parseErr != nil {
		return "", "", parseErr
	}

	bs := bytes.NewBufferString("")
	err := tmpl.Execute(bs, config)
	if err != nil {
		return "", "", err
	}
	return "appcontext/appcontext.go", bs.String(), nil
}

func getSubTemplate(basepath string, name string) (*template.Template, error) {
	var tmplS string
	name = strings.ToLower(name)
	if name == "db" {
		tmplS = dbTemplateString
	} else if name == "redis" {
		tmplS = redisTemplateString
	} else if name == "config" {
		tmplS = configTemplateString
	} else {
		return nil, errors.New("invalid app context type:" + name)
	}
	return template.New("context").Funcs(template.FuncMap{
		"ToCap":    utils.ToCap,
		"Basepath": func(args ...string) string { return basepath },
	}).Parse(tmplS)
}

func GetSubContexts(basepath string, subconfig domain.Context) (string, string, error) {
	tmpl, parseErr := getSubTemplate(basepath, subconfig.Name)
	if parseErr != nil {
		return "", "", parseErr
	}

	bs := bytes.NewBufferString("")
	err := tmpl.Execute(bs, subconfig)
	if err != nil {
		return "", "", err
	}
	filePath := fmt.Sprintf("appcontext/%s.go", subconfig.Name)
	return filePath, bs.String(), nil
}

// GetAppContext accepts whole app configuration
func GetAll(basepath string, config domain.Application) (map[string]string, error) {
	all := make(map[string]string)
	appContextConfig := config.AppContext
	configPath, configContent, err := GetAppContext(basepath, config)
	if err != nil {
		return nil, err
	}
	all[configPath] = configContent

	if len(appContextConfig.Subcontexts) <= 0 {
		return all, nil
	}
	subConfigs := appContextConfig.Subcontexts
	for _, subconfig := range subConfigs {
		sConfigPath, sConfigContent, sErr := GetSubContexts(basepath, subconfig)
		if sErr != nil {
			return nil, sErr
		}
		all[sConfigPath] = sConfigContent
	}
	return all, nil
}
