package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApphubApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "Required. The Application identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "attributes": {
        "computed": true,
        "description": "Consumer provided attributes.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "business_owners": [
                "list",
                [
                  "object",
                  {
                    "display_name": "string",
                    "email": "string"
                  }
                ]
              ],
              "criticality": [
                "list",
                [
                  "object",
                  {
                    "type": "string"
                  }
                ]
              ],
              "developer_owners": [
                "list",
                [
                  "object",
                  {
                    "display_name": "string",
                    "email": "string"
                  }
                ]
              ],
              "environment": [
                "list",
                [
                  "object",
                  {
                    "type": "string"
                  }
                ]
              ],
              "operator_owners": [
                "list",
                [
                  "object",
                  {
                    "display_name": "string",
                    "email": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Create time.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Optional. User-defined description of an Application.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Optional. User-defined name for the Application.",
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
        "description": "Part of 'parent'. See documentation of 'projectsId'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of an Application. Format:\n\"projects/{host-project-id}/locations/{location}/applications/{application-id}\"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "Scope of an application.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "type": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "Output only. Application state. \n Possible values:\n STATE_UNSPECIFIED\nCREATING\nACTIVE\nDELETING",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. A universally unique identifier (in UUID4 format) for the 'Application'.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Update time.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleApphubApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApphubApplication), &result)
	return &result
}
