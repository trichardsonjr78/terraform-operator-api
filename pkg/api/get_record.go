package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trichardsonjr78/terraform-operator-api/pkg/common/models"
)

// func (h handler) GetLog(c *gin.Context) {
// 	uuid := c.Param("id")

// 	var log models.TFOTaskLog

// 	if result := h.DB.First(&log, uuid); result.Error != nil {
// 		c.AbortWithError(http.StatusNotFound, result.Error)
// 		return
// 	}

// 	c.JSON(http.StatusOK, &log)
// }

func (h handler) GetLog(c *gin.Context) {
	uuid := c.Param("tfo_resource_uuid")
	var log models.TFOTaskLog

	if result := h.DB.First(&log, "tfo_resource_uuid = ?", uuid); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &log)
}
