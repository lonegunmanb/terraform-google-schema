package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsKeyRings = `{
  "block": {
    "attributes": {
      "filter": {
        "description": "\n\t\t\t\t\tThe filter argument is used to add a filter query parameter that limits which keys are retrieved by the data source: ?filter={{filter}}.\n\t\t\t\t\tExample values:\n\t\t\t\t\t\n\t\t\t\t\t* \"name:my-key-\" will retrieve key rings that contain \"my-key-\" anywhere in their name. Note: names take the form projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}.\n\t\t\t\t\t* \"name=projects/my-project/locations/global/keyRings/my-key-ring\" will only retrieve a key ring with that exact name.\n\t\t\t\t\t\n\t\t\t\t\t[See the documentation about using filters](https://cloud.google.com/kms/docs/sorting-and-filtering)\n\t\t\t\t",
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
      "key_rings": {
        "computed": true,
        "description": "A list of all the retrieved key rings",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsKeyRingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsKeyRings), &result)
	return &result
}
