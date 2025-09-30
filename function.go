package function

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

var bqContext BqContext
var u *bigquery.Uploader

func init() {
	gcpProjectID := os.Getenv("GCP_PROJECT_ID")
	datasetId := os.Getenv("BIGQUERY_DATASET_ID")
	bqContext.InitBigqueryClient(gcpProjectID, datasetId)

	// Create bigquery table if it doesn't exist
	bqTable := bqContext.Client.Dataset(datasetId).Table("brevo-test-consumer")
	tables, err := bqContext.ListTables()
	if err != nil {
		log.Printf("Error listing bigquery tables: %v", err)
		return
	}
	log.Printf("Tables: %v", tables)
	if !slices.Contains(tables, "brevo-test-consumer") {
		log.Printf("Creating bigquery table: %v", bqTable)
		schema, err := GenerateTableSchema(TransactionalEmailEvent{})
		if err != nil {
			log.Printf("Error generating bigquery schema: %v", err)
			return
		}
		bqTable := bqContext.Client.Dataset(datasetId).Table("brevo-test-consumer")
		err = bqTable.Create(bqContext.Ctx, &bigquery.TableMetadata{Schema: schema})
		if err != nil {
			log.Printf("Error creating bigquery table: %v", err)
			return
		}
	} else {
		log.Printf("Bigquery table already exists: %v", bqTable)
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

	name := string(msg.Message.Data) // Automatically decoded from base64.
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)

	var data EventDecode
	err := json.Unmarshal(msg.Message.Data, &data)
	if err != nil {
		log.Printf("Errorrrrr: %v\n", err)
		return err
	}

	log.Printf("Event Category: %s", data.EventCategory)
	log.Printf("Event Data: %v", data.Data)

	// Insert data into BigQuery
	if err := u.Put(bqContext.Ctx, data.Data); err != nil {
		log.Printf("Error inserting data into BigQuery: %v\n", err)
		return err
	}
	log.Printf("Successfully sent row to Bigquery: %v\n", data.Data)

	return nil
}
