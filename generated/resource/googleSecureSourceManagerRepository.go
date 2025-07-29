package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecureSourceManagerRepository = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the repository was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_policy": {
        "description": "The deletion policy for the repository. Setting 'ABANDON' allows the resource\nto be abandoned, rather than deleted. Setting 'DELETE' deletes the resource\nand all its contents. Setting 'PREVENT' prevents the resource from accidental deletion\nby erroring out during plan.\nDefault is 'DELETE'.  Possible values are:\n  * DELETE\n  * PREVENT\n  * ABANDON",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "Description of the repository, which cannot exceed 500 characters.",
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
      "instance": {
        "description": "The name of the instance in which the repository is hosted.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the Repository.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for the Repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_id": {
        "description": "The ID for the Repository.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Unique identifier of the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the repository was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "uris": {
        "computed": true,
        "description": "URIs for the repository.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "api": "string",
              "git_https": "string",
              "html": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "initial_config": {
        "block": {
          "attributes": {
            "default_branch": {
              "description": "Default branch name of the repository.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gitignores": {
              "description": "List of gitignore template names user can choose from.\nValid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "license": {
              "description": "License template name user can choose from.\nValid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "readme": {
              "description": "README template name.\nValid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Initial configurations for the repository.",
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

func GoogleSecureSourceManagerRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecureSourceManagerRepository), &result)
	return &result
}
