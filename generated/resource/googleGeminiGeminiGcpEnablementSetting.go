package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGeminiGeminiGcpEnablementSetting = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Create time stamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "disable_web_grounding": {
        "deprecated": true,
        "description": "Whether web grounding should be disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "enable_customer_data_sharing": {
        "description": "Whether customer data sharing should be enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gemini_gcp_enablement_setting_id": {
        "description": "Id of the Gemini Gcp Enablement setting.",
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
      "labels": {
        "description": "Labels as key value pairs.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. Name of the resource.\nFormat:projects/{project}/locations/{location}/geminiGcpEnablementSettings/{geminiGcpEnablementSetting}",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Update time stamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "web_grounding_type": {
        "description": "Web grounding type.\nPossible values:\nGROUNDING_WITH_GOOGLE_SEARCH\nWEB_GROUNDING_FOR_ENTERPRISE",
        "description_kind": "plain",
        "optional": true,
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

func GoogleGeminiGeminiGcpEnablementSettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGeminiGeminiGcpEnablementSetting), &result)
	return &result
}
