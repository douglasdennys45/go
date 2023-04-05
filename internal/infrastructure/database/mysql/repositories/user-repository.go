package repositories

import (
	"database/sql"

	"douglasdenny45.github.com/go/internal/domain/entities"
	"douglasdenny45.github.com/go/internal/domain/repositories"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) repositories.UserRepo {
	return &userRepository{conn: conn}
}

func (u *userRepository) Create(user *entities.User) error {
	stmt, err := u.conn.Prepare("INSERT INTO users (id, name, email, password, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	return err
}

func (u *userRepository) GetByID(id string) (*entities.User, error) {
	var user entities.User
	stmt, err := u.conn.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
