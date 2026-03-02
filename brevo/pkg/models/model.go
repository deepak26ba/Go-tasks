package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Connection struct {
	ApiKey string
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05.000Z",
		"2006-01-02 15:04:05",
	}

	var parseErr error

	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			ct.Time = t
			return nil
		}
		parseErr = err
	}

	return parseErr
}

type AllResponse struct {
	Count     int            `json:"count"`
	Templates []ReceivedData `json:"templates"`
}

type ReceivedData struct {
	CreatedAt   CustomTime  `json:"createdAt"`
	HtmlContent string      `json:"htmlContent"`
	Id          int         `json:"id"`
	IsActive    bool        `json:"isActive"`
	ModifiedAt  CustomTime  `json:"modifiedAt"`
	Name        string      `json:"name"`
	ReplyTo     string      `json:"replyTo"`
	Sender      EmailSender `json:"sender"`
	Subject     string      `json:"subject"`
	Tag         string      `json:"tag"`
	TestSent    bool        `json:"testSent"`
	ToField     string      `json:"toField"`
	DoiTemplate bool        `json:"doiTemplate"`
}

type EmailSender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    any    `json:"id"`
}

type EmailTemplate struct {
	gorm.Model
	Subject      string      `json:"subject" gorm:"size:255;not null"`
	TemplateName string      `json:"templateName" gorm:"size:100;uniqueIndex;not null"`
	Sender       EmailSender `json:"sender" `
	HTMLContent  string      `json:"htmlContent" gorm:"type:text"`
}



