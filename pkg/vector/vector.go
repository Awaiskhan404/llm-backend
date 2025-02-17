package vector

import "go.uber.org/fx"

// ======== EXPORTS ========

// Module exports services present
var Context = fx.Options(
	fx.Provide(SetVectorRoutes),
	fx.Provide(GetVectorController),
	fx.Provide(GetVectorService),
)