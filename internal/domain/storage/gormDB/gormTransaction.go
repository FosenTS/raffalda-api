package gormDB

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/internal/domain/storage/gormDB/scheme"
	"raffalda-api/pkg/advancedlog"
)

type GormTransactionRepository storage.Transaction

type gormTransactionRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewGormTransactionRepository(db *gorm.DB, log *logrus.Entry) (storage.Transaction, error) {

	logF := advancedlog.FunctionLog(log)
	err := db.AutoMigrate(&scheme.Transaction{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return &gormTransactionRepository{
		db:  db,
		log: log,
	}, nil
}

func (g *gormTransactionRepository) InsertTransaction(ctx context.Context, transaction *dto.TransactionCreate) error {

	err := g.db.Create(&scheme.Transaction{
		WarehausId:    transaction.WarehousesId,
		SoldPointId:   transaction.SoldPointId,
		MerchandiseId: transaction.MerchandiseId,
		Count:         transaction.Count,
	}).Error
	if err != nil {
		g.log.Errorln(err)
		return err
	}
	return nil
}

func (g *gormTransactionRepository) GetAllTransactions(ctx context.Context) ([]*entity.Transaction, error) {
	logF := advancedlog.FunctionLog(g.log)
	transactions := make([]*scheme.Transaction, 0)
	result := g.db.Find(&transactions)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	ts := make([]*entity.Transaction, 0)
	for _, transaction := range transactions {
		ts = append(ts, &entity.Transaction{
			Id:            transaction.Id,
			WarehouseId:   transaction.WarehausId,
			SoldPointId:   transaction.SoldPointId,
			MerchandiseId: transaction.MerchandiseId,
			Count:         transaction.Count,
		})
	}
	return nil, nil
}

func (g *gormTransactionRepository) DeleteTransaction(ctx context.Context, id uint) error {
	logF := advancedlog.FunctionLog(g.log)
	result := g.db.Delete(&scheme.Transaction{}, id)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}
	return nil
}

func (g *gormTransactionRepository) GetTransactionById(ctx context.Context, id uint) (*entity.Transaction, error) {
	logF := advancedlog.FunctionLog(g.log)
	transaction := new(scheme.Transaction)
	result := g.db.Where("id = ?", id).First(&transaction)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	return &entity.Transaction{
		Id:            transaction.Id,
		WarehouseId:   transaction.WarehausId,
		SoldPointId:   transaction.SoldPointId,
		MerchandiseId: transaction.MerchandiseId,
		Count:         transaction.Count,
	}, nil
}

func (g *gormTransactionRepository) GetTransactionByWarehausId(ctx context.Context, id uint) (*entity.Transaction, error) {
	logF := advancedlog.FunctionLog(g.log)
	transaction := new(scheme.Transaction)
	result := g.db.Where("warehaus_id = ?", id).First(&transaction)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	return &entity.Transaction{
		Id:            transaction.Id,
		WarehouseId:   transaction.WarehausId,
		SoldPointId:   transaction.SoldPointId,
		MerchandiseId: transaction.MerchandiseId,
		Count:         transaction.Count,
	}, nil
}

func (g *gormTransactionRepository) GetTransactionBySoldPointId(ctx context.Context, id uint) (*entity.Transaction, error) {
	logF := advancedlog.FunctionLog(g.log)
	transaction := new(scheme.Transaction)
	result := g.db.Where("sold_point_id = ?", id).First(&transaction)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	return &entity.Transaction{
		Id:            transaction.Id,
		WarehouseId:   transaction.WarehausId,
		SoldPointId:   transaction.SoldPointId,
		MerchandiseId: transaction.MerchandiseId,
		Count:         transaction.Count,
	}, nil
}

func (g *gormTransactionRepository) GetTransactionByMerchandiseId(ctx context.Context, id uint) (*entity.Transaction, error) {
	logF := advancedlog.FunctionLog(g.log)
	transaction := new(scheme.Transaction)
	result := g.db.Where("merchandise_id = ?", id).First(&transaction)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	return &entity.Transaction{
		Id:            transaction.Id,
		WarehouseId:   transaction.WarehausId,
		SoldPointId:   transaction.SoldPointId,
		MerchandiseId: transaction.MerchandiseId,
		Count:         transaction.Count,
	}, nil
}
