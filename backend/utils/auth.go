package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)


type claims struct {
	Username string `username:"foo"`
	jwt.RegisteredClaims
}

func ValidateToken(ctx *gin.Context,token string) error {
	claims := &claims{} 
	var jwtSignedKey = []byte("secret_key")
	tokenParse, err := jwt.ParseWithClaims(token, claims,
	func(t *jwt.Token)(interface{}, error){
		return jwtSignedKey,nil
	})
	if err != nil{
		if err == jwt.ErrSignatureInvalid{
			ctx.JSON(http.StatusUnauthorized,err)
			return err
		}
		if err != jwt.ErrSignatureInvalid{
			ctx.JSON(http.StatusBadRequest,err)
			return err
		}
	}
	if !tokenParse.Valid{
		ctx.JSON(http.StatusUnauthorized, "Token is invalid")
		return err
	}
	ctx.Next()
	return nil
}

func GetTokenAndVerify(ctx *gin.Context) error{
	authorizationHearderkey := ctx.GetHeader("authorization")
	fields := strings.Fields(authorizationHearderkey)
	err := ValidateToken(ctx, fields[1])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "errorResponse(err)")
		return err
	}
	return nil
}