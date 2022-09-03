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

func NewUsesThisReader(url string, filepath string, region string, recipient string, sender string) *SimpleUsesThisReader {
	remote := &RemoteUsesThisClient{
		url: url,
	}

	local := &LocalUsesThisClient{
		filepath: filepath,
	}

	notif := &notification.AWSSESNotificationClient{
		Region:    region,
		Recipient: recipient,
		Sender:    sender,
	}

	return &SimpleUsesThisReader{
		remote: remote,
		local:  local,
		notif:  notif,
	}
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
