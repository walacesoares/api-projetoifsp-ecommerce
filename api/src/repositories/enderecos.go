package repositories

import (
	"api/src/models"
	"database/sql"
)

//Endereços representa um repositório de endereços
type Enderecos struct {
	db *sql.DB
}

//NovoRepositorioDeEnderecos cria um repositório de endereços
func NovoRepositorioDeEnderecos(db *sql.DB) *Enderecos {
	return &Enderecos{db}
}

//Criar insere um endereco no banco de dados
func (repository Enderecos) Criar(endereco models.Endereco) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into endereco (rua, numero, cidade, bairro, estado, idcliente, idempresa) values(?,?,?,?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(endereco.Rua, endereco.Numero, endereco.Cidade, endereco.Bairro, endereco.Estado, endereco.IDCliente, endereco.IDEmpresa)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
