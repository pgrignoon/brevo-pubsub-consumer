package function

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

var bqContext BqContext
var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	gcpProjectID := os.Getenv("GCP_PROJECT_ID")
	datasetId := os.Getenv("BIGQUERY_DATASET_ID")
	bqContext.InitBigqueryClient(gcpProjectID, datasetId)
	for _, tableId := range []string{"transactional-email", "marketing-email", "marketing-sms", "transactional-sms"} {
		bqContext.CreateTableAndUploader(tableId)
	}
	functions.CloudEvent("RunPubSubConsumer", runPubSubConsumer)
}

type MessagePublishedData struct {
	Message PubSubMessage
}

type PubSubMessage struct {
	Data       []byte            `json:"data"`
	Attributes map[string]string `json:"attributes"`
}

// runPubSubConsumer consumes a CloudEvent message and extracts the Pub/Sub message.
func runPubSubConsumer(ctx context.Context, e event.Event) error {
	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	category := msg.Message.Attributes["category"]
	switch category {
	case "transactional-email":
		DecodeAndSend[TransactionalEmailEvent](msg.Message.Data, bqContext.Uploaders[category])
	case "marketing-email":
		DecodeAndSend[MarketingEmailEvent](msg.Message.Data, bqContext.Uploaders[category])
	case "marketing-sms":
		DecodeAndSend[MarketingSMSEvent](msg.Message.Data, bqContext.Uploaders[category])
	case "transactional-sms":
		DecodeAndSend[TransactionalSMSEvent](msg.Message.Data, bqContext.Uploaders[category])
	default:
		return fmt.Errorf("invalid category: ##%s##", category)
	}
	return nil
}
