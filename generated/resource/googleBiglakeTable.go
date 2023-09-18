package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBiglakeTable = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The creation time of the table. A timestamp in RFC3339 UTC\n\"Zulu\" format, with nanosecond resolution and up to nine fractional\ndigits. Examples: \"2014-10-02T15:01:23Z\" and\n\"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "database": {
        "description": "The id of the parent database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Output only. The deletion time of the table. Only set after the\ntable is deleted. A timestamp in RFC3339 UTC \"Zulu\" format, with\nnanosecond resolution and up to nine fractional digits. Examples:\n\"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "The checksum of a table object computed by the server based on the value\nof other fields. It may be sent on update requests to ensure the client\nhas an up-to-date value before proceeding. It is only checked for update\ntable operations.",
        "description_kind": "plain",
        "type": "string"
      },
      "expire_time": {
        "computed": true,
        "description": "Output only. The time when this table is considered expired. Only set\nafter the table is deleted. A timestamp in RFC3339 UTC \"Zulu\" format,\nwith nanosecond resolution and up to nine fractional digits. Examples:\n\"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Output only. The name of the Table. Format:\nprojects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "The database type. Possible values: [\"HIVE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last modification time of the table. A timestamp in\nRFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: \"2014-10-02T15:01:23Z\" and\n\"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "hive_options": {
        "block": {
          "attributes": {
            "parameters": {
              "description": "Stores user supplied Hive table parameters. An object containing a\nlist of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "table_type": {
              "description": "Hive table type. For example, MANAGED_TABLE, EXTERNAL_TABLE.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "storage_descriptor": {
              "block": {
                "attributes": {
                  "input_format": {
                    "description": "The fully qualified Java class name of the input format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "location_uri": {
                    "description": "Cloud Storage folder URI where the table data is stored, starting with \"gs://\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_format": {
                    "description": "The fully qualified Java class name of the output format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Stores physical storage information on the data.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Options of a Hive table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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

func GoogleBiglakeTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBiglakeTable), &result)
	return &result
}
