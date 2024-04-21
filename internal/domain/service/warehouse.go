package service

import (
	"context"
	"errors"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/pkg/advancedlog"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	errWarehouseOverflowing   = errors.New("warehouse overflowing")
	errWarehouseEmpty         = errors.New("warehouse empty")
	errMerchandiseEmpty       = errors.New("merchandise empty")
	errUnableToChangeCapacity = errors.New("unable to change capacity")
)

type Warehouse interface {
	StoreNewWarehouse(ctx context.Context, wC *dto.WarehouseCreate) error
	DeleteWarehouse(ctx context.Context, id uint) error
	UpdateWarehouse(ctx context.Context, w *dto.Warehouse) error
	GetAll(ctx context.Context) ([]*entity.Warehouse, error)
	GetById(ctx context.Context, id uint) (*entity.Warehouse, error)
	GetAllExpireStats(ctx context.Context) (*entity.WarehouseStatistic, error)
	GetExpireStatsByWarehouseId(ctx context.Context, id uint) (*entity.WarehouseStatistic, error)

	StoreWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandiseCreate) error
	UpdateWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandise) error
	GetAllAndMoreInfo(ctx context.Context) ([]*entity.WarehouseMoreInfo, error)
	GetWarehouseMerchandiseByWarehouseId(ctx context.Context, id uint) ([]*entity.MerchandiseMoreInfo, error)

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

func (wH *warehouse) DeleteWarehouse(ctx context.Context, id uint) error {
	logF := advancedlog.FunctionLog(wH.log)
	err := wH.warehouseStorage.DeleteWarehouse(ctx, id)
	if err != nil {
		logF.Errorln(err)
		return err
	}
	return nil
}

func (wH *warehouse) UpdateWarehouse(ctx context.Context, w *dto.Warehouse) error {
	logF := advancedlog.FunctionLog(wH.log)
	if w.Capacity > w.Volume {
		logF.Errorln(errUnableToChangeCapacity)
		return errUnableToChangeCapacity
	}
	err := wH.warehouseStorage.UpdateWarehouse(ctx, w)
	if err != nil {
		logF.Errorln(err)
		return err
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

func (wH *warehouse) GetAllExpireStats(ctx context.Context) (*entity.WarehouseStatistic, error) {
	logF := advancedlog.FunctionLog(wH.log)

	var DangerToSold, RecommendationToSold, NormalToSold uint = 0, 0, 0
	wM, err := wH.warehouseStorage.GetAllWarehouseMerchandise(ctx)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	for _, m := range wM {
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

		expireDays := manufactureDate.Sub(expireDate).Hours() / 24
		expirePercentage := manufactureDate.Sub(time.Now()).Hours() / 24 / expireDays * 100
		if expirePercentage < 10 {
			DangerToSold++
		} else if expirePercentage < 30 {
			RecommendationToSold++

		}
		NormalToSold++
	}

	return &entity.WarehouseStatistic{
		DangerToSold:         DangerToSold,
		RecommendationToSold: RecommendationToSold,
		NormalToSold:         NormalToSold,
	}, nil
}

func (wH *warehouse) GetExpireStatsByWarehouseId(ctx context.Context, id uint) (*entity.WarehouseStatistic, error) {
	logF := advancedlog.FunctionLog(wH.log)

	var DangerToSold, RecommendationToSold, NormalToSold uint = 0, 0, 0
	wM, err := wH.warehouseStorage.GetWarehouseMerchandiseByWarehouseId(ctx, id)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	for _, m := range wM {
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

		expireDays := manufactureDate.Sub(expireDate).Hours() / 24
		expirePercentage := manufactureDate.Sub(time.Now()).Hours() / 24 / expireDays * 100
		if expirePercentage < 10 {
			DangerToSold++
		} else if expirePercentage < 30 {
			RecommendationToSold++

		}
		NormalToSold++
	}

	return &entity.WarehouseStatistic{
		DangerToSold:         DangerToSold,
		RecommendationToSold: RecommendationToSold,
		NormalToSold:         NormalToSold,
	}, nil
}

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
		expireDays := manufactureDate.Sub(expireDate).Hours() / 24
		expirePercentage := manufactureDate.Sub(time.Now()).Hours() / 24 / expireDays * 100

		mmi = append(mmi, &entity.MerchandiseMoreInfo{
			Id:               m.Id,
			WarehouseId:      m.WarehouseId,
			WarehouseName:    m.WarehouseName,
			ProductName:      m.ProductName,
			ProductCost:      m.ProductCost,
			ManufactureDate:  m.ManufactureDate,
			ExpiryDate:       m.ExpireDate,
			SKU:              m.SKU,
			Quantity:         m.Quantity,
			Measure:          m.Measure,
			ExpirePercentage: uint(expirePercentage),
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

	warehouse, err := wH.warehouseStorage.GetWarehouseById(ctx, wM.WarehouseId)
	if err != nil {
		logF.Errorln(err)
		return err
	}

	warehouse.Volume += wM.Quantity
	if warehouse.Volume > warehouse.Capacity {
		logF.Errorln(errWarehouseOverflowing)
		return errWarehouseOverflowing
	}
	err = wH.warehouseStorage.UpdateWarehouse(ctx, &dto.Warehouse{
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

func (wH *warehouse) UpdateWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandise) error {
	logF := advancedlog.FunctionLog(wH.log)

	oldWarehouseMerchandise, err := wH.warehouseStorage.GetWarehouseMerchandiseById(ctx, wM.Id)
	if err != nil {
		logF.Errorln(err)
		return err
	}

	warehouse, err := wH.warehouseStorage.GetWarehouseById(ctx, wM.WarehouseId)
	if err != nil {
		logF.Errorln(err)
		return err
	}
	warehouse.Volume = warehouse.Volume - oldWarehouseMerchandise.Quantity + wM.Quantity
	if warehouse.Volume > warehouse.Capacity {
		logF.Errorln(errWarehouseOverflowing)
		return errWarehouseOverflowing
	}
	err = wH.warehouseStorage.UpdateWarehouse(ctx, &dto.Warehouse{
		Id:       warehouse.ID,
		Name:     warehouse.Name,
		Volume:   warehouse.Volume,
		Capacity: warehouse.Capacity,
	})
	if err != nil {
		logF.Errorln(err)
		return err
	}

	err = wH.warehouseStorage.UpdateWarehouseMerchandise(ctx, wM)
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

	for _, w := range warehouses {
		moreInfos, err := wH.warehouseStorage.GetWarehouseMerchandiseByWarehouseId(ctx, w.ID)
		if err != nil {
			logF.Errorln(err)
			return nil, err
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

func (wH *warehouse) GetWarehouseMerchandiseById(ctx context.Context, id uint) (*entity.WarehouseMerchandise, error) {
	logF := advancedlog.FunctionLog(wH.log)
	wm, err := wH.warehouseStorage.GetWarehouseMerchandiseById(ctx, id)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	return wm, nil
}

func (wH *warehouse) GetWarehouseMerchandiseByWarehouseId(ctx context.Context, id uint) ([]*entity.MerchandiseMoreInfo, error) {
	logF := advancedlog.FunctionLog(wH.log)
	wms, err := wH.warehouseStorage.GetWarehouseMerchandiseByWarehouseId(ctx, id)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	mmis := make([]*entity.MerchandiseMoreInfo, 0)
	for _, wm := range wms {
		manufactureDate, err := time.Parse("2006-01-02 15:04:05", wm.ManufactureDate)
		if err != nil {
			logF.Errorln(err)
			return nil, err
		}
		expireDate, err := time.Parse("2006-01-02 15:04:05", wm.ExpireDate)
		if err != nil {
			logF.Errorln(err)
			return nil, err
		}
		expireDays := manufactureDate.Sub(expireDate).Hours() / 24
		expirePercentage := manufactureDate.Sub(time.Now()).Hours() / 24 / expireDays * 100

		mmis = append(mmis, &entity.MerchandiseMoreInfo{
			Id:               wm.Id,
			WarehouseId:      wm.WarehouseId,
			WarehouseName:    wm.WarehouseName,
			ProductName:      wm.ProductName,
			ProductCost:      wm.ProductCost,
			ManufactureDate:  wm.ManufactureDate,
			ExpiryDate:       wm.ExpireDate,
			SKU:              wm.SKU,
			Quantity:         wm.Quantity,
			Measure:          wm.Measure,
			ExpirePercentage: uint(expirePercentage),
		})
	}

	return mmis, nil
}

func (wH *warehouse) DeleteWarehouseMerchandise(ctx context.Context, id uint) error {
	logF := advancedlog.FunctionLog(wH.log)
	err := wH.warehouseStorage.DeleteWarehouseMerchandiseById(ctx, id)
	if err != nil {
		logF.Errorln(err)
		return err
	}
	return nil
}
