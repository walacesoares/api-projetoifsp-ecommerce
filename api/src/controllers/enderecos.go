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

func CriarEndereco(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var endereco models.Endereco
	if erro = json.Unmarshal(corpoRequest, &endereco); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// if erro = endereco.Preparar(); erro != nil {
	// 	respostas.Erro(w, http.StatusBadRequest, erro)
	// 	return
	// }

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeEnderecos(db)
	endereco.IDEndereco, erro = repository.Criar(endereco)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, endereco)
}

func BuscarEnderecos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os endereços!"))
}

func AtualizarEndereco(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando endereço!"))
}

func DeletarEndereco(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando endereço!"))
}
