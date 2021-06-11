package makefile

import (
	"bytes"
	_ "embed"
	"github.com/gogen/domain"
	"github.com/gogen/utils"
	"text/template"
)

var (
	//go:embed makefile.tmpl
	makefileTemplateString string
)


func GenerateMakefile(basepath string, config domain.Application) (map[string]string, error){
	all := make(map[string]string)
	tmpl, parseErr := template.New("make").Funcs(template.FuncMap{
		"ToCap": utils.ToCap,
		"Basepath": func(arg0 string, args ...string) string { return basepath },
	}).Parse(makefileTemplateString)
	if parseErr != nil {
		return nil, parseErr
	}

	bs := bytes.NewBufferString("")
	err := tmpl.Execute(bs, config)
	if err != nil {
		return nil, err
	}
	all["Makefile"] =bs.String()
	return all,nil
}
