package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasTelefone = []Rota{
	{
		URI:                "/telefone",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCliente,
		RequerAutenticacao: false,
	},
	{
		URI:                "/telefone",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarClientes,
		RequerAutenticacao: false,
	},
	{
		URI:                "/telefone/{telefoneId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarClientes,
		RequerAutenticacao: false,
	},
	{
		URI:                "/telefone/{telefoneId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarCliente,
		RequerAutenticacao: false,
	},
	{
		URI:                "/telefone/{telefoneId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarCliente,
		RequerAutenticacao: false,
	},
}
