package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/pkg/advancedlog"
	"time"
)

type Transaction interface {
	InsertTransaction(ctx context.Context, transaction *dto.TransactionCreate) error
	GetAllTransactions(ctx context.Context) ([]*entity.Transaction, error)
	DeleteTransaction(ctx context.Context, id uint) error
	GetTransactionById(ctx context.Context, id uint) (*entity.TransactionInfo, error)
	GetTransactionByWarehousesId(ctx context.Context, id uint) (*entity.TransactionInfo, error)
	GetTransactionsByWarehousesId(ctx context.Context, id uint) ([]*entity.TransactionInfo, error)
	GetTransactionBySoldPointId(ctx context.Context, id uint) (*entity.TransactionInfo, error)
	GetTransactionByMerchandiseId(ctx context.Context, id uint) (*entity.TransactionInfo, error)

	GetTransactionsStatsByWarehousesId(ctx context.Context, id uint) (*entity.TransactionStats, error)
}

type transaction struct {
	transactionStorage storage.Transaction
	warehouseStorage   storage.Warehouse
	soldPointStorage   storage.SoldPoint

	log *logrus.Entry
}

func NewTransaction(transactionStorage storage.Transaction, warehouseStorage storage.Warehouse, soldPointStorage storage.SoldPoint, log *logrus.Entry) Transaction {
	return &transaction{transactionStorage: transactionStorage, warehouseStorage: warehouseStorage, soldPointStorage: soldPointStorage, log: log}
}

func (t *transaction) InsertTransaction(ctx context.Context, transaction *dto.TransactionCreate) error {
	logF := advancedlog.FunctionLog(t.log)
	err := t.transactionStorage.InsertTransaction(ctx, transaction)
	if err != nil {
		logF.Errorln(err)
		return err
	}

	warehouseMerchandise, err := t.warehouseStorage.GetWarehouseMerchandiseById(ctx, transaction.MerchandiseId)
	if err != nil {
		logF.Errorln(err)
		return err
	}

	warehouseMerchandise.Quantity -= transaction.Count
	if warehouseMerchandise.Quantity < 0 {
		logF.Errorln(errMerchandiseEmpty)
		return errMerchandiseEmpty
	}
	err = t.warehouseStorage.UpdateWarehouseMerchandise(ctx, &dto.WarehouseMerchandise{
		Id:              warehouseMerchandise.Id,
		WarehouseId:     warehouseMerchandise.WarehouseId,
		ProductName:     warehouseMerchandise.ProductName,
		ProductCost:     warehouseMerchandise.ProductCost,
		ManufactureDate: warehouseMerchandise.ManufactureDate,
		ExpireDate:      warehouseMerchandise.ExpireDate,
		SKU:             warehouseMerchandise.SKU,
		Quantity:        warehouseMerchandise.Quantity,
		Measure:         warehouseMerchandise.Measure,
	})
	if err != nil {
		logF.Errorln(err)
		return err
	}

	warehouse, err := t.warehouseStorage.GetWarehouseById(ctx, transaction.WarehousesId)
	if err != nil {
		logF.Errorln(err)
		return err
	}

	warehouse.Volume -= transaction.Count
	if warehouse.Volume < 0 {
		logF.Errorln(errWarehouseEmpty)
		return errWarehouseEmpty
	}
	err = t.warehouseStorage.UpdateWarehouse(ctx, &dto.Warehouse{
		Id:       warehouse.ID,
		Name:     warehouse.Name,
		Volume:   warehouse.Volume,
		Capacity: warehouse.Capacity,
	})
	if err != nil {
		logF.Errorln(err)
		return err
	}

	return nil
}

func (t *transaction) GetAllTransactions(ctx context.Context) ([]*entity.Transaction, error) {
	transactions, err := t.transactionStorage.GetAllTransactions(ctx)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}

	transactionsInfo := make([]*entity.TransactionInfo, 0)

	for _, transaction := range transactions {
		warehouse, err := t.warehouseStorage.GetWarehouseById(ctx, transaction.WarehouseId)
		if err != nil {
			t.log.Errorln(err)
			return nil, err
		}
		soldPoint, err := t.soldPointStorage.GetSoldPointById(ctx, transaction.SoldPointId)
		if err != nil {
			t.log.Errorln(err)
			return nil, err
		}
		warehouseMerchandises, err := t.warehouseStorage.GetWarehouseMerchandiseById(ctx, transaction.WarehouseId)
		if err != nil {
			t.log.Errorln(err)
			return nil, err
		}
		transactionsInfo = append(transactionsInfo, &entity.TransactionInfo{
			Id:               transaction.Id,
			WarehausId:       transaction.WarehouseId,
			SoldPointId:      transaction.SoldPointId,
			MerchandiseId:    transaction.MerchandiseId,
			Date:             transaction.Date,
			Count:            transaction.Count,
			WarehouseName:    warehouse.Name,
			SoldPointRegion:  soldPoint.Region,
			SoldPointName:    soldPoint.Name,
			SoldPointAddress: soldPoint.Address,
			ProductName:      warehouseMerchandises.ProductName,
			ProductCost:      warehouseMerchandises.ProductCost,
			ManufactureDate:  warehouseMerchandises.ManufactureDate,
			ExpireDate:       warehouseMerchandises.ExpireDate,
			SKU:              warehouseMerchandises.SKU,
			Quantity:         warehouseMerchandises.Quantity,
			Measure:          warehouseMerchandises.Measure,
		})
	}

	return transactions, nil
}

func (t *transaction) DeleteTransaction(ctx context.Context, id uint) error {
	err := t.transactionStorage.DeleteTransaction(ctx, id)
	if err != nil {
		t.log.Errorln(err)
		return err
	}
	return nil
}

func (t *transaction) GetTransactionById(ctx context.Context, id uint) (*entity.TransactionInfo, error) {
	transaction, err := t.transactionStorage.GetTransactionById(ctx, id)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	warehouse, err := t.warehouseStorage.GetWarehouseById(ctx, transaction.WarehouseId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	soldPoint, err := t.soldPointStorage.GetSoldPointById(ctx, transaction.SoldPointId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	warehouseMerchandises, err := t.warehouseStorage.GetWarehouseMerchandiseById(ctx, transaction.WarehouseId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	return &entity.TransactionInfo{
		Id:               transaction.Id,
		WarehausId:       transaction.WarehouseId,
		SoldPointId:      transaction.SoldPointId,
		MerchandiseId:    transaction.MerchandiseId,
		Date:             transaction.Date,
		Count:            transaction.Count,
		WarehouseName:    warehouse.Name,
		SoldPointRegion:  soldPoint.Region,
		SoldPointName:    soldPoint.Name,
		SoldPointAddress: soldPoint.Address,
		ProductName:      warehouseMerchandises.ProductName,
		ProductCost:      warehouseMerchandises.ProductCost,
		ManufactureDate:  warehouseMerchandises.ManufactureDate,
		ExpireDate:       warehouseMerchandises.ExpireDate,
		SKU:              warehouseMerchandises.SKU,
		Quantity:         warehouseMerchandises.Quantity,
		Measure:          warehouseMerchandises.Measure,
	}, nil
}

func (t *transaction) GetTransactionsByWarehousesId(ctx context.Context, id uint) ([]*entity.TransactionInfo, error) {
	transactions, err := t.transactionStorage.GetTransactionsByWarehousesId(ctx, id)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	transactionsInfo := make([]*entity.TransactionInfo, 0)
	for _, transaction := range transactions {
		warehouse, err := t.warehouseStorage.GetWarehouseById(ctx, transaction.WarehouseId)
		if err != nil {
			t.log.Errorln(err)
			return nil, err
		}
		soldPoint, err := t.soldPointStorage.GetSoldPointById(ctx, transaction.SoldPointId)
		if err != nil {
			t.log.Errorln(err)
			return nil, err
		}
		warehouseMerchandises, err := t.warehouseStorage.GetWarehouseMerchandiseById(ctx, transaction.WarehouseId)
		if err != nil {
			t.log.Errorln(err)
			return nil, err
		}

		transactionsInfo = append(transactionsInfo, &entity.TransactionInfo{
			Id:               transaction.Id,
			WarehausId:       transaction.WarehouseId,
			SoldPointId:      transaction.SoldPointId,
			MerchandiseId:    transaction.MerchandiseId,
			Date:             transaction.Date,
			Count:            transaction.Count,
			WarehouseName:    warehouse.Name,
			SoldPointRegion:  soldPoint.Region,
			SoldPointName:    soldPoint.Name,
			SoldPointAddress: soldPoint.Address,
			ProductName:      warehouseMerchandises.ProductName,
			ProductCost:      warehouseMerchandises.ProductCost,
			ManufactureDate:  warehouseMerchandises.ManufactureDate,
			ExpireDate:       warehouseMerchandises.ExpireDate,
			SKU:              warehouseMerchandises.SKU,
			Quantity:         warehouseMerchandises.Quantity,
			Measure:          warehouseMerchandises.Measure,
		})
	}

	return transactionsInfo, nil
}

func (t *transaction) GetTransactionsStatsByWarehousesId(ctx context.Context, id uint) (*entity.TransactionStats, error) {
	transactions, err := t.GetTransactionsByWarehousesId(ctx, id)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	var transactionStats = new(entity.TransactionStats)
	for _, tc := range transactions {
		date, err := time.Parse("2006-01-02 15:01:05", tc.Date)
		if err != nil {
			t.log.Errorln(err)
			return nil, err
		}
		switch date.Weekday() {
		case time.Monday:
			transactionStats.Monday++
		case time.Tuesday:
			transactionStats.Tuesday++
		case time.Wednesday:
			transactionStats.Wednesday++
		case time.Thursday:
			transactionStats.Thursday++
		case time.Friday:
			transactionStats.Friday++
		case time.Saturday:
			transactionStats.Saturday++
		case time.Sunday:
			transactionStats.Sunday++
		}
	}

	return transactionStats, nil
}

func (t *transaction) GetTransactionByWarehousesId(ctx context.Context, id uint) (*entity.TransactionInfo, error) {
	transaction, err := t.transactionStorage.GetTransactionByWarehousesId(ctx, id)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	warehouse, err := t.warehouseStorage.GetWarehouseById(ctx, transaction.WarehouseId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	soldPoint, err := t.soldPointStorage.GetSoldPointById(ctx, transaction.SoldPointId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	warehouseMerchandises, err := t.warehouseStorage.GetWarehouseMerchandiseById(ctx, transaction.WarehouseId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	return &entity.TransactionInfo{
		Id:               transaction.Id,
		WarehausId:       transaction.WarehouseId,
		SoldPointId:      transaction.SoldPointId,
		MerchandiseId:    transaction.MerchandiseId,
		Date:             transaction.Date,
		Count:            transaction.Count,
		WarehouseName:    warehouse.Name,
		SoldPointRegion:  soldPoint.Region,
		SoldPointName:    soldPoint.Name,
		SoldPointAddress: soldPoint.Address,
		ProductName:      warehouseMerchandises.ProductName,
		ProductCost:      warehouseMerchandises.ProductCost,
		ManufactureDate:  warehouseMerchandises.ManufactureDate,
		ExpireDate:       warehouseMerchandises.ExpireDate,
		SKU:              warehouseMerchandises.SKU,
		Quantity:         warehouseMerchandises.Quantity,
		Measure:          warehouseMerchandises.Measure,
	}, nil
}

func (t *transaction) GetTransactionBySoldPointId(ctx context.Context, id uint) (*entity.TransactionInfo, error) {
	transaction, err := t.transactionStorage.GetTransactionBySoldPointId(ctx, id)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	warehouse, err := t.warehouseStorage.GetWarehouseById(ctx, transaction.WarehouseId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	soldPoint, err := t.soldPointStorage.GetSoldPointById(ctx, transaction.SoldPointId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	warehouseMerchandises, err := t.warehouseStorage.GetWarehouseMerchandiseById(ctx, transaction.WarehouseId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	return &entity.TransactionInfo{
		Id:               transaction.Id,
		WarehausId:       transaction.WarehouseId,
		SoldPointId:      transaction.SoldPointId,
		MerchandiseId:    transaction.MerchandiseId,
		Date:             transaction.Date,
		Count:            transaction.Count,
		WarehouseName:    warehouse.Name,
		SoldPointRegion:  soldPoint.Region,
		SoldPointName:    soldPoint.Name,
		SoldPointAddress: soldPoint.Address,
		ProductName:      warehouseMerchandises.ProductName,
		ProductCost:      warehouseMerchandises.ProductCost,
		ManufactureDate:  warehouseMerchandises.ManufactureDate,
		ExpireDate:       warehouseMerchandises.ExpireDate,
		SKU:              warehouseMerchandises.SKU,
		Quantity:         warehouseMerchandises.Quantity,
		Measure:          warehouseMerchandises.Measure,
	}, nil
}

func (t *transaction) GetTransactionByMerchandiseId(ctx context.Context, id uint) (*entity.TransactionInfo, error) {
	transaction, err := t.transactionStorage.GetTransactionByMerchandiseId(ctx, id)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}

	warehouse, err := t.warehouseStorage.GetWarehouseById(ctx, transaction.WarehouseId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	soldPoint, err := t.soldPointStorage.GetSoldPointById(ctx, transaction.SoldPointId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	warehouseMerchandises, err := t.warehouseStorage.GetWarehouseMerchandiseById(ctx, transaction.WarehouseId)
	if err != nil {
		t.log.Errorln(err)
		return nil, err
	}
	return &entity.TransactionInfo{
		Id:               transaction.Id,
		WarehausId:       transaction.WarehouseId,
		SoldPointId:      transaction.SoldPointId,
		MerchandiseId:    transaction.MerchandiseId,
		Date:             transaction.Date,
		Count:            transaction.Count,
		WarehouseName:    warehouse.Name,
		SoldPointRegion:  soldPoint.Region,
		SoldPointName:    soldPoint.Name,
		SoldPointAddress: soldPoint.Address,
		ProductName:      warehouseMerchandises.ProductName,
		ProductCost:      warehouseMerchandises.ProductCost,
		ManufactureDate:  warehouseMerchandises.ManufactureDate,
		ExpireDate:       warehouseMerchandises.ExpireDate,
		SKU:              warehouseMerchandises.SKU,
		Quantity:         warehouseMerchandises.Quantity,
		Measure:          warehouseMerchandises.Measure,
	}, nil
}
