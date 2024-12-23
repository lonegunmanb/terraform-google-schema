package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDiscoveryEngineChatEngine = `{
  "block": {
    "attributes": {
      "chat_engine_metadata": {
        "computed": true,
        "description": "Additional information of the Chat Engine.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dialogflow_agent": "string"
            }
          ]
        ]
      },
      "collection_id": {
        "description": "The collection ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Timestamp the Engine was created at.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_store_ids": {
        "description": "The data stores associated with this engine. Multiple DataStores in the same Collection can be associated here. All listed DataStores must be 'SOLUTION_TYPE_CHAT'. Adding or removing data stores will force recreation.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description": "The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_id": {
        "description": "The ID to use for chat engine.",
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
      "industry_vertical": {
        "description": "The industry vertical that the chat engine registers. Vertical on Engine has to match vertical of the DataStore linked to the engine. Default value: \"GENERIC\" Possible values: [\"GENERIC\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique full resource name of the chat engine. Values are of the format\n'projects/{project}/locations/{location}/collections/{collection_id}/engines/{engine_id}'.\nThis field must be a UTF-8 encoded string with a length limit of 1024\ncharacters.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Timestamp the Engine was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "chat_engine_config": {
        "block": {
          "attributes": {
            "dialogflow_agent_to_link": {
              "description": "The resource name of an existing Dialogflow agent to link to this Chat Engine. Format: 'projects/\u003cProject_ID\u003e/locations/\u003cLocation_ID\u003e/agents/\u003cAgent_ID\u003e'.\nExactly one of 'agent_creation_config' or 'dialogflow_agent_to_link' must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "agent_creation_config": {
              "block": {
                "attributes": {
                  "business": {
                    "description": "Name of the company, organization or other entity that the agent represents. Used for knowledge connector LLM prompt and for knowledge search.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "default_language_code": {
                    "description": "The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language codes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "location": {
                    "description": "Agent location for Agent creation, currently supported values: global/us/eu, it needs to be the same region as the Chat Engine.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "time_zone": {
                    "description": "The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The configuration to generate the Dialogflow agent that is associated to this Engine.\nExactly one of 'agent_creation_config' or 'dialogflow_agent_to_link' must be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configurations for a chat Engine.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "common_config": {
        "block": {
          "attributes": {
            "company_name": {
              "description": "The name of the company, business or entity that is associated with the engine. Setting this may help improve LLM related features.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Common config spec that specifies the metadata of the engine.",
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

func GoogleDiscoveryEngineChatEngineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDiscoveryEngineChatEngine), &result)
	return &result
}
