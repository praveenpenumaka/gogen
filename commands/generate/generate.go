package generate

import (
	"encoding/json"
	"fmt"
	"github.com/gogen/domain"
	"io/ioutil"
	"os"
)


func getProjectConfig(cwd string) (*domain.Application, error) {
	projectFile := fmt.Sprintf("%s/project.json", cwd)
	project, err := ioutil.ReadFile(projectFile)
	if err != nil {
		return nil, err
	}
	m := &domain.Application{}
	json.Unmarshal(project, m)
	return m, nil
}

func All(cwd string, components ...string) error {
	app,err:=getProjectConfig(cwd)
	if err != nil {
		return err
	}

	eMap := GetMapper()

	if len(components) == 0 {
		components = GetAllComponents()
	}

	pPath := app.Project.Basepath

	for _, k := range components {
		if k != "makefile" && k != "main" && k != ".env.sample" {
			me := os.Mkdir(k,0774)
			if me != nil {
				fmt.Println(me)
			}
		}
		v := eMap[k]
		files,err := v(pPath,*app)
		if err != nil {
			return err
		}
		for fileName, fileContent := range files {
			e := ioutil.WriteFile(fileName,[]byte(fileContent),0774)
			if e != nil {
				return e
			}
			fmt.Println("Writing file:"+fileName)
		}
	}
	return nil
}