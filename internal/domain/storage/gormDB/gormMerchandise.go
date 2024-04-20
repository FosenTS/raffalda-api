package gormDB

import (
	"context"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/internal/domain/storage/gormDB/scheme"
	"raffalda-api/pkg/advancedlog"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MerchandiseRepository storage.Merchandise

type merchandiseRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewMerchandiseRepository(db *gorm.DB, log *logrus.Entry) (MerchandiseRepository, error) {
	logF := advancedlog.FunctionLog(log)

	err := db.AutoMigrate(&scheme.Merchandise{})
	if err != nil {
		logF.Errorln(err)
	}

	return &merchandiseRepository{db: db, log: log}, nil
}

func (mR *merchandiseRepository) InsertMerchandise(ctx context.Context, m *dto.MerchandiseCreate) error {
	logF := advancedlog.FunctionLog(mR.log)

	merchandiseF := scheme.Merchandise{
		ProductName:     m.ProductName,
		ProductCost:     m.ProductCost,
		ManufactureDate: m.ManufactureDate,
		ExpiryDate:      m.ExpiryDate,
		SKU:             m.SKU,
		StoreName:       m.StoreName,
		StoreAddress:    m.StoreAddress,
		Region:          m.Region,
		SaleDate:        m.SaleDate,
		QuantitySold:    m.QuantitySold,
		ProductAmount:   m.ProductAmount,
		ProductMeasure:  m.ProductMeasure,
		ProductVolume:   m.ProductVolume,
		Manufacturer:    m.Manufacturer,
	}

	result := mR.db.Create(&merchandiseF)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (mR *merchandiseRepository) BulkInsertMerchandise(ctx context.Context, ms []*dto.MerchandiseCreate) error {
	logF := advancedlog.FunctionLog(mR.log)

	merchandisesF := make([]*scheme.Merchandise, 0)
	for _, m := range ms {
		merchandiseF := &scheme.Merchandise{
			ProductName:     m.ProductName,
			ProductCost:     m.ProductCost,
			ManufactureDate: m.ManufactureDate,
			ExpiryDate:      m.ExpiryDate,
			SKU:             m.SKU,
			StoreName:       m.StoreName,
			StoreAddress:    m.StoreAddress,
			Region:          m.Region,
			SaleDate:        m.SaleDate,
			QuantitySold:    m.QuantitySold,
			ProductAmount:   m.ProductAmount,
			ProductMeasure:  m.ProductMeasure,
			ProductVolume:   m.ProductVolume,
			Manufacturer:    m.Manufacturer,
		}
		merchandisesF = append(merchandisesF, merchandiseF)
	}

	if result := mR.db.Create(&merchandisesF); result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (mR *merchandiseRepository) DeleteById(ctx context.Context, id uint) error {
	logF := advancedlog.FunctionLog(mR.log)

	result := mR.db.Delete(&scheme.Merchandise{}, id)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}
