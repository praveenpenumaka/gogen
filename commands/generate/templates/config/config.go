package config

import (
	_ "embed"
	"fmt"
	"github.com/gogen/domain"
	"github.com/gogen/services"
	"strings"
)

var (
	//go:embed config.tmpl
	configTemplateString string
	//go:embed subconfig.tmpl
	subConfigTemplateString string
)

func GenerateConfig(basepath string, config domain.Application) error {

	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "config",
		Template:   configTemplateString,
		Data:       config,
		FileName:   "config.go",
	}
	return g.Generate()
}

func GenerateSubconfig(basepath string, config domain.Config) error {
	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "config",
		Template:   subConfigTemplateString,
		Data:       config,
		FileName:   fmt.Sprintf("%s.go", strings.ToLower(config.Name)),
	}

	return g.Generate()
}

func Generate(basepath string, config domain.Application) (error error) {

	err := GenerateConfig(basepath, config)
	if err != nil {
		return err
	}

	subConfigs := config.Configs
	for _, subconfig := range subConfigs {
		sErr := GenerateSubconfig(basepath, subconfig)
		if sErr != nil {
			error = sErr
		}
	}

	return
}
