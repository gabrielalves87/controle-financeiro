package api

import (
	"database/sql"
	"net/http"

	db "github.com/gabrielalves87/controle-financeiro/db/sqlc"
	"github.com/gin-gonic/gin"
)


type createUserRequest struct{
	Username string `json:"username" binding:"reqired"` 
	Password string `json:"password" binding:"reqired"`
	Email    string `json:"email" binding:"reqired"`
	
}

func(server *Server) createUser(ctx *gin.Context){
	var req createUserRequest
	err := ctx.ShouldBindJSON(&req)
	if  err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
		Email: req.Email,
	}

	user, err:= server.store.CreateUser(ctx,arg)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK,user)

}


type getUserRequest struct{
	Username string `uri:"username" binding:"reqired"` 
}

func(server *Server) getUser(ctx *gin.Context){
	var req getUserRequest
	err := ctx.ShouldBindUri(&req)
	if  err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	user, err:= server.store.GetUser(ctx,req.Username)
	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,user)
}


type getUserbyIdRequest struct{
	ID int32 `uri:"id" binding:"reqired"` 
}

func(server *Server) getUserbyId(ctx *gin.Context){
	var req getUserbyIdRequest
	err := ctx.ShouldBindUri(&req)
	if  err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	user, err:= server.store.GetUserById(ctx,req.ID)
	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,user)
}