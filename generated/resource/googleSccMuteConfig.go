package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSccMuteConfig = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time at which the mute config was created. This field is set by\nthe server and will be ignored if provided on config creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the mute config.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "expiry_time": {
        "description": "Optional. The expiry of the mute config. Only applicable for dynamic configs.\nIf the expiry is set, when the config expires, it is removed from all findings.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to\nnine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter": {
        "description": "An expression that defines the filter to apply across create/update\nevents of findings. While creating a filter string, be mindful of\nthe scope in which the mute configuration is being created. E.g.,\nIf a filter contains project = X but is created under the\nproject = Y scope, it might not match any findings.",
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
      "most_recent_editor": {
        "computed": true,
        "description": "Email address of the user who last edited the mute config. This\nfield is set by the server and will be ignored if provided on\nconfig creation or update.",
        "description_kind": "plain",
        "type": "string"
      },
      "mute_config_id": {
        "description": "Unique identifier provided by the client within the parent scope.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the mute config. Its format is\norganizations/{organization}/muteConfigs/{configId},\nfolders/{folder}/muteConfigs/{configId},\nor projects/{project}/muteConfigs/{configId}",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "Resource name of the new mute configs's parent. Its format is\n\"organizations/[organization_id]\", \"folders/[folder_id]\", or\n\"projects/[project_id]\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "The type of the mute config, which determines what type of mute state the config affects. Default value: \"DYNAMIC\" Possible values: [\"MUTE_CONFIG_TYPE_UNSPECIFIED\", \"STATIC\", \"DYNAMIC\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The most recent time at which the mute config was\nupdated. This field is set by the server and will be ignored if\nprovided on config creation or update.",
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

func GoogleSccMuteConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSccMuteConfig), &result)
	return &result
}
