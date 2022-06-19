package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasCliente
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasEmpresa...)
	rotas = append(rotas, rotasEndereco...)
	rotas = append(rotas, rotasTelefone...)
	rotas = append(rotas, rotasOrdemServico...)
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
