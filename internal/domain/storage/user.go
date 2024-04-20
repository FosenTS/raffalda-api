package storage

import (
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage/dto"
)

type User interface {
	InsertUser(user *dto.UserCreate) error
	FindByLogin(login string) (*entity.User, error)
	DeleteByID(id uint) error
}
