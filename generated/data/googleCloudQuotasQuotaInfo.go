package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudQuotasQuotaInfo = `{
  "block": {
    "attributes": {
      "container_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dimensions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "dimensions_infos": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "applicable_locations": [
                "list",
                "string"
              ],
              "details": [
                "list",
                [
                  "object",
                  {
                    "value": "string"
                  }
                ]
              ],
              "dimensions": [
                "map",
                "string"
              ]
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
      "is_concurrent": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "is_fixed": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "is_precise": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "metric": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metric_display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metric_unit": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quota_display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "quota_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quota_increase_eligibility": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ineligibility_reason": "string",
              "is_eligible": "bool"
            }
          ]
        ]
      },
      "refresh_interval": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_request_quota_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudQuotasQuotaInfoSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudQuotasQuotaInfo), &result)
	return &result
}
