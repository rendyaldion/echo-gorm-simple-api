package repositories

import (
	"context"

	"github.com/rendyaldion/echo-api/models"
)

func (r *repositories) CreateUser(ctx context.Context, data models.User) error {
	tx := r.db.Begin()

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}

	err := tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repositories) GetUsers(ctx context.Context, users []models.User) ([]models.User, error) {
	tx := r.db.Begin()

	if err := tx.Find(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	err := tx.Commit().Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repositories) GetUser(ctx context.Context, user models.User, name string) (models.User, error) {
	tx := r.db.Begin()

	if err := tx.First(&user, "name = ?", name).Error; err != nil {
		tx.Rollback()
		return user, err
	}

	err := tx.Commit().Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositories) UpdateUser(ctx context.Context, user models.User, id int64) error {
	tx := r.db.Begin()

	if err := tx.Model(&user).Where("id = ?", id).Update("name", user.Name).Error; err != nil {
		tx.Rollback()
		return err
	}

	err := tx.Commit().Error
	if err != nil {
		return err
	}
	
	return nil
}

func (r *repositories) DeleteUser(ctx context.Context, user models.User, id int64) error {
	tx := r.db.Begin()

	if err := tx.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	err := tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}