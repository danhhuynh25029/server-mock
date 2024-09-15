// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sm/http"
	"sm/repository/mgo"
	"sm/service"
)

// Injectors from wire.go:

func initApp(client *mongo.Client) *http.Server {
	iRepository := mgo.NewMongo(client)
	iService := service.NewMockSvc(iRepository)
	server := http.NewServer(iService)
	return server
}
