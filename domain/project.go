package domain

type Project struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Basepath string `json:"basepath"`
}
