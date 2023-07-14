package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasEmpresa = []Rota{
	{
		URI:                "/criarempresa",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarEmpresa,
		RequerAutenticacao: false,
	},
	{
		URI:                "/buscarempresa/{empresaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarEmpresa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/buscarempresas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarEmpresas,
		RequerAutenticacao: true,
	},
	{
		URI:                "/atualizarempresa/{empresaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarEmpresa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/deletarempresa/{empresaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarEmpresa,
		RequerAutenticacao: true,
	},
}
