/*
Package Name: auth
File Name: auth_routes.go
Abstract: The routes for logging in and signing up.
*/
package auth

import "llm-backend/pkg/lib"

// ======== TYPES ========

// UserRoutes struct
type AuthRoutes struct {
	logger         lib.Logger
	router         *lib.Router
	authController AuthController
}

// ======== PUBLIC METHODS ========

// Returns an AuthRoutes struct.
func SetAuthRoutes(
	logger lib.Logger,
	router *lib.Router,
	authController AuthController,
) AuthRoutes {
	return AuthRoutes{
		router:         router,
		logger:         logger,
		authController: authController,
	}
}

// Setup the auth routes
func (route AuthRoutes) Setup() {
	route.logger.Info("Setting up [AUTH] routes.")
	route.router.POST("/login", route.authController.Login)
	route.router.POST("/signup", route.authController.Signup)
}
