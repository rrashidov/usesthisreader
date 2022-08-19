package usesthisreader

import (
	"errors"
	"testing"
)

var (
	ErrTestUsesThisClient = errors.New("test uses this client error")
	ErrNotificationError  = errors.New("notification sending error")
)

func TestUsesThisReader(t *testing.T) {

	latest_string := "Test Latest String"

	updated_string := "Test Updated String"

	err_client := &TestUsesThisClient{
		latest: latest_string,
		err:    ErrTestUsesThisClient,
	}

	ok_client := &TestUsesThisClient{
		latest: latest_string,
		err:    nil,
	}

	another_ok_client := &TestUsesThisClient{
		latest: updated_string,
		err:    nil,
	}

	ok_notification_client := &TestNotificationClient{}

	err_notification_client := &TestNotificationClient{
		err: ErrNotificationError,
	}

	tests := []struct {
		name                       string
		remote                     UsesThisClient
		local                      UsesThisClient
		notification_client        *TestNotificationClient
		expected_notification_sent bool
		expected_error             error
	}{
		{"Err remote client", err_client, ok_client, ok_notification_client, false, ErrTestUsesThisClient},
		{"Err local client", ok_client, err_client, ok_notification_client, false, ErrTestUsesThisClient},
		{"No change", ok_client, ok_client, ok_notification_client, false, nil},
		{"New article", ok_client, another_ok_client, ok_notification_client, true, nil},
		{"Error sending notification", ok_client, another_ok_client, err_notification_client, true, ErrNotificationError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &SimpleUsesThisReader{
				remote: tt.remote,
				local:  tt.local,
				notif:  tt.notification_client,
			}

			err := r.Run()

			if err != tt.expected_error {
				t.Errorf("Reader did not return expected erro; expected: %q, got: %q", tt.expected_error, err)
			}

			if tt.expected_notification_sent != tt.notification_client.sent {
				t.Errorf("Reader did not do proper notification; expected to be sent: %v, got: %v", tt.expected_notification_sent, tt.notification_client.sent)
			}
		})
	}
}

type TestUsesThisClient struct {
	latest string
	err    error
}

func (c TestUsesThisClient) GetLatest() (string, error) {
	return c.latest, c.err
}

func (c *TestUsesThisClient) Update(s string) error {
	c.latest = s

	return c.err
}

type TestNotificationClient struct {
	sent bool
	err  error
}

func (t *TestNotificationClient) Notify() error {
	t.sent = true

	return t.err
}
