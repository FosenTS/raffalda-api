package storage

import (
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage/dto"
)

type RefreshToken interface {
	InsertRefreshToken(tCreate *dto.RefreshTokenCreate) (*entity.RefreshToken, error)
	FindByToken(token string) (*entity.RefreshToken, error)
	DeleteByLogin(string) error
	DeleteByID(id uint) error
}
