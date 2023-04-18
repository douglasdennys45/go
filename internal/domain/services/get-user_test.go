package services_test

import (
	"github.com/douglasdennys45/go/internal/domain/repositories/mocks"
	"github.com/douglasdennys45/go/internal/domain/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupRepo(t *testing.T) *mocks.MockUserRepo {
	ctrl := gomock.NewController(t)
	return mocks.NewMockUserRepo(ctrl)
}

func TestNewGetUser(t *testing.T) {
	repo := setupRepo(t)
	repo.EXPECT().GetByID(gomock.Any()).Return(nil, nil)
	service := services.NewGetUser(repo)
	result, err := service.Execute("123")
	assert.Nil(t, err)
	assert.Nil(t, result)
}
