/*
Package Name: lib
File Name: lib.go
Abstract: The fx provider for allowing dependency injection
for the lib files.
*/
package lib

import "go.uber.org/fx"

// ======== EXPORTS ========

// Module exports dependency
var Module = fx.Options(
	fx.Provide(
		GetLogger,
		GetDatabase,
		GetRouter,
	),
)
