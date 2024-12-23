package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudQuotasQuotaPreference = `{
  "block": {
    "attributes": {
      "contact_email": {
        "description": "An email address that can be used for quota related communication between the Google Cloud and the user in case the Google Cloud needs further information to make a decision on whether the user preferred quota can be granted.\n\nThe Google account for the email address must have quota update permission for the project, folder or organization this quota preference is for.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Create time stamp.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.",
        "description_kind": "plain",
        "type": "string"
      },
      "dimensions": {
        "computed": true,
        "description": "The dimensions that this quota preference applies to. The key of the map entry is the name of a dimension, such as \"region\", \"zone\", \"network_id\", and the value of the map entry is the dimension value. If a dimension is missing from the map of dimensions, the quota preference applies to all the dimension values except for those that have other quota preferences configured for the specific value.\n\nNOTE: QuotaPreferences can only be applied across all values of \"user\" and \"resource\" dimension. Do not set values for \"user\" or \"resource\" in the dimension map.\n\nExample: '{\"provider\": \"Foo Inc\"}' where \"provider\" is a service specific dimension.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "etag": {
        "computed": true,
        "description": "The current etag of the quota preference. If an etag is provided on update and does not match the current server's etag of the quota preference, the request will be blocked and an ABORTED error will be returned. See https://google.aip.dev/134#etags for more details on etags.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_safety_checks": {
        "description": "The list of quota safety checks to be ignored. Default value: \"QUOTA_SAFETY_CHECK_UNSPECIFIED\" Possible values: [\"QUOTA_SAFETY_CHECK_UNSPECIFIED\", \"QUOTA_DECREASE_BELOW_USAGE\", \"QUOTA_DECREASE_PERCENTAGE_TOO_HIGH\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "justification": {
        "description": "The reason / justification for this quota preference.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the quota preference. Required except in the CREATE requests.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent": {
        "computed": true,
        "description": "The parent of the quota preference. Allowed parents are \"projects/[project-id / number]\" or \"folders/[folder-id / number]\" or \"organizations/[org-id / number]\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_id": {
        "computed": true,
        "description": "The id of the quota to which the quota preference is applied. A quota id is unique in the service.\nExample: 'CPUS-per-project-region'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Is the quota preference pending Google Cloud approval and fulfillment.",
        "description_kind": "plain",
        "type": "bool"
      },
      "service": {
        "computed": true,
        "description": "The name of the service to which the quota preference is applied.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Update time stamp.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "quota_config": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "The annotations map for clients to store small amounts of arbitrary data. Do not put PII or other sensitive information here. See https://google.aip.dev/128#annotations.\n\nAn object containing a list of \"key: value\" pairs. Example: '{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }'.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "granted_value": {
              "computed": true,
              "description": "Granted quota value.",
              "description_kind": "plain",
              "type": "string"
            },
            "preferred_value": {
              "description": "The preferred value. Must be greater than or equal to -1. If set to -1, it means the value is \"unlimited\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "request_origin": {
              "computed": true,
              "description": "The origin of the quota preference request.",
              "description_kind": "plain",
              "type": "string"
            },
            "state_detail": {
              "computed": true,
              "description": "Optional details about the state of this quota preference.",
              "description_kind": "plain",
              "type": "string"
            },
            "trace_id": {
              "computed": true,
              "description": "The trace id that the Google Cloud uses to provision the requested quota. This trace id may be used by the client to contact Cloud support to track the state of a quota preference request. The trace id is only produced for increase requests and is unique for each request. The quota decrease requests do not have a trace id.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "The preferred quota configuration.",
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

func GoogleCloudQuotasQuotaPreferenceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudQuotasQuotaPreference), &result)
	return &result
}
