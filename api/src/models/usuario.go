package models

import (
	"api/src/seguranca"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

//Usuario representa um usuário utilizando o sistema
type Usuario struct {
	IDUsuario uint64 `json:"idusuario, omitempty"`
	Nome      string `json:"nome,omitempty"`
	Email     string `json:"email,omitempty"`
	Senha     string `json:"senha,omitempty"`
	CPF       string `json:"cpf,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode ser em branco!")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode ser em branco!")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o formato do email inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode ser em branco!")
	}
	if usuario.CPF == "" {
		return errors.New("O CPF é obrigatório e não pode ser em branco!")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.CPF = strings.TrimSpace(usuario.CPF)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}
	return nil
}
