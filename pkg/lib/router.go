/*
Package Name: lib
File Name: logger.go
Abstract: The route handler.
*/
package lib

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ======== TYPES ========

// Router type
type Router = gin.Engine

// ======== METHODS ========

// GetRouter retrieves the router used by the API.
func GetRouter() *Router {

	// ======== ROUTER ========
	router := gin.New()

	// ======== LOGGER ========
	logger, _ := zap.NewProduction()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger, true))

	// ======== ERROR HANDLING ========

	// ======== SETTINGS ========
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)

	return router

}
