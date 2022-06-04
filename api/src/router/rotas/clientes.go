package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasCliente = []Rota{
	{
		URI:                "/clientes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCliente,
		RequerAutenticacao: false,
	},
	{
		URI:                "/clientes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarClientes,
		RequerAutenticacao: false,
	},
	{
		URI:                "/clientes/{clienteId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarClientes,
		RequerAutenticacao: false,
	},
	{
		URI:                "/clientes/{clienteId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarCliente,
		RequerAutenticacao: false,
	},
	{
		URI:                "/clientes/{clienteId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarCliente,
		RequerAutenticacao: false,
	},
}
