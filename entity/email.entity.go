package entity

const (
	// EmailCategorySendOTP define email category
	EmailCategorySendOTP = "send-otp"
)

// EmailPayload is the payload for sending email
type EmailPayload struct {
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

// NewEmailPayload is the constructor for EmailPayload
func NewEmailPayload(to, subject, content, category string) *EmailPayload {
	return &EmailPayload{
		To:       to,
		Subject:  subject,
		Content:  content,
		Category: category,
	}
}
