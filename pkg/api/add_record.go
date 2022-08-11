package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trichardsonjr78/terraform-operator-api/pkg/common/models"
)

type AddTFOResourceBody struct {
	UUID              string `json:"uuid" gorm:"primaryKey"`
	CreatedBy         string `json:"createdby"`
	CreatedAt         string `json:"createdat"`
	UpdatedBy         string `json:"updatedby"`
	UpdatedAt         string `json:"updatedat"`
	DeletedBy         string `json:"deletedby"`
	DeletedAt         string `json:"deleetedat"`
	NamespacedName    string `json:"namespacedname"`
	CurrentGeneration string `json:"currentgeneration"`
}

func (h handler) AddTfoResource(c *gin.Context) {
	body := AddTFOResourceBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var add_resource models.TFOResource

	add_resource.UUID = body.UUID
	add_resource.CreatedAt = body.CreatedAt
	add_resource.UpdatedBy = body.UpdatedBy
	add_resource.UpdatedAt = body.UpdatedAt
	add_resource.DeletedBy = body.DeletedBy
	add_resource.DeletedAt = body.DeletedAt
	add_resource.NamespacedName = body.NamespacedName
	add_resource.CurrentGeneration = body.CurrentGeneration

	if result := h.DB.Create(&add_resource); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &add_resource)
}
