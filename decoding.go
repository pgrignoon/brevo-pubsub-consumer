package function

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/bigquery"
)

/*
DecodeAndSend decodes the message using the event struct, converts it to a pubsub event struct, and sends it to BigQuery.
*/
func DecodeAndSend[T Event](msg []byte, uploader *bigquery.Uploader, ctx context.Context, datasetId, tableId string) error {
	var data T
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return err
	}
	// Insert data into BigQuery
	if err := uploader.Put(ctx, data.ToBigquery()); err != nil {
		return err
	}
	logger.Info("Successfully sent row to Bigquery", "data", data, "datasetId", datasetId, "tableId", tableId)
	return nil
}
