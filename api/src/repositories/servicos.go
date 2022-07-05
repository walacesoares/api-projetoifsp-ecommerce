package repositories

// import (
// 	"api/src/banco"
// 	m "api/src/models"
// 	"context"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// var collection = banco.GetCollection("services")
// var ctx = context.Background()

// func Create(servico m.Servico) error {

// 	var erro error
// 	_, erro = collection.InsertOne(ctx, servico)

// 	if erro != nil {
// 		return erro
// 	}
// 	return nil
// }

// func Read() (m.Servicos, error) {

// 	var servicos m.Servicos

// 	filter := bson.D{}
// 	cur, erro := collection.Find(ctx, filter)

// 	if erro != nil {

// 	}

// 	for cur.Next(ctx) {
// 		var servico m.Servico
// 		erro = cur.Decode(&servico)
// 		if erro != nil {
// 			return nil, erro
// 		}
// 		servicos = append(servicos, &servico)
// 	}

// 	return servicos, nil
// }
// func Update(servico m.Servico, servicoID string) error {

// 	var erro error

// 	oid, _ := primitive.ObjectIDFromHex(servicoID)

// 	filter := bson.M{"_id": oid}

// 	update := bson.M{
// 		"$set": bson.M{
// 			"nome":      servico.Nome,
// 			"custo":     servico.Custo,
// 			"prazo":     servico.Prazo,
// 			"descricao": servico.Descricao,
// 			"orcamento": servico.Orcamento,
// 		},
// 	}
// 	_, erro = collection.UpdateOne(ctx, filter, update)
// 	if erro != nil {
// 		return erro
// 	}

// 	return nil

// }
// func Delete(servicoID string) error {

// 	var erro error
// 	var oid primitive.ObjectID

// 	oid, erro = primitive.ObjectIDFromHex(servicoID)

// 	if erro != nil {
// 		return erro
// 	}

// 	filter := bson.M{"_id": oid}

// 	_, erro = collection.DeleteOne(ctx, filter)

// 	if erro != nil {
// 		return erro
// 	}

// 	return nil
// }
