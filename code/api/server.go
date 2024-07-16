package api
import (
	db "github.com/gabrielalves87/controle-financeiro/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct{
	store *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server{
	server := &Server{store:store}
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/:id", server.getUserbyId)
	server.router = router
	return server
}


func(server *Server) Start(address string) error {
	return server.router.Run()
}

func errorResponse(err error) gin.H{
	return gin.H{
		"api has error:": err.Error(),
	}
}