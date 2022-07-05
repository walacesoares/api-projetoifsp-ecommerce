package banco

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database = "projetoweb"
)

func GetCollection(collection string) *mongo.Collection {

	uri := fmt.Sprintf("mongodb+srv://admin:admin@cluster0.ctsu4.mongodb.net/?retryWrites=true&w=majority")
	// uri := fmt.Sprintf("mongodb+srv://admin:admin@cluster0.ctsu4.mongodb.net/?retryWrites=true&w=majority")

	client, erro := mongo.NewClient(options.Client().ApplyURI(uri))
	if erro != nil {
		panic(erro.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	erro = client.Connect(ctx)

	if erro != nil {
		panic(erro.Error())
	}

	return client.Database(database).Collection(collection)
}
