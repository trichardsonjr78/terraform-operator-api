package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trichardsonjr78/terraform-operator-api/pkg/common/models"
)

func (h handler) AddTfoResource(c *gin.Context) {
	tfo_resource := models.TFOResource{}

	// getting request's body
	if err := c.BindJSON(&tfo_resource); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.DB.Create(&tfo_resource); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &tfo_resource)
}
