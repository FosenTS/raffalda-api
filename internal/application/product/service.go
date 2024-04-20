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
	SoldPoint          service.SoldPoint
	Transaction        service.Transaction
	MarksService       service.Marks
	AnalyzeService     service.Analyze
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

	soldPoint := service.NewSoldPoint(storage.SoldPoint, log.WithField("location", "sold-point-service"))

	transaction := service.NewTransaction(storage.Transaction, storage.Warehouse, storage.SoldPoint, log.WithField("location", "transaction-service"))

	marksService := service.NewMarks(storage.SoldPoint, storage.Warehouse, storage.Marks, log.WithField("location", "marks-service"))

	analyzeService := service.NewAnalyze(storage.Warehouse, storage.Merchandise, storage.SoldPoint, storage.Transaction, log.WithField("location", "analyze-service"))

	return &Services{
		Auth:               authService,
		MerchandiseParser:  merchandiseParser,
		MerchandiseService: merchandise,
		WarehouseService:   warehouse,
		SoldPoint:          soldPoint,
		Transaction:        transaction,
		MarksService:       marksService,
		AnalyzeService:     analyzeService,
	}
}
