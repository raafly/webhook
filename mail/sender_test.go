package mail

import (
	"testing"

	"github.com/raaafly/powerup-client-service-golang/config"
)

func TestXSender(t *testing.T) {
	email := "lvidiamine@gmail.com"

	sender := NewGmailSender(
		config.NewAppConfig().Email.Sender, 
		config.NewAppConfig().Email.Adderss, 
		config.NewAppConfig().Email.Password,
	)
	to := []string{email}
	
	sender.SendEmail(to, nil, nil)
}