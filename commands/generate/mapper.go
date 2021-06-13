package generate

import (
	"github.com/gogen/commands/generate/templates/appcontext"
	"github.com/gogen/commands/generate/templates/config"
	"github.com/gogen/commands/generate/templates/controllers"
	"github.com/gogen/commands/generate/templates/env"
	"github.com/gogen/commands/generate/templates/makefile"
	"github.com/gogen/commands/generate/templates/models"
	"github.com/gogen/commands/generate/templates/project"
	"github.com/gogen/commands/generate/templates/router"
	"github.com/gogen/commands/generate/templates/start"
	"github.com/gogen/domain"
)

func GetMapper() map[string]func(basepath string, config domain.Application) (error error) {
	mapper := make(map[string]func(basepath string, config domain.Application) (error error))
	mapper["project"] = project.Generate
	mapper["models"] = models.Generate
	mapper["appcontext"] = appcontext.Generate
	mapper["configs"] = config.Generate
	mapper["controllers"] = controllers.Generate
	mapper["makefile"] = makefile.Generate
	mapper["router"] = router.Generate
	mapper["main"] = start.Generate
	mapper[".env.sample"] = env.Generate
	return mapper
}
