package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageManagedFolder = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the bucket that contains the managed folder.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp at which this managed folder was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "force_destroy": {
        "description": "Allows the deletion of a managed folder even if contains\nobjects. If a non-empty managed folder is deleted, any objects\nwithin the folder will remain in a simulated folder with the\nsame name.",
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
        "description": "The metadata generation of the managed folder.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the managed folder expressed as a path. Must include\ntrailing '/'. For example, 'example_dir/example_dir2/'.",
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
        "description": "The timestamp at which this managed folder was most recently updated.",
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

func GoogleStorageManagedFolderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageManagedFolder), &result)
	return &result
}
