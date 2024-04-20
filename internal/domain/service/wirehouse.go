package service

import (
	"context"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/pkg/advancedlog"
	"time"

	"github.com/sirupsen/logrus"
)

type Warehouse interface {
	StoreNewWarehouse(ctx context.Context, wC *dto.WarehouseCreate) error
	GetAll(ctx context.Context) ([]*entity.Warehouse, error)
	GetById(ctx context.Context, id uint) (*entity.Warehouse, error)

	StoreWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandiseCreate) error
	GetAllAndMoreInfo(ctx context.Context) ([]*entity.WarehouseMoreInfo, error)

	GetAllMerchandiseMoreInfo(ctx context.Context, num uint) ([]*entity.MerchandiseMoreInfo, error)
}

type warehouse struct {
	warehouseStorage   storage.Warehouse
	merchandiseStorage storage.Merchandise

	log *logrus.Entry
}

func NewWarehouse(warehouseStorage storage.Warehouse, merchandiseStorage storage.Merchandise, log *logrus.Entry) Warehouse {
	return &warehouse{
		warehouseStorage:   warehouseStorage,
		merchandiseStorage: merchandiseStorage,
		log:                log,
	}
}

func (wH *warehouse) StoreNewWarehouse(ctx context.Context, wC *dto.WarehouseCreate) error {
	logF := advancedlog.FunctionLog(wH.log)

	err := wH.warehouseStorage.InsertWarehouse(ctx, wC)
	if err != nil {
		logF.Errorln(err)
		return nil
	}

	return nil
}

func (wH *warehouse) GetAll(ctx context.Context) ([]*entity.Warehouse, error) {
	logF := advancedlog.FunctionLog(wH.log)

	warehouses, err := wH.warehouseStorage.GetAllWarehouse(ctx)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return warehouses, nil
}

const paginationStep uint = 20

func (wH *warehouse) GetAllMerchandiseMoreInfo(ctx context.Context, num uint) ([]*entity.MerchandiseMoreInfo, error) {
	logF := advancedlog.FunctionLog(wH.log)

	wM, err := wH.warehouseStorage.GetAllWarehouseMerchandise(ctx)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	var backSteps uint = uint(0)
	var steps uint = uint(len(wM))
	if uint(len(wM)) > paginationStep {
		steps = num * paginationStep
		backSteps = steps - paginationStep
	}

	mmi := make([]*entity.MerchandiseMoreInfo, 0)

	for _, m := range wM[backSteps:steps] {
		manufactureDate, err := time.Parse("2006-01-02 15:04:05", m.ManufactureDate)
		if err != nil {
			logF.Errorln(err)
			return nil, err
		}
		expireDate, err := time.Parse("2006-01-02 15:04:05", m.ExpireDate)
		if err != nil {
			logF.Errorln(err)
			return nil, err
		}
		expireDays := int(manufactureDate.Sub(expireDate).Hours() / 24)

		mmi = append(mmi, &entity.MerchandiseMoreInfo{
			Id:              m.Id,
			WarehouseId:     m.WarehouseId,
			ProductName:     m.ProductName,
			ProductCost:     m.ProductCost,
			ManufactureDate: m.ManufactureDate,
			ExpiryDate:      m.ExpireDate,
			SKU:             m.SKU,
			Quantity:        m.Quantity,
			Measure:         m.Measure,
			ExpireProcent:   uint(expireDays),
		})
	}

	return mmi, nil
}

func (wH *warehouse) GetById(ctx context.Context, id uint) (*entity.Warehouse, error) {

	logF := advancedlog.FunctionLog(wH.log)

	wirehouse, err := wH.warehouseStorage.GetWarehouseById(ctx, id)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return wirehouse, nil
}

func (wH *warehouse) StoreWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandiseCreate) error {
	logF := advancedlog.FunctionLog(wH.log)
	err := wH.warehouseStorage.InsertWarehouseMerchandise(ctx, wM)
	if err != nil {
		logF.Errorln(err)
		return err
	}

	return nil
}

func (wH *warehouse) GetAllAndMoreInfo(ctx context.Context) ([]*entity.WarehouseMoreInfo, error) {
	logF := advancedlog.FunctionLog(wH.log)
	wmis := make([]*entity.WarehouseMoreInfo, 0)
	warehouses, err := wH.warehouseStorage.GetAllWarehouse(ctx)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	ms, err := wH.warehouseStorage.GetAllWarehouseMerchandise(ctx)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	for _, w := range warehouses {
		moreInfos := make([]*entity.WarehouseMerchandise, 0)
		for _, m := range ms {
			if w.ID == m.WarehouseId {
				moreInfos = append(moreInfos, m)
			}
		}
		wmis = append(wmis, &entity.WarehouseMoreInfo{
			Id:           w.ID,
			Name:         w.Name,
			Volume:       w.Volume,
			Capacity:     w.Capacity,
			Merchandises: moreInfos,
		})
	}
	return wmis, nil
}
