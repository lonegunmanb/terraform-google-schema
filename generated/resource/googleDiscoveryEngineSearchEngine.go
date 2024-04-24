package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDiscoveryEngineSearchEngine = `{
  "block": {
    "attributes": {
      "collection_id": {
        "description": "The collection ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Timestamp the Engine was created at.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_store_ids": {
        "description": "The data stores associated with this engine. For SOLUTION_TYPE_SEARCH type of engines, they can only associate with at most one data store.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description": "Required. The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_id": {
        "description": "Unique ID to use for Search Engine App.",
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
      "industry_vertical": {
        "description": "The industry vertical that the engine registers. The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to GENERIC. Vertical on Engine has to match vertical of the DataStore liniked to the engine. Default value: \"GENERIC\" Possible values: [\"GENERIC\", \"MEDIA\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique full resource name of the search engine. Values are of the format\n'projects/{project}/locations/{location}/collections/{collection_id}/engines/{engine_id}'.\nThis field must be a UTF-8 encoded string with a length limit of 1024\ncharacters.",
        "description_kind": "plain",
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
        "description": "Timestamp the Engine was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "common_config": {
        "block": {
          "attributes": {
            "company_name": {
              "description": "The name of the company, business or entity that is associated with the engine. Setting this may help improve LLM related features.cd",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Common config spec that specifies the metadata of the engine.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "search_engine_config": {
        "block": {
          "attributes": {
            "search_add_ons": {
              "description": "The add-on that this search engine enables. Possible values: [\"SEARCH_ADD_ON_LLM\"]",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "search_tier": {
              "description": "The search feature tier of this engine. Defaults to SearchTier.SEARCH_TIER_STANDARD if not specified. Default value: \"SEARCH_TIER_STANDARD\" Possible values: [\"SEARCH_TIER_STANDARD\", \"SEARCH_TIER_ENTERPRISE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configurations for a Search Engine.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleDiscoveryEngineSearchEngineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDiscoveryEngineSearchEngine), &result)
	return &result
}
