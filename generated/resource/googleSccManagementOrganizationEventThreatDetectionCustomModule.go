package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSccManagementOrganizationEventThreatDetectionCustomModule = `{
  "block": {
    "attributes": {
      "config": {
        "description": "Config for the module. For the resident module, its config value is defined at this level.\nFor the inherited module, its config value is inherited from the ancestor module.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human readable name to be displayed for the module.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enablement_state": {
        "description": "The state of enablement for the module at the given level of the hierarchy. Possible values: [\"ENABLED\", \"DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_editor": {
        "computed": true,
        "description": "The editor that last updated the custom module",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "Location ID of the parent organization. Only global is supported at the moment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Event Threat Detection custom module.\nIts format is \"organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{eventThreatDetectionCustomModule}\".",
        "description_kind": "plain",
        "type": "string"
      },
      "organization": {
        "description": "Numerical ID of the parent organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which the custom module was last updated.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and\nup to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
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

func GoogleSccManagementOrganizationEventThreatDetectionCustomModuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSccManagementOrganizationEventThreatDetectionCustomModule), &result)
	return &result
}
