package controllers

import "net/http"

func CriarTelefone(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando endereço!"))
}

func BuscarTelefones(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os endereços!"))
}

func AtualizarTelefone(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando endereço!"))
}

func DeletarTelefone(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando endereço!"))
}
