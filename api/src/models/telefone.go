package models

import "errors"

type Telefone struct {
	IDTelefone  uint64 `json:id_telefone,omitempty`
	Pessoal     string `json:pessoal,omitempty`
	Residencial string `json:residencial,omitempty`
	IDCliente   uint64 `json:id_cliente,omitempty`
	IDEmpresa   uint64 `json:id_empresa,omitempty`
}

func (telefone *Telefone) Preparar() error {
	if erro := telefone.validar(); erro != nil {
		return erro
	}
	return nil
}

func (telefone *Telefone) validar() error {
	if telefone.Pessoal == "" {
		return errors.New("O telefone pessoal é obrigatório e não pode ser em branco!")
	}
	if telefone.Residencial == "" {
		return errors.New("O telefone residencial é obrigatório e não pode ser em branco!")
	}
	return nil
}
