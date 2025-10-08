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
	err := bqContext.InitBigqueryClient(gcpProjectID)
	if err != nil {
		panic("Error initializing bigquery client: " + err.Error())
	}
	err = bqContext.CreateTablesAndUploaders()
	if err != nil {
		panic("Error creating tables and uploaders: " + err.Error())
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
	// Extract the category, datasetId, and tableId from the attributes
	category, ok := msg.Message.Attributes["category"]
	if !ok {
		return fmt.Errorf("category not found in attributes")
	}
	datasetId, ok := msg.Message.Attributes["target-dataset"]
	if !ok {
		return fmt.Errorf("target-dataset not found in attributes")
	}
	tableId, ok := msg.Message.Attributes["target-table"]
	if !ok {
		return fmt.Errorf("target-table not found in attributes")
	}
	// Switch on the category and send the data to the appropriate table
	switch category {
	case "transactional-email":
		uploader, ok := bqContext.Uploaders[fmt.Sprintf("%s.%s", datasetId, tableId)]
		if !ok {
			return fmt.Errorf("uploader not found for datasetId: %s, tableId: %s", datasetId, tableId)
		}
		err := DecodeAndSend[TransactionalEmailEvent](msg.Message.Data, uploader, bqContext.Ctx, datasetId, tableId)
		if err != nil {
			return fmt.Errorf("error decoding and sending transactional email event to table: %s.%s: %w", datasetId, tableId, err)
		}
	case "marketing-email":
		uploader, ok := bqContext.Uploaders[fmt.Sprintf("%s.%s", datasetId, tableId)]
		if !ok {
			return fmt.Errorf("uploader not found for datasetId: %s, tableId: %s", datasetId, tableId)
		}
		err := DecodeAndSend[MarketingEmailEvent](msg.Message.Data, uploader, bqContext.Ctx, datasetId, tableId)
		if err != nil {
			return fmt.Errorf("error decoding and sending marketing email event to table: %s.%s: %w", datasetId, tableId, err)
		}
	case "marketing-sms":
		uploader, ok := bqContext.Uploaders[fmt.Sprintf("%s.%s", datasetId, tableId)]
		if !ok {
			return fmt.Errorf("uploader not found for datasetId: %s, tableId: %s", datasetId, tableId)
		}
		err := DecodeAndSend[MarketingSMSEvent](msg.Message.Data, uploader, bqContext.Ctx, datasetId, tableId)
		if err != nil {
			return fmt.Errorf("error decoding and sending marketing sms event to table: %s.%s: %w", datasetId, tableId, err)
		}
	case "transactional-sms":
		uploader, ok := bqContext.Uploaders[fmt.Sprintf("%s.%s", datasetId, tableId)]
		if !ok {
			return fmt.Errorf("uploader not found for datasetId: %s, tableId: %s", datasetId, tableId)
		}
		err := DecodeAndSend[TransactionalSMSEvent](msg.Message.Data, uploader, bqContext.Ctx, datasetId, tableId)
		if err != nil {
			return fmt.Errorf("error decoding and sending transactional sms event to table: %s.%s: %w", datasetId, tableId, err)
		}
	default:
		return fmt.Errorf("invalid category: ##%s##", category)
	}
	return nil
}
