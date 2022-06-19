package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasEmpresa = []Rota{
	{
		URI:                "/empresas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarEmpresa,
		RequerAutenticacao: false,
	},
	{
		URI:                "/empresas/{empresaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarEmpresa,
		RequerAutenticacao: false,
	},
	{
		URI:                "/empresas/{empresaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarEmpresa,
		RequerAutenticacao: false,
	},
	{
		URI:                "/empresas/{empresaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarEmpresa,
		RequerAutenticacao: false,
	},
}
