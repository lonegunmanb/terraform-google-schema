package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBackupDrBackup = `{
  "block": {
    "attributes": {
      "backup_vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backups": {
        "computed": true,
        "description": "List of all backups under data source.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_id": "string",
              "backup_vault_id": "string",
              "data_source_id": "string",
              "location": "string",
              "name": "string"
            }
          ]
        ]
      },
      "data_source_id": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of resource",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBackupDrBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrBackup), &result)
	return &result
}
