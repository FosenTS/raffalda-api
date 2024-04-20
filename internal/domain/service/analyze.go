package service

import (
	"github.com/sirupsen/logrus"
	"raffalda-api/internal/domain/storage"
)

type Analyze interface {
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
