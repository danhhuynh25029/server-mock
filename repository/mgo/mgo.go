package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sm/model"
	"sm/repository"
)

type Mongo struct {
	db *mongo.Client
}

func NewMongo(client *mongo.Client) repository.IRepository {
	return &Mongo{
		db: client,
	}
}

func (m *Mongo) Connect() error {
	return nil
}

func (m *Mongo) GetMockData(ctx context.Context, id string) error {
	return nil
}

func (m *Mongo) InsertMock(ctx context.Context, api model.MockApi) error {
	_, err := m.db.Database("mock").Collection("api").InsertOne(ctx, api)
	return err
}

func (m *Mongo) GetMock(ctx context.Context, id string) (*model.MockApi, error) {
	var api model.MockApi
	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = m.db.Database("mock").Collection("api").FindOne(ctx, bson.M{"_id": objId}).Decode(&api)
	if err != nil {
		return nil, err
	}
	return &api, nil

}
