package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeSecurityAction = `{
  "block": {
    "attributes": {
      "api_proxies": {
        "description": "If unset, this would apply to all proxies in the environment.\nIf set, this action is enforced only if at least one proxy in the repeated\nlist is deployed at the time of enforcement. If set, several restrictions are enforced on SecurityActions.\nThere can be at most 100 enabled actions with proxies set in an env.\nSeveral other restrictions apply on conditions and are detailed later.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The create time for this SecurityAction.\nUses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9 fractional digits.\nOffsets other than \"Z\" are also accepted. Examples: \"2014-10-02T15:01:23Z\", \"2014-10-02T15:01:23.045123456Z\" or \"2014-10-02T15:01:23+05:30\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional user provided description of the SecurityAction.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "env_id": {
        "description": "The Apigee environment that this security action applies to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expire_time": {
        "description": "The expiration for this SecurityAction.\nUses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9\nfractional digits. Offsets other than \"Z\" are also accepted.\nExamples: \"2014-10-02T15:01:23Z\", \"2014-10-02T15:01:23.045123456Z\" or \"2014-10-02T15:01:23+05:30\".",
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
      "org_id": {
        "description": "The organization that this security action applies to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_action_id": {
        "description": "The ID to use for the SecurityAction, which will become the final component of the action's resource name.\nThis value should be 0-61 characters, and valid format is (^a-z?$).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "description": "Only an ENABLED SecurityAction is enforced. An ENABLED SecurityAction past its expiration time will not be enforced. Possible values: [\"ENABLED\", \"DISABLED\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ttl": {
        "description": "The TTL for this SecurityAction.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The update time for this SecurityAction. This reflects when this SecurityAction changed states.\nUses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9 fractional digits.\nOffsets other than \"Z\" are also accepted. Examples: \"2014-10-02T15:01:23Z\", \"2014-10-02T15:01:23.045123456Z\" or \"2014-10-02T15:01:23+05:30\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "allow": {
        "block": {
          "description": "Allow a request through if it matches this SecurityAction.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "condition_config": {
        "block": {
          "attributes": {
            "access_tokens": {
              "description": "A list of accessTokens. Limit 1000 per action.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "api_keys": {
              "description": "A list of API keys. Limit 1000 per action.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "api_products": {
              "description": "A list of API Products. Limit 1000 per action.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "asns": {
              "description": "A list of ASN numbers to act on, e.g. 23. https://en.wikipedia.org/wiki/Autonomous_system_(Internet)\nThis uses int64 instead of uint32 because of https://linter.aip.dev/141/forbidden-types.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "bot_reasons": {
              "description": "A list of Bot Reasons. Current options: Flooder, Brute Guessor, Static Content Scraper,\nOAuth Abuser, Robot Abuser, TorListRule, Advanced Anomaly Detection, Advanced API Scraper,\nSearch Engine Crawlers, Public Clouds, Public Cloud AWS, Public Cloud Azure, and Public Cloud Google.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "developer_apps": {
              "description": "A list of developer apps. Limit 1000 per action.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "developers": {
              "description": "A list of developers. Limit 1000 per action.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "http_methods": {
              "description": "Act only on particular HTTP methods. E.g. A read-only API can block POST/PUT/DELETE methods.\nAccepted values are: GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE and PATCH.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "ip_address_ranges": {
              "description": "A list of IP addresses. This could be either IPv4 or IPv6. Limited to 100 per action.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "region_codes": {
              "description": "A list of countries/region codes to act on, e.g. US. This follows https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "user_agents": {
              "description": "A list of user agents to deny. We look for exact matches. Limit 50 per action.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "A valid SecurityAction must contain at least one condition.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "deny": {
        "block": {
          "attributes": {
            "response_code": {
              "description": "The HTTP response code if the Action = DENY.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Deny a request through if it matches this SecurityAction.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "flag": {
        "block": {
          "block_types": {
            "headers": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The header name to be sent to the target.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The header value to be sent to the target.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A list of HTTP headers to be sent to the target in case of a FLAG SecurityAction.\nLimit 5 headers per SecurityAction.\nAt least one is mandatory.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Flag a request through if it matches this SecurityAction.",
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

func GoogleApigeeSecurityActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeSecurityAction), &result)
	return &result
}
