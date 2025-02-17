/*
Package Name: main
File Name: main.go
Abstract: The entry point of the project and the source code of the
executable that will be used for initializing the API.
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"llm-backend/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

// ======== PRIVATE METHODS ========

// isValidEnvironment checks whether the environment flag is
// a valid environment name.
func isValidEnvironment(environment *string) bool {
	switch *environment {
	case
		"development",
		"production",
		"test":
		return true
	}
	return false
}

// ======== ENTRY POINT ========
func main() {

	//	======== CHECK ENVIRONMENT ========
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	//	The only available environments are "production", "development",
	//	and "test". If any other environment is provided the api should
	//	exit.
	if !isValidEnvironment(environment) {
		fmt.Println("The environment is not valid!")
		os.Exit(1)
	}

	// Set the 'ENVIRONMENT' value to the flag passed so that
	// we can check the state wherever we want.
	os.Setenv("ENVIRONMENT", *environment)

	// ======== CONFIG FILES ========
	// Load the corresponding environment file
	godotenv.Load("configs/.env." + *environment)

	// Load the main file
	godotenv.Load("configs/.env")

	// ======== DISCLAIMER ========
	fmt.Printf("Welcome to %s %s; Written by %s\n", os.Getenv("APP_NAME"), os.Getenv("APP_VERSION"), os.Getenv("APP_AUTHOR"))

	// ======== DEPENDENCY INJECTION ========
	// The api is divided following the next structure:
	// - Bootstrap: bundles all the dependency injection under one `fx.Options` variable for cleaner code.

	// If the environment is production, set the gin
	// environment to 'release'.
	if *environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// fx.NopLogger disables the logger
	fx.New(
		bootstrap.Module,
		fx.NopLogger,
	).Run()
}
