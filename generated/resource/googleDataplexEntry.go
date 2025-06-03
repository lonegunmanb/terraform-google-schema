package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataplexEntry = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the Entry was created in Dataplex.",
        "description_kind": "plain",
        "type": "string"
      },
      "entry_group_id": {
        "description": "The entry group id of the entry group the entry will be created in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "entry_id": {
        "description": "The entry id of the entry.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "entry_type": {
        "description": "The relative resource name of the entry type that was used to create this entry, in the format projects/{project_number}/locations/{locationId}/entryTypes/{entryTypeId}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fully_qualified_name": {
        "description": "A name for the entry that can be referenced by an external system. For more information, see https://cloud.google.com/dataplex/docs/fully-qualified-names.\nThe maximum size of the field is 4000 characters.",
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
      "location": {
        "description": "The location where entry will be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The relative resource name of the entry, in the format projects/{project_number}/locations/{locationId}/entryGroups/{entryGroupId}/entries/{entryId}.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_entry": {
        "description": "The resource name of the parent entry, in the format projects/{project_number}/locations/{locationId}/entryGroups/{entryGroupId}/entries/{entryId}.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time when the entry was last updated in Dataplex.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "aspects": {
        "block": {
          "attributes": {
            "aspect_key": {
              "description": "Depending on how the aspect is attached to the entry, the format of the aspect key can be one of the following:\n\nIf the aspect is attached directly to the entry: {project_number}.{locationId}.{aspectTypeId}\nIf the aspect is attached to an entry's path: {project_number}.{locationId}.{aspectTypeId}@{path}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "aspect": {
              "block": {
                "attributes": {
                  "aspect_type": {
                    "computed": true,
                    "description": "The resource name of the type used to create this Aspect.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "create_time": {
                    "computed": true,
                    "description": "The time when the Aspect was created.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data": {
                    "description": "The content of the aspect in JSON form, according to its aspect type schema. The maximum size of the field is 120KB (encoded as UTF-8).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description": "The path in the entry under which the aspect is attached.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "update_time": {
                    "computed": true,
                    "description": "The time when the Aspect was last modified.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "A nested object resource.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The aspects that are attached to the entry.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "entry_source": {
        "block": {
          "attributes": {
            "create_time": {
              "description": "The time when the resource was created in the source system.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description": "A description of the data resource. Maximum length is 2,000 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "display_name": {
              "description": "A user-friendly display name. Maximum length is 500 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "description": "User-defined labels. The maximum size of keys and values is 128 characters each.\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "location": {
              "computed": true,
              "description": "Location of the resource in the source system. You can search the entry by this location.\nBy default, this should match the location of the entry group containing this entry.\nA different value allows capturing the source location for data external to Google Cloud.",
              "description_kind": "plain",
              "type": "string"
            },
            "platform": {
              "description": "The platform containing the source system. Maximum length is 64 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource": {
              "description": "The name of the resource in the source system. Maximum length is 4,000 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "system": {
              "description": "The name of the source system. Maximum length is 64 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_time": {
              "description": "The time when the resource was last updated in the source system.\nIf the entry exists in the system and its EntrySource has updateTime populated,\nfurther updates to the EntrySource of the entry must provide incremental updates to its updateTime.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ancestors": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The name of the ancestor resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "The type of the ancestor resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "A nested object resource.",
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

func GoogleDataplexEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataplexEntry), &result)
	return &result
}
