package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryAnalyticsHubListingSubscription = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Timestamp when the subscription was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_exchange_id": {
        "description": "The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modify_time": {
        "computed": true,
        "description": "Timestamp when the subscription was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "linked_dataset_map": {
        "computed": true,
        "description": "Output only. Map of listing resource names to associated linked resource,\ne.g. projects/123/locations/US/dataExchanges/456/listings/789 -\u003e projects/123/datasets/my_dataset",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "linked_dataset": "string",
              "listing": "string",
              "resource_name": "string"
            }
          ]
        ]
      },
      "linked_resources": {
        "computed": true,
        "description": "Output only. Linked resources created in the subscription. Only contains values if state = STATE_ACTIVE.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "linked_dataset": "string",
              "listing": "string"
            }
          ]
        ]
      },
      "listing_id": {
        "description": "The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The name of the location of the data exchange. Distinct from the location of the destination data set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_linked_dataset_query_user_email": {
        "computed": true,
        "description": "Output only. By default, false. If true, the Subscriber agreed to the email sharing mandate that is enabled for Listing.",
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the subscription. e.g. \"projects/myproject/locations/US/subscriptions/123\"",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_display_name": {
        "computed": true,
        "description": "Display name of the project of this subscription.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_id": {
        "computed": true,
        "description": "Organization of the project this subscription belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "Listing shared asset type.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Current state of the subscription.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscriber_contact": {
        "computed": true,
        "description": "Email of the subscriber.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_id": {
        "computed": true,
        "description": "The subscription id used to reference the subscription.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "destination_dataset": {
        "block": {
          "attributes": {
            "description": {
              "description": "A user-friendly description of the dataset.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "friendly_name": {
              "description": "A descriptive name for the dataset.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "description": "The labels associated with this dataset. You can use these to\norganize and group your datasets.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "location": {
              "description": "The geographic location where the dataset should reside.\nSee https://cloud.google.com/bigquery/docs/locations for supported locations.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "dataset_reference": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "description": "A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The ID of the project containing this dataset.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A reference that identifies the destination dataset.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The destination dataset for this subscription.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBigqueryAnalyticsHubListingSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryAnalyticsHubListingSubscription), &result)
	return &result
}
