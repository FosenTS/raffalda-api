package product

import (
	"raffalda-api/internal/application/config"
	"raffalda-api/internal/domain/service"
	"raffalda-api/pkg/ajwt"
	"raffalda-api/pkg/passlib"

	"github.com/sirupsen/logrus"
)

type Services struct {
	Auth               service.Auth
	MerchandiseParser  service.MerchandiseParser
	MerchandiseService service.MerchandiseService
	WarehouseService   service.Warehouse
}

func NewServices(
	storage *Storage,
	authConfig config.AuthConfig,
	merchandiseParserConfig config.MerchandiseParserConfig,
	log *logrus.Entry,
) *Services {
	hashManager := passlib.NewHashManager(authConfig.Salt)
	jwtManager := ajwt.NewJWTManager(hashManager, authConfig.SecretJWTKey, authConfig.JwtLiveTime, authConfig.RefreshLiveTime)

	authService := service.NewAuth(storage.User, storage.RefreshTokens, hashManager, jwtManager, authConfig, log.WithField("location", "auth-service"))

	merchandiseParser := service.NewMerchandiseParser(merchandiseParserConfig)

	merchandise := service.NewMerchandiseService(storage.Merchandise, log.WithField("location", "merchandise-service"))

	warehouse := service.NewWarehouse(storage.Warehouse, storage.Merchandise, log.WithField("location", "warehouse-service"))

	return &Services{
		Auth:               authService,
		MerchandiseParser:  merchandiseParser,
		MerchandiseService: merchandise,
		WarehouseService:   warehouse,
	}
}
