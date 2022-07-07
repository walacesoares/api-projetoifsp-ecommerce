package main

import (
	"api/src/config"
	"api/src/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {

	config.Carregar()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ctsu4.mongodb.net/?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d ", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
