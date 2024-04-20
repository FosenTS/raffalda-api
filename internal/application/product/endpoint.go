package product

import "github.com/gofiber/fiber/v2"

type Endpoint struct {
	*Controller
}

func NewEndpoint(controller *Controller) *Endpoint {
	return &Endpoint{Controller: controller}
}

func (e *Endpoint) ConfigureFiber(r *fiber.App) {
	e.Controller.ConfigureFiber(r)
}

func (e *Endpoint) ListenMetrics() error {
	return e.Controller.ListenMetrics()

}
