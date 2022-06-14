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

func AtualizarTelefone(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando telefone!"))
}

func DeletarTelefone(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando telefone!"))
}
