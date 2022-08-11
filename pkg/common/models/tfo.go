package models

import "gorm.io/gorm"

type TFOTaskLog struct {
	gorm.Model
	TaskType        string
	Generation      string
	Rerun           int
	Message         string
	TFOResource     TFOResource
	TFOResourceUUID string `json:"tfo_resource_uuid"`
	LineNo          string
}

type TFOResource struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
	DeletedBy string
	DeletedAt string

	// NamespacedName comprises a resource name, with a mandatory namespace,
	// rendered as "<namespace>/<name>".
	NamespacedName string

	CurrentGeneration string
}
