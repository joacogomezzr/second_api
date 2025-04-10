package dependencies

import (
	"newapi/src/redomendaciones/application"
	"newapi/src/redomendaciones/infrastructure/adapters"
	"newapi/src/redomendaciones/infrastructure/controllers"
	"newapi/src/redomendaciones/infrastructure/routers"
	"github.com/gin-gonic/gin"
)

func RecomendacionSetup(r *gin.Engine) {
	ps, err := adapters.NewMySQL()
	if err != nil {
		panic(err)
	}
	recomendationRegister:= application.NewRegisterRecomendacionUseCase(ps)
	recomendationController:= controllers.NewRegisterRecomendationController(recomendationRegister)

	routers.RecomendationRouter(r,recomendationController)
	



	 
}