// domain/entities/Connection.go
package entities

import "github.com/gorilla/websocket"

type Connection struct {
	ID   int64
	Conn *websocket.Conn
}