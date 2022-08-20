package usesthisreader

import "rrashidov/usesthisreader/internal/notification"

type UsesThisReader interface {
	Run() error
}

type SimpleUsesThisReader struct {
	remote UsesThisClient
	local  UsesThisClient
	notif  notification.NotificationClient
}

func NewUsesThisReader(url string, filepath string) *SimpleUsesThisReader {
	return nil
}

func (r SimpleUsesThisReader) Run() error {

	latest_remote, err := r.remote.GetLatest()

	if err != nil {
		return err
	}

	latest_local, err := r.local.GetLatest()

	if err != nil {
		return err
	}

	if latest_local != latest_remote {
		r.local.Update(latest_remote)

		return r.notif.Notify()
	}

	return nil
}
