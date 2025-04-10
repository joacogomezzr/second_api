package adapters

import (
	"log"
	"newapi/src/redomendaciones/domain/entities"
)

type NotificationRepository struct {
	connections []*entities.Connection 
}

// RegisterConnection guarda una conexión en la lista
func (r *NotificationRepository) RegisterConnection(connection *entities.Connection) error {
	r.connections = append(r.connections, connection) 
	log.Println("Conexión registrada")
	return nil
}

// SendNotification envía un mensaje a todos los usuarios conectados
func (r *NotificationRepository) SendNotification(message string) error {
	for _, connection := range r.connections {
		err := connection.Conn.WriteMessage(1, []byte(message)) // Enviar mensaje
		if err != nil {
			log.Printf("Error al enviar notificación: %v", err)
		} else {
			log.Println("Notificación enviada")
		}
	}
	return nil
}