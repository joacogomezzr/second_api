package routers

import (
	"newapi/src/redomendaciones/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RecomendationRouter(r *gin.Engine, recomendationController *controllers.RegisterRecomendationController, notificationController *controllers.NotificationController) {
	v1 := r.Group("/v1/recomendaciones")
	{
		v1.POST("/", recomendationController.RegisterRecomendation)
		v1.GET("/ws", notificationController.WebSocketHandler) 
	}
}