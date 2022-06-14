package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//conexao com o banco
	StringConexaoMongo = ""

	//Porta que a api vai estar rodando
	PortaMongo = 0
)

//Carregar vai inicializar as variaveis de ambiente
func CarregarMongo() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		PortaMongo = 9000
	}

	StringConexaoMongo = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
