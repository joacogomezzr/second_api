package controllers

import (
	"net/http"
	"newapi/src/redomendaciones/application"
	"newapi/src/redomendaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

type RegisterRecomendationController struct {
	useCase *application.RegisterRecomendacionUseCase
}

func NewRegisterRecomendationController(useCase *application.RegisterRecomendacionUseCase) *RegisterRecomendationController {
	return &RegisterRecomendationController{useCase: useCase}
}

func (controller *RegisterRecomendationController) RegisterRecomendation(g *gin.Context) {
	var recomendacion *entities.Recomendacion

	if err := g.ShouldBindJSON(&recomendacion); err != nil {
		g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	register, errToRegister := controller.useCase.Execute(recomendacion.IdBook)
	if errToRegister != nil {
		g.JSON(http.StatusNotAcceptable, gin.H{"error": errToRegister.Error()})
		return
	}
	g.JSON(http.StatusCreated, register)



}
