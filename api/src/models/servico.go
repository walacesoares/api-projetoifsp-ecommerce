package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Servico struct {
	IDServico primitive.ObjectID `bson:"_id,omitempty" json:"idservico, omitempty"`
	Nome      string             `bson:"nome,omitempty" json:"nome,omitempty"`
	Custo     string             `bson:"custo,omitempty" json:"custo,omitempty"`
	Prazo     string             `bson: "prazo,omitempty" json:"prazo,omitempty"`
	Descricao string             `bson:"descricao,omitempty" json:"descricao,omitempty"`
	Orcamento bool               `bson:"orcamento,omitempty" json:"orcamento,omitempty"`
	// Data

}

type Servicos []*Servico
