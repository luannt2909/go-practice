package cachefx

import "go.uber.org/fx"

var Module = fx.Provide(provideCache)