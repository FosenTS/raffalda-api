package storage

import (
	"context"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage/dto"
)

type Marks interface {
	GetAllMarks(ctx context.Context) ([]*entity.Mark, error)
	GetMarkByObjectIdAndType(ctx context.Context, id uint, t string) (*entity.Mark, error)
	InsertMark(ctx context.Context, m *dto.MarkCreate) error
	DeleteMark(ctx context.Context, id uint) error
	GetMarkByObjectId(ctx context.Context, id uint) (*entity.Mark, error)
	GetMarkById(ctx context.Context, id uint) (*entity.Mark, error)
}
