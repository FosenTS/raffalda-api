package storage

import (
	"context"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage/dto"
)

type SoldPoint interface {
	InsertSoldPoint(ctx context.Context, sP *dto.SoldPointCreate) error
	GetAllSoldPoints(ctx context.Context) ([]*entity.SoldPoint, error)
	GetSoldPointById(ctx context.Context, id uint) (*entity.SoldPoint, error)
}
