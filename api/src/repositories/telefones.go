package repositories

import (
	"api/src/models"
	"database/sql"
)

//Telefones representa um repositório de telefones
type Telefones struct {
	db *sql.DB
}

//NovoRepositorioDeTelefones cria um repositório de telefones
func NovoRepositorioDeTelefones(db *sql.DB) *Telefones {
	return &Telefones{db}
}

//Criar insere um telefone no banco de dados
func (repository Telefones) Criar(telefone models.Telefone) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into telefone (pessoal,residencial,idcliente,idempresa) values(?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(telefone.Pessoal, telefone.Residencial, telefone.IDCliente, telefone.IDEmpresa)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
