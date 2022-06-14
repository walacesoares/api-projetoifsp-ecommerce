package models

type Telefone struct {
	IDTelefone  uint64 `json:id_telefone,omitempty`
	Pessoal     string `json:pessoal,omitempty`
	Residencial string `json:residencial,omitempty`
	IDCliente   uint64 `json:id_cliente,omitempty`
	IDEmpresa   uint64 `json:id_empresa,omitempty`
}
