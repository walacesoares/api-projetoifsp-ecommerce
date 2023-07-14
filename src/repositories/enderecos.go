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
		"insert into endereco (rua, numero, cidade, bairro, estado, idempresa) values(?,?,?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(endereco.Rua, endereco.Numero, endereco.Cidade, endereco.Bairro, endereco.Estado, endereco.IDEmpresa)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repository Enderecos) BuscarPorID(IDEndereco uint64) (models.Endereco, error) {
	linhas, erro := repository.db.Query(
		"select idendereco, rua, numero, cidade, bairro, estado from endereco where idendereco = ?",
		IDEndereco,
	)
	if erro != nil {
		return models.Endereco{}, erro
	}

	defer linhas.Close()

	var endereco models.Endereco

	if linhas.Next() {
		if erro = linhas.Scan(
			&endereco.IDEndereco,
			&endereco.Rua,
			&endereco.Numero,
			&endereco.Cidade,
			&endereco.Bairro,
			&endereco.Estado,
		); erro != nil {
			return models.Endereco{}, erro
		}
	}
	return endereco, nil
}

func (repository Enderecos) Atualizar(IDEndereco uint64, endereco models.Endereco) error {
	statement, erro := repository.db.Prepare(
		"update endereco set rua = ?, numero = ? , cidade = ?, bairro = ?, estado = ? where id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(endereco.Rua, endereco.Numero, endereco.Cidade, endereco.Bairro, endereco.Estado, IDEndereco); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Enderecos) Deletar(IDEndereco uint64) error {
	statement, erro := repositorio.db.Prepare("delete from endereco where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(IDEndereco); erro != nil {
		return erro
	}
	return nil
}
