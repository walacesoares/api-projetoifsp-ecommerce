package models

import "errors"

type Empresa struct {
	IDEmpresa     uint64 `json:id_empresa,omitempty`
	CNPJ          string `json:cnpj,omitempty`
	Razao_social  string `json:razao_social,omitempty`
	Nome_fantasia string `json:nome_fantasia,omitempty`
	IDUsuario     string `json:id_usuario,omitempty`
}

func (empresa *Empresa) Preparar() error {
	if erro := empresa.validar(); erro != nil {
		return erro
	}
	return nil
}

func (empresa *Empresa) validar() error {
	if empresa.CNPJ == "" {
		return errors.New("O CNPJ é obrigatório e não pode ser em branco!")
	}
	return nil
}
