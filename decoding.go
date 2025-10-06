package function

import (
	"encoding/json"

	"cloud.google.com/go/bigquery"
)

/*
DecodeAndSend decodes the message using the event struct, converts it to a pubsub event struct, and sends it to BigQuery.
*/
func DecodeAndSend[T Event](msg []byte, uploader *bigquery.Uploader) error {
	var data T
	err := json.Unmarshal(msg, &data)
	if err != nil {
		logger.Error("Error unmarshalling event data", "error", err)
		return err
	}
	// Insert data into BigQuery
	if err := uploader.Put(bqContext.Ctx, data.ToBigquery()); err != nil {
		logger.Error("Error inserting data into BigQuery", "error", err)
		return err
	}
	logger.Info("Successfully sent row to Bigquery", "data", data)
	return nil
}
