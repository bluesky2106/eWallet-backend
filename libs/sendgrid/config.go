package sendgrid

// Config : sendgrid configurations
type Config struct {
	APIKey      string `json:"apiKey"`
	SenderEmail string `json:"senderEmail"`
	SenderName  string `json:"senderName"`
	CCEmail     string `json:"ccEmail"`
	CCName      string `json:"ccName"`
}
