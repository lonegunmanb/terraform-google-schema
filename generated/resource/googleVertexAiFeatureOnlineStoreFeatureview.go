package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiFeatureOnlineStoreFeatureview = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp of when the featureOnlinestore was created in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "feature_online_store": {
        "description": "The name of the FeatureOnlineStore to use for the featureview.",
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
      "labels": {
        "description": "A set of key/value label pairs to assign to this FeatureView.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.",
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
      "region": {
        "description": "The region for the resource. It should be the same as the featureonlinestore region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "big_query_source": {
        "block": {
          "attributes": {
            "entity_id_columns": {
              "description": "Columns to construct entityId / row keys. Start by supporting 1 only.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "uri": {
              "description": "The BigQuery view URI that will be materialized on each sync trigger based on FeatureView.SyncConfig.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "feature_registry_source": {
        "block": {
          "attributes": {
            "project_number": {
              "description": "The project number of the parent project of the feature Groups.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "feature_groups": {
              "block": {
                "attributes": {
                  "feature_group_id": {
                    "description": "Identifier of the feature group.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "feature_ids": {
                    "description": "Identifiers of features under the feature group.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "List of features that need to be synced to Online Store.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configures the features from a Feature Registry source that need to be loaded onto the FeatureOnlineStore.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sync_config": {
        "block": {
          "attributes": {
            "cron": {
              "computed": true,
              "description": "Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs.\nTo explicitly set a timezone to the cron tab, apply a prefix in the cron tab: \"CRON_TZ=${IANA_TIME_ZONE}\" or \"TZ=${IANA_TIME_ZONE}\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.",
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

func GoogleVertexAiFeatureOnlineStoreFeatureviewSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiFeatureOnlineStoreFeatureview), &result)
	return &result
}
