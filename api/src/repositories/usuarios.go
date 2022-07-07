package repositories

import (
	"api/src/models"
	"database/sql"
)

//Usuarios representa um repositório de usuarios
type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repository Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into usuario (nome, email, senha) values(?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repository Usuarios) BuscarPorID(IDUsuario uint64) (models.Usuario, error) {
	linhas, erro := repository.db.Query(
		"select idusuario, nome, email, cpf from usuario where idusuario = ?",
		IDUsuario,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.IDUsuario,
			&usuario.Nome,
			&usuario.Email,
			&usuario.CPF,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repository Usuarios) Atualizar(IDUsuario uint64, usuario models.Usuario) error {
	statement, erro := repository.db.Prepare(
		"update usuario set nome = ?, email = ? where idusuario = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, IDUsuario); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) Deletar(IDUsuario uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuario where idusuario = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(IDUsuario); erro != nil {
		return erro
	}
	return nil
}

func (repository Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repository.db.Query("select idusuario, senha from usuario where email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.IDUsuario, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repository Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repository.db.Query("select senha from usuario where idusuario = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

func (repository Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repository.db.Prepare("update usuario set senha = ? where idusuario= ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
