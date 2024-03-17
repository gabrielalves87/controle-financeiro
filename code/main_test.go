package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielalves87/controle-financeiro/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasdeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasdeTeste()
	r.GET("/alunos/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/alunos/gabriel", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)

	// }
}
