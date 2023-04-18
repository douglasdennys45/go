package services_test

import (
	"github.com/douglasdennys45/go/internal/domain/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAddUser(t *testing.T) {
	repo := setupRepo(t)
	repo.EXPECT().Create(gomock.Any()).Return(nil)
	service := services.NewAddUser(repo)
	err := service.Execute("Douglas", "douglasdennys45@gmail.com", "123")
	assert.Nil(t, err)
}
