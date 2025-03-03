package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageAnywhereCache = `{
  "block": {
    "attributes": {
      "admission_policy": {
        "description": "The cache admission policy dictates whether a block should be inserted upon a cache miss. Default value: \"admit-on-first-miss\" Possible values: [\"admit-on-first-miss\", \"admit-on-second-miss\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "anywhere_cache_id": {
        "computed": true,
        "description": "The ID of the Anywhere cache instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "description": "A reference to Bucket resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The creation time of the cache instance in RFC 3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pending_update": {
        "computed": true,
        "description": "True if the cache instance has an active Update long-running operation.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of the cache instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "ttl": {
        "description": "The TTL of all cache entries in whole seconds. e.g., \"7200s\". It defaults to '86400s'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The modification time of the cache instance metadata in RFC 3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "description": "The zone in which the cache instance needs to be created. For example, 'us-central1-a.'",
        "description_kind": "plain",
        "required": true,
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

func GoogleStorageAnywhereCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageAnywhereCache), &result)
	return &result
}
