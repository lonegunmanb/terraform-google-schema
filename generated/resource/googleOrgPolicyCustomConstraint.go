package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOrgPolicyCustomConstraint = `{
  "block": {
    "attributes": {
      "action_type": {
        "description": "The action to take if the condition is met. Possible values: [\"ALLOW\", \"DENY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "condition": {
        "description": "A CEL condition that refers to a supported service resource, for example 'resource.management.autoUpgrade == false'. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "A human-friendly description of the constraint to display as an error message when the policy is violated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "A human-friendly name for the constraint.",
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
      "method_types": {
        "description": "A list of RESTful methods for which to enforce the constraint. Can be 'CREATE', 'UPDATE', or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Immutable. The name of the custom constraint. This is unique within the organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The parent of the resource, an organization. Format should be 'organizations/{organization_id}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_types": {
        "description": "Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, 'container.googleapis.com/NodePool'.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The timestamp representing when the constraint was last updated.",
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

func GoogleOrgPolicyCustomConstraintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOrgPolicyCustomConstraint), &result)
	return &result
}
