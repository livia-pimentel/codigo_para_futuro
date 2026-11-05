package main

import (
	"net/http"

	"github.com/colibriproject-dev/colibri-sdk-go"
	"github.com/colibriproject-dev/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/colibriproject-dev/colibri-sdk-go/pkg/web/restserver"
	"github.com/livia-pimentel/codigo-para-o-futuro/oportunidades"
)

func main() {

	// Função para listar oportunidades
	listarOportunidadesFn := func (wctx restserver.WebContext) {
		result, err := sqlDB.NewQuery[oportunidades.Oportunidade](
			wctx.Context(),
			"SELECT id, titulo, empresa, categoria, ativa FROM oportunidades",
		).Many()
		if err != nil {
			wctx.ErrorResponse(http.StatusInternalServerError, err)
		}

		wctx.JsonResponse(http.StatusOK, result)
	}

	// Função para criar oportunidade
	cadastrarOportunidadesFn := func (wctx restserver.WebContext) {
		var body oportunidades.CriarOportunidadeRequest
		if err := wctx.DecodeBody(&body); err != nil {
			wctx.ErrorResponse(http.StatusBadRequest, err)
			return
		}

		if err := sqlDB.NewStatement(
			wctx.Context(),
			"INSERT INTO oportunidades (titulo, empresa, categoria) VALUES ($1, $2, $3)",
			body.Titulo,
			body.Empresa,
			body.Categoria,
		).Execute(); err != nil {
			wctx.ErrorResponse(http.StatusInternalServerError, err)
			return
		}
		
		wctx.EmptyResponse(http.StatusCreated)
	}

	// Rotas
	rotasDaMinhaAplicacao := []restserver.Route{
		{
			URI: "oportunidades",
			Method: http.MethodGet,
			Function: listarOportunidadesFn,
			Prefix: restserver.NoPrefix,
		},
		{
			URI: "oportunidades",
			Method: http.MethodPost,
			Function: cadastrarOportunidadesFn,
			Prefix: restserver.NoPrefix,
		},
	}

	colibri.InitializeApp()
	sqlDB.Initialize()

	restserver.AddRoutes(rotasDaMinhaAplicacao)
	restserver.ListenAndServe()
}