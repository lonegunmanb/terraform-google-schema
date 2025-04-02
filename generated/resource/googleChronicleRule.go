package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleChronicleRule = `{
  "block": {
    "attributes": {
      "allowed_run_frequencies": {
        "computed": true,
        "description": "Output only. The run frequencies that are allowed for the rule.\nPopulated in BASIC view and FULL view.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "author": {
        "computed": true,
        "description": "Output only. The author of the rule. Extracted from the meta section of text.\nPopulated in BASIC view and FULL view.",
        "description_kind": "plain",
        "type": "string"
      },
      "compilation_diagnostics": {
        "computed": true,
        "description": "Output only. A list of a rule's corresponding compilation diagnostic messages\nsuch as compilation errors and compilation warnings.\nPopulated in FULL view.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "message": "string",
              "position": [
                "list",
                [
                  "object",
                  {
                    "end_column": "number",
                    "end_line": "number",
                    "start_column": "number",
                    "start_line": "number"
                  }
                ]
              ],
              "severity": "string",
              "uri": "string"
            }
          ]
        ]
      },
      "compilation_state": {
        "computed": true,
        "description": "Output only. The current compilation state of the rule.\nPopulated in FULL view.\nPossible values:\nCOMPILATION_STATE_UNSPECIFIED\nSUCCEEDED\nFAILED",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The timestamp of when the rule was created.\nPopulated in FULL view.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_tables": {
        "computed": true,
        "description": "Output only. Resource names of the data tables used in this rule.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "deletion_policy": {
        "description": "Policy to determine if the rule should be deleted forcefully.\nIf deletion_policy = \"FORCE\", any retrohunts and any detections associated with the rule\nwill also be deleted. If deletion_policy = \"DEFAULT\", the call will only succeed if the\nrule has no associated retrohunts, including completed retrohunts, and no\nassociated detections. Regardless of this field's value, the rule\ndeployment associated with this rule will also be deleted.\nPossible values: DEFAULT, FORCE",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Output only. Display name of the rule.\nPopulated in BASIC view and FULL view.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "The etag for this rule.\nIf this is provided on update, the request will succeed if and only if it\nmatches the server-computed value, and will fail with an ABORTED error\notherwise.\nPopulated in BASIC view and FULL view.",
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
      "metadata": {
        "computed": true,
        "description": "Output only. Additional metadata specified in the meta section of text.\nPopulated in FULL view.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Full resource name for the rule. This unique identifier is generated using values provided for the URL parameters.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/rules/{rule}",
        "description_kind": "plain",
        "type": "string"
      },
      "near_real_time_live_rule_eligible": {
        "computed": true,
        "description": "Output only. Indicate the rule can run in near real time live rule.\nIf this is true, the rule uses the near real time live rule when the run\nfrequency is set to LIVE.",
        "description_kind": "plain",
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reference_lists": {
        "computed": true,
        "description": "Output only. Resource names of the reference lists used in this rule.\nPopulated in FULL view.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "revision_create_time": {
        "computed": true,
        "description": "Output only. The timestamp of when the rule revision was created.\nPopulated in FULL, REVISION_METADATA_ONLY views.",
        "description_kind": "plain",
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "Output only. The revision ID of the rule.\nA new revision is created whenever the rule text is changed in any way.\nFormat: v_{10 digits}_{9 digits}\nPopulated in REVISION_METADATA_ONLY view and FULL view.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description": "Rule Id is the ID of the Rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "description": "Resource name of the DataAccessScope bound to this rule.\nPopulated in BASIC view and FULL view.\nIf reference lists are used in the rule, validations will be performed\nagainst this scope to ensure that the reference lists are compatible with\nboth the user's and the rule's scopes.\nThe scope should be in the format:\n\"projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "severity": {
        "computed": true,
        "description": "Severity represents the severity level of the rule.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "display_name": "string"
            }
          ]
        ]
      },
      "text": {
        "description": "The YARA-L content of the rule.\nPopulated in FULL view.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "Possible values:\nRULE_TYPE_UNSPECIFIED\nSINGLE_EVENT\nMULTI_EVENT",
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

func GoogleChronicleRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleChronicleRule), &result)
	return &result
}
