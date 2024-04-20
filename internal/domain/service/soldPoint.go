package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
)

type SoldPoint interface {
	StoreSoldPoint(ctx context.Context, sP *dto.SoldPointCreate) error
	GetAllSoldPoints(ctx context.Context) ([]*entity.SoldPoint, error)
	GetSoldPointById(ctx context.Context, id uint) (*entity.SoldPoint, error)
}

type soldPoint struct {
	soldPointStorage storage.SoldPoint

	log *logrus.Entry
}

func NewSoldPoint(soldPointStorage storage.SoldPoint, log *logrus.Entry) SoldPoint {
	return &soldPoint{soldPointStorage: soldPointStorage, log: log}
}

func (s *soldPoint) StoreSoldPoint(ctx context.Context, sP *dto.SoldPointCreate) error {
	err := s.soldPointStorage.InsertSoldPoint(ctx, sP)
	if err != nil {
		s.log.Errorln(err)
		return err
	}
	return nil
}

func (s *soldPoint) GetAllSoldPoints(ctx context.Context) ([]*entity.SoldPoint, error) {
	sps, err := s.soldPointStorage.GetAllSoldPoints(ctx)
	if err != nil {
		s.log.Errorln(err)
		return nil, err
	}
	return sps, nil
}

func (s *soldPoint) GetSoldPointById(ctx context.Context, id uint) (*entity.SoldPoint, error) {
	sp, err := s.soldPointStorage.GetSoldPointById(ctx, id)
	if err != nil {
		s.log.Errorln(err)
		return nil, err
	}
	return sp, nil
}
