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

func (repository Telefones) BuscarPorID(IDTelefone uint64) (models.Telefone, error) {
	linhas, erro := repository.db.Query(
		"select  idtelefone, pessoal, residencial from telefone where id = ?",
		IDTelefone,
	)
	if erro != nil {
		return models.Telefone{}, erro
	}

	defer linhas.Close()

	var telefone models.Telefone

	if linhas.Next() {
		if erro = linhas.Scan(
			&telefone.IDTelefone,
			&telefone.Pessoal,
			&telefone.Residencial,
		); erro != nil {
			return models.Telefone{}, erro
		}
	}
	return telefone, nil
}

func (repository Telefones) Atualizar(IDTelefone uint64, telefone models.Telefone) error {
	statement, erro := repository.db.Prepare(
		"update telefone set pessoal = ?, residencial = ? where id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(telefone.Pessoal, telefone.Residencial, IDTelefone); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Telefones) Deletar(IDTelefone uint64) error {
	statement, erro := repositorio.db.Prepare("delete from telefone where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(IDTelefone); erro != nil {
		return erro
	}
	return nil
}
