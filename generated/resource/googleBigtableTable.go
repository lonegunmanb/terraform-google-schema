package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigtableTable = `{
  "block": {
    "attributes": {
      "change_stream_retention": {
        "computed": true,
        "description": "Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, currently deletion protection will be set to UNPROTECTED as it is the API default value.",
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
      "instance_name": {
        "description": "The name of the Bigtable instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "split_keys": {
        "description": "A list of predefined keys to split the table on. !\u003e Warning: Modifying the split_keys of an existing table will cause Terraform to delete/recreate the entire google_bigtable_table resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "automated_backup_policy": {
        "block": {
          "attributes": {
            "frequency": {
              "computed": true,
              "description": "How frequently automated backups should occur.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retention_period": {
              "computed": true,
              "description": "How long the automated backups should be retained.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Defines an automated backup policy for a table, specified by Retention Period and Frequency. To _create_ a table with automated backup disabled, omit this argument. To disable automated backup on an _existing_ table that has automated backup enabled, set both Retention Period and Frequency to \"0\". If this argument is not provided in the configuration on update, the resource's automated backup policy will _not_ be modified.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "column_family": {
        "block": {
          "attributes": {
            "family": {
              "description": "The name of the column family.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "The type of the column family.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A group of columns within a table which share a common configuration. This can be specified multiple times.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBigtableTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigtableTable), &result)
	return &result
}
