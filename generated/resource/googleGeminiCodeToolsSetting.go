package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGeminiCodeToolsSetting = `{
  "block": {
    "attributes": {
      "code_tools_setting_id": {
        "description": "Id of the Code Tools Setting.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Create time stamp.",
        "description_kind": "plain",
        "type": "string"
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
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. Name of the resource.\nFormat:projects/{project}/locations/{location}/codeToolsSettings/{codeToolsSetting}",
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
      }
    },
    "block_types": {
      "enabled_tool": {
        "block": {
          "attributes": {
            "account_connector": {
              "description": "Link to the Dev Connect Account Connector that holds the user credentials.\nprojects/{project}/locations/{location}/accountConnectors/{account_connector_id}",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "handle": {
              "description": "Handle used to invoke the tool.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tool": {
              "description": "Link to the Tool",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "uri_override": {
              "description": "Overridden URI, if allowed by Tool.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "config": {
              "block": {
                "attributes": {
                  "key": {
                    "description": "Key of the configuration item.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value of the configuration item.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Configuration parameters for the tool.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Represents the full set of enabled tools.",
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

func GoogleGeminiCodeToolsSettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGeminiCodeToolsSetting), &result)
	return &result
}
