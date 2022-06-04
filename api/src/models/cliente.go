package models

//Usuario representa um usuário utilizando o sistema
type Cliente struct {
	IDCliente uint64 `json:id_cliente,omitempty`
	CPF       string `json:cpf,omitempty`
	IDUsuario string `json:id_usuario,omitempty`
}
