package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApihubApiHubInstance = `{
  "block": {
    "attributes": {
      "api_hub_instance_id": {
        "description": "Optional. Identifier to assign to the Api Hub instance. Must be unique within\nscope of the parent resource. If the field is not provided,\nsystem generated id will be used.\n\nThis value should be 4-40 characters, and valid characters\nare '/a-z[0-9]-_/'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Creation timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. Description of the ApiHub instance.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Optional. Instance labels to represent user-provided metadata.\nRefer to cloud documentation on labels for more details.\nhttps://cloud.google.com/compute/docs/labeling-resources\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. Format:\n'projects/{project}/locations/{location}/apiHubInstances/{apiHubInstance}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The current state of the ApiHub instance.\nPossible values:\nSTATE_UNSPECIFIED\nINACTIVE\nCREATING\nACTIVE\nUPDATING\nDELETING\nFAILED",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "Output only. Extra information about ApiHub instance state. Currently the message\nwould be populated when state is 'FAILED'.",
        "description_kind": "plain",
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
        "description": "Output only. Last update timestamp.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "config": {
        "block": {
          "attributes": {
            "cmek_key_name": {
              "description": "Optional. The Customer Managed Encryption Key (CMEK) used for data encryption.\nThe CMEK name should follow the format of\n'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)',\nwhere the location must match the instance location.\nIf the CMEK is not provided, a GMEK will be created for the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disable_search": {
              "description": "Optional. If true, the search will be disabled for the instance. The default value\nis false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "encryption_type": {
              "computed": true,
              "description": "Optional. Encryption type for the region. If the encryption type is CMEK, the\ncmek_key_name must be provided. If no encryption type is provided,\nGMEK will be used.\nPossible values:\nENCRYPTION_TYPE_UNSPECIFIED\nGMEK\nCMEK",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vertex_location": {
              "description": "Optional. The name of the Vertex AI location where the data store is stored.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Available configurations to provision an ApiHub Instance.",
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

func GoogleApihubApiHubInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApihubApiHubInstance), &result)
	return &result
}
