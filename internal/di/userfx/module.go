package userfx

import "go.uber.org/fx"

var Module = fx.Provide(
	provideUserDBRepository,
	provideUserRepository)
