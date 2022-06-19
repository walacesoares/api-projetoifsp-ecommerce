package models

import "errors"

//Cliente representa um cliente utilizando o sistema
type Cliente struct {
	IDCliente uint64 `json:id_cliente,omitempty`
	CPF       string `json:cpf,omitempty`
	IDUsuario string `json:id_usuario,omitempty`
}

func (cliente *Cliente) Preparar() error {
	if erro := cliente.validar(); erro != nil {
		return erro
	}
	return nil
}

func (cliente *Cliente) validar() error {
	if cliente.CPF == "" {
		return errors.New("O CPF é obrigatório e não pode ser em branco!")
	}
	return nil
}
