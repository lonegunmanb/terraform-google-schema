package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleParameterManagerParameters = `{
  "block": {
    "attributes": {
      "filter": {
        "description": "Filter string, adhering to the rules in List-operation filtering. List only parameters matching the filter. \nIf filter is empty, all parameters are listed.",
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

func GoogleParameterManagerParametersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleParameterManagerParameters), &result)
	return &result
}
