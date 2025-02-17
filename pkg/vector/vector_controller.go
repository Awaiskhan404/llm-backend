package vector

import (
	"llm-backend/pkg/common"
	"llm-backend/pkg/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ======== TYPES ========

// VectorController struct
type VectorController struct {
	logger lib.Logger
	service VectorRepository
}

type VectorBody struct {
	Name string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"omitempty"`
	ConnectionString string `json:"connection_string" form:"connection_string" binding:"required"`
}

// Returns a VectorController struct.
func GetVectorController(logger lib.Logger, service VectorRepository) VectorController {
	return VectorController{
		logger: logger,
		service: service,
	}
}

// Hello endpoint - example function
func (controller VectorController) Hello(ctx *gin.Context) {
	controller.logger.Info("Hello from VectorController")
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello from Vector!"})
}

// Create endpoint - example function
func (controller VectorController) Create(ctx *gin.Context) {
	

	body := VectorBody{}


	if errors := common.Validation.ValidateBody(ctx, &body); errors != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors)
		return
	}

	object, err := controller.service.CreateVector(body.Name, body.Description, body.ConnectionString)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error while creating vector"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Vector created successfully",
		"vector": object,
	})
}


func (controller VectorController) GetAll(ctx *gin.Context) {
	vectors, err := controller.service.GetAllVectors()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error while fetching vectors"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"vectors": vectors,
	})
}