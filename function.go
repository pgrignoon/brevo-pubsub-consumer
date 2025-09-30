package function

import (
	"context"
	"encoding/json"
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
	bqContext.CreateTableAndUploader("brevo-test-consumer")
	functions.CloudEvent("HelloPubSub", helloPubSub)
}

type MessagePublishedData struct {
	Message PubSubMessage
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

// helloPubSub consumes a CloudEvent message and extracts the Pub/Sub message.
func helloPubSub(ctx context.Context, e event.Event) error {
	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	var data EventDecode
	err := json.Unmarshal(msg.Message.Data, &data)
	if err != nil {
		logger.Error("Error unmarshalling event data", "error", err)
		return err
	}

	// Insert data into BigQuery
	if err := bqContext.Uploaders["brevo-test-consumer"].Put(bqContext.Ctx, data.Data); err != nil {
		logger.Error("Error inserting data into BigQuery", "error", err)
		return err
	}
	logger.Info("Successfully sent row to Bigquery", "data", data.Data)

	return nil
}
