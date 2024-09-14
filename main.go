package main

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"sm/config"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("cannot load config : %v", err)
	}
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		log.Fatalf("cannot connect mongo : %v", err)
	}
	app := initApp(client)
	app.Run()
}
