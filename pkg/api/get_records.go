package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trichardsonjr78/terraform-operator-api/pkg/common/models"
)

func (h handler) GetRecords(c *gin.Context) {
	var Records []models.TFOResource

	if result := h.DB.Find(&Records); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &Records)
}
