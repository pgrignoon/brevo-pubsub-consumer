package function

import "cloud.google.com/go/bigquery"

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

type MarketingEmailEvent struct {
	Event        bigquery.NullString     `json:"event"`
	Email        bigquery.NullString     `json:"email"`
	Id           bigquery.NullInt64      `json:"id"`
	DateSent     bigquery.NullString     `json:"date_sent"`
	DateEvent    bigquery.NullString     `json:"date_event"`
	TSSent       bigquery.NullInt64      `json:"ts_sent"`
	TSEvent      bigquery.NullInt64      `json:"ts_event"`
	CampId       bigquery.NullInt64      `json:"camp_id"`
	CampaignName bigquery.NullString     `json:"campaign_name"`
	Reason       bigquery.NullString     `json:"reason"`
	TS           bigquery.NullInt64      `json:"ts"`
	Tag          bigquery.NullString     `json:"tag"`
	SegmentIds   []int                   `json:"segment_ids"`
	Url          bigquery.NullString     `json:"url"`
	SendingIP    bigquery.NullString     `json:"sending_ip"`
	ListId       []int                   `json:"list_id"`
	Key          bigquery.NullString     `json:"key"`
	Date         bigquery.NullString     `json:"date"`
	Content      []MarketingEmailContent `json:"content"`
}

type MarketingEmailContent struct {
	Name      bigquery.NullString `json:"name"`
	LastName  bigquery.NullString `json:"last_name"`
	WorkPhone bigquery.NullString `json:"work_phone"`
}

type MarketingSMSEvent struct {
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
	MessageId        bigquery.NullInt64   `json:"message_id"`
}

type TransactionalSMSEvent struct {
	Id              bigquery.NullInt64          `json:"id"`
	To              bigquery.NullString         `json:"to"`
	SMSCount        bigquery.NullInt64          `json:"sms_count"`
	CreditsUsed     bigquery.NullFloat64        `json:"credits_used"`
	MessageId       bigquery.NullInt64          `json:"message_id"`
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
