package repository

import (
	"context"
	"finora-wealthlab/internal/model"
	"finora-wealthlab/pkg/database"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *model.User) error {
	query := `
	INSERT INTO users (id, email, password)
	VALUES ($1, $2, $3)
	`

	_, err := database.DB.Exec(context.Background(),
		query,
		user.ID,
		user.Email,
		user.Password,
	)

	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `
	SELECT id, email, password
	FROM users
	WHERE email = $1 AND deleted_at IS NULL
	`

	row := database.DB.QueryRow(context.Background(), query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
