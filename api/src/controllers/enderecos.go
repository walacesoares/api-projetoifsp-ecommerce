package controllers

import "net/http"

func CriarEndereco(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando endereço!"))
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
