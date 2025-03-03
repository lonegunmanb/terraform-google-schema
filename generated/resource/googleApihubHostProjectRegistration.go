package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApihubHostProjectRegistration = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which the host project registration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "gcp_project": {
        "description": "Required. Immutable. Google cloud project name in the format: \"projects/abc\" or \"projects/123\".\nAs input, project name with either project id or number are accepted.\nAs output, this field will contain project number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_project_registration_id": {
        "description": "Required. The ID to use for the Host Project Registration, which will become the\nfinal component of the host project registration's resource name. The ID\nmust be the same as the Google cloud project specified in the\nhost_project_registration.gcp_project field.",
        "description_kind": "plain",
        "required": true,
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
        "description": "Identifier. The name of the host project registration.\nFormat:\n\"projects/{project}/locations/{location}/hostProjectRegistrations/{host_project_registration}\".",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
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

func GoogleApihubHostProjectRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApihubHostProjectRegistration), &result)
	return &result
}
