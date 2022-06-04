package controllers

import "net/http"

func CriarEmpresa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando empresa!"))
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
