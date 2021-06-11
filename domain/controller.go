package domain

type Controller struct {
	Name           string `json:"name"`
	Ctype          string `json:"type" default:"default"`
	DoNotOverwrite bool   `json:"donotoverwrite"`
}
