package function

import "cloud.google.com/go/bigquery"

type EventDecode struct {
	EventCategory string                  `json:"EventCategory"`
	Data          TransactionalEmailEvent `json:"Data"`
}

type TransactionalEmailEvent struct {
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
	TemplateName  bigquery.NullString `json:"template_name"`
	Tag           bigquery.NullString `json:"tag"`
	Message       bigquery.NullString `json:"message"`
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
