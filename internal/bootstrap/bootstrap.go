/*
Package Name: bootstrap
File Name: bootstrap.go
Abstract: Wrapper for invoking all the module's dependencies and starting
the API by loading the essential initial components that allow it to run
and perform more complex tasks.
*/
package bootstrap

import (
	"context"
	"fmt"
	"os"

	"llm-backend/internal/middlewares"
	"llm-backend/pkg/auth"
	"llm-backend/pkg/lib"
	"llm-backend/pkg/users"
	"llm-backend/pkg/vector"

	"go.uber.org/fx"
)

// ======== PRIVATE METHODS ========

// registerHooks registers a lifecycle hook that starts the API and logs a message
// when the app is stopped.
func registerHooks(
	lifecycle fx.Lifecycle,
	router *lib.Router,
	logger lib.Logger,
	routes Routes,
	middlewares middlewares.Middlewares,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				// Log the start of the application with the configured host and port
				logger.Info(
					fmt.Sprintf(
						"Starting application in %s:%s",
						os.Getenv("APP_HOST"),
						os.Getenv("APP_PORT"),
					),
				)

				// ======== SET UP COMPONENTS ========
				// Perform any necessary setup or initialization tasks for the middlewares
				middlewares.Setup()

				// Perform any necessary setup or initialization tasks for the routes
				routes.Setup()

				// Start the router by running it in a separate goroutine
				go router.Run(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))

				return nil
			},
			OnStop: func(ctx context.Context) error {
				// Log the stop of the application and any associated error
				logger.Fatal(
					fmt.Sprintf(
						"Stopping application. Error: %s", ctx.Err(),
					),
				)

				return nil
			},
		},
	)
}

// ======== EXPORTS ========

// Module exports for fx
var Module = fx.Options(
	// Module exports
	lib.Module,
	middlewares.Module,

	// Context exports
	users.Context,
	auth.Context,
	vector.Context,
	// Bootstrap exports
	fx.Provide(GetRoutes),

	// Methods
	fx.Invoke(registerHooks),
)
