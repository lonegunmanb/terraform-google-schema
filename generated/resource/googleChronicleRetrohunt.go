package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleChronicleRetrohunt = `{
  "block": {
    "attributes": {
      "execution_interval": {
        "computed": true,
        "description": "Represents a time interval, encoded as a Timestamp start (inclusive) and a\nTimestamp end (exclusive).\n\nThe start must be less than or equal to the end.\nWhen the start equals the end, the interval is empty (matches no time).\nWhen both start and end are unspecified, the interval matches any time.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "end_time": "string",
              "start_time": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "The unique identifier for the Chronicle instance, which is the same as the customer ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the resource. This is the geographical region where the Chronicle instance resides, such as \"us\" or \"europe-west2\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the retrohunt.\nRetrohunt is the child of a rule revision. {rule} in the format below is\nstructured as {rule_id@revision_id}.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/rules/{rule}/retrohunts/{retrohunt}",
        "description_kind": "plain",
        "type": "string"
      },
      "progress_percentage": {
        "computed": true,
        "description": "Output only. Percent progress of the retrohunt towards completion, from 0.00 to 100.00.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retrohunt": {
        "computed": true,
        "description": "The retrohunt ID of the Retrohunt. A retrohunt is an execution of a Rule over a time range in the past.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule": {
        "description": "The Rule ID of the rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The state of the retrohunt.\nPossible values:\nRUNNING\nDONE\nCANCELLED\nFAILED",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "process_interval": {
        "block": {
          "attributes": {
            "end_time": {
              "description": "Exclusive end of the interval.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "start_time": {
              "description": "Inclusive start of the interval.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Represents a time interval, encoded as a Timestamp start (inclusive) and a\nTimestamp end (exclusive).\n\nThe start must be less than or equal to the end.\nWhen the start equals the end, the interval is empty (matches no time).\nWhen both start and end are unspecified, the interval matches any time.",
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

func GoogleChronicleRetrohuntSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleChronicleRetrohunt), &result)
	return &result
}
