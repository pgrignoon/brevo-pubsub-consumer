package function

import (
	"encoding/json"

	"cloud.google.com/go/bigquery"
)

func DecodeAndSend[T any](msg []byte, uploader *bigquery.Uploader) error {
	var data T
	err := json.Unmarshal(msg, &data)
	if err != nil {
		logger.Error("Error unmarshalling event data", "error", err)
		return err
	}
	// Insert data into BigQuery
	if err := uploader.Put(bqContext.Ctx, data); err != nil {
		logger.Error("Error inserting data into BigQuery", "error", err)
		return err
	}
	logger.Info("Successfully sent row to Bigquery", "data", data)
	return nil
}
