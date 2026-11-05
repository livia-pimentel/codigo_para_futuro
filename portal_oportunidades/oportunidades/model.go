package oportunidades

type Oportunidade struct {
	ID int `json:"id"`
	Titulo string `json:"titulo"`
	Empresa  string `json:"empresa"`
	Categoria  string `json:"categoria"`
	Ativa  bool `json:"ativa"`
}

type CriarOportunidadeRequest struct {
	Titulo string `json:"titulo"`
	Empresa  string `json:"empresa"`
	Categoria  string `json:"categoria"`
}