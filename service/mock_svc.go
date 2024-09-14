package service

import (
	"context"
	"sm/repository"
)

type MockSVC struct {
	repo repository.IRepository
}

func NewMockSvc(repo repository.IRepository) IService {
	return &MockSVC{repo: repo}
}

func (m *MockSVC) GetMockData(ctx context.Context, id string) error {
	return nil
}
