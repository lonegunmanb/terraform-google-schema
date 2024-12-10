package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamPrincipalAccessBoundaryPolicy = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "User defined annotations. See https://google.aip.dev/148#annotations\nfor more details such as format and size limitations\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time when the principal access boundary policy was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The description of the principal access boundary policy. Must be less than or equal to 63 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "etag": {
        "computed": true,
        "description": "The etag for the principal access boundary. If this is provided on update, it must match the server's etag.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location the principal access boundary policy is in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the principal access boundary policy.  The following format is supported:\n 'organizations/{organization_id}/locations/{location}/principalAccessBoundaryPolicies/{policy_id}'",
        "description_kind": "plain",
        "type": "string"
      },
      "organization": {
        "description": "The parent organization of the principal access boundary policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_access_boundary_policy_id": {
        "description": "The ID to use to create the principal access boundary policy.\nThis value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, hyphens, or dots. Pattern, /a-z{2,62}/.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. The globally unique ID of the principal access boundary policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time when the principal access boundary policy was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "details": {
        "block": {
          "attributes": {
            "enforcement_version": {
              "computed": true,
              "description": "The version number that indicates which Google Cloud services\nare included in the enforcement (e.g. \\\"latest\\\", \\\"1\\\", ...). If empty, the\nPAB policy version will be set to the current latest version, and this version\nwon't get updated when new versions are released.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "rules": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "The description of the principal access boundary policy rule. Must be less than or equal to 256 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "effect": {
                    "description": "The access relationship of principals to the resources in this rule.\nPossible values: ALLOW",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "resources": {
                    "description": "A list of Cloud Resource Manager resources. The resource\nand all the descendants are included. The number of resources in a policy\nis limited to 500 across all rules.\nThe following resource types are supported:\n* Organizations, such as '//cloudresourcemanager.googleapis.com/organizations/123'.\n* Folders, such as '//cloudresourcemanager.googleapis.com/folders/123'.\n* Projects, such as '//cloudresourcemanager.googleapis.com/projects/123'\nor '//cloudresourcemanager.googleapis.com/projects/my-project-id'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "A list of principal access boundary policy rules. The number of rules in a policy is limited to 500.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Principal access boundary policy details",
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

func GoogleIamPrincipalAccessBoundaryPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamPrincipalAccessBoundaryPolicy), &result)
	return &result
}
