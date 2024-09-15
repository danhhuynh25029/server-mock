package model

import "go.mongodb.org/mongo-driver/v2/bson"

type MockApi struct {
	ID       bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID     string        `json:"uuid,omitempty" bson:"uuid"`
	Method   string        `json:"method,omitempty" bson:"method,omitempty"`
	Path     string        `json:"path,omitempty" bson:"path,omitempty"`
	Request  string        `json:"request,omitempty" bson:"request,omitempty"`
	Response string        `json:"response,omitempty" bson:"response,omitempty"`
}
