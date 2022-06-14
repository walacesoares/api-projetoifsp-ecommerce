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
	{
		URI:                "/ordensservicos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarOrdensServicos,
		RequerAutenticacao: false,
	},
	{
		URI:                "/ordensservicos/{ordensservicos}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarOrdemServico,
		RequerAutenticacao: false,
	},
	{
		URI:                "/ordensservicos/{ordensservicosId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarOrdemServico,
		RequerAutenticacao: false,
	},
	{
		URI:                "/ordensservicos/{ordensservicosId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarOrdemServico,
		RequerAutenticacao: false,
	},
}
