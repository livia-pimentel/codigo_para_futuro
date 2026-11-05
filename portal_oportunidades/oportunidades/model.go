package oportunidades

type Oportunidade struct {
	ID int `json:"id"`
	Titulo string `json:"titulo"`
	Empresa  string `json:"empresa"`
	Categoria  string `json:"categoria"`
	Ativa  bool `json:"ativa"`
}

type CriarOportunidadeRequest struct {
	Titulo string `json:"titulo" validate:"required,min=3,max=120"`
	Empresa  string `json:"empresa" validate:"required"`
	Categoria  string `json:"categoria" validate:"required"`
}