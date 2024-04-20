package product

import (
	"raffalda-api/internal/application/config"
	"raffalda-api/internal/infrastructure/controllers/fiberHTTP"
	apihandler "raffalda-api/internal/infrastructure/controllers/fiberHTTP/handlers/apiHandler"
	authhandler "raffalda-api/internal/infrastructure/controllers/fiberHTTP/handlers/authHandler"
	"raffalda-api/internal/infrastructure/controllers/fiberHTTP/middleware"
	"raffalda-api/internal/infrastructure/controllers/metrics"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	fiberController fiberHTTP.FiberController
	metrics         metrics.Metrics
}

func NewController(gateway *Gateway, metricsConfig config.MetricsConfig, log *logrus.Entry) *Controller {
	middlew := middleware.NewMiddleware(gateway.Services.Auth, log.WithField("location", "middleware"))
	return &Controller{
		fiberController: fiberHTTP.NewFiberController(authhandler.NewHandlerAuth(gateway.Services.Auth, middlew, log.WithField("location", "handler-auth")), apihandler.NewHandlerApi(gateway.Services.WarehouseService, gateway.Services.MerchandiseParser, gateway.Services.MerchandiseService, gateway.Services.SoldPoint, gateway.Services.Transaction, gateway.Services.AnalyzeService, log.WithField("location", "handler-api")), middlew),
		metrics:         metrics.NewMetrics(metricsConfig, log.WithField("location", "metrics-listener")),
	}
}

func (c *Controller) ConfigureFiber(r *fiber.App) error {
	c.fiberController.RegisterRoutes(r)

	return nil
}

func (c *Controller) ListenMetrics() error {
	return c.metrics.Listen()
}
