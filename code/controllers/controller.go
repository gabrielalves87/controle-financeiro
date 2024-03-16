package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// @BasePath /

// ExibeTodosAlunos godoc
// @Summary ExibeTodosAlunos example
// @Schemes
// @Description do pinExibeTodosAlunosg
// @Tags ExibeTodosAlunos
// @Accept json
// @Produce json
// @Success 200 {string} ExibeTodosAlunos
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Gab Alves",
	})

}

// @BasePath /

// Saudacao godoc
// @Summary Saudacao example
// @Schemes
// @Description do Saudacao
// @Tags Saudacao
// @Accept json
// @Param name path string true "Nome"
// @Produce json
// @Success 200 {string} Saudacao
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos/{name} [get]
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API": "E ai " + nome + ", tudo blz",
	})
}
