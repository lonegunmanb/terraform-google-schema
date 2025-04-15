package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMlEngineModel = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The description specified for the model when it was created.",
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
        "description": "One or more labels that you can add, to organize your models.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name specified for the model.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "online_prediction_console_logging": {
        "description": "If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "online_prediction_logging": {
        "description": "If true, online prediction access logs are sent to StackDriver Logging.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "regions": {
        "description": "The list of regions where the model is going to be deployed.\nCurrently only one region per model is supported",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "default_version": {
        "block": {
          "attributes": {
            "name": {
              "description": "The name specified for the version when it was created.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The default version of the model. This version will be used to handle\nprediction requests that do not specify a version.",
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
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleMlEngineModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMlEngineModel), &result)
	return &result
}
