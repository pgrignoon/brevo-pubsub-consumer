package function

import "cloud.google.com/go/bigquery"

/*
TransactionalEmailEvent is a struct that represents a transactional email event.
Documentation: https://developers.brevo.com/docs/transactional-webhooks#transactional-email
*/
type TransactionalEmailEvent struct {
	Event         *string   `json:"event"`
	Email         *string   `json:"email"`
	Id            *int64    `json:"id"`
	Date          *string   `json:"date"`
	TS            *int64    `json:"ts"`
	MessageId     *string   `json:"message-id"`
	TSEvent       *int64    `json:"ts_event"`
	Subject       *string   `json:"subject"`
	XMailinCustom *string   `json:"X-Mailin-custom"`
	SendingIP     *string   `json:"sending_ip"`
	TSEpoch       *int64    `json:"ts_epoch"`
	TemplateId    *int64    `json:"template_id"`
	Tag           *string   `json:"tag"`
	Status        *string   `json:"status"`
	Reason        *string   `json:"reason"`
	Tags          *[]string `json:"tags"`
	Link          *string   `json:"link"`
	UserAgent     *string   `json:"user_agent"`
	DeviceUsed    *string   `json:"device_used"`
	MirrorLink    *string   `json:"mirror_link"`
	ContactId     *int64    `json:"contact_id"`
	SenderEmail   *string   `json:"sender_email"`
}

/*
TransactionalEmailEventBigquery is a struct that represents a transactional email event in the bigquery format.
*/
type TransactionalEmailEventBigquery struct {
	Event         bigquery.NullString `json:"event"`
	Email         bigquery.NullString `json:"email"`
	Id            bigquery.NullInt64  `json:"id"`
	Date          bigquery.NullString `json:"date"`
	TS            bigquery.NullInt64  `json:"ts"`
	MessageId     bigquery.NullString `json:"message-id"`
	TSEvent       bigquery.NullInt64  `json:"ts_event"`
	Subject       bigquery.NullString `json:"subject"`
	XMailinCustom bigquery.NullString `json:"X-Mailin-custom"`
	SendingIP     bigquery.NullString `json:"sending_ip"`
	TSEpoch       bigquery.NullInt64  `json:"ts_epoch"`
	TemplateId    bigquery.NullInt64  `json:"template_id"`
	Tag           bigquery.NullString `json:"tag"`
	Status        bigquery.NullString `json:"status"`
	Reason        bigquery.NullString `json:"reason"`
	Tags          []string            `json:"tags"`
	Link          bigquery.NullString `json:"link"`
	UserAgent     bigquery.NullString `json:"user_agent"`
	DeviceUsed    bigquery.NullString `json:"device_used"`
	MirrorLink    bigquery.NullString `json:"mirror_link"`
	ContactId     bigquery.NullInt64  `json:"contact_id"`
	SenderEmail   bigquery.NullString `json:"sender_email"`
}

var TransactionalEmailEventBigqueryDescription = map[string]string{
	"Event":         "Event type",
	"Email":         "Recipient of message",
	"Id":            "Webhook id",
	"Date":          "date sent listed in YEAR-MONTH-DAY, HOUR:MINUTE:SECOND in your timezone",
	"TS":            "Timestamp in seconds of when event occurred",
	"TSEvent":       "It is the time at which the callback is sent to client in Unix format",
	"Subject":       "Message subject",
	"XMailinCustom": "Custom added header",
	"Status":        "Status of the event",
	"MessageId":     "Internal message id",
	"SendingIP":     "IP used to send message",
	"TSEpoch":       "Time stamp in milliseconds UTC of when message was sent",
	"TemplateId":    "Internal id of the template",
	"Tag":           "Tag you might have used to identify your message",
	"Reason":        "Reason for the event",
	"Tags":          "Tags you might have used to identify your message",
	"Link":          "URL accessed by recipient",
	"UserAgent":     "Information about the user’s browser or email client",
	"DeviceUsed":    "Details about the device from which the action originated",
	"MirrorLink":    "A preview link corresponding to the UI logs for the event",
	"ContactId":     "Brevo identifier for an existing contact. If contact is not present, return 0",
	"SenderEmail":   "Email address of the sender",
}

func (t TransactionalEmailEvent) ToBigquery() any {
	var tags []string
	if t.Tags != nil {
		tags = *t.Tags
	}
	return TransactionalEmailEventBigquery{
		Event:         toNullString(t.Event),
		Email:         toNullString(t.Email),
		Id:            toNullInt64(t.Id),
		Date:          toNullString(t.Date),
		TS:            toNullInt64(t.TS),
		MessageId:     toNullString(t.MessageId),
		TSEvent:       toNullInt64(t.TSEvent),
		Subject:       toNullString(t.Subject),
		XMailinCustom: toNullString(t.XMailinCustom),
		SendingIP:     toNullString(t.SendingIP),
		TSEpoch:       toNullInt64(t.TSEpoch),
		TemplateId:    toNullInt64(t.TemplateId),
		Tag:           toNullString(t.Tag),
		Status:        toNullString(t.Status),
		Reason:        toNullString(t.Reason),
		Tags:          tags,
		Link:          toNullString(t.Link),
		UserAgent:     toNullString(t.UserAgent),
		DeviceUsed:    toNullString(t.DeviceUsed),
		MirrorLink:    toNullString(t.MirrorLink),
		ContactId:     toNullInt64(t.ContactId),
		SenderEmail:   toNullString(t.SenderEmail),
	}
}
