package dependencies

import (
	"newapi/src/redomendaciones/application"
	"newapi/src/redomendaciones/infrastructure/adapters"
	"newapi/src/redomendaciones/infrastructure/controllers"
	"newapi/src/redomendaciones/infrastructure/routers"

	"github.com/gin-gonic/gin"
	"fmt"
	"os"

)

func RecomendacionSetup(r *gin.Engine) {
	// Configurar el repositorio MySQL
	ps, err := adapters.NewMySQL()
	if err != nil {
		panic(err) // Maneja errores apropiadamente en producción
	}
	
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
    rabbitmqPass := os.Getenv("RABBITMQ_PASS")
    rabbitmqHost := os.Getenv("RABBITMQ_HOST")
    rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPass, rabbitmqHost, rabbitmqPort)
	rabbitmqConn := adapters.NewRabbitMQAdapter(connStr)

	// Configurar el repositorio para WebSocket (notificaciones)
	notificationRepo := &adapters.NotificationRepository{}

	// Crear casos de uso
	recomendationRegister := application.NewRegisterRecomendacionUseCase(ps, notificationRepo, rabbitmqConn)

	// Crear controladores
	recomendationController := controllers.NewRegisterRecomendationController(recomendationRegister)
	notificationController := &controllers.NotificationController{
		UseCase: &application.NotificationUseCase{Repo: notificationRepo},
	}

	// Registrar rutas
	routers.RecomendationRouter(r, recomendationController, notificationController)
}