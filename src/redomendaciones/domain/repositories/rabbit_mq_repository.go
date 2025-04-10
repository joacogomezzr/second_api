package repositories

import "newapi/src/redomendaciones/domain/entities"

type RabbitMQRepository interface {
	CreateMessageRabbit(messageBook *entities.Recomendacion) error
}