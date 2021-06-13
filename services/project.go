package services

import (
	"encoding/json"
	"fmt"
	"github.com/gogen/domain"
	"io/ioutil"
)

func GetProjectConfig(basepath string) (*domain.Application, error) {
	projectFile := fmt.Sprintf("%s/project.json", basepath)
	project, err := ioutil.ReadFile(projectFile)
	if err != nil {
		return nil, err
	}
	m := &domain.Application{}
	json.Unmarshal(project, m)
	return m, nil
}

func WriteProjectConfig(basepath string, m *domain.Application) error {
	projectFile := fmt.Sprintf("%s/project.json", basepath)
	bs, e := json.MarshalIndent(m, "", " ")
	if e != nil {
		return e
	}
	return ioutil.WriteFile(projectFile, bs, 0755)
}
