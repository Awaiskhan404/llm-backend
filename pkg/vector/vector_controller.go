package vector

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"llm-backend/pkg/lib"
)

// ======== TYPES ========

// VectorController struct
type VectorController struct {
	logger lib.Logger
}

// Returns a VectorController struct.
func GetVectorController(logger lib.Logger) VectorController {
	return VectorController{
		logger: logger,
	}
}

// Hello endpoint - example function
func (controller VectorController) Hello(ctx *gin.Context) {
	controller.logger.Info("Hello from VectorController")
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello from Vector!"})
}