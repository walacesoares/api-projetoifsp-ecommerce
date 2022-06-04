package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarCliente(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var cliente models.Cliente
	if erro = json.Unmarshal(corpoRequest, &cliente); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NovoRepositorioDeClientes(db)
	clienteId, erro := repository.Criar(cliente)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", clienteId)))
}

func BuscarClientes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os clientes!"))
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
