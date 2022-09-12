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

func (h handler) GetLogByGeneration(c *gin.Context) {
	uuid := c.Param("tfo_resource_uuid")
	generation := c.Param("generation")
	var gnerationLog models.TFOTaskLog

	if result := h.DB.Where("generation = ? AND tfo_resource_uuid = ?", &generation, &uuid).First(&gnerationLog); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &gnerationLog)

}

func (h handler) GetUuidByClusterID(c *gin.Context) {
	clusterName := c.Param("cluster_id")
	var clusterInfo models.TFOResource

	if result := h.DB.Where("cluster_id = ?", clusterName).First(&clusterInfo); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &clusterInfo)

}
