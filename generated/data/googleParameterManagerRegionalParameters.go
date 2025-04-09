package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleParameterManagerRegionalParameters = `{
  "block": {
    "attributes": {
      "filter": {
        "description": "Filter string, adhering to the rules in List-operation filtering. List only parameters matching the filter. \nIf filter is empty, all regional parameters are listed from specific location.",
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "effective_labels": [
                "map",
                "string"
              ],
              "format": "string",
              "kms_key": "string",
              "labels": [
                "map",
                "string"
              ],
              "location": "string",
              "name": "string",
              "parameter_id": "string",
              "policy_member": [
                "list",
                [
                  "object",
                  {
                    "iam_policy_name_principal": "string",
                    "iam_policy_uid_principal": "string"
                  }
                ]
              ],
              "project": "string",
              "terraform_labels": [
                "map",
                "string"
              ],
              "update_time": "string"
            }
          ]
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleParameterManagerRegionalParametersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleParameterManagerRegionalParameters), &result)
	return &result
}
