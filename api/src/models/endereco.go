package models

import "errors"

type Endereco struct {
	IDEndereco uint64 `json:id_endereco,omitempty`
	Rua        string `json:rua,omitempty`
	Numero     string `json:numero,omitempty`
	Cidade     string `json:cidade,omitempty`
	Bairro     string `json:bairro,omitempty`
	Estado     string `json:estado,omitempty`
	IDCliente  uint64 `json:id_cliente `
	IDEmpresa  uint64 `json:id_empresa`
}

func (endereco *Endereco) Preparar() error {
	if erro := endereco.validar(); erro != nil {
		return erro
	}
	return nil
}

func (endereco *Endereco) validar() error {
	if endereco.Rua == "" {
		return errors.New("A rua é obrigatório e não pode ser em branco!")
	}
	if endereco.Numero == "" {
		return errors.New("O número é obrigatório e não pode ser em branco!")
	}
	if endereco.Cidade == "" {
		return errors.New("A cidade é obrigatório e não pode ser em branco!")
	}
	if endereco.Bairro == "" {
		return errors.New("O bairro é obrigatório e não pode ser em branco!")
	}
	if endereco.Estado == "" {
		return errors.New("O estado é obrigatório e não pode ser em branco!")
	}
	return nil
}
