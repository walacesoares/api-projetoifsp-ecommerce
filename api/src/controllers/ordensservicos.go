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

func CriarOrdemServico(w http.ResponseWriter, r *http.Request) {
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

// func BuscarOrdensServicos(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Buscando ordens de serviços!"))
// }

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

	repository := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repository.Deletar(ordemServicoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
