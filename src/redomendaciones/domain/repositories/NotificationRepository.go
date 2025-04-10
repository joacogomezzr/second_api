package repositories

import "newapi/src/redomendaciones/domain/entities"

type NotificationRepository interface {
	RegisterConnection(id int64, conn *entities.Connection) error
	SendNotification(id int64, message string) error
}
