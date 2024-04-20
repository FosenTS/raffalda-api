package gormDB

import (
	"errors"
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/internal/domain/storage/gormDB/scheme"
	"raffalda-api/pkg/advancedlog"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RefreshTokenRepository storage.RefreshToken

type refreshTokenRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewRefreshTokenRepository(db *gorm.DB, log *logrus.Entry) (RefreshTokenRepository, error) {
	logF := advancedlog.FunctionLog(log)

	err := db.AutoMigrate(&scheme.RefreshToken{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	return &refreshTokenRepository{db: db, log: log}, nil
}

func (rR *refreshTokenRepository) InsertRefreshToken(tCreate *dto.RefreshTokenCreate) (*entity.RefreshToken, error) {
	logF := advancedlog.FunctionLog(rR.log)
	refreshTokenF := scheme.RefreshToken{
		Token:              tCreate.Token,
		Login:              tCreate.Login,
		ExpirationTimeUnix: tCreate.ExpirationTimeUnix,
		CreateTimeUnix:     tCreate.CreateTimeUnix,
	}
	result := rR.db.Create(&refreshTokenF)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.RefreshToken{
		ID:                 refreshTokenF.ID,
		Token:              refreshTokenF.Token,
		Login:              refreshTokenF.Login,
		ExpirationTimeUnix: refreshTokenF.ExpirationTimeUnix,
		CreateTimeUnix:     refreshTokenF.CreateTimeUnix,
	}, nil
}

func (rR *refreshTokenRepository) FindByToken(token string) (*entity.RefreshToken, error) {
	logF := advancedlog.FunctionLog(rR.log)

	var refreshToken *scheme.RefreshToken

	result := rR.db.Where("token = ?", token).First(&refreshToken)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.RefreshToken{
		ID:                 refreshToken.ID,
		Token:              refreshToken.Token,
		Login:              refreshToken.Login,
		ExpirationTimeUnix: refreshToken.ExpirationTimeUnix,
		CreateTimeUnix:     refreshToken.CreateTimeUnix,
	}, nil
}

func (rR *refreshTokenRepository) DeleteByLogin(login string) error {
	logF := advancedlog.FunctionLog(rR.log)

	result := rR.db.Where("login = ?", login).Delete(&scheme.RefreshToken{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (rR *refreshTokenRepository) DeleteByID(id uint) error {
	logF := advancedlog.FunctionLog(rR.log)

	result := rR.db.Delete(&scheme.RefreshToken{}, id)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}
