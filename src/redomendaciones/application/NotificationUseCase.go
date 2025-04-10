// application/NotificationUseCase.go
package application

import (
	"newapi/src/redomendaciones/domain/entities"
	"newapi/src/redomendaciones/domain/repositories"

	"github.com/gorilla/websocket"
)

type NotificationUseCase struct {
	Repo repositories.NotificationRepository
}

func (u *NotificationUseCase) RegisterConnection(id int64, conn *websocket.Conn) error {
	connection := &entities.Connection{ID: id, Conn: conn}
	return u.Repo.RegisterConnection(id, connection)
}

// NotifyBook envía una notificación al ID del libro especificado
func (u *NotificationUseCase) NotifyBook(id int64, message string) error {
	return u.Repo.SendNotification(id, message)
}