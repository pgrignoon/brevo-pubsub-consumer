package function

import "cloud.google.com/go/bigquery"

/*
MarketingSMSEvent is a struct that represents a marketing SMS event.
Documentation: https://developers.brevo.com/docs/marketing-webhooks#marketing-sms
*/
type MarketingSMSEvent struct {
	Id               *int64    `json:"id"`
	To               *string   `json:"to"`
	SMSCount         *int64    `json:"sms_count"`
	CreditsUsed      *float64  `json:"credits_used"`
	RemainingCredits *float64  `json:"remaining_credits"`
	MsgStatus        *string   `json:"msg_status"`
	Date             *string   `json:"date"`
	Type             *string   `json:"type"`
	CampaignId       *int64    `json:"campaign_id"`
	Status           *string   `json:"status"`
	Description      *string   `json:"description"`
	TSEvent          *int64    `json:"ts_event"`
	Tag              *[]string `json:"tag"`
	ErrorCode        *int64    `json:"error_code"`
	Reply            *string   `json:"reply"`
	BounceType       *string   `json:"bounce_type"`
	MessageId        *int64    `json:"message_id"`
}

/*
MarketingSMSEventBigquery is a struct that represents a marketing SMS event in the bigquery format.
*/
type MarketingSMSEventBigquery struct {
	Id               bigquery.NullInt64   `json:"id"`
	To               bigquery.NullString  `json:"to"`
	SMSCount         bigquery.NullInt64   `json:"sms_count"`
	CreditsUsed      bigquery.NullFloat64 `json:"credits_used"`
	RemainingCredits bigquery.NullFloat64 `json:"remaining_credits"`
	MsgStatus        bigquery.NullString  `json:"msg_status"`
	Date             bigquery.NullString  `json:"date"`
	Type             bigquery.NullString  `json:"type"`
	CampaignId       bigquery.NullInt64   `json:"campaign_id"`
	Status           bigquery.NullString  `json:"status"`
	Description      bigquery.NullString  `json:"description"`
	TSEvent          bigquery.NullInt64   `json:"ts_event"`
	Tag              []string             `json:"tag"`
	ErrorCode        bigquery.NullInt64   `json:"error_code"`
	Reply            bigquery.NullString  `json:"reply"`
	BounceType       bigquery.NullString  `json:"bounce_type"`
	MessageId        bigquery.NullInt64   `json:"messageId"`
}

var MarketingSMSEventBigqueryDescription = map[string]string{
	"Id":               "Unique id generated for each payload",
	"To":               "Mobile number",
	"SMSCount":         "Number of SMS sent",
	"CreditsUsed":      "Credits deducted",
	"RemainingCredits": "Remaining balance credit",
	"MsgStatus":        "Status of the message (sent, delivered, soft_bounce, hard_bounce)",
	"Date":             "Time at which the event is generated",
	"Type":             "Type of sms(marketing/transactional)",
	"CampaignId":       "Campaign id for campaign sms",
	"Status":           "Status of the event",
	"Description":      "Bounce Reason for Failed to deliver message, description of the event",
	"TSEvent":          "Timestamp in seconds of when event occurred",
	"Tag":              "Internal tag of campaign",
	"ErrorCode":        "Error code",
	"Reply":            "Reply to the message",
	"BounceType":       "Bounce type",
	"MessageId":        "Internal id of message",
}

func (m MarketingSMSEvent) ToBigquery() any {
	var tags []string
	if m.Tag != nil {
		tags = *m.Tag
	}
	return MarketingSMSEventBigquery{
		Id:               toNullInt64(m.Id),
		To:               toNullString(m.To),
		SMSCount:         toNullInt64(m.SMSCount),
		CreditsUsed:      toNullFloat64(m.CreditsUsed),
		RemainingCredits: toNullFloat64(m.RemainingCredits),
		MsgStatus:        toNullString(m.MsgStatus),
		Date:             toNullString(m.Date),
		Type:             toNullString(m.Type),
		CampaignId:       toNullInt64(m.CampaignId),
		Status:           toNullString(m.Status),
		Description:      toNullString(m.Description),
		TSEvent:          toNullInt64(m.TSEvent),
		Tag:              tags,
		ErrorCode:        toNullInt64(m.ErrorCode),
		Reply:            toNullString(m.Reply),
		BounceType:       toNullString(m.BounceType),
		MessageId:        toNullInt64(m.MessageId),
	}
}
