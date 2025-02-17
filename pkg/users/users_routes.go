/*
Package Name: users
File Name: users_routes.go
Abstract: The file containing all the user routes.
*/
package users

import (
	"llm-backend/internal/middlewares"

	"llm-backend/pkg/lib"
)

// ======== TYPES ========

// UsersRoutes struct
type UsersRoutes struct {
	logger          lib.Logger
	router          *lib.Router
	usersController UsersController
	authMiddleware  middlewares.AuthMiddleware
}

// ======== PUBLIC METHODS ========

// Returns a UserRoutes struct.
func SetUsersRoutes(
	logger lib.Logger,
	router *lib.Router,
	usersController UsersController,
	authMiddleware middlewares.AuthMiddleware,
) UsersRoutes {
	return UsersRoutes{
		logger:          logger,
		router:          router,
		usersController: usersController,
		authMiddleware:  authMiddleware,
	}
}

// Setup the user routes
func (route UsersRoutes) Setup() {
	route.logger.Info("Setting up [USERS] routes.")
	api := route.router.Group("/users").Use(route.authMiddleware.Handler())
	{
		api.GET("/", route.usersController.GetAll)
		api.GET("/:id", route.usersController.Get)
	}
}
