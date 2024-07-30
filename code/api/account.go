package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"github.com/gabrielalves87/controle-financeiro/utils"
	db "github.com/gabrielalves87/controle-financeiro/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	UserID       int32     `json:"user_id" binding:"required"`
	CategoriesID int32     `json:"categories_id"  binding:"required"`
	Title        string    `json:"title" binding:"required"`
	Type         string    `json:"type" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Value        int32     `json:"value" binding:"required"`
	Date         time.Time `json:"date" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	err := utils.GetTokenAndVerify(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req createAccountRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	var categoryID = req.CategoriesID
	var accountType = req.Type
	
	category, err := server.store.GetCategory(ctx, categoryID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	categoryTypeIsDifferentofAccountType := category.Type != accountType
	if categoryTypeIsDifferentofAccountType {
		ctx.JSON(http.StatusBadRequest, "Account type different of Category type")
		return
	 }

	
	arg := db.CreateAccountParams{
		UserID:       req.UserID,
		CategoriesID: categoryID,
		Title:        req.Title,
		Type:         accountType,
		Description:  req.Description,
		Value:        req.Value,
		Date:         req.Date,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, account)

}

type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	err := utils.GetTokenAndVerify(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}	
	var req getAccountRequest
	err = ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	accounts, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type deleteAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	err := utils.GetTokenAndVerify(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req deleteAccountRequest
	err = ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	err = server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateAccountRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	err := utils.GetTokenAndVerify(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}	
	var req updateAccountRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.UpdateAccountParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Value:       req.Value,
	}

	account, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, account)

}

type getAccountsRequest struct {
	UserID       int32     `json:"user_id" binding:"required"`
	Type         string    `json:"type"`
	CategoriesID int32     `json:"categories_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
}

func (server *Server) getAccounts(ctx *gin.Context) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	err := utils.GetTokenAndVerify(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req getAccountsRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.GetAccountsParams{
		UserID:       req.UserID,
		CategoriesID: sql.NullInt32{
			Int32: req.CategoriesID,
			Valid: req.CategoriesID > 0,
		},
		Type:         req.Type,
		Title:        req.Title,
		Description:  req.Description,
		Date:         sql.NullTime{
			Time: req.Date,
			Valid: !req.Date.IsZero(),
		},
	}
	accounts, err := server.store.GetAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type getAccountGraphRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
    Type   string `uri:"type"`
}

func (server *Server) getAccountGraph(ctx *gin.Context) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	err := utils.GetTokenAndVerify(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}	
	var req getAccountGraphRequest
	err = ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.GetAccountsGraphParams{
		UserID: req.UserID,
		Type  : req.Type,
	}

	countgraph, err := server.store.GetAccountsGraph(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, countgraph)
}

type getAccountReportRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
    Type   string `uri:"type"`
}

func (server *Server) getAccountReport(ctx *gin.Context) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	err := utils.GetTokenAndVerify(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}	
	var req getAccountReportRequest
	err = ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.GetAccountsReportsParams{
		UserID: req.UserID,
		Type  : req.Type,
	}

	sumreports, err := server.store.GetAccountsReports(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sumreports)
}