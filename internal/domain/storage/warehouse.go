package storage

import (
	"context"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage/dto"
)

type Warehouse interface {
	InsertWarehouse(ctx context.Context, w *dto.WarehouseCreate) error
	GetWarehouseById(ctx context.Context, id uint) (*entity.Warehouse, error)
	GetAllWarehouse(ctx context.Context) ([]*entity.Warehouse, error)
	DeleteWarehouse(ctx context.Context, id uint) error

	InsertWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandiseCreate) error
	GetAllWarehouseMerchandise(ctx context.Context) ([]*entity.WarehouseMerchandise, error)
	DeleteWarehouseMerchandiseById(ctx context.Context, id uint) error
}
