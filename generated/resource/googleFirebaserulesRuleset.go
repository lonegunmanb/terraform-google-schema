package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaserulesRuleset = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Time the ` + "`" + `Ruleset` + "`" + ` was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description": "Output only. The metadata for this ruleset.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "services": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "Output only. Name of the ` + "`" + `Ruleset` + "`" + `. The ruleset_id is auto generated by the service. Format: ` + "`" + `projects/{project_id}/rulesets/{ruleset_id}` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "source": {
        "block": {
          "attributes": {
            "language": {
              "description": "` + "`" + `Language` + "`" + ` of the ` + "`" + `Source` + "`" + ` bundle. If unspecified, the language will default to ` + "`" + `FIREBASE_RULES` + "`" + `. Possible values: LANGUAGE_UNSPECIFIED, FIREBASE_RULES, EVENT_FLOW_TRIGGERS",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "files": {
              "block": {
                "attributes": {
                  "content": {
                    "description": "Textual Content.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "fingerprint": {
                    "description": "Fingerprint (e.g. github sha) associated with the ` + "`" + `File` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "File name.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "` + "`" + `File` + "`" + ` set constituting the ` + "`" + `Source` + "`" + ` bundle.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "` + "`" + `Source` + "`" + ` for the ` + "`" + `Ruleset` + "`" + `.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleFirebaserulesRulesetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaserulesRuleset), &result)
	return &result
}
