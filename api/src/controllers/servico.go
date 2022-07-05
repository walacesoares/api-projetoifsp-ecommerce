package controllers

import (
	"api/src/banco"
	"api/src/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var collection = banco.GetCollection("servico")
var ctx = context.Background()

func CreateServicoEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var servico models.Servico
	_ = json.NewDecoder(request.Body).Decode(&servico)
	result, _ := collection.InsertOne(ctx, servico)
	json.NewEncoder(response).Encode(result)
}
func GetServicosEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var servicos models.Servicos
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
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	oid, _ := primitive.ObjectIDFromHex(params["id"])
	var servico models.Servico
	filter := bson.M{"_id": oid}
	err := collection.FindOne(ctx, filter).Decode(&servico)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(servico)
}

func UpdateServicoEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	oid, _ := primitive.ObjectIDFromHex(params["id"])
	var servico models.Servico
	filter := bson.M{"_id": oid}
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
	resultado, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", resultado.DeletedCount)
}

// func Create(service m.Servico) error {

// 	erro := repositories.Create(service)

// 	if erro != nil {
// 		return erro
// 	}
// 	return nil
// }

// func Read() (m.Servicos, error) {

// 	services, erro := repositories.Read()

// 	if erro != nil {
// 		return nil, erro
// 	}
// 	return services, nil
// }
// func Update(servico m.Servico, servicoID string) error {

// 	erro := repositories.Update(servico, servicoID)

// 	if erro != nil {
// 		return erro
// 	}

// 	return nil
// }
// func Delete(servicoID string) error {

// 	erro := repositories.Delete(servicoID)

// 	if erro != nil {
// 		return erro
// 	}
// 	return nil
// }
