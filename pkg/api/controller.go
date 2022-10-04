package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/api/v1/")
	routes.GET("/", h.GetRecords)
	routes.GET("/clusters", h.GetClusters)
	routes.GET("/cluster/:cluster_id", h.GetClustersResources)
	routes.GET("/cluster_id/:cluster_id", h.GetUuidByClusterID)
	routes.GET("/resource/:resource_uuid/logs/generation/:generation", h.GetClustersResourcesLogs)
	//routes.GET("/cluster/:cluster_id/resource/:resource_uuid/logs/generation/:generation", h.GetUuidByClusterID)
	routes.GET("/cluster_name/:cluster_name", h.GeIdByClusterName)
	routes.GET("/:tfo_resource_uuid", h.GetLog)
	routes.GET("/logsByGeneration/:tfo_resource_uuid/:generation", h.GetLogByGeneration)
	routes.GET("/queryRerun/:tfo_resource_uuid/:task_type/:rerun_value/:generation", h.GetRerunByNumber)
	//routes.GET("/highestRerun/:tfo_resource_uuid/:task_type", h.GetHighestRerunLog)
	routes.GET("/highestRerun/:generation", h.GetHighestRerunLog)
}
