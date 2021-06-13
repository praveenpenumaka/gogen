package project

import (
	_ "embed"
	"github.com/gogen/domain"
	"github.com/gogen/services"
)

var (
	//go:embed project.tmpl
	TemplateString string
)

func Generate(basepath string, config domain.Application) (error error) {

	g := services.Generator{
		Basepath:   basepath,
		Template:   TemplateString,
		Data:       config.Project,
		ModulePath: "",
		FileName:   "package.json",
	}

	err := g.Generate()
	if err != nil {
		error = err
	}
	return
}
