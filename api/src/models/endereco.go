package models

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
