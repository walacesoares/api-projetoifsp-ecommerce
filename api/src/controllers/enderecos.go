package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func BuscarEndereco(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	enderecoID, erro := strconv.ParseUint(parametros["enderecoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeEnderecos(db)
	endereco, erro := repository.BuscarPorID(enderecoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, endereco)
}

func AtualizarEndereco(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	enderecoID, erro := strconv.ParseUint(parametros["enderecoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var endereco models.Endereco
	if erro = json.Unmarshal(corpoRequisicao, &endereco); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = endereco.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeEnderecos(db)
	if erro = repository.Atualizar(enderecoID, endereco); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarEndereco(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	enderecoID, erro := strconv.ParseUint(parametros["enderecoId"], 10, 64)

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeEnderecos(db)
	if erro = repository.Deletar(enderecoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
