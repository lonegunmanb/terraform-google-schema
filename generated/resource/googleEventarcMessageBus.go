package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEventarcMessageBus = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. Resource annotations.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "crypto_key_name": {
        "description": "Optional. Resource name of a KMS crypto key (managed by the user) used to\nencrypt/decrypt their event data.\n\nIt must match the pattern\n'projects/*/locations/*/keyRings/*/cryptoKeys/*'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Optional. Resource display name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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
      "etag": {
        "computed": true,
        "description": "Output only. This checksum is computed by the server based on the value of other\nfields, and might be sent only on update and delete requests to ensure that\nthe client has an up-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Optional. Resource labels.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
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
      "message_bus_id": {
        "description": "Required. The user-provided ID to be assigned to the MessageBus. It should match the\nformat '^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. Resource name of the form\nprojects/{project}/locations/{location}/messageBuses/{message_bus}",
        "description_kind": "plain",
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
      },
      "uid": {
        "computed": true,
        "description": "Output only. Server assigned unique identifier for the channel. The value is a UUID4\nstring and guaranteed to remain unchanged until the resource is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last-modified time.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "logging_config": {
        "block": {
          "attributes": {
            "log_severity": {
              "computed": true,
              "description": "Optional. The minimum severity of logs that will be sent to Stackdriver/Platform\nTelemetry. Logs at severitiy ≥ this value will be sent, unless it is NONE. Possible values: [\"NONE\", \"DEBUG\", \"INFO\", \"NOTICE\", \"WARNING\", \"ERROR\", \"CRITICAL\", \"ALERT\", \"EMERGENCY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The configuration for Platform Telemetry logging for Eventarc Advanced\nresources.",
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

func GoogleEventarcMessageBusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEventarcMessageBus), &result)
	return &result
}
