package service

import (
	"context"
	"raffalda-api/internal/domain/service/sto"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/pkg/advancedlog"

	"github.com/sirupsen/logrus"
)

type MerchandiseService interface {
	StoreBulkParse(ctx context.Context, merchandises []*sto.MerchandiseParse) error
}

type merchandiseService struct {
	merchandiseStorage storage.Merchandise

	log *logrus.Entry
}

func NewMerchandiseService(merchandiseStorage storage.Merchandise, log *logrus.Entry) MerchandiseService {
	return &merchandiseService{
		merchandiseStorage: merchandiseStorage,
		log:                log,
	}
}

func (mS *merchandiseService) StoreBulkParse(ctx context.Context, merchandises []*sto.MerchandiseParse) error {
	logF := advancedlog.FunctionLog(mS.log)

	logF.Debugln("Start storing merchandises")

	// Segmentation bulk
	chunkSize := 300
	for start := 0; start < len(merchandises); start += chunkSize {
		ms := make([]*dto.MerchandiseCreate, 0)

		end := start + chunkSize
		if end > len(merchandises) {
			end = len(merchandises)
		}

		for _, m := range merchandises[start:end] {
			ms = append(ms, &dto.MerchandiseCreate{
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
			})
		}

		err := mS.merchandiseStorage.BulkInsertMerchandise(ctx, ms)
		if err != nil {
			logF.Errorln(err)
			return err
		}
	}

	logF.Debugln("Complete storing merchandises")

	return nil
}
