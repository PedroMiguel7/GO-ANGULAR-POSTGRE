package pessoas

import (
	"net/http"

	"github.com/PedroMiguel7/GO-ANGULAR-POSTGRE/pkg/common/models"
	"github.com/gin-gonic/gin"
)


func (h handler) GetPerson(c *gin.Context) {
	id := c.Param("id")

	var pessoa models.Pessoa
	var equipe models.Equipe

	if result := h.DB.First(&pessoa, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	id_equipe := pessoa.EquipeID

	if result := h.DB.First(&equipe, id_equipe); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	pessoa.Equipe = equipe

	c.JSON(http.StatusOK, &pessoa)
}