package storage

import (
	"context"
	"raffalda-api/internal/domain/storage/dto"
)

type Merchandise interface {
	InsertMerchandise(ctx context.Context, m *dto.MerchandiseCreate) error
	BulkInsertMerchandise(ctx context.Context, ms []*dto.MerchandiseCreate) error

	DeleteById(ctx context.Context, id uint) error
}
