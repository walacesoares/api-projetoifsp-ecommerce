package rotas

import (
	"api/src/middlewares"
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
	rotas = append(rotas, rotasUsuario...)
	rotas = append(rotas, rotasServico...)
	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}
