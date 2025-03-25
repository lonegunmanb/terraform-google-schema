package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterRoutePolicy = `{
  "block": {
    "attributes": {
      "fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.  Used\ninternally during updates.",
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
        "description": "Name of the route policy. This policy's name, which must be a resource ID segment and unique within all policies owned by the Router",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where the router and NAT reside.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router": {
        "description": "The name of the Cloud Router in which this route policy will be configured.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "This is policy's type, which is one of IMPORT or EXPORT Possible values: [\"ROUTE_POLICY_TYPE_IMPORT\", \"ROUTE_POLICY_TYPE_EXPORT\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "terms": {
        "block": {
          "attributes": {
            "priority": {
              "description": "The evaluation priority for this term, which must be between 0 (inclusive) and 231 (exclusive), and unique within the list.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "actions": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "Description of the expression",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expression": {
                    "description": "Textual representation of an expression in Common Expression\nLanguage syntax.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "location": {
                    "description": "String indicating the location of the expression for error\nreporting, e.g. a file name and a position in the file",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "title": {
                    "description": "Title for the expression, i.e. a short string describing its\npurpose.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "'CEL expressions to evaluate to modify a route when this term matches.'\\",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "match": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "Description of the expression",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expression": {
                    "description": "Textual representation of an expression in Common Expression Language syntax.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "location": {
                    "description": "String indicating the location of the expression for error reporting, e.g. a file name and a position in the file",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "title": {
                    "description": "Title for the expression, i.e. a short string describing its purpose.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "CEL expression evaluated against a route to determine if this term applies (see Policy Language). When not set, the term applies to all routes.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "List of terms (the order in the list is not important, they are evaluated in order of priority).",
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

func GoogleComputeRouterRoutePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterRoutePolicy), &result)
	return &result
}
