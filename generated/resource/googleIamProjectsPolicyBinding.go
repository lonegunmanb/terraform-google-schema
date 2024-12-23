package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamProjectsPolicyBinding = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size limitations\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time when the policy binding was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Optional. The description of the policy binding. Must be less than or equal to 63 characters.",
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
        "description": "Optional. The etag for the policy binding. If this is provided on update, it must match the server's etag.",
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
        "description": "The location of the Policy Binding",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the policy binding in the format '{binding_parent/locations/{location}/policyBindings/{policy_binding_id}'",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description": "Required. Immutable. The resource name of the policy to be bound. The binding parent and policy must belong to the same Organization (or Project).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_binding_id": {
        "description": "The Policy Binding ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_kind": {
        "description": "Immutable. The kind of the policy to attach in this binding. This\nfield must be one of the following:  - Left empty (will be automatically set\nto the policy kind) - The input policy kind   Possible values:  POLICY_KIND_UNSPECIFIED PRINCIPAL_ACCESS_BOUNDARY ACCESS",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_uid": {
        "computed": true,
        "description": "Output only. The globally unique ID of the policy to be bound.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. The globally unique ID of the policy binding. Assigned when the policy binding is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time when the policy binding was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "condition": {
        "block": {
          "attributes": {
            "description": {
              "description": "Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression": {
              "description": "Textual representation of an expression in Common Expression Language syntax.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "location": {
              "description": "Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "title": {
              "description": "Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Represents a textual expression in the Common Expression Language\n(CEL) syntax. CEL is a C-like expression language. The syntax and semantics of\nCEL are documented at https://github.com/google/cel-spec.\nExample (Comparison):\ntitle: \\\"Summary size limit\\\"\ndescription: \\\"Determines if a summary is less than 100 chars\\\"\nexpression: \\\"document.summary.size() \u003c 100\\\"\nExample\n(Equality):\ntitle: \\\"Requestor is owner\\\"\ndescription: \\\"Determines if requestor is the document owner\\\"\nexpression: \\\"document.owner == request.auth.claims.email\\\"  Example\n(Logic):\ntitle: \\\"Public documents\\\"\ndescription: \\\"Determine whether the document should be publicly visible\\\"\nexpression: \\\"document.type != 'private' \u0026\u0026 document.type != 'internal'\\\"\nExample (Data Manipulation):\ntitle: \\\"Notification string\\\"\ndescription: \\\"Create a notification string with a timestamp.\\\"\nexpression: \\\"'New message received at ' + string(document.create_time)\\\"\nThe exact variables and functions that may be referenced within an expression are\ndetermined by the service that evaluates it. See the service documentation for\nadditional information.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "target": {
        "block": {
          "attributes": {
            "principal_set": {
              "description": "Required. Immutable. The resource name of the policy to be bound.\nThe binding parent and policy must belong to the same Organization (or Project).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Target is the full resource name of the resource to which the policy will be bound. Immutable once set.",
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

func GoogleIamProjectsPolicyBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamProjectsPolicyBinding), &result)
	return &result
}
