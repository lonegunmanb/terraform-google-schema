package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudQuotasQuotaInfos = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quota_infos": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "container_type": "string",
              "dimensions": [
                "list",
                "string"
              ],
              "dimensions_infos": [
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
              ],
              "is_concurrent": "bool",
              "is_fixed": "bool",
              "is_precise": "bool",
              "metric": "string",
              "metric_display_name": "string",
              "metric_unit": "string",
              "name": "string",
              "quota_display_name": "string",
              "quota_id": "string",
              "quota_increase_eligibility": [
                "list",
                [
                  "object",
                  {
                    "ineligibility_reason": "string",
                    "is_eligible": "bool"
                  }
                ]
              ],
              "refresh_interval": "string",
              "service": "string",
              "service_request_quota_uri": "string"
            }
          ]
        ]
      },
      "service": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudQuotasQuotaInfosSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudQuotasQuotaInfos), &result)
	return &result
}
