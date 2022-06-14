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
		"insert into ordemservico (nome,custo,prazo,descricao,idcliente,idservico,idorcamento) values(?,?,?,?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(ordemServico.Nome, ordemServico.Custo, ordemServico.Custo, ordemServico.Descricao, ordemServico.IDCliente)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
