package function

import "cloud.google.com/go/bigquery"

/*
MarketingEmailEvent is a struct that represents a marketing email event.
Documentation: https://developers.brevo.com/docs/marketing-webhooks#marketing-email
*/
type MarketingEmailEvent struct {
	Event        *string                  `json:"event"`
	Email        *string                  `json:"email"`
	Id           *int64                   `json:"id"`
	DateSent     *string                  `json:"date_sent"`
	DateEvent    *string                  `json:"date_event"`
	TSSent       *int64                   `json:"ts_sent"`
	TSEvent      *int64                   `json:"ts_event"`
	CampId       *int64                   `json:"camp_id"`
	CampaignName *string                  `json:"campaign_name"`
	Reason       *string                  `json:"reason"`
	TS           *int64                   `json:"ts"`
	Tag          *string                  `json:"tag"`
	SegmentIds   *[]int64                 `json:"segment_ids"`
	Url          *string                  `json:"url"`
	SendingIP    *string                  `json:"sending_ip"`
	ListId       *[]int64                 `json:"list_id"`
	Key          *string                  `json:"key"`
	Date         *string                  `json:"date"`
	Content      *[]MarketingEmailContent `json:"content"`
}

type MarketingEmailContent struct {
	Name      *string `json:"name"`
	LastName  *string `json:"last_name"`
	WorkPhone *string `json:"work_phone"`
}

/*
MarketingEmailEventBigquery is a struct that represents a marketing email event in the bigquery format.
*/
type MarketingEmailEventBigquery struct {
	Event        bigquery.NullString             `json:"event"`
	Email        bigquery.NullString             `json:"email"`
	Id           bigquery.NullInt64              `json:"id"`
	DateSent     bigquery.NullString             `json:"date_sent"`
	DateEvent    bigquery.NullString             `json:"date_event"`
	TSSent       bigquery.NullInt64              `json:"ts_sent"`
	TSEvent      bigquery.NullInt64              `json:"ts_event"`
	CampId       bigquery.NullInt64              `json:"camp_id"`
	CampaignName bigquery.NullString             `json:"campaign_name"`
	Reason       bigquery.NullString             `json:"reason"`
	TS           bigquery.NullInt64              `json:"ts"`
	Tag          bigquery.NullString             `json:"tag"`
	SegmentIds   []int64                         `json:"segment_ids"`
	Url          bigquery.NullString             `json:"url"`
	SendingIP    bigquery.NullString             `json:"sending_ip"`
	ListId       []int64                         `json:"list_id"`
	Key          bigquery.NullString             `json:"key"`
	Date         bigquery.NullString             `json:"date"`
	Content      []MarketingEmailContentBigquery `json:"content"`
}

type MarketingEmailContentBigquery struct {
	Name      bigquery.NullString `json:"name"`
	LastName  bigquery.NullString `json:"last_name"`
	WorkPhone bigquery.NullString `json:"work_phone"`
}

var MarketingEmailEventBigqueryDescription = map[string]string{
	"Event":        "The event type",
	"Email":        "Recipient email",
	"Id":           "Internal id of webhook",
	"DateSent":     "Date the campaign was sent (year-month-day, hour:minute:second)",
	"DateEvent":    "Date the event occurred (year-month-day, hour:minute:second)",
	"TSSent":       "Timestamp in seconds of when campaign was sent",
	"TSEvent":      "Timestamp in seconds of when event occurred",
	"CampId":       "Internal id of campaign",
	"CampaignName": "Internal name of campaign",
	"Reason":       "The reason the event occurred",
	"TS":           "Timestamp in seconds of when event occurred",
	"Tag":          "Internal tag of campaign",
	"SegmentIds":   "Newly added fields for mails that are sent to a segment",
	"Url":          "URL clicked",
	"SendingIP":    "IP used to send message",
	"ListId":       "Array of ids",
	"Key":          "Internal Key",
	"Date":         "Date the event occurred (year-month-day, hour:minute:second)",
	"Content":      "Full contact information with updates",
}

func (m MarketingEmailEvent) ToBigquery() any {
	var segmentIds []int64
	if m.SegmentIds != nil {
		segmentIds = *m.SegmentIds
	}
	var listId []int64
	if m.ListId != nil {
		listId = *m.ListId
	}
	var content []MarketingEmailContentBigquery
	if m.Content != nil {
		for _, c := range *m.Content {
			content = append(content, MarketingEmailContentBigquery{
				Name:      toNullString(c.Name),
				LastName:  toNullString(c.LastName),
				WorkPhone: toNullString(c.WorkPhone),
			})
		}
	}
	return MarketingEmailEventBigquery{
		Event:        toNullString(m.Event),
		Email:        toNullString(m.Email),
		Id:           toNullInt64(m.Id),
		DateSent:     toNullString(m.DateSent),
		DateEvent:    toNullString(m.DateEvent),
		TSSent:       toNullInt64(m.TSSent),
		TSEvent:      toNullInt64(m.TSEvent),
		CampId:       toNullInt64(m.CampId),
		CampaignName: toNullString(m.CampaignName),
		Reason:       toNullString(m.Reason),
		TS:           toNullInt64(m.TS),
		Tag:          toNullString(m.Tag),
		SegmentIds:   segmentIds,
		Url:          toNullString(m.Url),
		SendingIP:    toNullString(m.SendingIP),
		ListId:       listId,
		Key:          toNullString(m.Key),
		Date:         toNullString(m.Date),
		Content:      content,
	}
}
