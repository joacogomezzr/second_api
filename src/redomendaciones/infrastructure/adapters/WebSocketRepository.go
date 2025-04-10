// infrastructure/adapters/NotificationRepository.go
package adapters

import (
	"errors"
	"log"
	"sync"
	"newapi/src/redomendaciones/domain/entities"
)

type NotificationRepository struct {
	connections sync.Map 
}

// RegisterConnection guarda una conexión asociada a un ID
func (r *NotificationRepository) RegisterConnection(id int64, connection *entities.Connection) error {
	r.connections.Store(id, connection)
	log.Printf("Conexión registrada para ID: %d", id)
	return nil
}

// SendNotification envía un mensaje a un ID específico
func (r *NotificationRepository) SendNotification(id int64, message string) error {
	value, ok := r.connections.Load(id)
	if !ok {
		return errors.New("no se encontró una conexión para el ID especificado")
	}

	conn := value.(*entities.Connection).Conn 
	err := conn.WriteMessage(1, []byte(message)) 
	if err != nil {
		return err
	}

	log.Printf("Notificación enviada al ID: %d", id)
	return nil
}