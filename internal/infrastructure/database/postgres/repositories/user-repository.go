package repositories

import (
	gopostgres "github.com/douglasdennys45/go-postgres/repositories"
	"github.com/douglasdennys45/go/internal/domain/entities"
	"github.com/douglasdennys45/go/internal/domain/repositories"
	"github.com/douglasdennys45/go/internal/infrastructure/database/postgres"
)

type userRepository struct {
}

func NewUserRepository() repositories.UserRepo {
	return &userRepository{}
}

func (u *userRepository) Create(user *entities.User) error {
	tx := postgres.NewPostgreSQLConnect().GetTX()
	repo := gopostgres.NewPostgresRepository(tx)
	_, err := repo.Insert(postgres.NewPostgreSQLConnect().GetContext(), "INSERT INTO users (id, name, email, password, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6)", user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (u *userRepository) GetByID(id string) (*entities.User, error) {
	tx := postgres.NewPostgreSQLConnect().GetTX()
	repo := gopostgres.NewPostgresRepository(tx)
	var user entities.User
	_, err := repo.Query(postgres.NewPostgreSQLConnect().GetContext(), "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, err
}
