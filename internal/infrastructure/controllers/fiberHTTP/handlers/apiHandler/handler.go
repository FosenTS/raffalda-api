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
	log                *logrus.Entry
}

func NewHandlerApi(warehouse service.Warehouse, merchandiseParser service.MerchandiseParser, merchandiseService service.MerchandiseService, log *logrus.Entry) fiberHTTP.HandlerFiber {
	return &handlerApi{
		warehousService:    warehouse,
		merchandiseParser:  merchandiseParser,
		merchandiseService: merchandiseService,
		log:                log,
	}
}
