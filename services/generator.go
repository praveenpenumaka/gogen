package services

import (
	"bytes"
	"github.com/gogen/utils"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Generator struct {
	Basepath   string
	ModulePath string
	Template   string
	Data       interface{}
	FileName   string
}

func (g *Generator) EnsureModule() {
	if g.ModulePath == "" {
		return
	}
	_ = os.Mkdir(g.ModulePath, 0774)
}

func (g *Generator) Generate() error {
	g.EnsureModule()
	tmpl, parseErr := template.New(g.FileName).Funcs(template.FuncMap{
		"ToCap":       utils.ToCap,
		"ToLowerCase": strings.ToLower,
		"ToUpperCase": strings.ToUpper,
		"Basepath":    func(arg0 string, args ...string) string { return g.Basepath },
	}).Parse(g.Template)
	if parseErr != nil {
		return parseErr
	}
	bs := bytes.NewBufferString("")
	err := tmpl.Execute(bs, g.Data)
	if err != nil {
		return err
	}
	filePath := g.ModulePath + "/" + g.FileName
	if g.ModulePath == "" {
		filePath = g.FileName
	}
	return ioutil.WriteFile(filePath, bs.Bytes(), 0744)
}
