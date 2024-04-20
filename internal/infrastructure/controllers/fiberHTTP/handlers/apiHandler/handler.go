package apihandler

import (
	"raffalda-api/internal/domain/service"
	"raffalda-api/internal/infrastructure/controllers/fiberHTTP"

	"github.com/sirupsen/logrus"
)

type handlerApi struct {
	warehousService    service.Warehouse
	merchandiseParser  service.MerchandiseParser
	merchandiseService service.MerchandiseService
	soldPointService   service.SoldPoint
	transactionService service.Transaction
	marksService       service.Marks
	analyzeService     service.Analyze
	log                *logrus.Entry
}

func NewHandlerApi(warehouse service.Warehouse, merchandiseParser service.MerchandiseParser, merchandiseService service.MerchandiseService, soldPointService service.SoldPoint, transactionService service.Transaction, marksService service.Marks, analyzeService service.Analyze, log *logrus.Entry) fiberHTTP.HandlerFiber {
	return &handlerApi{
		warehousService:    warehouse,
		merchandiseParser:  merchandiseParser,
		merchandiseService: merchandiseService,
		soldPointService:   soldPointService,
		transactionService: transactionService,
		marksService:       marksService,
		analyzeService:     analyzeService,
		log:                log,
	}
}
