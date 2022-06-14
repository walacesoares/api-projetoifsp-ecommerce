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

func CriarEmpresa(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var empresa models.Empresa
	if erro = json.Unmarshal(corpoRequest, &empresa); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = empresa.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NovoRepositorioDeEmpresas(db)
	empresa.IDEmpresa, erro = repository.Criar(empresa)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, empresa)
}

func BuscarEmpresas(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todas as empresas!"))
}

func BuscarEmpresa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando uma empresa!"))
}

func AtualizarEmpresa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando empresa!"))
}

func DeletarEmpresa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando empresa!"))
}
