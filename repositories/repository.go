package repositories

import (
	"context"

	"github.com/rendyaldion/echo-api/models"
	"gorm.io/gorm"
)

type repositories struct {
	db *gorm.DB
}

type Repositories interface {
	CreateUser(ctx context.Context, data models.User) error
	GetUsers(ctx context.Context, users []models.User) ([]models.User, error)
	GetUser(ctx context.Context, user models.User, name string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User, id int64) error
	DeleteUser(ctx context.Context, user models.User, id int64) error
}

func NewRepository(db *gorm.DB) Repositories {
	return &repositories{db:db}
}