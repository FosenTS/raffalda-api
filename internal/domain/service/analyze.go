package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
)

const (
	enStorageLess20 = "There's less than 20% space left in the vault"
	ruStorageLess20 = "В хранилище осталось меньше 20% места"
)

type Analyze interface {
	GeneralWarehouseAnalyze(ctx context.Context) ([]*entity.GeneralWarehouseNotify, error)
}

type analyze struct {
	warehouseStorage   storage.Warehouse
	merchandiseStorage storage.Merchandise
	soldPointStorage   storage.SoldPoint
	transactionStorage storage.Transaction

	log *logrus.Entry
}

func NewAnalyze(warehouseStorage storage.Warehouse, merchandiseStorage storage.Merchandise, soldPointStorage storage.SoldPoint, transactionStorage storage.Transaction, log *logrus.Entry) Analyze {
	return &analyze{warehouseStorage: warehouseStorage, merchandiseStorage: merchandiseStorage, soldPointStorage: soldPointStorage, transactionStorage: transactionStorage, log: log}
}

func (s *analyze) GeneralWarehouseAnalyze(ctx context.Context) ([]*entity.GeneralWarehouseNotify, error) {

	// warehouse
	ws, err := s.warehouseStorage.GetAllWarehouse(ctx)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	notifys := make([]*entity.GeneralWarehouseNotify, 0)
	for _, w := range ws {
		precentile := w.Volume / w.Capacity * 100
		if precentile < 20 {
			notifys = append(notifys, &entity.GeneralWarehouseNotify{
				WarehouseName: w.Name,
				ProblemInfo:   ruStorageLess20,
			})
		}
	}
	return notifys, nil
}
