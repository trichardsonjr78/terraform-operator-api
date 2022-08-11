package models

import "gorm.io/gorm"

type TFSetupLog struct {
	gorm.Model
	Generation      string
	Message         string
	TFOResource     TFOResource
	TFOResourceUUID string `json:"tfo_resource_uuid"`
	LineNo          string
}

type TFInitLog struct {
	gorm.Model
	Generation      string
	Message         string
	TFOResource     TFOResource
	TFOResourceUUID string `json:"tfo_resource_uuid"`
	LineNo          string
}

type TFPlanLog struct {
	gorm.Model
	Generation      string
	Message         string
	TFOResource     TFOResource
	TFOResourceUUID string `json:"tfo_resource_uuid"`
	LineNo          string
}

type TFApplyLog struct {
	gorm.Model
	Generation      string
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
