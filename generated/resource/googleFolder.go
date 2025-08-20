package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFolder = `{
  "block": {
    "attributes": {
      "configured_capabilities": {
        "computed": true,
        "description": "A list of capabilities that are configured for this folder.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "description": "When the field is set to true or unset in Terraform state, a terraform apply or terraform destroy that would delete the instance will fail. When the field is set to false, deleting the instance is allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "folder_id": {
        "computed": true,
        "description": "The folder id from the name \"folders/{folder_id}\"",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lifecycle_state": {
        "computed": true,
        "description": "The lifecycle state of the folder such as ACTIVE or DELETE_REQUESTED.",
        "description_kind": "plain",
        "type": "string"
      },
      "management_project": {
        "computed": true,
        "description": "The Management Project associated with the folder's configured capabilities.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Folder. Its format is folders/{folder_id}.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The resource name of the parent Folder or Organization. Must be of the form folders/{folder_id} or organizations/{org_id}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description": "A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when empty. This field is only set at create time and modifying this field after creation will trigger recreation. To apply tags to an existing resource, see the google_tags_tag_value resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
            "read": {
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

func GoogleFolderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFolder), &result)
	return &result
}
