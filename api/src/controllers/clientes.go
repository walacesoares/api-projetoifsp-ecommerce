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

func BuscarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um cliente!"))
}

func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	// parametros := mux.Vars(r)
	// clienteID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	// if erro != nil {
	// 	respostas.Erro(w, http.StatusInternalServerError, erro)
	// 	return
	// }

	// corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	// if erro != nil {
	// 	respostas.Erro(w, http.StatusUnprocessableEntity, erro)
	// 	return
	// }

	// var cliente models.Cliente
	// if erro = json.Unmarshal(corpoRequisicao, &cliente); erro != nil {
	// 	respostas.Erro(w, http.StatusBadRequest, erro)
	// 	return
	// }
}

func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando cliente!"))
}
