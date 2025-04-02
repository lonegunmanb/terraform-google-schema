package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleParameterManagerParameterVersionRender = `{
  "block": {
    "attributes": {
      "disabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parameter": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameter_data": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "parameter_version_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rendered_parameter_data": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleParameterManagerParameterVersionRenderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleParameterManagerParameterVersionRender), &result)
	return &result
}
