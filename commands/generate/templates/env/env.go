package env

import (
	_ "embed"
	"github.com/gogen/domain"
	"github.com/gogen/services"
)

var (
	//go:embed env.tmpl
	envTemplateString string
)

func Generate(basepath string, config domain.Application) error {

	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "",
		Template:   envTemplateString,
		Data:       config,
		FileName:   ".env.sample",
	}

	return g.Generate()
}
