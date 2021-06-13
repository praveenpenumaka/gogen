package models

import (
	_ "embed"
	"fmt"
	"github.com/gogen/domain"
	"github.com/gogen/services"
	"strings"
)

var (
	//go:embed model.tmpl
	modelTemplateString string
)

func Generate(basepath string, config domain.Application) (error error) {

	for _, model := range config.Models {
		g := services.Generator{
			Basepath:   basepath,
			Template:   modelTemplateString,
			Data:       model,
			ModulePath: "models",
			FileName:   fmt.Sprintf("%s.go", strings.ToLower(model.Name)),
		}
		err := g.Generate()
		if err != nil {
			error = err
		}
	}
	return
}
