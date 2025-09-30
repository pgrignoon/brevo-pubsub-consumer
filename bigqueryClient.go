package function

import (
	"context"
	"fmt"
	"slices"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

/*
BqContext struct to hold the BigQuery client and configuration (projectId, datasetId)
*/
type BqContext struct {
	Ctx       context.Context
	Client    *bigquery.Client
	ProjectId string `json:"projectId"`
	DatasetId string `json:"datasetId"`
	Uploaders map[string]*bigquery.Uploader
}

/*
Initialise the BigQuery client to perform BigQuery operations
*/
func (bqContext *BqContext) InitBigqueryClient(projectId, datasetId string) error {
	var err error
	bqContext.Ctx = context.Background()
	bqContext.DatasetId = datasetId
	bqContext.ProjectId = projectId
	bqContext.Client, err = bigquery.NewClient(bqContext.Ctx, bqContext.ProjectId)
	if err != nil {
		return err
	}
	bqContext.Uploaders = make(map[string]*bigquery.Uploader)
	return nil
}

/*
Create the bigquery table and uploader if it doesn't exist, by listing the existing tables in the dataset
*/
func (bqContext *BqContext) CreateTableAndUploader(tableId string) error {
	bqTable := bqContext.Client.Dataset(bqContext.DatasetId).Table(tableId)
	tables, err := bqContext.ListTables()
	if err != nil {
		return err
	}
	if !slices.Contains(tables, tableId) {
		logger.Info("Creating bigquery table", "table", bqTable)
		var schema bigquery.Schema
		switch tableId {
		case "transactional-email":
			schema, err = GenerateTableSchema(TransactionalEmailEvent{})
		case "marketing-email":
			schema, err = GenerateTableSchema(MarketingEmailEvent{})
		case "marketing-sms":
			schema, err = GenerateTableSchema(MarketingSMSEvent{})
		case "transactional-sms":
			schema, err = GenerateTableSchema(TransactionalSMSEvent{})
		default:
			return fmt.Errorf("table %s not found", tableId)
		}
		if err != nil {
			return err
		}
		bqTable := bqContext.Client.Dataset(bqContext.DatasetId).Table(tableId)
		err = bqTable.Create(bqContext.Ctx, &bigquery.TableMetadata{Schema: schema})
		if err != nil {
			return err
		}
	} else {
		logger.Info(fmt.Sprintf("Bigquery table %v already exists", "brevo-test-consumer"), "table", bqTable)
	}
	bqContext.Uploaders[tableId] = bqTable.Uploader()
	return nil
}

/*
List all the tables in the dataset from the BqContext object
*/
func (bqContext *BqContext) ListTables() ([]string, error) {
	table := make([]string, 1)
	ts := bqContext.Client.Dataset(bqContext.DatasetId).Tables(bqContext.Ctx)
	for {
		t, err := ts.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return table, err
		}
		table = append(table, t.TableID)
	}
	return table, nil
}

/*
Generate the BigQuery schema for the table, and remove the required constraint for all fields
*/
func GenerateTableSchema(model any) (bigquery.Schema, error) {
	schema, err := bigquery.InferSchema(model)
	if err != nil {
		return schema, err
	}
	for i := range schema {
		schema[i].Required = false
	}
	return schema, nil
}
