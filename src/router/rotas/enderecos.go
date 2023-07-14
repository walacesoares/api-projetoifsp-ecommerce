package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasEndereco = []Rota{
	{
		URI:                "/criarendereco",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarEndereco,
		RequerAutenticacao: true,
	},
	{
		URI:                "/buscarendereco",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarEndereco,
		RequerAutenticacao: true,
	},
	{
		URI:                "/atualizarendereco/{enderecoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarEndereco,
		RequerAutenticacao: true,
	},
	{
		URI:                "/deletarendereco/{enderecoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarEndereco,
		RequerAutenticacao: true,
	},
}
