package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEventarcGoogleApiSource = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Resource annotations.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "crypto_key_name": {
        "description": "Resource name of a KMS crypto key (managed by the user) used to\nencrypt/decrypt their event data.\n\nIt must match the pattern\n'projects/*/locations/*/keyRings/*/cryptoKeys/*'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination": {
        "description": "Destination is the message bus that the GoogleApiSource is delivering to.\nIt must be point to the full resource name of a MessageBus. Format:\n\"projects/{PROJECT_ID}/locations/{region}/messagesBuses/{MESSAGE_BUS_ID)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "Resource display name.",
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
        "description": "This checksum is computed by the server based on the value of other\nfields, and might be sent only on update and delete requests to ensure that\nthe client has an up-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "google_api_source_id": {
        "description": "The user-provided ID to be assigned to the GoogleApiSource. It should match\nthe format '^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$'.",
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
        "description": "Resource labels.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
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
        "description": "Resource name of the form\nprojects/{project}/locations/{location}/googleApiSources/{google_api_source}",
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
        "description": "Server assigned unique identifier for the channel. The value is a UUID4\nstring and guaranteed to remain unchanged until the resource is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The last-modified time.",
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
              "description": "The minimum severity of logs that will be sent to Stackdriver/Platform\nTelemetry. Logs at severitiy ≥ this value will be sent, unless it is NONE. Possible values: [\"NONE\", \"DEBUG\", \"INFO\", \"NOTICE\", \"WARNING\", \"ERROR\", \"CRITICAL\", \"ALERT\", \"EMERGENCY\"]",
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

func GoogleEventarcGoogleApiSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEventarcGoogleApiSource), &result)
	return &result
}
