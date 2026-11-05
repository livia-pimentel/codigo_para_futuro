package oportunidades

import "fmt"

func ImprimirLista(oportunidade []Oportunidade) {
	for _, valor := range oportunidade {
		fmt.Println(valor)
	}
}

func CriarNovaOportunidade(titulo string, statusOportunidade bool) Oportunidade{
	return Oportunidade{
		Titulo: titulo,
		Ativa: statusOportunidade,
	}
}

func ImprimirSoOportunidadesAtivas(oportunidade []Oportunidade) {
	for _, valor := range oportunidade {
		if valor.Ativa == true {
			fmt.Println(valor)
		}
	}
}