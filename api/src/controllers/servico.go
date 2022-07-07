package controllers

import (
	"api/src/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func CreateServico(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var servico models.Servico
	json.NewDecoder(request.Body).Decode(&servico)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	collection := client.Database("projetoweb").Collection("servico")
	result, _ := collection.InsertOne(ctx, servico)
	json.NewEncoder(response).Encode(result)
}

func GetServicosEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var servicos models.Servicos
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	collection := client.Database("projetoweb").Collection("servico")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var servico models.Servico
		cursor.Decode(&servico)
		servicos = append(servicos, &servico)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(servicos)
}
func GetServicoEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	oid, _ := primitive.ObjectIDFromHex(params["servicoid"])
	var servico models.Servico
	filter := bson.M{"_id": oid}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	collection := client.Database("projetoweb").Collection("servico")
	err := collection.FindOne(ctx, filter).Decode(&servico)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(servico)
}

func UpdateServicoEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	oid, _ := primitive.ObjectIDFromHex(params["servicoid"])
	var servico models.Servico
	filter := bson.M{"_id": oid}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	collection := client.Database("projetoweb").Collection("servico")
	update := bson.M{
		"$set": bson.M{
			"nome":      servico.Nome,
			"custo":     servico.Custo,
			"prazo":     servico.Prazo,
			"descricao": servico.Descricao,
			"orcamento": servico.Orcamento,
		},
	}
	resultado, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", resultado.ModifiedCount)
}

func DeleteServicoEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	oid, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": oid}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	collection := client.Database("projetoweb").Collection("servico")
	resultado, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", resultado.DeletedCount)
}
