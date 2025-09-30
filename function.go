package function

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"slices"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

var bqContext BqContext
var u *bigquery.Uploader
var logger *slog.Logger

func init() {
	gcpProjectID := os.Getenv("GCP_PROJECT_ID")
	datasetId := os.Getenv("BIGQUERY_DATASET_ID")
	bqContext.InitBigqueryClient(gcpProjectID, datasetId)

	// Create bigquery table if it doesn't exist
	bqTable := bqContext.Client.Dataset(datasetId).Table("brevo-test-consumer")
	tables, err := bqContext.ListTables()
	if err != nil {
		logger.Error("Error listing bigquery tables", "error", err)
		return
	}
	if !slices.Contains(tables, "brevo-test-consumer") {
		logger.Info("Creating bigquery table", "table", bqTable)
		schema, err := GenerateTableSchema(TransactionalEmailEvent{})
		if err != nil {
			logger.Error("Error generating bigquery schema", "error", err)
			return
		}
		bqTable := bqContext.Client.Dataset(datasetId).Table("brevo-test-consumer")
		err = bqTable.Create(bqContext.Ctx, &bigquery.TableMetadata{Schema: schema})
		if err != nil {
			logger.Error("Error creating bigquery table", "error", err)
			return
		}
	} else {
		logger.Info(fmt.Sprintf("Bigquery table %v already exists", "brevo-test-consumer"), "table", bqTable)
	}
	u = bqTable.Uploader()
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
	if err := u.Put(bqContext.Ctx, data.Data); err != nil {
		logger.Error("Error inserting data into BigQuery", "error", err)
		return err
	}
	logger.Info("Successfully sent row to Bigquery", "data", data.Data)

	return nil
}
