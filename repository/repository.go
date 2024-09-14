package repository

import "context"

type IRepository interface {
	GetMockData(ctx context.Context, id string) error
}
