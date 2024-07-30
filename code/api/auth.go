package api

import (
	"bytes"
	"crypto/sha512"
	"database/sql"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type claims struct {
	Username string `username:"foo"`
	jwt.RegisteredClaims
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hashedInput := sha512.Sum512_256([]byte(req.Password))
	// Bcrypt terminates at NULL bytes, so we need to trim these away
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPasswordInput := string(trimmedHash)
	reqPasswordInbytes := []byte(preparedPasswordInput)
	userHashedFromDb := user.Password
	userHashedinBytes := []byte(userHashedFromDb)
	err = bcrypt.CompareHashAndPassword(userHashedinBytes, reqPasswordInbytes)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "User or Passoword wrong")
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)

	var jwtSignedKey = []byte("secret_key")
	
	claims := &claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "gofinance",

		},
	}
	
	
	
	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	stringToken, err := generatedToken.SignedString(jwtSignedKey)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, stringToken)

}
