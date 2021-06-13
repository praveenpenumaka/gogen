package router

import (
	_ "embed"
	"fmt"
	"github.com/gogen/domain"
	"github.com/gogen/services"
	"strings"
)

var (
	//go:embed crud.tmpl
	crudTemplateString string
	//go:embed router.tmpl
	routerTemplateString string
)

func GenerateRouter(basepath string, config domain.Application) error {

	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "server",
		Template:   routerTemplateString,
		Data:       config,
		FileName:   "router.go",
	}

	return g.Generate()
}

func GenerateCruds(basepath string, config domain.Crud) error {

	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "server",
		Template:   crudTemplateString,
		Data:       config,
		FileName:   fmt.Sprintf("%s.go", strings.ToLower(config.Name)),
	}

	return g.Generate()
}

func Generate(basepath string, config domain.Application) (error error) {
	routerContextConfig := config.Router
	err := GenerateRouter(basepath, config)
	if err != nil {
		return err
	}

	subConfigs := routerContextConfig.Cruds
	for _, subconfig := range subConfigs {
		sErr := GenerateCruds(basepath, subconfig)
		if sErr != nil {
			error = sErr
		}
	}

	return
}
