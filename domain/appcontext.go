package domain

import (
	_ "embed"
)

type Context struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AppContext struct {
	Subcontexts []Context `json:"subcontexts"`
}

func (a AppContext) HasDB() bool {
	for _, subcontext := range a.Subcontexts {
		if subcontext.Name == "DB" {
			return true
		}
	}
	return false
}

func (a *AppContext) AddSubContext(c *Context) error {
	a.Subcontexts = append(a.Subcontexts, *c)
	return nil
}
