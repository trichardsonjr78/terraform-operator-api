package api

import (
	"net/http"

	"github.com/GalleyBytes/terraform-operator-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusInfo StatusInfo  `json:"StatusInfo"`
	Data       interface{} `json:"Data"`
}

type StatusInfo struct {
	StatusCode int64  `json:"StatusCode"`
	Message    string `json:"Message"`
}

func response(httpstatus int64, message string, results interface{}) *Response {
	resp := Response{
		StatusInfo: StatusInfo{
			StatusCode: httpstatus,
			Message:    message,
		},
		Data: results,
	}
	return &resp
}

func (h handler) GetLog(c *gin.Context) {
	uuid := c.Param("tfo_resource_uuid")
	var log models.TFOTaskLog

	if result := h.DB.First(&log, "tfo_resource_uuid = ?", uuid); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		resp := response(http.StatusNotFound, "tfo resource uuid not found", log)
		c.JSON(int(resp.StatusInfo.StatusCode), &resp)
		return
	}
	resp := response(http.StatusOK, "tfo resource uuid found", log)
	c.JSON(int(resp.StatusInfo.StatusCode), &resp)
}

func (h handler) GetLogByGeneration(c *gin.Context) {
	uuid := c.Param("tfo_resource_uuid")
	generation := c.Param("generation")
	var generationLogs []models.TFOTaskLog

	if result := h.DB.Where("generation = ? AND tfo_resource_uuid = ?", &generation, &uuid).Find(&generationLogs); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		resp := response(http.StatusNotFound, "generation logs not found", generationLogs)
		c.JSON(int(resp.StatusInfo.StatusCode), &resp)
		return
	}

	resp := response(http.StatusOK, "generation logs found", generationLogs)
	c.JSON(int(resp.StatusInfo.StatusCode), &resp)

}

func (h handler) GetDistinctGeneration(c *gin.Context) {
	uuid := c.Param("resource_uuid")
	var generation []int
	if result := h.DB.Raw("SELECT DISTINCT generation FROM tfo_task_logs WHERE tfo_resource_uuid = ?", &uuid).Scan(&generation); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &generation)
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
	var records []models.TFOResource

	if result := h.DB.Find(&records); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &records)
}

func (h handler) GetClusters(c *gin.Context) {
	var clusters []models.Cluster

	if result := h.DB.Find(&clusters); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &clusters)
}

func (h handler) GetClustersResources(c *gin.Context) {
	var resources []models.TFOResource
	clusterID := c.Param("cluster_id")

	if result := h.DB.Where("cluster_id = ?", clusterID).Find(&resources); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &resources)
}

func (h handler) GetResourceByUUID(c *gin.Context) {
	var tfoResource models.TFOResource
	uuid := c.Param("resource_uuid")

	if result := h.DB.First(&tfoResource, "uuid = ?", uuid); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tfoResource)
}

// func (h handler) GetResourceLogsByUUID(c *gin.Context) {
// 	var tfoResource models.TFOTaskLog
// 	uuid := c.Param("resource_uuid")

// 	if result := h.DB.First(&tfoResource, "uuid = ?", uuid); result.Error != nil {
// 		c.AbortWithError(http.StatusNotFound, result.Error)
// 		return
// 	}

// 	c.JSON(http.StatusOK, &tfoResource)
// }

func (h handler) GetClustersResourcesLogs(c *gin.Context) {
	var logs []models.TFOTaskLog
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

	if result := h.DB.Where("tfo_resource_uuid = ? AND generation = ?", &uuid, &generation).Find(&logs); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &logs)
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
