package domain

type Crud struct {
	Name string `json:"name"`
}

type Upload struct {
	Name string `json:"name"`
	Directory string `json:"directory"`
}

type Static struct {
	Name string `json:"name"`
	Directory string `json:"directory"`
}

type Router struct {
	Cruds []Crud `json:"cruds"`
	Uploads []Upload `json:"uploads"`
	Statics []Static `json:"statics"`
}

func (r *Router) AddCrud(c *Crud) {
	r.Cruds = append(r.Cruds, *c)
}
