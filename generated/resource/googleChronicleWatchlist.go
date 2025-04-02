package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleChronicleWatchlist = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Time the watchlist was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. Description of the watchlist.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Required. Display name of the watchlist.\nNote that it must be at least one character and less than 63 characters\n(https://google.aip.dev/148).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_count": {
        "computed": true,
        "description": "Count of different types of entities in the watchlist.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "asset": "number",
              "user": "number"
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
      "instance": {
        "description": "The unique identifier for the Chronicle instance, which is the same as the customer ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the resource. This is the geographical region where the Chronicle instance resides, such as \"us\" or \"europe-west2\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "multiplying_factor": {
        "description": "Optional. Weight applied to the risk score for entities\nin this watchlist.\nThe default is 1.0 if it is not specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Identifier. Resource name of the watchlist. This unique identifier is generated using values provided for the URL parameters.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/watchlists/{watchlist}",
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
        "description": "Output only. Time the watchlist was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "watchlist_id": {
        "computed": true,
        "description": "Optional. The ID to use for the watchlist,\nwhich will become the final component of the watchlist's resource name.\nThis value should be 4-63 characters, and valid characters\nare /a-z-/.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "entity_population_mechanism": {
        "block": {
          "block_types": {
            "manual": {
              "block": {
                "description": "Entities are added manually.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Mechanism to populate entities in the watchlist.",
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
      },
      "watchlist_user_preferences": {
        "block": {
          "attributes": {
            "pinned": {
              "description": "Optional. Whether the watchlist is pinned on the dashboard.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "A collection of user preferences for watchlist UI configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleChronicleWatchlistSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleChronicleWatchlist), &result)
	return &result
}
