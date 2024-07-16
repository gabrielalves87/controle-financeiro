package routes

import (
	"github.com/gabrielalves87/controle-financeiro/controllers"
	docs "github.com/gabrielalves87/controle-financeiro/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handlerequest() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:nome", controllers.Saudacao)
	r.POST("/user", controllers.CriaNovoUsuario)
	r.POST("/categorie", controllers.CriaNovaCategoria)
	r.Run("0.0.0.0:3000")
	// err := r.Run("0.0.0.0:3000")
	// if err != nil {
	// 	log.Fatalf("Failed to run server: %v", err)
	// }
}
