package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasEndereco = []Rota{
	{
		URI:                "/enderecos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarEndereco,
		RequerAutenticacao: false,
	},
	{
		URI:                "/enderecos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarEnderecos,
		RequerAutenticacao: false,
	},
	{
		URI:                "/enderecos/{enderecoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarEndereco,
		RequerAutenticacao: false,
	},
	{
		URI:                "/enderecos/{enderecoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarEndereco,
		RequerAutenticacao: false,
	},
}
