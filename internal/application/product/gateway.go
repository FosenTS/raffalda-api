package product

type Gateway struct {
	*Services
}

func NewGateway(services *Services) *Gateway {
	return &Gateway{
		Services: services,
	}
}
