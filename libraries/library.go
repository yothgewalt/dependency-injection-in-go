package libraries

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewEnvironment),
	fx.Provide(NewLogger),
	fx.Provide(NewHandler),
)
