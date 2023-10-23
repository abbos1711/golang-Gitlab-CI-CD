package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/tizim-back/api/models"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

// Users
func (u *userRepo) GetUserByUserName(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT 
			id,
			name,
			password
		FROM users WHERE name = $1
	`
	var user models.User
	err := u.db.QueryRow(context.Background(), query, username).Scan(
		&user.Id,
		&user.UserName,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
