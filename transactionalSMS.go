package function

import "cloud.google.com/go/bigquery"

/*
TransactionalSMSEvent is a struct that represents a transactional SMS event.
Documentation: https://developers.brevo.com/docs/transactional-webhooks#transactional-sms
*/
type TransactionalSMSEvent struct {
	Id              *int64             `json:"id"`
	To              *string            `json:"to"`
	SMSCount        *int64             `json:"sms_count"`
	CreditsUsed     *float64           `json:"credits_used"`
	MessageId       *int64             `json:"message_id"`
	RemainingCredit *float64           `json:"remaining_credit"`
	MsgStatus       *string            `json:"msg_status"`
	Date            *string            `json:"date"`
	Type            *string            `json:"type"`
	Reference       *map[string]string `json:"reference"`
	Status          *string            `json:"status"`
	Description     *string            `json:"description"`
	TSEvent         *int64             `json:"ts_event"`
	Tag             *[]string          `json:"tag"`
	ErrorCode       *int64             `json:"error_code"`
	Reply           *string            `json:"reply"`
	BounceType      *string            `json:"bounce_type"`
}

/*
TransactionalSMSEventBigquery is a struct that represents a transactional SMS event in the bigquery format.
*/
type TransactionalSMSEventBigquery struct {
	Id              bigquery.NullInt64          `json:"id"`
	To              bigquery.NullString         `json:"to"`
	SMSCount        bigquery.NullInt64          `json:"sms_count"`
	CreditsUsed     bigquery.NullFloat64        `json:"credits_used"`
	MessageId       bigquery.NullInt64          `json:"messageId"`
	RemainingCredit bigquery.NullFloat64        `json:"remaining_credit"`
	MsgStatus       bigquery.NullString         `json:"msg_status"`
	Date            bigquery.NullString         `json:"date"`
	Type            bigquery.NullString         `json:"type"`
	Reference       []TransactionalSMSReference `json:"reference"`
	Status          bigquery.NullString         `json:"status"`
	Description     bigquery.NullString         `json:"description"`
	TSEvent         bigquery.NullInt64          `json:"ts_event"`
	Tag             []string                    `json:"tag"`
	ErrorCode       bigquery.NullInt64          `json:"error_code"`
	Reply           bigquery.NullString         `json:"reply"`
	BounceType      bigquery.NullString         `json:"bounce_type"`
}

type TransactionalSMSReference struct {
	Key   bigquery.NullString `json:"key"`
	Value bigquery.NullString `json:"value"`
}

var TransactionalSMSEventBigqueryDescription = map[string]string{
	"Id":              "The id is the webhook ID, so it will remain the same for webhook payloads sent to a single webhook URL",
	"To":              "Mobile number of the recipient",
	"SMSCount":        "Number of SMS sent",
	"CreditsUsed":     "Credits deducted",
	"MessageId":       "Message id for Transactional SMS",
	"RemainingCredit": "Remaining balance credit",
	"MsgStatus":       "Status of the message (sent, delivered, soft_bounce, hard_bounce)",
	"Date":            "Time at which the event is generated",
	"Type":            "Type of sms(marketing/transactional)",
	"Reference":       "Id generated for every Transactional SMS",
	"Status":          "Status of the event",
	"Description":     "Bounce Reason for Failed to deliver message, or description of the event",
	"TSEvent":         "It is the time at which the callback is sent to client in Unix format",
	"Tag":             "SMS tag if the client has used any",
	"ErrorCode":       "Error code",
	"Reply":           "Reply to the message",
	"BounceType":      "Bounce type",
}

func (t TransactionalSMSEvent) ToBigquery() any {
	reference := []TransactionalSMSReference{}
	if t.Reference != nil {
		for k, v := range *t.Reference {
			reference = append(reference, TransactionalSMSReference{Key: bigquery.NullString{StringVal: k, Valid: true}, Value: bigquery.NullString{StringVal: v, Valid: true}})
		}
	}
	var tags []string
	if t.Tag != nil {
		tags = *t.Tag
	}
	return TransactionalSMSEventBigquery{
		Id:              toNullInt64(t.Id),
		To:              toNullString(t.To),
		SMSCount:        toNullInt64(t.SMSCount),
		CreditsUsed:     toNullFloat64(t.CreditsUsed),
		MessageId:       toNullInt64(t.MessageId),
		RemainingCredit: toNullFloat64(t.RemainingCredit),
		MsgStatus:       toNullString(t.MsgStatus),
		Date:            toNullString(t.Date),
		Type:            toNullString(t.Type),
		Reference:       reference,
		Status:          toNullString(t.Status),
		Description:     toNullString(t.Description),
		TSEvent:         toNullInt64(t.TSEvent),
		Tag:             tags,
		ErrorCode:       toNullInt64(t.ErrorCode),
		Reply:           toNullString(t.Reply),
		BounceType:      toNullString(t.BounceType),
	}
}
