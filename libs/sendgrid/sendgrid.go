package sendgrid

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// MailManager : mail manager
type MailManager struct {
	conf *Config
}

// NewMailManager : init MailManager
func NewMailManager(conf *config.Config) *MailManager {
	return &MailManager{
		conf: &Config{
			APIKey:      conf.Sendgrid.APIKey,
			SenderEmail: conf.Sendgrid.SenderEmail,
			SenderName:  conf.Sendgrid.SenderName,
			CCEmail:     conf.Sendgrid.CCEmail,
			CCName:      conf.Sendgrid.CCName,
		},
	}
}

// SendEmail : email info
func (mm *MailManager) SendEmail(data *pb.EmailInfo) error {
	m := mail.NewV3Mail()
	p := mail.NewPersonalization()

	// 1. set from email
	mm.setFromEmail(m, data)
	// 2. set to email
	mm.setToEmails(p, data)

	// 3. set template
	m.SetTemplateID(data.GetTemplateId())

	// 4. set body for template
	for k, v := range data.GetData() {
		p.SetDynamicTemplateData(k, v)
	}

	m.AddPersonalizations(p)

	// 5. call sendgrid
	request := sendgrid.GetRequest(mm.conf.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (mm *MailManager) setFromEmail(m *mail.SGMailV3, data *pb.EmailInfo) {
	var (
		fromEmail *mail.Email
	)
	if len(data.GetSenderEmail()) == 0 {
		fromEmail = mail.NewEmail(mm.conf.SenderName, mm.conf.SenderEmail)
	} else {
		fromEmail = mail.NewEmail(data.GetSenderName(), data.GetSenderEmail())
	}
	m.SetFrom(fromEmail)
}

func (mm *MailManager) setToEmails(p *mail.Personalization, data *pb.EmailInfo) {
	arrTo := data.GetReceivers()
	for _, receiver := range arrTo {
		toEmail := mail.NewEmail(receiver.GetToName(), receiver.GetToEmail())
		p.AddTos(toEmail)
	}
}
