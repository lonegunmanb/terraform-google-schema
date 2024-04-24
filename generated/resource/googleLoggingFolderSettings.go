package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingFolderSettings = `{
  "block": {
    "attributes": {
      "disable_default_sink": {
        "computed": true,
        "description": "If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "folder": {
        "description": "The folder for which to retrieve settings.",
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
      "kms_key_name": {
        "computed": true,
        "description": "The resource name for the configured Cloud KMS key.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_service_account_id": {
        "computed": true,
        "description": "The service account that will be used by the Log Router to access your Cloud KMS key.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_service_account_id": {
        "computed": true,
        "description": "The service account for the given container. Sinks use this service account as their writerIdentity if no custom service account is provided.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the settings.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_location": {
        "computed": true,
        "description": "The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.",
        "description_kind": "plain",
        "optional": true,
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

func GoogleLoggingFolderSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingFolderSettings), &result)
	return &result
}
