package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/PedroMiguel7/GO-ANGULAR-POSTGRE/pkg/common/models"
)

func (h handler) GetTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &task)
}