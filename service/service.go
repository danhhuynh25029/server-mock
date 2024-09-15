package service

import (
	"context"
	"sm/model"
)

type IService interface {
	GetMockData(ctx context.Context, id string) error
	InsertMockApi(ctx context.Context, api model.MockApi) error
	GetMockApi(ctx context.Context, id string) (*model.MockApi, error)
}
