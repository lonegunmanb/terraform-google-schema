package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaseDataConnectService = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. Stores small amounts of arbitrary data.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. [Output only] Create time stamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_policy": {
        "description": "The deletion policy for the database. Setting the field to FORCE allows the\nService to be deleted even if a Schema or Connector is present. By default,\nthe Service deletion will only succeed when no Schema or Connectors are\npresent.\nPossible values: DEFAULT, FORCE",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Optional. Mutable human-readable name. 63 character limit.",
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
        "description": "Output only. This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.\n[AIP-154](https://google.aip.dev/154)",
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
        "description": "Optional. Labels as key value pairs.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The region in which the service resides, e.g. \"us-central1\" or \"asia-east1\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The relative resource name of the Firebase Data Connect service, in the\nformat:\n'''\nprojects/{project}/locations/{location}/services/{service}\n'''\nNote that the service ID is specific to Firebase Data Connect and does not\ncorrespond to any of the instance IDs of the underlying data source\nconnections.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Output only. A field that if true, indicates that the system is working update the\nservice.",
        "description_kind": "plain",
        "type": "bool"
      },
      "service_id": {
        "description": "Required. The ID to use for the service, which will become the final component of the\nservice's resource name.",
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
      "uid": {
        "computed": true,
        "description": "Output only. System-assigned, unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. [Output only] Update time stamp.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleFirebaseDataConnectServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaseDataConnectService), &result)
	return &result
}
