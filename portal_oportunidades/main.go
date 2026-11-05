package main

import (
	"net/http"

	"github.com/colibriproject-dev/colibri-sdk-go"
	"github.com/colibriproject-dev/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/colibriproject-dev/colibri-sdk-go/pkg/web/restserver"
	"github.com/livia-pimentel/codigo-para-o-futuro/oportunidades"
)

func main() {

	listaOportunidade := []oportunidades.Oportunidade{}

	listaOportunidade = append(listaOportunidade, oportunidades.CriarNovaOportunidade("Curso de Java", true))
	listaOportunidade = append(listaOportunidade, oportunidades.CriarNovaOportunidade("Curso de C++", true))

	println("####### TODAS AS OPORTUNIDADES #######")
	oportunidades.ImprimirLista(listaOportunidade)
	println("\n####### OPORTUNIDADES ATIVAS #######")
	oportunidades.ImprimirSoOportunidadesAtivas(listaOportunidade)

	listarOportunidades := func (wctx restserver.WebContext) {
	wctx.JsonResponse(http.StatusOK, listaOportunidade)
	}

	
	rotasDaMinhaAplicacao := []restserver.Route{
		{
			URI: "oportunidades",
			Method: http.MethodGet,
			Function: listarOportunidades,
			Prefix: restserver.NoPrefix,
		},
	}

	colibri.InitializeApp()
	sqlDB.Initialize()

	restserver.AddRoutes(rotasDaMinhaAplicacao)
	restserver.ListenAndServe()
}