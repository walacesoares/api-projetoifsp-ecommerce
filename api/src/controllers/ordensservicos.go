package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarOrdemServico(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var ordemServico models.OrdemServico
	if erro = json.Unmarshal(corpoRequest, &ordemServico); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = ordemServico.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairEmpresaID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível criar uma ordem de serviço que não seja para você"))
		return
	}

	ordemServico.IDUsuario = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeOrdensServicos(db)
	ordemServico.IDOrdemServico, erro = repository.Criar(ordemServico)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, ordemServico)
}

func BuscarOrdemServico(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ordemServicoID, erro := strconv.ParseUint(parametros["ordemServicoId"], 10, 64)
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

	repository := repositories.NovoRepositorioDeOrdensServicos(db)
	ordemServico, erro := repository.BuscarPorID(ordemServicoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, ordemServico)
}

func BuscarOrdensServicos(w http.ResponseWriter, r *http.Request) {

}

// func AtualizarOrdemServico(w http.ResponseWriter, r *http.Request) {
// 	parametros := mux.Vars(r)
// 	ordemServicoID, erro := strconv.ParseUint(parametros["ordemServicoId"], 10, 64)
// 	if erro != nil {
// 		respostas.Erro(w, http.StatusInternalServerError, erro)
// 		return
// 	}

// 	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
// 	if erro != nil {
// 		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
// 		return
// 	}

// 	var ordemServico models.OrdemServico
// 	if erro = json.Unmarshal(corpoRequisicao, &ordemServico); erro != nil {
// 		respostas.Erro(w, http.StatusBadRequest, erro)
// 		return
// 	}

// 	if erro = ordemServico.Preparar(); erro != nil {
// 		respostas.Erro(w, http.StatusBadRequest, erro)
// 		return
// 	}

// 	db, erro := banco.Conectar()
// 	if erro != nil {
// 		respostas.Erro(w, http.StatusInternalServerError, erro)
// 		return
// 	}

// 	defer db.Close()

// 	repository := repositories.NovoRepositorioDeOrdensServicos(db)
// 	if erro = repository.Atualizar(ordemServicoID, ordemServico); erro != nil {
// 		respostas.Erro(w, http.StatusInternalServerError, erro)
// 		return
// 	}

// 	respostas.JSON(w, http.StatusNoContent, nil)
// }

func DeletarOrdemServico(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ordemServicoID, erro := strconv.ParseUint(parametros["ordemServicoId"], 10, 64)

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeOrdensServicos(db)
	if erro = repository.Deletar(ordemServicoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
