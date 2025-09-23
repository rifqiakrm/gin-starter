// Package helper provides utilities
package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"gin-starter/common/errors"
)

// EmailPayload is the payload for sending email
type EmailPayload struct {
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

// ConstructEmailPayload is a function to construct an email payload
func ConstructEmailPayload(templatePath, receiver, subject, category string, data map[string]interface{}) (*EmailPayload, error) {
	var body bytes.Buffer

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(fmt.Errorf("failed to load email template: %w", err))
		return nil, errors.ErrInternalServerError.Error()
	}

	err = t.Execute(&body, data)
	if err != nil {
		log.Println(fmt.Errorf("failed to execute email data: %w", err))
		return nil, errors.ErrInternalServerError.Error()
	}

	emailPayload := &EmailPayload{
		To:       receiver,
		Subject:  subject,
		Content:  body.String(),
		Category: category,
	}

	return emailPayload, nil
}
