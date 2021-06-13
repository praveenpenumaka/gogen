package start

import (
	_ "embed"
	"github.com/gogen/domain"
	"github.com/gogen/services"
)

var (
	//go:embed main.tmpl
	mainTemplateString string
)

func Generate(basepath string, config domain.Application) error {
	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "",
		Template:   mainTemplateString,
		Data:       config,
		FileName:   "main.go",
	}
	return g.Generate()
}
