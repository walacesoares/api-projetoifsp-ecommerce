package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasServico = []Rota{
	{
		URI:                "/criarservico",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateServicoEndpoint,
		RequerAutenticacao: false,
	},
	{
		URI:                "/buscarservico/{servicoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetServicoEndpoint,
		RequerAutenticacao: false,
	},
	{
		URI:                "/atualizarservico/{servicoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateServicoEndpoint,
		RequerAutenticacao: false,
	},
	{
		URI:                "/deletarservico/{servicoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteServicoEndpoint,
		RequerAutenticacao: false,
	},
}
