package vector

import (
	"llm-backend/internal/middlewares"
	"llm-backend/pkg/lib"
)

// ======== TYPES ========

// VectorRoutes struct
type VectorRoutes struct {
	logger           lib.Logger
	router           *lib.Router
	vectorController VectorController
	authMiddleware   middlewares.AuthMiddleware
}

// Returns a VectorRoutes struct.
func SetVectorRoutes(
	logger lib.Logger,
	router *lib.Router,
	controller VectorController,
	authMiddleware middlewares.AuthMiddleware,
) VectorRoutes {
	return VectorRoutes{
		logger:           logger,
		router:           router,
		vectorController: controller,
		authMiddleware:   authMiddleware,
	}
}

// Setup the vector routes
func (route VectorRoutes) Setup() {
	route.logger.Info("Setting up [Vector] routes.")
	api := route.router.Group("/vector").Use(route.authMiddleware.Handler())
	{
		api.GET("/", route.vectorController.GetAll)
		api.POST("/", route.vectorController.Create)
		api.GET("/:id", route.vectorController.Get)
		api.PUT("/:id", route.vectorController.Update)
	}
}