package repository

import (
	"context"
	"sm/model"
)

type IRepository interface {
	GetMockData(ctx context.Context, id string) error
	InsertMock(ctx context.Context, api model.MockApi) error
	GetMock(ctx context.Context, id string) (*model.MockApi, error)
}
