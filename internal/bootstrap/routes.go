/*
Package Name: bootstrap
File Name: routes.go
Abstract: The wrapper for setting up all the routes.
*/
package bootstrap

import (
	"llm-backend/pkg/auth"
	"llm-backend/pkg/users"
)

// ======== TYPES ========

// Route interface
type Route interface {
	Setup()
}

// Routes contains multiple routes
type Routes []Route

// ======== PUBLIC METHODS ========

// GetRoutes provides all the routes
func GetRoutes(
	userRoutes users.UsersRoutes,
	authRoutes auth.AuthRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
	}
}

// Sets up all the routes
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
