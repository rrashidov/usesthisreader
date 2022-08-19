package notification

type NotificationClient interface {
	Notify() error
}
