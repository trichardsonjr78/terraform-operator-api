package api

import (
	"net/http"

	"github.com/GalleyBytes/terraform-operator-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

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
	var gnerationLogs []models.TFOTaskLog

	if result := h.DB.Where("generation = ? AND tfo_resource_uuid = ?", &generation, &uuid).Find(&gnerationLogs); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &gnerationLogs)

}

func (h handler) GetDistinctGeneration(c *gin.Context) {
	uuid := c.Param("resource_uuid")
	var gen []int
	if result := h.DB.Raw("SELECT DISTINCT generation FROM tfo_task_logs WHERE tfo_resource_uuid = ?", &uuid).Scan(&gen); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &gen)
}

func (h handler) GetUuidByClusterID(c *gin.Context) {
	clusterID := c.Param("cluster_id")
	var clusterIdInfo models.TFOResource

	if result := h.DB.Where("cluster_id = ?", clusterID).First(&clusterIdInfo); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &clusterIdInfo)

}

func (h handler) GeIdByClusterName(c *gin.Context) {
	clusterName := c.Param("cluster_name")
	var clusterNameInfo models.Cluster

	if result := h.DB.Where("name = ?", clusterName).First(&clusterNameInfo); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &clusterNameInfo)

}

func (h handler) GetRecords(c *gin.Context) {
	var Records []models.TFOResource

	if result := h.DB.Find(&Records); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &Records)
}

func (h handler) GetClusters(c *gin.Context) {
	var Clusters []models.Cluster

	if result := h.DB.Find(&Clusters); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &Clusters)
}

func (h handler) GetClustersResources(c *gin.Context) {
	var Resources []models.TFOResource
	clusterID := c.Param("cluster_id")

	if result := h.DB.Where("cluster_id = ?", clusterID).Find(&Resources); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &Resources)
}

func (h handler) GetClustersResourcesLogs(c *gin.Context) {
	var Logs []models.TFOTaskLog
	var tfoResource models.TFOResource
	//clusterID := c.Param("cluster")
	generation := c.Param("generation")
	uuid := c.Param("resource_uuid")

	// if result := h.DB.Where("cluster_id = ?", clusterID).Find(&tfoResource); result.Error != nil {
	// 	c.AbortWithError(http.StatusNotFound, result.Error)
	// 	return
	// }
	// uuid := tfoResource.UUID

	if generation == "latest" {
		if result := h.DB.First(&tfoResource, "uuid = ?", &uuid); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		generation = tfoResource.CurrentGeneration
	}

	if result := h.DB.Where("tfo_resource_uuid = ? AND generation = ?", &uuid, &generation).Find(&Logs); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &Logs)
}

func (h handler) GetRerunByNumber(c *gin.Context) {
	uuid := c.Param("tfo_resource_uuid")
	taskType := c.Param("task_type")
	generation := c.Param("generation")
	var rerunNumbers []models.TFOTaskLog
	rerunValue := c.Param("rerun_value")
	var tfoResource models.TFOResource
	//todo query for latest

	if generation == "latest" {
		if result := h.DB.First(&tfoResource, "uuid = ?", &uuid); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		generation = tfoResource.CurrentGeneration
	}

	if result := h.DB.Where("task_type = ? AND tfo_resource_uuid = ? AND rerun = ? AND generation = ?", &taskType, &uuid, &rerunValue, &generation).Find(&rerunNumbers); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &rerunNumbers)
}

func (h handler) GetHighestRerunLog(c *gin.Context) {
	generation := c.Param("generation")
	var maxRerun int

	if result := h.DB.Raw("SELECT MAX(rerun) FROM tfo_task_logs WHERE generation = ?", &generation).Scan(&maxRerun); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &maxRerun)
}
