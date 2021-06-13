package makefile

import (
	_ "embed"
	"github.com/gogen/domain"
	"github.com/gogen/services"
)

var (
	//go:embed makefile.tmpl
	makefileTemplateString string
)

func Generate(basepath string, config domain.Application) error {

	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "",
		Template:   makefileTemplateString,
		Data:       config,
		FileName:   "Makefile",
	}
	return g.Generate()
}
