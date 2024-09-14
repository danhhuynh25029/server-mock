package service

import "context"

type IService interface {
	GetMockData(ctx context.Context, id string) error
}
