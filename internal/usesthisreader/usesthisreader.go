package usesthisreader

import (
	"fmt"
	"rrashidov/usesthisreader/internal/notification"
)

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
		fmt.Printf("Problem retrieving remote latest interview: %q\n", err.Error())

		return err
	}

	fmt.Printf("Latest remote interview: %q\n", latest_remote)

	latest_local, err := r.local.GetLatest()

	if err != nil {
		fmt.Printf("Problem retrieving local latest interview: %q\n", latest_local)

		return err
	}

	fmt.Printf("Latest local interview: %q\n", latest_local)

	if latest_local != latest_remote {
		fmt.Println("New remote interview found. Notify")

		r.local.Update(latest_remote)

		return r.notif.Notify()
	} else {
		fmt.Println("No new interview found")
	}

	return nil
}
