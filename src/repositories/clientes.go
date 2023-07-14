package repositories

import (
	"api/src/models"
	"database/sql"
)

//Clientes representa um repositório de clientes
type Clientes struct {
	db *sql.DB
}

//NovoRepositorioDeCliente cria um repositório de clientes
func NovoRepositorioDeClientes(db *sql.DB) *Clientes {
	return &Clientes{db}
}

//Criar insere um cliente no banco de dados
func (repository Clientes) Criar(cliente models.Cliente) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into cliente (cpf, idusuario) values(?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(cliente.CPF, cliente.IDUsuario)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
