package application

import (
	"log"
	"newapi/src/redomendaciones/domain"
	"newapi/src/redomendaciones/domain/entities"
	"newapi/src/redomendaciones/domain/repositories"
)

type RegisterRecomendacionUseCase struct {
	recomendacionesRepository domain.IRecomendaciones
	notificationRepository   domain.INotification
	rabbit repositories.RabbitMQRepository
}

func NewRegisterRecomendacionUseCase(
	recomendacionesRepository domain.IRecomendaciones,
	notificationRepository domain.INotification,
	rabbit repositories.RabbitMQRepository,
) *RegisterRecomendacionUseCase {
	return &RegisterRecomendacionUseCase{
		recomendacionesRepository: recomendacionesRepository,
		notificationRepository:    notificationRepository,
		rabbit: rabbit,

	}
}

func (useCase *RegisterRecomendacionUseCase) Execute(idBook int32) (*entities.Recomendacion, error) {
	recomendacion := entities.NewRecomendacion(idBook)

	err := useCase.recomendacionesRepository.Register(recomendacion)
	if err != nil {
		return nil, err
	}
	err = useCase.rabbit.CreateMessageRabbit(recomendacion)
	if err!= nil {
		return nil, err
	}

	message := "¡El libro fue marcado como recommendable con éxito y se ha agregado a la lista!"
	err = useCase.notificationRepository.SendNotification(int64(idBook), message)
	if err != nil {
		log.Printf("Error enviando la notificación: %v", err)
	} else {
		log.Printf("Notificación enviada al ID del libro: %d", idBook)
	}

	return recomendacion, nil
}