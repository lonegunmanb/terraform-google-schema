package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetappVolumeQuotaRule = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Create time of the quota rule. A timestamp in RFC3339 UTC \"Zulu\" format. Examples: \"2023-06-22T09:13:01.617Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description for the quota rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_limit_mib": {
        "description": "The maximum allowed capacity in MiB.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
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
        "description": "Labels as key value pairs of the quota rule. Example: '{ \"owner\": \"Bob\", \"department\": \"finance\", \"purpose\": \"testing\" }'.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Loction of the quotaRule. QuotaRules are child resources of volumes and live in the same location.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the quotaRule.",
        "description_kind": "plain",
        "required": true,
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
        "description": "The state of the quota rule. Possible Values : [STATE_UNSPECIFIED, CREATING, UPDATING, READY, DELETING, ERROR]",
        "description_kind": "plain",
        "type": "string"
      },
      "state_details": {
        "computed": true,
        "description": "State details of the quota rule",
        "description_kind": "plain",
        "type": "string"
      },
      "target": {
        "description": "The quota rule applies to the specified user or group.\nValid targets for volumes with NFS protocol enabled:\n  - UNIX UID for individual user quota\n  - UNIX GID for individual group quota\nValid targets for volumes with SMB protocol enabled:\n  - Windows SID for individual user quota\nLeave empty for default quotas",
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
      "type": {
        "description": "Types of Quota Rule. Possible values: [\"INDIVIDUAL_USER_QUOTA\", \"INDIVIDUAL_GROUP_QUOTA\", \"DEFAULT_USER_QUOTA\", \"DEFAULT_GROUP_QUOTA\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_name": {
        "description": "Name of the volume to create the quotaRule in.",
        "description_kind": "plain",
        "required": true,
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

func GoogleNetappVolumeQuotaRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetappVolumeQuotaRule), &result)
	return &result
}
