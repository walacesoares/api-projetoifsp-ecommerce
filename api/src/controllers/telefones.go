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

func CriarTelefone(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var telefone models.Telefone
	if erro = json.Unmarshal(corpoRequest, &telefone); erro != nil {
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

	repository := repositories.NovoRepositorioDeTelefones(db)
	telefone.IDTelefone, erro = repository.Criar(telefone)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, telefone)
}

func BuscarTelefones(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os telefones!"))
}

func BuscarTelefone(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	telefoneID, erro := strconv.ParseUint(parametros["telefoneId"], 10, 64)
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

	repository := repositories.NovoRepositorioDeTelefones(db)
	telefone, erro := repository.BuscarPorID(telefoneID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, telefone)
}

func AtualizarTelefone(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	telefoneID, erro := strconv.ParseUint(parametros["telefoneId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var telefone models.Telefone
	if erro = json.Unmarshal(corpoRequisicao, &telefone); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = telefone.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeTelefones(db)
	if erro = repository.Atualizar(telefoneID, telefone); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarTelefone(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	telefoneID, erro := strconv.ParseUint(parametros["telefoneId"], 10, 64)

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeTelefones(db)
	if erro = repository.Deletar(telefoneID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
