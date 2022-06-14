package models

type OrdemServico struct {
	IDOrdemServico uint64 `json:id_ordemservico,omitempty`
	Nome           string `json:nome,omitempty`
	Custo          string `json:custo,omitempty`
	Prazo          string `json:prazo,omitempty`
	Descricao      string `json:descricao,omitempty`
	IDCliente      uint64 `json:id_cliente,omitempty`
	IDServico      uint64 `json:id_servico,omitempty`
	IDOrcamento    uint64 `json:id_orcamento,omitempty`
}
