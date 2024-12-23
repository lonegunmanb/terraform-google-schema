package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerInstanceConfig = `{
  "block": {
    "attributes": {
      "base_config": {
        "computed": true,
        "description": "Base configuration name, e.g. nam3, based on which this configuration is created.\nOnly set for user managed configurations.\nbaseConfig must refer to a configuration of type GOOGLE_MANAGED in the same project as this configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "config_type": {
        "computed": true,
        "description": "Output only. Whether this instance config is a Google or User Managed Configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The name of this instance configuration as it appears in UIs.",
        "description_kind": "plain",
        "required": true,
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "An object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "A unique identifier for the instance configuration. Values are of the\nform projects/\u003cproject\u003e/instanceConfigs/[a-z][-a-z0-9]*",
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
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "replicas": {
        "block": {
          "attributes": {
            "default_leader_location": {
              "description": "If true, this location is designated as the default leader location where\nleader replicas are placed.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "location": {
              "description": "The location of the serving resources, e.g. \"us-central1\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Indicates the type of replica.  See the [replica types\ndocumentation](https://cloud.google.com/spanner/docs/replication#replica_types)\nfor more details. Possible values: [\"READ_WRITE\", \"READ_ONLY\", \"WITNESS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The geographic placement of nodes in this instance configuration and their replication properties.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
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

func GoogleSpannerInstanceConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSpannerInstanceConfig), &result)
	return &result
}
