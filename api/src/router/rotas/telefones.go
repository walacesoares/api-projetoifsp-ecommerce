package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasTelefone = []Rota{
	{
		URI:                "/telefones",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTelefone,
		RequerAutenticacao: false,
	},
	{
		URI:                "/telefones",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTelefones,
		RequerAutenticacao: false,
	},
	{
		URI:                "/telefones/{telefoneId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTelefone,
		RequerAutenticacao: false,
	},
	{
		URI:                "/telefones/{telefoneId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarTelefone,
		RequerAutenticacao: false,
	},
	{
		URI:                "/telefones/{telefoneId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarTelefone,
		RequerAutenticacao: false,
	},
}
