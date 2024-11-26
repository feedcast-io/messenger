package messenger

import (
	"github.com/feedcast-io/messenger/types"
	"testing"
)

func TestSendEmail(t *testing.T) {
	service := NewService("feedcast-preprod")

	if e := service.SendEmail("romain@orixamedia.com", "FR", types.EmailTemplateForgotPassword, map[string]any{
		"url":      "http://www.google.fr",
		"customer": "TEST CUSTOMER",
	}); e != nil {
		t.Error(e)
	}
}
