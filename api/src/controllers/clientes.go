package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarCliente(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cliente models.Cliente
	if erro = json.Unmarshal(corpoRequest, &cliente); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = cliente.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeClientes(db)
	cliente.IDCliente, erro = repository.Criar(cliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, cliente)
}

func BuscarClientes(w http.ResponseWriter, r *http.Request) {
	// nomeOuNick := strings.ToLower(r.URL.Query().Get("cliente"))

	// db, erro := banco.Conectar()
	// if erro != nil {
	// 	respostas.Erro(w, http.StatusInternalServerError, erro)
	// 	return
	// }
	// defer db.Close()

	// repositorio := repositories.NovoRepositorioDeClientes(db)
	// clientes, erro := repositories.Buscar(nomeOuNick)
}

func BuscarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um cliente!"))
}

func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando cliente!"))
}

func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando cliente!"))
}
