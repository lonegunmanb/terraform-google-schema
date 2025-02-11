package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleColabRuntime = `{
  "block": {
    "attributes": {
      "auto_upgrade": {
        "description": "Triggers an upgrade anytime the runtime is started if it is upgradable.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "The description of the Runtime.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_state": {
        "description": "Desired state of the Colab Runtime. Set this field to 'RUNNING' to start the runtime, and 'STOPPED' to stop it.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Required. The display name of the Runtime.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expiration_time": {
        "computed": true,
        "description": "Output only. Timestamp when this NotebookRuntime will be expired.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_upgradable": {
        "computed": true,
        "description": "Output only. Checks if the NotebookRuntime is upgradable.",
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "description": "The location for the resource: https://cloud.google.com/colab/docs/locations",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the Runtime",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notebook_runtime_type": {
        "computed": true,
        "description": "Output only. The type of the notebook runtime.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_user": {
        "description": "The user email of the NotebookRuntime.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The state of the runtime.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "notebook_runtime_template_ref": {
        "block": {
          "attributes": {
            "notebook_runtime_template": {
              "description": "The resource name of the NotebookRuntimeTemplate based on which a NotebookRuntime will be created.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "'Runtime specific information used for NotebookRuntime creation.'",
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

func GoogleColabRuntimeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleColabRuntime), &result)
	return &result
}
