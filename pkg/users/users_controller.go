/*
Package Name: users
File Name: users_controller.go
Abstract: The user controller for performing operations after a route is called.
*/
package users

import (
	"errors"
	"net/http"
	"strconv"

	"llm-backend/pkg/lib"

	"github.com/gin-gonic/gin"
)

// ======== TYPES ========

// UsersController data type
type UsersController struct {
	// service domains.UserService
	logger  lib.Logger
	service UsersRepository
}

// ======== METHODS ========

// Creates a new user controller and exposes its routes
// to the router.
func GetUsersController(logger lib.Logger, service UsersRepository) UsersController {
	return UsersController{
		logger:  logger,
		service: service,
	}
}

func (controller UsersController) Get(ctx *gin.Context) {
	// Get the id from the context
	idParam := ctx.Param("id")
	controller.logger.Info("[GET] Getting user with id", idParam)

	// ======== TYPE CONVERSION ========
	// Convert the id from string to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("The id must be an int."))
		return
	}

	// ======== RETRIEVE USER ========
	internalUser, err := controller.service.GetUserById(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// The controller.service.GetUser(int) function returns a models.InternalUser
	// struct, which contains the password. To avoid exposing this data to the
	// end user, we must convert the internal user to a public user as follows:
	publicUser := internalUser.ToPublic()

	// We can now return the user
	ctx.JSON(http.StatusOK, publicUser)
}

func (controller UsersController) GetAll(ctx *gin.Context) {
	controller.logger.Info("[GET] Getting all users.")

	// ======== RETRIEVE USER ========
	internalUsers, err := controller.service.GetUsers()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// The controller.service.GetUser(int) function returns a models.InternalUser
	// struct, which contains the password. To avoid exposing this data to the
	// end user, we must convert the internal user to a public user as follows:
	publicUsers := make([]PublicUser, len(internalUsers))
	for user := range internalUsers {
		publicUsers[user] = internalUsers[user].ToPublic()
	}

	// We can now return the user
	ctx.JSON(http.StatusOK, publicUsers)
}
