package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gogen/domain"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

const initData = `{
	"project":{
		"name": "{{ .Name }}",
		"version": "0.0.1",
		"basepath": "{{ .BasePath }}"
	},
	"models": [],
	"configs": [],
	"appcontext": {},
	"controllers": [],
	"router": {}
}`

func Project(basePath, name string) error {
	finalPath := fmt.Sprintf("%s/%s", basePath, name)
	direrr := os.Mkdir(finalPath, 0755)
	if direrr != nil {
		return direrr
	}
	projectMeta := fmt.Sprintf("%s/project.json", finalPath)

	tmpl, parseErr := template.New("projectmeta").Parse(initData)
	if parseErr != nil {
		return parseErr
	}

	bs := bytes.NewBufferString("")
	m := make(map[string]string)
	m["Name"] = name
	goPath := os.Getenv("GOPATH")+"/src/"
	spl := strings.Split(finalPath,goPath)
	m["BasePath"] = strings.TrimSpace(spl[1])
	err := tmpl.Execute(bs, m)
	if err != nil {
		return nil
	}
	fmt.Println("Creating meta file at :" + projectMeta)
	return ioutil.WriteFile(projectMeta, bs.Bytes(), 0755)
}

func getProjectConfig(basepath string) (*domain.Application, error) {
	projectFile := fmt.Sprintf("%s/project.json", basepath)
	project, err := ioutil.ReadFile(projectFile)
	if err != nil {
		return nil, err
	}
	m := &domain.Application{}
	json.Unmarshal(project, m)
	return m, nil
}

func writeProjectConfig(basepath string, m *domain.Application) error {
	projectFile := fmt.Sprintf("%s/project.json", basepath)
	bs, e := json.MarshalIndent(m, "", " ")
	if e != nil {
		return e
	}
	return ioutil.WriteFile(projectFile, bs, 0755)
}

func Model(basePath, name string, params map[string]string) error {

	m, err := getProjectConfig(basePath)
	if err != nil {
		return err
	}

	var paramList []domain.ModelParam
	for key, value := range params {
		param := domain.ModelParam{
			Name: key,
			Type: value,
		}
		paramList = append(paramList, param)
	}

	newModel := domain.Model{
		Name:   name,
		Params: paramList,
	}

	m.AddModel(&newModel)

	return writeProjectConfig(basePath, m)
}

func CRUD(basePath, model string) error {
	m, err := getProjectConfig(basePath)
	if err != nil {
		return err
	}

	newCrud := domain.Crud{Name: model}

	m.AddCrud(&newCrud)

	return writeProjectConfig(basePath, m)
}

func AppContext(basepath, name, ntype string) error  {
	m, err := getProjectConfig(basepath)
	if err != nil {
		return err
	}

	c := domain.Context{
		Name:name,
		Type:ntype,
	}
	m.AddContext(&c)

	return writeProjectConfig(basepath, m)
}

func Controller(basePath, model string) error {
	m, err := getProjectConfig(basePath)
	if err != nil {
		return err
	}

	newController := domain.Controller{Name: model}

	m.AddController(&newController)

	return writeProjectConfig(basePath, m)
}

func Config(basePath, model string, params map[string]string) error {
	m, err := getProjectConfig(basePath)
	if err != nil {
		return err
	}

	var paramList []domain.ConfigField
	for key, value := range params {
		param := domain.ConfigField{
			Name: key,
			Type: value,
		}
		paramList = append(paramList, param)
	}

	newConfig := domain.Config{
		Name:   model,
		Fields: paramList,
	}

	m.AddConfig(&newConfig)

	return writeProjectConfig(basePath, m)

}
