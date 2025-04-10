package repositories

import "newapi/src/redomendaciones/domain/entities"

type NotificationRepository interface {
	RegisterConnection(conn *entities.Connection) error
	SendNotification(message string) error
}
