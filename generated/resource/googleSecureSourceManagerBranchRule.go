package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecureSourceManagerBranchRule = `{
  "block": {
    "attributes": {
      "allow_stale_reviews": {
        "description": "Determines if allow stale reviews or approvals before merging to the branch.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "branch_rule_id": {
        "description": "The ID for the BranchRule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time the BranchRule was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "description": "Determines if the branch rule is disabled or not.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_pattern": {
        "description": "The BranchRule matches branches based on the specified regular expression. Use .* to match all branches.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the Repository.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "minimum_approvals_count": {
        "description": "The minimum number of approvals required for the branch rule to be matched.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "minimum_reviews_count": {
        "description": "The minimum number of reviews required for the branch rule to be matched.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The resource name for the BranchRule.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_id": {
        "description": "The ID for the Repository.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "require_comments_resolved": {
        "description": "Determines if require comments resolved before merging to the branch.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "require_linear_history": {
        "description": "Determines if require linear history before merging to the branch.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "require_pull_request": {
        "description": "Determines if the branch rule requires a pull request or not.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "uid": {
        "computed": true,
        "description": "Unique identifier of the BranchRule.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the BranchRule was updated in UTC.",
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

func GoogleSecureSourceManagerBranchRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecureSourceManagerBranchRule), &result)
	return &result
}
