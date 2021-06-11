package generate

import (
	"github.com/gogen/commands/generate/templates/appcontext"
	"github.com/gogen/commands/generate/templates/config"
	"github.com/gogen/commands/generate/templates/controllers"
	"github.com/gogen/commands/generate/templates/env"
	"github.com/gogen/commands/generate/templates/makefile"
	"github.com/gogen/commands/generate/templates/models"
	"github.com/gogen/commands/generate/templates/router"
	"github.com/gogen/commands/generate/templates/start"
	"github.com/gogen/domain"
)

func GetAllComponents() []string {
	return []string {
		"models",
		"appcontext",
		"config",
		"controllers",
		"makefile",
		"server",
		"main",
		".env.sample",
	}
}
func GetMapper() map[string]func(basepath string, config domain.Application) (map[string]string,error) {
	mapper := make(map[string]func(basepath string, config domain.Application) (map[string]string,error))
	mapper["models"] = models.GenerateModels
	mapper["appcontext"] = appcontext.GetAll
	mapper["config"] = config.GetAll
	mapper["controllers"] = controllers.GetAll
	mapper["makefile"] = makefile.GenerateMakefile
	mapper["server"] = router.GetAll
	mapper["main"] = start.GenerateMain
	mapper[".env.sample"] = env.GetEnv
	return mapper
}
