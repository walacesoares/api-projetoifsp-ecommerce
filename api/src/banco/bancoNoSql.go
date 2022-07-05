package banco

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	_ "go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const connectingString = "mongodb+srv://admin:1a2b3c4d@cluster0.ctsu4.mongodb.net/?retryWrites=true&w=majority"
// const dbName = "projetoweb"
// const colName = "usuario"

// var collection *mongo.Collection

// func Init() {
// 	clientOption := options.Client().ApplyURI(connectingString)

// 	client, erro := mongo.Connect(context.TODO(), clientOption)

// 	if erro != nil {
// 		log.Fatal(erro)
// 	}

// 	fmt.Println("MongoDB conectado com sucesso!")

// 	collection = client.Database(dbName).Collection(colName)

// 	fmt.Println("Collection instance is ready!")
// }
