package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryAnalyticsHubListing = `{
  "block": {
    "attributes": {
      "allow_only_metadata_sharing": {
        "description": "If true, the listing is only available to get the resource metadata. Listing is non subscribable.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "categories": {
        "description": "Categories of the listing. Up to two categories are allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "commercial_info": {
        "computed": true,
        "description": "Commercial info contains the information about the commercial data products associated with the listing.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloud_marketplace": [
                "list",
                [
                  "object",
                  {
                    "commercial_state": "string",
                    "service": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "data_exchange_id": {
        "description": "The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delete_commercial": {
        "description": "If the listing is commercial then this field must be set to true, otherwise a failure is thrown. This acts as a safety guard to avoid deleting commercial listings accidentally.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "discovery_type": {
        "computed": true,
        "description": "Specifies the type of discovery on the discovery page. Cannot be set for a restricted listing. Note that this does not control the visibility of the exchange/listing which is defined by IAM permission. Possible values: [\"DISCOVERY_TYPE_PRIVATE\", \"DISCOVERY_TYPE_PUBLIC\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (\u0026) and can't start or end with spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "documentation": {
        "description": "Documentation describing the listing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "icon": {
        "description": "Base64 encoded image representing the listing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listing_id": {
        "description": "The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The name of the location this data exchange listing.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_linked_dataset_query_user_email": {
        "description": "If true, subscriber email logging is enabled and all queries on the linked dataset will log the email address of the querying user. Once enabled, this setting cannot be turned off.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the listing. e.g. \"projects/myproject/locations/US/dataExchanges/123/listings/456\"",
        "description_kind": "plain",
        "type": "string"
      },
      "primary_contact": {
        "description": "Email or URL of the primary point of contact of the listing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_access": {
        "description": "Email or URL of the request access of the listing. Subscribers can use this reference to request access.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Current state of the listing.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "bigquery_dataset": {
        "block": {
          "attributes": {
            "dataset": {
              "description": "Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "selected_resources": {
              "block": {
                "attributes": {
                  "table": {
                    "description": "Format: For table: projects/{projectId}/datasets/{datasetId}/tables/{tableId} Example:\"projects/test_project/datasets/test_dataset/tables/test_table\"",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Resource in this dataset that is selectively shared. This field is required for data clean room exchanges.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Shared dataset i.e. BigQuery dataset source.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "data_provider": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the data provider.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "primary_contact": {
              "description": "Email or URL of the data provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Details of the data provider who owns the source data.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "publisher": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the listing publisher.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "primary_contact": {
              "description": "Email or URL of the listing publisher.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Details of the publisher who owns the listing and who can share the source data.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "pubsub_topic": {
        "block": {
          "attributes": {
            "data_affinity_regions": {
              "description": "Region hint on where the data might be published. Data affinity regions are modifiable.\nSee https://cloud.google.com/about/locations for full listing of possible Cloud regions.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "topic": {
              "description": "Resource name of the Pub/Sub topic source for this listing. e.g. projects/myproject/topics/topicId",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Pub/Sub topic source.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "restricted_export_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "If true, enable restricted export.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "restrict_direct_table_access": {
              "computed": true,
              "description": "If true, restrict direct table access(read api/tabledata.list) on linked table.",
              "description_kind": "plain",
              "type": "bool"
            },
            "restrict_query_result": {
              "description": "If true, restrict export of query result derived from restricted linked dataset table.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "If set, restricted export configuration will be propagated and enforced on the linked dataset.",
          "description_kind": "plain"
        },
        "max_items": 1,
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
            },
            "update": {
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

func GoogleBigqueryAnalyticsHubListingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryAnalyticsHubListing), &result)
	return &result
}
