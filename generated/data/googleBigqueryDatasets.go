package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryDatasets = `{
  "block": {
    "attributes": {
      "datasets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dataset_id": "string",
              "friendly_name": "string",
              "labels": [
                "map",
                "string"
              ],
              "location": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description": "The ID of the project in which the datasets are located. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBigqueryDatasetsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryDatasets), &result)
	return &result
}
