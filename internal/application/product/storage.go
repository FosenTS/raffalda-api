package product

import (
	"raffalda-api/internal/domain/storage"
	"raffalda-api/internal/domain/storage/gormDB"
	"raffalda-api/pkg/advancedlog"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Storage struct {
	User          storage.User
	RefreshTokens storage.RefreshToken
	Merchandise   storage.Merchandise
	Warehouse     storage.Warehouse
	SoldPoint     storage.SoldPoint
	Transaction   storage.Transaction
	Marks         storage.Marks
}

func NewStorage(db *gorm.DB, log *logrus.Entry) (*Storage, error) {
	logF := advancedlog.FunctionLog(log)
	userStorage, err := gormDB.NewUserRepository(db, log.WithField("location", "gorm-user-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	refreshTokenStorage, err := gormDB.NewRefreshTokenRepository(db, log.WithField("location", "gorm-refresh-token-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	merchandiseStorage, err := gormDB.NewMerchandiseRepository(db, log.WithField("location", "gorm-merchandise-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	warehouseStorage, err := gormDB.NewWarehouseRepository(db, log.WithField("location", "gorm-warehouse-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	soldPointStorage, err := gormDB.NewGormSoldPointRepository(db, log.WithField("location", "gorm-sold-point-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	transactionStorage, err := gormDB.NewGormTransactionRepository(db, log.WithField("location", "gorm-transaction-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	marksStorage, err := gormDB.NewMarksRepository(db, log.WithField("location", "gorm-marks-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return &Storage{
		User:          userStorage,
		RefreshTokens: refreshTokenStorage,
		Merchandise:   merchandiseStorage,
		Warehouse:     warehouseStorage,
		SoldPoint:     soldPointStorage,
		Marks:         marksStorage,
		Transaction:   transactionStorage,
	}, nil
}
