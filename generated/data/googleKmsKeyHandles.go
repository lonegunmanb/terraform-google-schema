package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsKeyHandles = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_handles": {
        "computed": true,
        "description": "A list of all the retrieved key handles",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key": "string",
              "name": "string",
              "resource_type_selector": "string"
            }
          ]
        ]
      },
      "location": {
        "description": "The canonical id for the location. For example: \"us-east1\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "Project ID of the project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type_selector": {
        "description": "\n\t\t\t\t\tThe resource_type_selector argument is used to add a filter query parameter that limits which key handles are retrieved by the data source: ?filter=resource_type_selector=\"{{resource_type_selector}}\".\n\t\t\t\t\tExample values:\n\t\t\t\t\t* resource_type_selector=\"{SERVICE}.googleapis.com/{TYPE}\".\n\t\t\t\t\t[See the documentation about using filters](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyHandles/list)\n\t\t\t\t",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsKeyHandlesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsKeyHandles), &result)
	return &result
}
