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

type MarksRepository storage.Marks

type marksRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewMarksRepository(db *gorm.DB, log *logrus.Entry) (MarksRepository, error) {
	logF := advancedlog.FunctionLog(log)
	err := db.AutoMigrate(&scheme.Mark{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return &marksRepository{db: db, log: log}, nil
}

func (mR *marksRepository) InsertMark(ctx context.Context, m *dto.MarkCreate) error {
	logF := advancedlog.FunctionLog(mR.log)

	result := mR.db.Create(scheme.Mark{
		Type:      m.Type,
		ObjectId:  m.ObjectId,
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
	})
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}
	return nil
}

func (mR *marksRepository) GetMarkByObjectId(ctx context.Context, id uint) (*entity.Mark, error) {
	logF := advancedlog.FunctionLog(mR.log)
	m := new(scheme.Mark)
	result := mR.db.Where("object_id = ?", id).First(&m)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	return &entity.Mark{
		Type:      m.Type,
		ObjectId:  m.ObjectId,
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
	}, nil
}

func (mR *marksRepository) GetMarkById(ctx context.Context, id uint) (*entity.Mark, error) {
	logF := advancedlog.FunctionLog(mR.log)
	m := new(scheme.Mark)
	result := mR.db.Where("id = ?", id).First(&m)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	return &entity.Mark{
		Type:      m.Type,
		ObjectId:  m.ObjectId,
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
	}, nil
}

func (mR *marksRepository) GetMarkByObjectIdAndType(ctx context.Context, id uint, t string) (*entity.Mark, error) {
	logF := advancedlog.FunctionLog(mR.log)
	m := new(scheme.Mark)
	result := mR.db.Where("object_id = ? AND type = ?", id, t).First(&m)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	return &entity.Mark{
		Type:      m.Type,
		ObjectId:  m.ObjectId,
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
	}, nil
}

func (mR *marksRepository) DeleteMark(ctx context.Context, id uint) error {
	logF := advancedlog.FunctionLog(mR.log)
	result := mR.db.Where("id = ?", id).Delete(&scheme.Mark{})
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}
	return nil
}

func (mR *marksRepository) GetAllMarks(ctx context.Context) ([]*entity.Mark, error) {
	logF := advancedlog.FunctionLog(mR.log)
	m := make([]*scheme.Mark, 0)
	result := mR.db.Find(&m)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}
	marks := make([]*entity.Mark, 0)
	for _, mark := range m {
		marks = append(marks, &entity.Mark{
			Type:      mark.Type,
			ObjectId:  mark.ObjectId,
			Latitude:  mark.Latitude,
			Longitude: mark.Longitude,
		})
	}
	return marks, nil
}
