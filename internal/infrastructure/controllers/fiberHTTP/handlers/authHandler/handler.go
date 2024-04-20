package authhandler

import (
	"raffalda-api/internal/domain/service"
	"raffalda-api/internal/infrastructure/controllers/fiberHTTP"
	"raffalda-api/internal/infrastructure/controllers/fiberHTTP/middleware"

	"github.com/sirupsen/logrus"
)

type handlerAuth struct {
	authService service.Auth

	middleware middleware.Middleware

	log *logrus.Entry
}

func NewHandlerAuth(authService service.Auth, middleware middleware.Middleware, log *logrus.Entry) fiberHTTP.HandlerFiber {
	return &handlerAuth{
		authService: authService,
		middleware:  middleware,
		log:         log,
	}
}
