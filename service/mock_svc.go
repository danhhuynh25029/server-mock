package service

import (
	"context"
	"sm/model"
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

func (m *MockSVC) InsertMockApi(ctx context.Context, api model.MockApi) error {
	return m.repo.InsertMock(ctx, api)
}

func (m *MockSVC) GetMockApi(ctx context.Context, id string) (*model.MockApi, error) {
	return m.repo.GetMock(ctx, id)
}

func (m *MockSVC) GenerateDataFormMock(ctx context.Context, path string) error {
	return nil
}
