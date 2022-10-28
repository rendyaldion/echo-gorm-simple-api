package repositories

import (
	"context"

	"github.com/rendyaldion/echo-api/models"
)

func (r *repositories) CreateUser(ctx context.Context, data models.User) error {
	tx := r.db.Begin()

	err := tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	_ = tx.Commit().Error

	return nil
}

func (r *repositories) GetUsers(ctx context.Context, users []models.User) ([]models.User, error) {
	tx := r.db.Begin()

	if err := tx.Find(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	_ = tx.Commit().Error
	return users, nil
}

func (r *repositories) GetUser(ctx context.Context, user models.User, name string) (models.User, error) {
	tx := r.db.Begin()

	if err := tx.First(&user, "name = ?", name).Error; err != nil {
		tx.Rollback()
		return user, err
	}

	_ = tx.Commit()

	return user, nil
}

func (r *repositories) UpdateUser(ctx context.Context, user models.User, id int64) error {
	tx := r.db.Begin()

	if err := tx.Model(&user).Where("id = ?", id).Update("name", user.Name).Error; err != nil {
		tx.Rollback()
		return err
	}

	_ = tx.Commit()

	return nil
}

func (r *repositories) DeleteUser(ctx context.Context, user models.User, id int64) error {
	tx := r.db.Begin()

	if err := tx.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	_ = tx.Commit()

	return nil
}