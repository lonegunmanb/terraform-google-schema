package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleChronicleReferenceList = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Required. A user-provided description of the reference list.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Output only. The unique display name of the reference list.",
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
      "location": {
        "description": "The location of the resource. This is the geographical region where the Chronicle instance resides, such as \"us\" or \"europe-west2\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. The resource name of the reference list.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/referenceLists/{reference_list}",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reference_list_id": {
        "description": "Required. The ID to use for the reference list. This is also the display name for\nthe reference list. It must satisfy the following requirements:\n- Starts with letter.\n- Contains only letters, numbers and underscore.\n- Has length \u003c 256.\n- Must be unique.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision_create_time": {
        "computed": true,
        "description": "Output only. The timestamp when the reference list was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_associations_count": {
        "computed": true,
        "description": "Output only. The count of self-authored rules using the reference list.",
        "description_kind": "plain",
        "type": "number"
      },
      "rules": {
        "computed": true,
        "description": "Output only. The resource names for the associated self-authored Rules that use this\nreference list.\nThis is returned only when the view is REFERENCE_LIST_VIEW_FULL.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "scope_info": {
        "computed": true,
        "description": "ScopeInfo specifies the scope info of the reference list.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "reference_list_scope": [
                "list",
                [
                  "object",
                  {
                    "scope_names": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "syntax_type": {
        "description": "Possible values:\nREFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING\nREFERENCE_LIST_SYNTAX_TYPE_REGEX\nREFERENCE_LIST_SYNTAX_TYPE_CIDR",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "entries": {
        "block": {
          "attributes": {
            "value": {
              "description": "Required. The value of the entry. Maximum length is 512 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Required. The entries of the reference list.\nWhen listed, they are returned in the order that was specified at creation\nor update. The combined size of the values of the reference list may not\nexceed 6MB.\nThis is returned only when the view is REFERENCE_LIST_VIEW_FULL.",
          "description_kind": "plain"
        },
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

func GoogleChronicleReferenceListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleChronicleReferenceList), &result)
	return &result
}
