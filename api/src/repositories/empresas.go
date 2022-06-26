package repositories

import (
	"api/src/models"
	"database/sql"
)

//Empresas representa um repositório de empresas
type Empresas struct {
	db *sql.DB
}

//NovoRepositorioDeEmpresas cria um repositório de empresas
func NovoRepositorioDeEmpresas(db *sql.DB) *Empresas {
	return &Empresas{db}
}

//Criar insere uma empresa no banco de dados
func (repository Empresas) Criar(empresa models.Empresa) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into empresa (cnpj, razao_social, nome_fantasia, idusuario) values(?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(empresa.CNPJ, empresa.Razao_social, empresa.Nome_fantasia, empresa.IDUsuario)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repository Empresas) BuscarPorID(IDEmpresa uint64) (models.Empresa, error) {
	linhas, erro := repository.db.Query(
		"select idempresa,cnpJ, razao_social, nome_fantasia, idusuario from empresa where idempresa = ?",
		IDEmpresa,
	)
	if erro != nil {
		return models.Empresa{}, erro
	}

	defer linhas.Close()

	var empresa models.Empresa

	if linhas.Next() {
		if erro = linhas.Scan(
			&empresa.IDEmpresa,
			&empresa.CNPJ,
			&empresa.Razao_social,
			&empresa.Nome_fantasia,
			&empresa.IDUsuario,
		); erro != nil {
			return models.Empresa{}, erro
		}
	}
	return empresa, nil
}

func (repository Empresas) Atualizar(IDEmpresa uint64, empresa models.Empresa) error {
	statement, erro := repository.db.Prepare(
		"update empresa set cnpj = ?, razao_social = ?, nome_fantasia = ? where idempresa = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(empresa.CNPJ, empresa.Razao_social, empresa.Nome_fantasia, IDEmpresa); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Empresas) Deletar(IDEmpresa uint64) error {
	statement, erro := repositorio.db.Prepare("delete from empresa where idempresa = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(IDEmpresa); erro != nil {
		return erro
	}
	return nil
}
