package fiberHTTP

import (
	"raffalda-api/internal/infrastructure/controllers/fiberHTTP/middleware"

	"github.com/gofiber/fiber/v2"
)

type FiberController interface {
	RegisterRoutes(r *fiber.App)
}

type fiberController struct {
	authHandler HandlerFiber
	apiHanlder  HandlerFiber

	middleware middleware.Middleware
}

type HandlerFiber interface {
	RegisterGroup(g fiber.Router)
}

func NewFiberController(authHandler, apiHandler HandlerFiber, middleware middleware.Middleware) FiberController {
	return &fiberController{authHandler: authHandler, apiHanlder: apiHandler, middleware: middleware}
}

func (fC *fiberController) RegisterRoutes(app *fiber.App) {

	// policyChecker := fC.middleware.CreatePolicyFunc()

	authGroup := app.Group("/auth")
	// apiGroup := app.Group("/api", policyChecker)
	apiGroup := app.Group("/api")

	fC.authHandler.RegisterGroup(authGroup)
	fC.apiHanlder.RegisterGroup(apiGroup)
}
