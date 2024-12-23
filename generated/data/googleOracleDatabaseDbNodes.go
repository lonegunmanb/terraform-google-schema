package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseDbNodes = `{
  "block": {
    "attributes": {
      "cloud_vm_cluster": {
        "description": "vmcluster",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "properties": [
                "list",
                [
                  "object",
                  {
                    "db_node_storage_size_gb": "number",
                    "db_server_ocid": "string",
                    "hostname": "string",
                    "memory_size_gb": "number",
                    "ocid": "string",
                    "ocpu_count": "number",
                    "state": "string",
                    "total_cpu_core_count": "number"
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

func GoogleOracleDatabaseDbNodesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseDbNodes), &result)
	return &result
}
