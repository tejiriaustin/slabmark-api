package messaging

import (
	"github.com/mailjet/mailjet-apiv3-go/v4"
)

type MailjetClient struct {
	from   Author
	client *mailjet.Client
}

var _ Messaging = (*MailjetClient)(nil)

func NewMailjetClient(publicKey, privateKey string) *MailjetClient {
	return &MailjetClient{
		client: mailjet.NewMailjetClient(publicKey, privateKey),
	}
}

func (m *MailjetClient) Push(mail Mail) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: m.from.email,
				Name:  m.from.name,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: mail.Email,
					Name:  mail.Name,
				},
			},
			Subject:  mail.Subject,
			TextPart: mail.TextPart,
			HTMLPart: mail.HtmlPart,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := m.client.SendMailV31(&messages)
	if err != nil {
		return nil
	}
	return nil
}
