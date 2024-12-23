package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseDbServers = `{
  "block": {
    "attributes": {
      "cloud_exadata_infrastructure": {
        "description": "exadata",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_servers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "display_name": "string",
              "properties": [
                "list",
                [
                  "object",
                  {
                    "db_node_ids": [
                      "list",
                      "string"
                    ],
                    "db_node_storage_size_gb": "number",
                    "max_db_node_storage_size_gb": "number",
                    "max_memory_size_gb": "number",
                    "max_ocpu_count": "number",
                    "memory_size_gb": "number",
                    "ocid": "string",
                    "ocpu_count": "number",
                    "state": "string",
                    "vm_count": "number"
                  }
                ]
              ]
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
      "location": {
        "description": "location",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "The ID of the project in which the dataset is located. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleOracleDatabaseDbServersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseDbServers), &result)
	return &result
}
