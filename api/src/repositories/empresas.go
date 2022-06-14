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
