/*
Package Name: users
File Name: users.go
Abstract: Wrapper for exposing to fx all the components of the 'users' context.
*/
package auth

import "go.uber.org/fx"

// ======== EXPORTS ========

// Module exports services present
var Context = fx.Options(
	fx.Provide(GetAuthController),
	fx.Provide(GetAuthService),
	fx.Provide(SetAuthRoutes),
)
