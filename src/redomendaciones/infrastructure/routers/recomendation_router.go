package routers

import (
	"newapi/src/redomendaciones/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RecomendationRouter(r *gin.Engine, recomendationController *controllers.RegisterRecomendationController) {
	// r.POST("/recomendacion", recomendationController.RegisterRecomendation)
	v1:= r.Group("/v1/recomendaciones")
	{
		v1.POST("/",recomendationController.RegisterRecomendation )
	}
}