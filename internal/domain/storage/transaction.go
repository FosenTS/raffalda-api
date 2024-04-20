package storage

import (
	"context"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage/dto"
)

type Transaction interface {
	InsertTransaction(ctx context.Context, transaction *dto.TransactionCreate) error
	GetAllTransactions(ctx context.Context) ([]*entity.Transaction, error)
	DeleteTransaction(ctx context.Context, id uint) error
	GetTransactionById(ctx context.Context, id uint) (*entity.Transaction, error)
	GetTransactionByWarehousesId(ctx context.Context, id uint) (*entity.Transaction, error)
	GetTransactionsByWarehousesId(ctx context.Context, id uint) ([]*entity.Transaction, error)
	GetTransactionBySoldPointId(ctx context.Context, id uint) (*entity.Transaction, error)
	GetTransactionByMerchandiseId(ctx context.Context, id uint) (*entity.Transaction, error)
}
