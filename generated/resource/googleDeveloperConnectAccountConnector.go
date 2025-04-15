package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDeveloperConnectAccountConnector = `{
  "block": {
    "attributes": {
      "account_connector_id": {
        "description": "Required. The ID to use for the AccountConnector, which will become the final\ncomponent of the AccountConnector's resource name. Its format should adhere\nto https://google.aip.dev/122#resource-id-segments Names must be unique\nper-project per-location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "annotations": {
        "description": "Optional. Allows users to store small amounts of arbitrary data.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The timestamp when the userConnection was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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
        "description": "Optional. Labels as key value pairs\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the userConnection, in the format\n'projects/{project}/locations/{location}/accountConnectors/{account_connector_id}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "oauth_start_uri": {
        "computed": true,
        "description": "Output only. Start OAuth flow by clicking on this URL.",
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
        "description": "Output only. The timestamp when the userConnection was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "provider_oauth_config": {
        "block": {
          "attributes": {
            "scopes": {
              "description": "Required. User selected scopes to apply to the Oauth config\nIn the event of changing scopes, user records under AccountConnector will\nbe deleted and users will re-auth again.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "system_provider_id": {
              "description": "List of providers that are owned by Developer Connect.\n\nPossible values:\nGITHUB\nGITLAB\nGOOGLE\nSENTRY\nROVO",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "ProviderOAuthConfig is the OAuth config for a provider.",
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

func GoogleDeveloperConnectAccountConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDeveloperConnectAccountConnector), &result)
	return &result
}
