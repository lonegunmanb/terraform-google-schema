package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsAutokeyConfig = `{
  "block": {
    "attributes": {
      "etag": {
        "computed": true,
        "description": "The etag of the AutokeyConfig for optimistic concurrency control.",
        "description_kind": "plain",
        "type": "string"
      },
      "folder": {
        "description": "The folder for which to retrieve config.",
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
      "key_project": {
        "computed": true,
        "description": "The target key project for a given folder where KMS Autokey will provision a\nCryptoKey for any new KeyHandle the Developer creates. Should have the form\n'projects/\u003cproject_id_or_number\u003e'.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsAutokeyConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsAutokeyConfig), &result)
	return &result
}
