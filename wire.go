package main

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sm/http"
	"sm/repository/mgo"
	"sm/service"
)

func initApp(client *mongo.Client) *http.Server {
	wire.Build(
		mgo.NewMongo,
		service.NewMockSvc,
		http.NewServer,
	)
	return &http.Server{}
}
