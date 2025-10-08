package function

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/bigquery"
)

/*
DecodeAndSend decodes the message using the event struct, converts it to a pubsub event struct, and sends it to BigQuery.
*/
func DecodeAndSend[T Event](msg []byte, uploader *bigquery.Uploader, ctx context.Context) (T, error) {
	var data T
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return data, err
	}
	// Insert data into BigQuery
	if err := uploader.Put(ctx, data.ToBigquery()); err != nil {
		return data, err
	}
	return data, nil
}
