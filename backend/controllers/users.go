package controllers

import (
	"net/http"

	"github.com/gabrielalves87/controle-financeiro/database"
	"github.com/gabrielalves87/controle-financeiro/models"
	"github.com/gin-gonic/gin"
)

func CriaNovoUsuario(c *gin.Context) {
	var usuario models.User
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&usuario)
	c.JSON(http.StatusOK, usuario)
}
