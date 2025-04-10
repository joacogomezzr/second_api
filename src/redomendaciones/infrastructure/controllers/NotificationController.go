// infrastructure/controllers/NotificationController.go
package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"newapi/src/redomendaciones/application"
)

type NotificationController struct {
	UseCase *application.NotificationUseCase
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// WebSocketHandler establece una conexión WebSocket para un ID de libro
func (c *NotificationController) WebSocketHandler(ctx *gin.Context) {
	// Obtener el ID del libro desde los parámetros de la URL
	// Actualizar la conexión a WebSocket
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("Error al actualizar la conexión a WebSocket: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la conexión"})
		return
	}

	// Registrar la conexión en el caso de uso
	err = c.UseCase.RegisterConnection(conn)
	if err != nil {
		log.Printf("Error al registrar la conexión: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar la conexión"})
		return
	}

	log.Printf("Conexión WebSocket establecida")
}