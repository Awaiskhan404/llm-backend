/*
Package Name: users
File Name: users.go
Abstract: Wrapper for exposing to fx all the components of the 'users' context.
*/
package users

import "go.uber.org/fx"

// ======== EXPORTS ========

// Module exports services present
var Context = fx.Options(
	fx.Provide(GetUsersController),
	fx.Provide(GetUsersService),
	fx.Provide(SetUsersRoutes),
)
