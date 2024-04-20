package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/pkg/advancedlog"
)

type Marks interface {
	GetAllMarks(ctx context.Context) ([]*entity.Mark, error)
	GetMarkByObjectIdAndType(ctx context.Context, id uint, t string) (*entity.Mark, error)
	InsertMark(ctx context.Context, m *dto.MarkCreate) error
	DeleteMark(ctx context.Context, id uint) error
	GetMarkByObjectId(ctx context.Context, id uint) (*entity.Mark, error)
}

type marks struct {
	soldPoint storage.SoldPoint
	warehouse storage.Warehouse

	marksStorage storage.Marks

	log *logrus.Entry
}

func NewMarks(soldPoint storage.SoldPoint, warehouse storage.Warehouse, marksStorage storage.Marks, log *logrus.Entry) Marks {
	return &marks{
		soldPoint:    soldPoint,
		warehouse:    warehouse,
		marksStorage: marksStorage,
		log:          log,
	}
}

func (m *marks) InsertMark(ctx context.Context, mC *dto.MarkCreate) error {
	logF := advancedlog.FunctionLog(m.log)

	err := m.marksStorage.InsertMark(ctx, mC)
	if err != nil {
		logF.Errorln(err)
		return err
	}
	return nil
}

func (m *marks) DeleteMark(ctx context.Context, id uint) error {
	logF := advancedlog.FunctionLog(m.log)
	err := m.marksStorage.DeleteMark(ctx, id)
	if err != nil {
		logF.Errorln(err)
		return err
	}
	return nil
}

func (m *marks) GetMarkByObjectId(ctx context.Context, id uint) (*entity.Mark, error) {
	logF := advancedlog.FunctionLog(m.log)
	mark, err := m.marksStorage.GetMarkByObjectId(ctx, id)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	return mark, nil
}

func (m *marks) GetMarkByObjectIdAndType(ctx context.Context, id uint, t string) (*entity.Mark, error) {
	logF := advancedlog.FunctionLog(m.log)
	mark, err := m.marksStorage.GetMarkByObjectIdAndType(ctx, id, t)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	return mark, nil
}

func (m *marks) GetAllMarks(ctx context.Context) ([]*entity.Mark, error) {
	logF := advancedlog.FunctionLog(m.log)
	marks, err := m.marksStorage.GetAllMarks(ctx)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	return marks, nil
}
