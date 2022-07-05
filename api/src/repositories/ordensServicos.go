package repositories

import (
	"api/src/models"
	"database/sql"
)

//ORdensServicos representa um repositório de ordem de serviço
type OrdensServicos struct {
	db *sql.DB
}

//NovoRepositorioDeOrdensServicos cria um repositório de ordem de serviço
func NovoRepositorioDeOrdensServicos(db *sql.DB) *OrdensServicos {
	return &OrdensServicos{db}
}

//Criar insere uma ordem de serviço no banco de dados
func (repository OrdensServicos) Criar(ordemServico models.OrdemServico) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into ordemservico (nome,custo,prazo,descricao,idusuario,idempresa) values(?,?,?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(ordemServico.Nome, ordemServico.Custo, ordemServico.Custo, ordemServico.Descricao, ordemServico.IDUsuario, ordemServico.IDEmpresa)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repository OrdensServicos) BuscarPorID(IDOrdemServico uint64) (models.OrdemServico, error) {
	linhas, erro := repository.db.Query(
		"select idordemservico, nome, tipo, custo, prazo, descricao from ordemservico where idordemservico = ?",
		IDOrdemServico,
	)
	if erro != nil {
		return models.OrdemServico{}, erro
	}

	defer linhas.Close()

	var ordemServico models.OrdemServico

	if linhas.Next() {
		if erro = linhas.Scan(
			&ordemServico.IDOrdemServico,
			&ordemServico.Nome,
			&ordemServico.Tipo,
			&ordemServico.Custo,
			&ordemServico.Prazo,
			&ordemServico.Descricao,
		); erro != nil {
			return models.OrdemServico{}, erro
		}
	}
	return ordemServico, nil
}

func (repositorio OrdensServicos) Deletar(IDOrdemServico uint64) error {
	statement, erro := repositorio.db.Prepare("delete from ordemservico where idordemservico = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(IDOrdemServico); erro != nil {
		return erro
	}
	return nil
}
