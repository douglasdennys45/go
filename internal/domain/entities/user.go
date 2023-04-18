package entities

import (
	"time"

	"github.com/douglasdennys45/go/pkg/uuid"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewUser(name, email, password string) (*User, error) {
	user := User{
		ID:        uuid.NewUUID(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := user.isValidate()
	return &user, err
}

func (u *User) isValidate() error {
	return nil
}
