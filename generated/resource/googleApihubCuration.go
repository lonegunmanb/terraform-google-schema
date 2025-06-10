package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApihubCuration = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time at which the curation was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "curation_id": {
        "description": "The ID to use for the curation resource, which will become the final\ncomponent of the curations's resource name. This field is optional.\n\n* If provided, the same will be used. The service will throw an error if\nthe specified ID is already used by another curation resource in the API\nhub.\n* If not provided, a system generated ID will be used.\n\nThis value should be 4-500 characters, and valid characters\nare /a-z[0-9]-_/.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the curation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the curation.",
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
      "last_execution_error_code": {
        "computed": true,
        "description": "The error code of the last execution of the curation. The error code is\npopulated only when the last execution state is failed.\nPossible values:\nERROR_CODE_UNSPECIFIED\nINTERNAL_ERROR\nUNAUTHORIZED",
        "description_kind": "plain",
        "type": "string"
      },
      "last_execution_error_message": {
        "computed": true,
        "description": "Error message describing the failure, if any, during the last execution of\nthe curation.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_execution_state": {
        "computed": true,
        "description": "The last execution state of the curation.\nPossible values:\nLAST_EXECUTION_STATE_UNSPECIFIED\nSUCCEEDED\nFAILED",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the curation.\n\nFormat:\n'projects/{project}/locations/{location}/curations/{curation}'",
        "description_kind": "plain",
        "type": "string"
      },
      "plugin_instance_actions": {
        "computed": true,
        "description": "The plugin instances and associated actions that are using the curation.\nNote: A particular curation could be used by multiple plugin instances or\nmultiple actions in a plugin instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action_id": "string",
              "plugin_instance": "string"
            }
          ]
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which the curation was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "endpoint": {
        "block": {
          "block_types": {
            "application_integration_endpoint_details": {
              "block": {
                "attributes": {
                  "trigger_id": {
                    "description": "The API trigger ID of the Application Integration workflow.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "uri": {
                    "description": "The endpoint URI should be a valid REST URI for triggering an Application\nIntegration.\nFormat:\n'https://integrations.googleapis.com/v1/{name=projects/*/locations/*/integrations/*}:execute'\nor\n'https://{location}-integrations.googleapis.com/v1/{name=projects/*/locations/*/integrations/*}:execute'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The details of the Application Integration endpoint to be triggered for\ncuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The endpoint to be triggered for curation.\nThe endpoint will be invoked with a request payload containing\nApiMetadata.\nResponse should contain curated data in the form of\nApiMetadata.",
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

func GoogleApihubCurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApihubCuration), &result)
	return &result
}
