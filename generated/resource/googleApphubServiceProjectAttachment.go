package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApphubServiceProjectAttachment = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Create time.",
        "description_kind": "plain",
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
        "description": "\"Identifier. The resource name of a ServiceProjectAttachment. Format:\\\"projects/{host-project-id}/locations/global/serviceProjectAttachments/{service-project-id}.\\\" \"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_project": {
        "description": "\"Immutable. Service project name in the format: \\\"projects/abc\\\"\nor \\\"projects/123\\\". As input, project name with either project id or number\nare accepted. As output, this field will contain project number. \"",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_project_attachment_id": {
        "description": "Required. The service project attachment identifier must contain the project_id of the service project specified in the service_project_attachment.service_project field. Hint: \"projects/{project_id}\"",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "ServiceProjectAttachment state.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. A globally unique identifier (in UUID4 format) for the 'ServiceProjectAttachment'.",
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

func GoogleApphubServiceProjectAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApphubServiceProjectAttachment), &result)
	return &result
}
