package handler

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewHealthHandler,
		NewTodoHandler,
	),
	fx.Invoke(
		RegisterRoutes,
	),
)
