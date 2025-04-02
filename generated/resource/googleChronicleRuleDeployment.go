package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleChronicleRuleDeployment = `{
  "block": {
    "attributes": {
      "alerting": {
        "description": "Whether detections resulting from this deployment should be considered\nalerts.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "archive_time": {
        "computed": true,
        "description": "Output only. The timestamp when the rule deployment archive state was last set to true. If the rule deployment's current archive state is not set to true, the field will be empty.",
        "description_kind": "plain",
        "type": "string"
      },
      "archived": {
        "description": "The archive state of the rule deployment.\nCannot be set to true unless enabled is set to false i.e.\narchiving requires a two-step process: first, disable the rule by\nsetting 'enabled' to false, then set 'archive' to true.\nIf set to true, alerting will automatically be set to false.\nIf currently set to true, enabled, alerting, and run_frequency cannot be\nupdated.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "consumer_rules": {
        "computed": true,
        "description": "Output only. The names of the associated/chained consumer rules. Rules are considered\nconsumers of this rule if their rule text explicitly filters on this rule's ruleid.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/rules/{rule}",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "enabled": {
        "description": "Whether the rule is currently deployed continuously against incoming data.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "execution_state": {
        "computed": true,
        "description": "The execution state of the rule deployment.\nPossible values:\nDEFAULT\nLIMITED\nPAUSED",
        "description_kind": "plain",
        "type": "string"
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
      "last_alert_status_change_time": {
        "computed": true,
        "description": "Output only. The timestamp when the rule deployment alert state was lastly changed. This is filled regardless of the current alert state.E.g. if the current alert status is false, this timestamp will be the timestamp when the alert status was changed to false.",
        "description_kind": "plain",
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
        "description": "The resource name of the rule deployment.\nNote that RuleDeployment is a child of the overall Rule, not any individual\nrevision, so the resource ID segment for the Rule resource must not\nreference a specific revision.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/rules/{rule}/deployment",
        "description_kind": "plain",
        "type": "string"
      },
      "producer_rules": {
        "computed": true,
        "description": "Output only. The names of the associated/chained producer rules. Rules are considered\nproducers for this rule if this rule explicitly filters on their ruleid.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/rules/{rule}",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "project": {
        "computed": true,
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
      "run_frequency": {
        "description": "The run frequency of the rule deployment.\nPossible values:\nLIVE\nHOURLY\nDAILY",
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

func GoogleChronicleRuleDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleChronicleRuleDeployment), &result)
	return &result
}
