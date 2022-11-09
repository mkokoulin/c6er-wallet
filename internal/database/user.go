package database

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	customerrors "github.com/mkokoulin/c6er-wallet.git/internal/errors"
	"github.com/mkokoulin/c6er-wallet.git/internal/models"
)

func (db *Database) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	err := db.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var exists bool

		err := tx.Model(&models.User{}).Select("count(*) > 0").Where("login = ?", user.Login).Find(&exists).Error
		if err != nil {
			return customerrors.NewErrorWithDB(err, "an unknown error occurred during checking the user")
		}

		if exists {
			return customerrors.NewErrorWithDB(errors.New(""), "user already exists")
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return customerrors.NewErrorWithDB(err, "an unknown error occurred during generation password")
		}

		user.Password = string(hash)

		if err := tx.Create(&user).Error; err != nil {
			return customerrors.NewErrorWithDB(err, "an unknown error occurred during user creation")
		}

		return nil
	})

	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *Database) CheckUserPassword(ctx context.Context, user models.User) (string, error) {
	var result models.User

	err := db.Conn.WithContext(ctx).Model(&models.User{}).Where("login = ?", user.Login).First(&result).Error
	if err != nil {
		return result.Login, customerrors.NewErrorWithDB(err, "user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return result.Login, customerrors.NewErrorWithDB(err, "an unknown error occurred during generation password")
	}

	return result.ID, nil
}