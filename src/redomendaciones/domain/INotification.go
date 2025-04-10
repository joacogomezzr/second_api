
package domain

type INotification interface {
	SendNotification(id int64, message string) error
}