package models

import "errors"

type OrdemServico struct {
	IDOrdemServico uint64 `json:id_ordemservico,omitempty`
	Nome           string `json:nome,omitempty`
	Tipo           string `json:tipo,omitempty`
	Custo          string `json:custo,omitempty`
	Prazo          string `json:prazo,omitempty`
	Descricao      string `json:descricao,omitempty`
	IDCliente      uint64 `json:id_cliente,omitempty`
	IDServico      uint64 `json:id_servico,omitempty`
	IDOrcamento    uint64 `json:id_orcamento,omitempty`
}

func (ordemServico *OrdemServico) Preparar() error {
	if erro := ordemServico.validar(); erro != nil {
		return erro
	}
	return nil
}

func (ordemServico *OrdemServico) validar() error {
	if ordemServico.Nome == "" {
		return errors.New("O nome é obrigatório e não pode ser em branco!")
	}
	if ordemServico.Tipo == "" {
		return errors.New("O tipo é obrigatório e não pode ser em branco!")
	}
	if ordemServico.Custo == "" {
		return errors.New("O custo é obrigatório e não pode ser em branco!")
	}
	if ordemServico.Prazo == "" {
		return errors.New("O prazo é obrigatório e não pode ser em branco!")
	}
	if ordemServico.Descricao == "" {
		return errors.New("A descrição é obrigatório e não pode ser em branco!")
	}

	return nil
}
