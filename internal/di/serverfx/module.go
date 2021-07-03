package serverfx

import "go.uber.org/fx"

var Module = fx.Provide(
	provideUserController,
	provideUserRouter)
