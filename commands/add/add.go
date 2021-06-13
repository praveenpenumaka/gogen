package add

import (
	"errors"
	"github.com/gogen/domain"
	"github.com/gogen/services"
)

func Model(basePath, name string, params map[string]string) error {

	m, err := services.GetProjectConfig(basePath)
	if err != nil {
		return err
	}

	for _, model := range m.Models {
		if model.Name == name {
			return errors.New("model already exists, regenerate for generating code")
		}
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

	return services.WriteProjectConfig(basePath, m)
}

func CRUD(basePath, model string) error {
	m, err := services.GetProjectConfig(basePath)
	if err != nil {
		return err
	}

	for _, crud := range m.Router.Cruds {
		if crud.Name == model {
			return errors.New("crud already exists")
		}
	}
	newCrud := domain.Crud{Name: model}

	m.AddCrud(&newCrud)

	return services.WriteProjectConfig(basePath, m)
}

func AppContext(basepath, name, ntype string) error {
	m, err := services.GetProjectConfig(basepath)
	if err != nil {
		return err
	}

	for _, subcontext := range m.AppContext.Subcontexts {
		if subcontext.Name == name {
			return errors.New("appcontext already exists")
		}
	}
	c := domain.Context{
		Name: name,
		Type: ntype,
	}
	m.AddContext(&c)

	return services.WriteProjectConfig(basepath, m)
}

func Controller(basePath, model string) error {
	m, err := services.GetProjectConfig(basePath)
	if err != nil {
		return err
	}

	for _, controller := range m.Controllers {
		if controller.Name == model {
			return errors.New("controller already exists")
		}
	}
	newController := domain.Controller{Name: model, Ctype: "default", DoNotOverwrite: false}

	m.AddController(&newController)

	return services.WriteProjectConfig(basePath, m)
}

func AuthController(basePath string) error {
	m, err := services.GetProjectConfig(basePath)
	if err != nil {
		return err
	}

	for _, controller := range m.Controllers {
		if controller.Ctype == "auth" {
			return errors.New("auth controller already exists")
		}
	}

	newController := domain.Controller{Name: "auth", Ctype: "auth", DoNotOverwrite: false}

	m.AddController(&newController)

	return services.WriteProjectConfig(basePath, m)
}

func Config(basePath, model string, params map[string]string) error {
	m, err := services.GetProjectConfig(basePath)
	if err != nil {
		return err
	}

	for _, config := range m.Configs {
		if config.Name == model {
			return errors.New("config already exists")
		}
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

	return services.WriteProjectConfig(basePath, m)

}
