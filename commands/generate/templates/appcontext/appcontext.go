package appcontext

import (
	_ "embed"
	"github.com/gogen/domain"
	"github.com/gogen/services"
	"log"
	"strings"
)

var (
	//go:embed appcontext.tmpl
	appContextTemplateString string
	//go:embed config.tmpl
	configTemplateString string
	//go:embed db.tmpl
	dbTemplateString string
	//go:embed redis.tmpl
	redisTemplateString string
)

// GenerateAppContext accepts whole app configuration
func GenerateAppContext(basepath string, config domain.Application) error {
	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "appcontext",
		Template:   appContextTemplateString,
		Data:       config,
		FileName:   "appcontext.go",
	}
	return g.Generate()
}

func getSubTemplate(name string) string {
	name = strings.ToLower(name)
	if name == "db" {
		return dbTemplateString
	} else if name == "redis" {
		return redisTemplateString
	} else if name == "config" {
		return configTemplateString
	}

	panic("invalid app context type")
	return ""
}

func GenerateSubContext(basepath string, subconfig domain.Context) error {
	tmpl := getSubTemplate(subconfig.Type)
	log.Println("generating sub context " + subconfig.Name)
	g := services.Generator{
		Basepath:   basepath,
		ModulePath: "appcontext",
		Template:   tmpl,
		Data:       subconfig,
		FileName:   subconfig.Name,
	}
	return g.Generate()
}

// GenerateAppContext accepts whole app configuration
func Generate(basepath string, config domain.Application) (error error) {
	appContextConfig := config.AppContext
	err := GenerateAppContext(basepath, config)
	if err != nil {
		return err
	}

	if len(appContextConfig.Subcontexts) <= 0 {
		return nil
	}
	subConfigs := appContextConfig.Subcontexts
	for _, subconfig := range subConfigs {
		sErr := GenerateSubContext(basepath, subconfig)
		if sErr != nil {
			error = sErr
		}
	}
	return
}
