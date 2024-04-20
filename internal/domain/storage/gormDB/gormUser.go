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

type UserRepository storage.User

type userRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewUserRepository(db *gorm.DB, log *logrus.Entry) (UserRepository, error) {
	logF := advancedlog.FunctionLog(log)

	err := db.AutoMigrate(&scheme.User{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	err = db.AutoMigrate(&scheme.RefreshToken{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	return &userRepository{db: db, log: log}, nil
}

func (aR *userRepository) InsertUser(user *dto.UserCreate) error {
	logF := advancedlog.FunctionLog(aR.log)
	result := aR.db.Create(&scheme.User{Login: user.Login, Password: user.Password, Permission: uint(user.Permission)})
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (aR *userRepository) FindByLogin(login string) (*entity.User, error) {
	logF := advancedlog.FunctionLog(aR.log)
	var user *scheme.User
	result := aR.db.Where("login = ?", login).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.User{
		ID:         user.ID,
		Login:      user.Login,
		Password:   user.Password,
		Permission: user.Permission,
	}, nil
}

func (aR *userRepository) DeleteByID(id uint) error {
	logF := advancedlog.FunctionLog(aR.log)

	result := aR.db.Delete(&scheme.User{}, id)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}
