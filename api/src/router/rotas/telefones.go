package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasTelefone = []Rota{
	{
		URI:                "/criartelefone",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTelefone,
		RequerAutenticacao: true,
	},
	{
		URI:                "/buscartelefones",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTelefones,
		RequerAutenticacao: true,
	},
	{
		URI:                "/buscartelefone/{telefoneId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTelefone,
		RequerAutenticacao: true,
	},
	{
		URI:                "/atualizartelefone/{telefoneId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarTelefone,
		RequerAutenticacao: true,
	},
	{
		URI:                "/deletartelefone/{telefoneId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarTelefone,
		RequerAutenticacao: true,
	},
}
