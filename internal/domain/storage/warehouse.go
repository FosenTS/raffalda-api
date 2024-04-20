package storage

import (
	"context"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage/dto"
)

type Warehouse interface {
	InsertWarehouse(ctx context.Context, w *dto.WarehouseCreate) error
	UpdateWarehouse(ctx context.Context, w *dto.Warehouse) error
	GetWarehouseById(ctx context.Context, id uint) (*entity.Warehouse, error)
	GetAllWarehouse(ctx context.Context) ([]*entity.Warehouse, error)
	DeleteWarehouse(ctx context.Context, id uint) error

	InsertWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandiseCreate) error
	UpdateWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandise) error
	GetAllWarehouseMerchandise(ctx context.Context) ([]*entity.WarehouseMerchandise, error)
	GetWarehouseMerchandiseById(ctx context.Context, id uint) (*entity.WarehouseMerchandise, error)
	GetWarehouseMerchandiseByWarehouseId(ctx context.Context, id uint) ([]*entity.WarehouseMerchandise, error)
	DeleteWarehouseMerchandiseById(ctx context.Context, id uint) error
}
