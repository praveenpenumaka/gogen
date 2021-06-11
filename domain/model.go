package domain

type ModelParam struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	Labels           string `json:"labels"`
	NotRequiredField bool   `json:"notrequired"`
}

func (mp ModelParam) Required() bool {
	return !mp.NotRequiredField
}

type Model struct {
	Name   string       `json:"name"`
	Params []ModelParam `json:"params"`
}
