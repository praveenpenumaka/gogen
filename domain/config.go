package domain

import (
	_ "embed"
)

type ConfigField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Config struct {
	Name   string        `json:"name"`
	Fields []ConfigField `json:"fields"`
}