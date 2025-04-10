
package domain

type INotification interface {
	SendNotification(message string) error
}