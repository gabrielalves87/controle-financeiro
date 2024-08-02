package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gabrielalves87/controle-financeiro/database"
	"github.com/gabrielalves87/controle-financeiro/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CriaNovaCategoria(c *gin.Context) {
	var categoria models.Categorie
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	result := database.DB.Create(&categoria)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ErrForeignKeyViolated"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error})
		return

	}
	log.Println(result.Error)
	log.Println(result.RowsAffected)
	c.JSON(http.StatusOK, categoria)
}
