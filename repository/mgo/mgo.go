package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sm/repository"
)

type Mongo struct {
	db *mongo.Client
}

func NewMongo() repository.IRepository {
	return &Mongo{}
}

func (m *Mongo) Connect() error {
	return nil
}

func (m *Mongo) GetMockData(ctx context.Context, id string) error {
	return nil
}
