package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasOrdemServico = []Rota{
	{
		URI:                "/criarordemservico",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarOrdemServico,
		RequerAutenticacao: true,
	},
	{
		URI:                "/buscarordensservicos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarOrdensServicos,
		RequerAutenticacao: true,
	},
	{
		URI:                "/buscarordemservico/{ordensservicoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarOrdemServico,
		RequerAutenticacao: true,
	},
	// {
	// 	URI:                "/ordensservico/{ordensservicoId}",
	// 	Metodo:             http.MethodPut,
	// 	Funcao:             controllers.AtualizarOrdemServico,
	// 	RequerAutenticacao: true,
	// },
	{
		URI:                "/deletarordemservico/{ordensservicoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarOrdemServico,
		RequerAutenticacao: true,
	},
}
