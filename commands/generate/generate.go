package generate

import (
	"github.com/gogen/services"
	"log"
)

func All(cwd string, components ...string) error {
	app, err := services.GetProjectConfig(cwd)
	if err != nil {
		return err
	}

	eMap := GetMapper()

	if len(components) == 0 {
		components = app.GetComponents()
	}

	pPath := app.Project.Basepath

	for _, k := range components {
		log.Println("generating " + k)
		v := eMap[k]
		err := v(pPath, *app)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return nil
}
