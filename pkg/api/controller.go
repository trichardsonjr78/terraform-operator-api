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

	routes := r.Group("/api")
	routes.POST("/", h.AddTfoResource)
	routes.GET("/", h.GetRecords)
	routes.GET("/cluster_id/:cluster_id", h.GetUuidByClusterID)
	routes.GET("/:tfo_resource_uuid", h.GetLog)
	routes.GET("/:tfo_resource_uuid/:generation", h.GetLogByGeneration)
}
