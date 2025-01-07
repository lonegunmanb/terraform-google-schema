package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageFolder = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the bucket that contains the folder.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp at which this folder was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "force_destroy": {
        "description": "If set to true, items within folder if any will be force destroyed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metageneration": {
        "computed": true,
        "description": "The metadata generation of the folder.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the folder expressed as a path. Must include\ntrailing '/'. For example, 'example_dir/example_dir2/', 'example@#/', 'a-b/d-f/'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp at which this folder was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleStorageFolderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageFolder), &result)
	return &result
}
