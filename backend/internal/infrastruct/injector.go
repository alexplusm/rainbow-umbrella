package infrastruct

import (
	"rainbow-umbrella/internal/controllers"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/services"
)

type IInjector interface {
	InjectUserController() interfaces.IUserController
}

type injector struct {
}

func NewInjector() IInjector {
	return &injector{}
}

func (i injector) InjectUserController() interfaces.IUserController {
	return controllers.NewUserController(i.injectUserService())
}

// ---

func (i injector) injectUserService() interfaces.IUserService {
	return services.NewUserService()
}
