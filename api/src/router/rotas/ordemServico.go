package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasOrdemServico = []Rota{
	{
		URI:                "/ordensservicos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarOrdemServico,
		RequerAutenticacao: false,
	},
	// {
	// 	URI:                "/ordensservicos",
	// 	Metodo:             http.MethodGet,
	// 	Funcao:             controllers.BuscarOrdensServicos,
	// 	RequerAutenticacao: false,
	// },
	{
		URI:                "/ordensservico/{ordensservico}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarOrdemServico,
		RequerAutenticacao: false,
	},
	// {
	// 	URI:                "/ordensservico/{ordensservicoId}",
	// 	Metodo:             http.MethodPut,
	// 	Funcao:             controllers.AtualizarOrdemServico,
	// 	RequerAutenticacao: false,
	// },
	{
		URI:                "/ordensservico/{ordensservicoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarOrdemServico,
		RequerAutenticacao: false,
	},
}
