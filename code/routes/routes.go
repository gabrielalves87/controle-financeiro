package routes

import (
	"github.com/gabrielalves87/controle-financeiro/controllers"
	"github.com/gin-gonic/gin"
)

func Handlerequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:nome", controllers.Saudacao)
	r.POST("/user", controllers.CriaNovoUsuario)
	r.POST("/categorie", controllers.CriaNovaCategoria)
	r.Run("0.0.0.0:3000")
}
