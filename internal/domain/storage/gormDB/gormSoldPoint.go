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

type GormSoldPointRepository storage.SoldPoint

type gormSoldPointRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewGormSoldPointRepository(db *gorm.DB, log *logrus.Entry) (GormSoldPointRepository, error) {
	logF := advancedlog.FunctionLog(log)

	err := db.AutoMigrate(&scheme.SoldPoint{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return &gormSoldPointRepository{db: db, log: log}, nil
}

func (sR *gormSoldPointRepository) InsertSoldPoint(ctx context.Context, sP *dto.SoldPointCreate) error {
	logF := advancedlog.FunctionLog(sR.log)

	sPF := scheme.SoldPoint{
		Region:  sP.Region,
		Name:    sP.Name,
		Address: sP.Address,
	}
	result := *sR.db.Create(&sPF)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (sR *gormSoldPointRepository) GetAllSoldPoints(ctx context.Context) ([]*entity.SoldPoint, error) {
	logF := advancedlog.FunctionLog(sR.log)
	soldPoints := make([]*scheme.SoldPoint, 0)
	result := sR.db.Find(soldPoints)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, nil
	}
	sPs := make([]*entity.SoldPoint, 0)
	for _, sP := range soldPoints {
		sPs = append(sPs, &entity.SoldPoint{
			ID:      sP.ID,
			Region:  sP.Region,
			Name:    sP.Name,
			Address: sP.Address,
		})
	}

	return sPs, nil
}

func (sR *gormSoldPointRepository) GetSoldPointById(ctx context.Context, id uint) (*entity.SoldPoint, error) {
	logF := advancedlog.FunctionLog(sR.log)
	soldPoint := new(scheme.SoldPoint)

	result := sR.db.Where("id = ?", id)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.SoldPoint{
		ID:      soldPoint.ID,
		Region:  soldPoint.Region,
		Name:    soldPoint.Name,
		Address: soldPoint.Address,
	}, nil
}
