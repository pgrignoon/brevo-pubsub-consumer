# Brevo Pub/Sub to BigQuery Consumer

## Overview

This project is a Google Cloud Function written in Go that acts as a consumer for Brevo (formerly Sendinblue) webhook events. It receives events via Google Cloud Pub/Sub, processes them, and streams them into corresponding tables in Google BigQuery.

The primary use case is to create a data pipeline for Brevo marketing and transactional events, allowing for data analysis, monitoring, and building dashboards on top of Brevo event data.

## Features

- **Event-Driven Architecture**: Consumes events from Pub/Sub for a scalable and decoupled system.
- **Multiple Event Types**: Supports various Brevo event types, including transactional emails, marketing emails, transactional SMS, and marketing SMS.
- **Dynamic Table Routing**: Uses a flexible JSON configuration to map different event types to specific BigQuery datasets and tables. This allows a single function deployment to handle multiple projects or event categories seamlessly.
- **Automatic Table Creation**: If a target BigQuery table does not exist, the function automatically creates it based on a predefined schema for the event type.
- **Robust Data Handling**: Transforms incoming event payloads into a BigQuery-compatible format, correctly handling nullable fields.
- **Extensible**: Designed to be easily extendable to support new Brevo event types.

## How It Works

1.  **Event Trigger**: A Brevo webhook sends an event to a webhook endpoint (not part of this project) that publishes the event to a Google Cloud Pub/Sub topic.
2.  **Function Invocation**: The Pub/Sub message triggers this Google Cloud Function. The message is expected to have specific attributes: `category`, `target-dataset`, and `target-table`.
3.  **Initialization (`init`)**: On a cold start, the function:
    a. Reads environment variables for the GCP Project ID and the path to the configuration file.
    b. Loads the table mapping configuration from `config.json`.
    c. Initializes the Google BigQuery client.
    d. For each table defined in the configuration, it checks if the table exists. If not, it creates it with the appropriate schema.
    e. Creates a BigQuery `Uploader` instance for each table to efficiently stream data.
4.  **Message Processing**:
    a. The function extracts the event payload and the attributes from the Pub/Sub message.
    b. The `category` attribute determines the type of event (e.g., `transactional-email`).
    c. The `target-dataset` and `target-table` attributes determine the destination BigQuery table.
    d. Based on the category, the function decodes the JSON payload into the corresponding Go struct.
    e. The data is converted into a BigQuery-compatible format.
    f. The data is uploaded to the target BigQuery table using the uploader.

## Configuration

The function is configured through a combination of environment variables and a JSON configuration file.

### Environment Variables

-   `GCP_PROJECT_ID`: The ID of your Google Cloud Project where the function and BigQuery datasets reside.
-   `CONFIG_FILE_PATH`: The path to the configuration file (e.g., `config.json`).

### `config.json`

This file defines the mapping between event categories and BigQuery tables. It allows for flexible routing of events.

**Structure**:

```json
{
    "tables": [
        {
            "datasetId": "your_dataset_id",
            "tableId": "your_table_id",
            "eventCategory": "event_category_name"
        }
    ]
}
```

-   `datasetId`: The BigQuery dataset ID.
-   `tableId`: The BigQuery table ID.
-   `eventCategory`: A string that identifies the event type. This must match the `category` attribute in the Pub/Sub message.

**Example**:

```json
{
    "tables": [
        {
            "datasetId": "brevo_events",
            "tableId": "transactional_emails",
            "eventCategory": "transactional-email"
        },
        {
            "datasetId": "brevo_events",
            "tableId": "marketing_sms_campaign_1",
            "eventCategory": "marketing-sms"
        },
        {
            "datasetId": "another_project_dataset",
            "tableId": "marketing_sms_campaign_2",
            "eventCategory": "marketing-sms"
        }
    ]
}
```

In this example, `transactional-email` events are routed to the `transactional_emails` table in the `brevo_events` dataset. `marketing-sms` events can be routed to two different tables based on the Pub/Sub message attributes.

## Deployment

This function is designed to be deployed as a 2nd generation Google Cloud Function.

1.  **Prerequisites**:
    -   Google Cloud SDK (`gcloud`) installed and configured.
    -   A Pub/Sub topic.
    -   A service account for the function with permissions for Pub/Sub (subscriber) and BigQuery (data editor, job user).

2.  **Deploy Command**: TODO

## How to Add a New Event Type

To add support for a new Brevo event type, follow these steps:

1.  **Create a new Go file** for the event (e.g., `newEventType.go`).
2.  **Define two structs**:
    -   `NewEventTypeEvent`: Represents the JSON structure of the webhook payload from Brevo. Use pointers for all fields to handle missing values.
    -   `NewEventTypeEventBigquery`: Represents the BigQuery schema. Use `bigquery.Null*` types for nullable fields.
3.  **Implement the `Event` interface**: Create a `ToBigquery()` method for your `NewEventTypeEvent` struct that converts it to the `NewEventTypeEventBigquery` struct.
4.  **Update `function.go`**: Add a new `case` in the `switch` statement in `runPubSubConsumer` for your new event category.
5.  **Update `bigqueryClient.go`**: Add a new `case` in the `switch` statement in `CreateTablesAndUploaders` to handle automatic table creation for the new event type.
6.  **Update `config.json`**: Add a new entry for your event type, mapping it to a dataset and table.
7.  **Redeploy** the Cloud Function.
