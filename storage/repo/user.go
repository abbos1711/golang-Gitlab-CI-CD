package repo

import (
	"context"
	"gitlab.com/tizim-back/api/models"
)


type UserStorageI interface {
	GetUserByUserName(ctx context.Context, username string) (user *models.User, err error)
}
