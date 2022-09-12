package models

import "gorm.io/gorm"

type TFOTaskLog struct {
	gorm.Model
	TaskType        string `json:"taskType"`
	Generation      string `json:"generation"`
	Rerun           int    `json:"rerun"`
	Message         string `json:"message"`
	TFOResource     TFOResource
	TFOResourceUUID string `json:"tfo_resource_uuid"`
	LineNo          string `json:"lineNo"`
}

type TFOResource struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	CreatedBy string `json:"createdby"`
	CreatedAt string `json:"createdat"`
	UpdatedBy string `json:"updatedby"`
	UpdatedAt string `json:"updatedat"`
	DeletedBy string `json:"deletedby"`
	DeletedAt string `json:"deleetedat"`

	// foreign key to a cluster
	Cluster   Cluster
	ClusterID string `json:"cluster_id"`

	CurrentGeneration string `json:"currentgeneration"`
}

type Cluster struct {
	ClusterID string `json:"cluster_id" gorm:"primaryKey"`
}
