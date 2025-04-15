package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirestoreIndex = `{
  "block": {
    "attributes": {
      "api_scope": {
        "description": "The API scope at which a query is run. Default value: \"ANY_API\" Possible values: [\"ANY_API\", \"DATASTORE_MODE_API\", \"MONGODB_COMPATIBLE_API\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "collection": {
        "description": "The collection being indexed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database": {
        "description": "The Firestore database id. Defaults to '\"(default)\"'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "density": {
        "computed": true,
        "description": "The density configuration for this index. Possible values: [\"SPARSE_ALL\", \"SPARSE_ANY\", \"DENSE\"]",
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
      "multikey": {
        "description": "Optional. Whether the index is multikey. By default, the index is not multikey. For non-multikey indexes, none of the paths in the index definition reach or traverse an array, except via an explicit array index. For multikey indexes, at most one of the paths in the index definition reach or traverse an array, except via an explicit array index. Violations will result in errors. Note this field only applies to indexes with MONGODB_COMPATIBLE_API ApiScope.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "A server defined name for this index. Format:\n'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_scope": {
        "description": "The scope at which a query is run. Default value: \"COLLECTION\" Possible values: [\"COLLECTION\", \"COLLECTION_GROUP\", \"COLLECTION_RECURSIVE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "fields": {
        "block": {
          "attributes": {
            "array_config": {
              "description": "Indicates that this field supports operations on arrayValues. Only one of 'order', 'arrayConfig', and\n'vectorConfig' can be specified. Possible values: [\"CONTAINS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "field_path": {
              "description": "Name of the field.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "order": {
              "description": "Indicates that this field supports ordering by the specified order or comparing using =, \u003c, \u003c=, \u003e, \u003e=.\nOnly one of 'order', 'arrayConfig', and 'vectorConfig' can be specified. Possible values: [\"ASCENDING\", \"DESCENDING\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "vector_config": {
              "block": {
                "attributes": {
                  "dimension": {
                    "description": "The resulting index will only include vectors of this dimension, and can be used for vector search\nwith the same dimension.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "flat": {
                    "block": {
                      "description": "Indicates the vector index is a flat index.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Indicates that this field supports vector search operations. Only one of 'order', 'arrayConfig', and\n'vectorConfig' can be specified. Vector Fields should come after the field path '__name__'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The fields supported by this index. The last non-stored field entry is\nalways for the field path '__name__'. If, on creation, '__name__' was not\nspecified as the last field, it will be added automatically with the same\ndirection as that of the last field defined. If the final field in a\ncomposite index is not directional, the '__name__' will be ordered\n'\"ASCENDING\"' (unless explicitly specified otherwise).",
          "description_kind": "plain"
        },
        "min_items": 1,
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

func GoogleFirestoreIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirestoreIndex), &result)
	return &result
}
