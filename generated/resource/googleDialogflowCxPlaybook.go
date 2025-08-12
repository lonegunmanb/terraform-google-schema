package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxPlaybook = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp of initial playbook creation.\n\nUses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9 fractional digits. Offsets other than \"Z\" are also accepted. Examples: \"2014-10-02T15:01:23Z\", \"2014-10-02T15:01:23.045123456Z\" or \"2014-10-02T15:01:23+05:30\".",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the playbook, unique within an agent.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "goal": {
        "description": "High level description of the goal the playbook intend to accomplish. A goal should be concise since it's visible to other playbooks that may reference this playbook.",
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
      "name": {
        "computed": true,
        "description": "The unique identifier of the Playbook.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/playbooks/\u003cPlaybook ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a Playbook for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "playbook_type": {
        "description": "Type of the playbook. Possible values: [\"PLAYBOOK_TYPE_UNSPECIFIED\", \"TASK\", \"ROUTINE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "referenced_flows": {
        "computed": true,
        "description": "The resource name of flows referenced by the current playbook in the instructions.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "referenced_playbooks": {
        "computed": true,
        "description": "The resource name of other playbooks referenced by the current playbook in the instructions.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "referenced_tools": {
        "description": "The resource name of tools referenced by the current playbook in the instructions. If not provided explicitly, they are will be implied using the tool being referenced in goal and steps.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "token_count": {
        "computed": true,
        "description": "Estimated number of tokes current playbook takes when sent to the LLM.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Last time the playbook version was updated.\n\nUses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9 fractional digits. Offsets other than \"Z\" are also accepted. Examples: \"2014-10-02T15:01:23Z\", \"2014-10-02T15:01:23.045123456Z\" or \"2014-10-02T15:01:23+05:30\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "instruction": {
        "block": {
          "attributes": {
            "guidelines": {
              "description": "General guidelines for the playbook. These are unstructured instructions that are not directly part of the goal, e.g. \"Always be polite\". It's valid for this text to be long and used instead of steps altogether.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "steps": {
              "block": {
                "attributes": {
                  "steps": {
                    "description": "Sub-processing needed to execute the current step.\n\nThis field uses JSON data as a string. The value provided must be a valid JSON representation documented in [Step](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.playbooks#step).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "text": {
                    "description": "Step instruction in text format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Ordered list of step by step execution instructions to accomplish target goal.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Instruction to accomplish target goal.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "llm_model_settings": {
        "block": {
          "attributes": {
            "model": {
              "description": "The selected LLM model.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prompt_text": {
              "description": "The custom prompt to use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Llm model settings for the playbook.",
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

func GoogleDialogflowCxPlaybookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxPlaybook), &result)
	return &result
}
