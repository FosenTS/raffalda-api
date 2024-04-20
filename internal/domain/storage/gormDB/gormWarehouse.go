package gormDB

import (
	"context"
	"errors"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/internal/domain/storage/gormDB/scheme"
	"raffalda-api/pkg/advancedlog"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WarehouseRepository storage.Warehouse

type warehouseRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewWarehouseRepository(db *gorm.DB, log *logrus.Entry) (WarehouseRepository, error) {
	logF := advancedlog.FunctionLog(log)

	err := db.AutoMigrate(&scheme.Warehouse{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	err = db.AutoMigrate(&scheme.WarehouseMerchandise{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return &warehouseRepository{db: db, log: log}, nil
}

func (wR *warehouseRepository) InsertWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandiseCreate) error {
	logF := advancedlog.FunctionLog(wR.log)

	wMF := scheme.WarehouseMerchandise{
		WarehouseId:     wM.WarehouseId,
		ProductName:     wM.ProductName,
		ProductCost:     wM.ProductCost,
		ManufactureDate: wM.ManufactureDate,
		ExpiryDate:      wM.ExpiryDate,
		SKU:             wM.SKU,
		Quantity:        wM.Quantity,
		Measure:         wM.Measure,
	}

	result := wR.db.Create(&wMF)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (wR *warehouseRepository) UpdateWarehouse(ctx context.Context, w *dto.Warehouse) error {
	logF := advancedlog.FunctionLog(wR.log)

	wF := scheme.Warehouse{
		Name:     w.Name,
		Volume:   w.Volume,
		Capacity: w.Capacity,
	}

	result := wR.db.Model(&scheme.Warehouse{}).Where("id = ?", w.Id).Updates(wF)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (wR *warehouseRepository) UpdateWarehouseMerchandise(ctx context.Context, wM *dto.WarehouseMerchandise) error {
	logF := advancedlog.FunctionLog(wR.log)

	wMF := scheme.WarehouseMerchandise{
		WarehouseId:     wM.WarehouseId,
		ProductName:     wM.ProductName,
		ProductCost:     wM.ProductCost,
		ManufactureDate: wM.ManufactureDate,
		ExpiryDate:      wM.ExpireDate,
		SKU:             wM.SKU,
		Quantity:        wM.Quantity,
		Measure:         wM.Measure,
	}

	result := wR.db.Model(&scheme.WarehouseMerchandise{}).Where("id = ?", wM.Id).Updates(wMF)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (wR *warehouseRepository) GetWarehouseMerchandiseById(ctx context.Context, id uint) (*entity.WarehouseMerchandise, error) {
	logF := advancedlog.FunctionLog(wR.log)
	var wM *scheme.WarehouseMerchandise

	result := wR.db.Where("id = ?", id).First(&wM)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.WarehouseMerchandise{
		Id:              wM.Id,
		WarehouseId:     wM.WarehouseId,
		ProductName:     wM.ProductName,
		ProductCost:     wM.ProductCost,
		ManufactureDate: wM.ManufactureDate,
		ExpireDate:      wM.ExpiryDate,
		SKU:             wM.SKU,
		Quantity:        wM.Quantity,
		Measure:         wM.Measure,
	}, nil
}

func (wR *warehouseRepository) GetWarehouseMerchandiseByWarehouseId(ctx context.Context, id uint) ([]*entity.WarehouseMerchandise, error) {
	logF := advancedlog.FunctionLog(wR.log)

	wMCs := make([]*scheme.WarehouseMerchandise, 0)
	result := wR.db.Where("warehouse_id = ?", id).Find(&wMCs)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	wMs := make([]*entity.WarehouseMerchandise, 0)
	for _, wM := range wMCs {
		wMs = append(wMs, &entity.WarehouseMerchandise{
			Id:              wM.Id,
			WarehouseId:     wM.WarehouseId,
			ProductName:     wM.ProductName,
			ProductCost:     wM.ProductCost,
			ManufactureDate: wM.ManufactureDate,
			ExpireDate:      wM.ExpiryDate,
			SKU:             wM.SKU,
			Quantity:        wM.Quantity,
			Measure:         wM.Measure,
		})
	}

	return wMs, nil
}

func (wR *warehouseRepository) GetAllWarehouseMerchandise(ctx context.Context) ([]*entity.WarehouseMerchandise, error) {
	logF := advancedlog.FunctionLog(wR.log)

	wMCs := make([]*scheme.WarehouseMerchandise, 0)

	result := wR.db.Find(&wMCs)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}

	wMs := make([]*entity.WarehouseMerchandise, 0)
	for _, wM := range wMCs {
		var warehouse *scheme.Warehouse
		result := wR.db.Where("id = ?", wM.Id).First(&warehouse)
		if result.Error != nil {
			logF.Errorln(result.Error)
			return nil, result.Error
		}
		wMs = append(wMs, &entity.WarehouseMerchandise{
			Id:              wM.Id,
			WarehouseId:     wM.WarehouseId,
			ProductName:     wM.ProductName,
			ProductCost:     wM.ProductCost,
			ManufactureDate: wM.ManufactureDate,
			ExpireDate:      wM.ExpiryDate,
			SKU:             wM.SKU,
			Quantity:        wM.Quantity,
			Measure:         wM.Measure,
		})
	}

	return wMs, nil
}

func (wR *warehouseRepository) DeleteWarehouseMerchandiseById(ctx context.Context, id uint) error {
	logF := advancedlog.FunctionLog(wR.log)

	result := wR.db.Delete(&scheme.WarehouseMerchandise{}, id)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (wR *warehouseRepository) InsertWarehouse(ctx context.Context, w *dto.WarehouseCreate) error {
	logF := advancedlog.FunctionLog(wR.log)

	wF := scheme.Warehouse{
		Name:     w.Name,
		Volume:   w.Volume,
		Capacity: w.Capacity,
	}

	result := wR.db.Create(&wF)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (wR *warehouseRepository) GetWarehouseById(ctx context.Context, id uint) (*entity.Warehouse, error) {
	logF := advancedlog.FunctionLog(wR.log)
	var wirehouse *scheme.Warehouse

	result := wR.db.Where("id = ?", id).First(&wirehouse)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.Warehouse{
		ID:       wirehouse.ID,
		Name:     wirehouse.Name,
		Volume:   wirehouse.Volume,
		Capacity: wirehouse.Capacity,
	}, nil
}

func (wR *warehouseRepository) GetAllWarehouse(ctx context.Context) ([]*entity.Warehouse, error) {
	logF := advancedlog.FunctionLog(wR.log)

	warehouses := make([]*scheme.Warehouse, 0)
	result := wR.db.Find(&warehouses)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	ws := make([]*entity.Warehouse, 0)
	for _, w := range warehouses {
		ws = append(ws, &entity.Warehouse{
			ID:       w.ID,
			Name:     w.Name,
			Volume:   w.Volume,
			Capacity: w.Capacity,
		})
	}

	return ws, nil
}

func (wR *warehouseRepository) DeleteWarehouse(ctx context.Context, id uint) error {
	logF := advancedlog.FunctionLog(wR.log)

	result := wR.db.Delete(&scheme.Warehouse{}, id)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}
