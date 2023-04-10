package repositories

import (
	"douglasdenny45.github.com/go/internal/domain/entities"
	"douglasdenny45.github.com/go/internal/domain/repositories"
	"douglasdenny45.github.com/go/internal/infrastructure/database/postgres"
)

type userRepository struct {
}

func NewUserRepository() repositories.UserRepo {
	return &userRepository{}
}

func (u *userRepository) Create(user *entities.User) error {
	tx := postgres.NewPostgreSQLConnect().GetTX()
	ctx := postgres.NewPostgreSQLConnect().GetContext()
	_, err := tx.ExecContext(
		ctx,
		"INSERT INTO users (id, name, email, password, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

func (u *userRepository) GetByID(id string) (*entities.User, error) {
	tx := postgres.NewPostgreSQLConnect().GetTX()
	ctx := postgres.NewPostgreSQLConnect().GetContext()
	var user entities.User
	err := tx.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
