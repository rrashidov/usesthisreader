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

	ok_notification_client := &TestNotificationClient{}

	err_notification_client := &TestNotificationClient{
		err: ErrNotificationError,
	}

	tests := []struct {
		name                       string
		remote_value               string
		remote_err                 error
		local_value                string
		local_err                  error
		notification_client        *TestNotificationClient
		expected_notification_sent bool
		expected_error             error
	}{
		{"Err remote client", latest_string, ErrTestUsesThisClient, latest_string, nil, ok_notification_client, false, ErrTestUsesThisClient},
		{"Err local client", latest_string, nil, latest_string, ErrTestUsesThisClient, ok_notification_client, false, ErrTestUsesThisClient},
		{"No change", latest_string, nil, latest_string, nil, ok_notification_client, false, nil},
		{"New article", latest_string, nil, updated_string, nil, ok_notification_client, true, nil},
		{"Error sending notification", latest_string, nil, updated_string, nil, err_notification_client, true, ErrNotificationError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			remote_client := &TestUsesThisClient{
				latest: tt.remote_value,
				err:    tt.remote_err,
			}

			local_client := &TestUsesThisClient{
				latest: tt.local_value,
				err:    tt.local_err,
			}

			r := &SimpleUsesThisReader{
				remote: remote_client,
				local:  local_client,
				notif:  tt.notification_client,
			}

			err := r.Run()

			if err != tt.expected_error {
				t.Errorf("Reader did not return expected erro; expected: %q, got: %q", tt.expected_error, err)
			}

			if tt.expected_notification_sent != tt.notification_client.sent {
				t.Errorf("Reader did not do proper notification; expected to be sent: %v, got: %v", tt.expected_notification_sent, tt.notification_client.sent)
			}

			if tt.expected_notification_sent {
				remote_latest, _ := remote_client.GetLatest()
				local_latest, _ := local_client.GetLatest()

				if remote_latest != local_latest {
					t.Errorf("Latest value not updated in local; expected: %q, got: %q", remote_latest, local_latest)
					return
				}
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
