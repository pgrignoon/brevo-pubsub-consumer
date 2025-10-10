package function

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	ProjectId string  `json:"projectId"`
	Tables    []Table `json:"tables"`
	Uploaders map[string]*bigquery.Uploader
}

type Table struct {
	Source        string `json:"source"`
	DatasetId     string `json:"datasetId"`
	TableId       string `json:"tableId"`
	EventCategory string `json:"eventCategory"`
}

/*
Initialise the BigQuery client to perform BigQuery operations
*/
func (bqContext *BqContext) InitBigqueryClient(projectId, configFilePath string) error {
	var err error
	err = bqContext.LoadTablesFromConfig(configFilePath)
	if err != nil {
		return err
	}
	bqContext.Ctx = context.Background()
	bqContext.ProjectId = projectId
	bqContext.Client, err = bigquery.NewClient(bqContext.Ctx, bqContext.ProjectId)
	if err != nil {
		return err
	}
	bqContext.Uploaders = make(map[string]*bigquery.Uploader)
	logger.Info("Tables loaded from config.json", "tables", bqContext.Tables)
	return nil
}

/*
Load the tables from the config.json file
*/
func (bqContext *BqContext) LoadTablesFromConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open configuration file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %v", err)
	}
	err = json.Unmarshal(bytes, &bqContext)
	if err != nil {
		return fmt.Errorf("failed to parse configuration file: %v", err)
	}
	return nil
}

/*
Create the bigquery table and uploader if it doesn't exist, by listing the existing tables in the dataset
*/
func (bqContext *BqContext) CreateTablesAndUploaders() error {
	for _, table := range bqContext.Tables {
		bqTable := bqContext.Client.Dataset(table.DatasetId).Table(table.TableId)
		tables, err := bqContext.ListTables(table.DatasetId)
		if err != nil {
			return err
		}
		if !slices.Contains(tables, table.TableId) {
			logger.Info("Creating bigquery table", "table", bqTable)
			var schema bigquery.Schema
			switch table.EventCategory {
			case "transactional-email":
				schema, err = GenerateTableSchema(TransactionalEmailEventBigquery{}, TransactionalEmailEventBigqueryDescription)
			case "marketing-email":
				schema, err = GenerateTableSchema(MarketingEmailEventBigquery{}, MarketingEmailEventBigqueryDescription)
			case "marketing-sms":
				schema, err = GenerateTableSchema(MarketingSMSEventBigquery{}, MarketingSMSEventBigqueryDescription)
			case "transactional-sms":
				schema, err = GenerateTableSchema(TransactionalSMSEventBigquery{}, TransactionalSMSEventBigqueryDescription)
			default:
				return fmt.Errorf("schema not found for event category %s", table.EventCategory)
			}
			if err != nil {
				return err
			}
			err = bqTable.Create(bqContext.Ctx, &bigquery.TableMetadata{Schema: schema})
			if err != nil {
				return err
			}
		} else {
			logger.Info("Bigquery table already exists", "table", bqTable)
		}
		bqContext.Uploaders[fmt.Sprintf("%s.%s", table.DatasetId, table.TableId)] = bqTable.Uploader()
		logger.Info("Uploader created", "source", table.Source, "datasetId", table.DatasetId, "tableId", table.TableId)
	}
	return nil
}

/*
List all the tables in the dataset from the BqContext object
*/
func (bqContext *BqContext) ListTables(datasetId string) ([]string, error) {
	table := make([]string, 1)
	ts := bqContext.Client.Dataset(datasetId).Tables(bqContext.Ctx)
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
func GenerateTableSchema(model any, descriptions map[string]string) (bigquery.Schema, error) {
	schema, err := bigquery.InferSchema(model)
	if err != nil {
		return schema, err
	}
	// Remove the required constraint for all fields
	for i := range schema {
		schema[i].Required = false
	}
	// Add the descriptions to the fields
	for _, field := range schema {
		if desc, ok := descriptions[field.Name]; ok {
			field.Description = desc
		}
	}
	return schema, nil
}

/*
Get the target table for the source
*/
func (bqContext *BqContext) GetTargetTable(source string) (string, string, error) {
	for _, table := range bqContext.Tables {
		if table.Source == source {
			return table.DatasetId, table.TableId, nil
		}
	}
	return "", "", fmt.Errorf("table not found for source: %s", source)
}
