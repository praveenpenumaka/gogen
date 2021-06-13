package controllers

import (
	_ "embed"
	"fmt"
	"github.com/gogen/domain"
	"github.com/gogen/services"
	"strings"
)

var (
	//go:embed controller.tmpl
	controllerTemplateString string

	//go:embed auth.tmpl
	authControllerTemplateString string
)

func getTemplate(subtype string) string {
	if subtype == "auth" {
		return authControllerTemplateString
	}
	return controllerTemplateString
}

func Generate(basepath string, config domain.Application) (error error) {

	for _, controller := range config.Controllers {
		if !controller.DoNotOverwrite {
			g := services.Generator{
				Basepath:   basepath,
				ModulePath: "controllers",
				Template:   getTemplate(controller.Ctype),
				Data:       controller,
				FileName:   fmt.Sprintf("%s.go", strings.ToLower(controller.Name)),
			}
			err := g.Generate()
			if err != nil {
				error = err
			}
		}
	}

	return
}
