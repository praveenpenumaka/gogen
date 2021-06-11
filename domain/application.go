package domain

import "errors"

type Application struct {
	Project     Project      `json:"project"`
	Models      []Model      `json:"models"`
	Router      Router       `json:"router"`
	Controllers []Controller `json:"controllers"`
	Configs     []Config     `json:"configs"`
	AppContext  AppContext   `json:"appcontext"`
}

func (a *Application) AddModel(m *Model) error {
	if m == nil {
		return errors.New("Empty model")
	}
	a.Models = append(a.Models, *m)
	return nil
}

func (a *Application) AddCrud(m *Crud) error {
	if m == nil {
		return errors.New("Empty crud")
	}
	a.Router.AddCrud(m)
	return nil
}

func (a *Application) AddController(c *Controller) error {
	if c == nil {
		return errors.New("empty controller")
	}
	a.Controllers = append(a.Controllers, *c)
	return nil
}

func (a *Application) AddConfig(c *Config) error {
	if c == nil {
		return errors.New("empty config")
	}
	a.Configs = append(a.Configs, *c)
	return nil
}

func (a *Application) AddContext(c *Context) error {
	if c == nil {
		return errors.New("empty context ")
	}
	return a.AppContext.AddSubContext(c)
}
