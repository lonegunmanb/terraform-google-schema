package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxGenerativeSettings = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "language_code": {
        "description": "Language for this settings.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the generativeSettings.\nFormat: projects/\u003cProjectID\u003e/locations/\u003cLocationID\u003e/agents/\u003cAgentID\u003e/generativeSettings.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a flow for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "fallback_settings": {
        "block": {
          "attributes": {
            "selected_prompt": {
              "description": "Display name of the selected prompt.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "prompt_templates": {
              "block": {
                "attributes": {
                  "display_name": {
                    "description": "Prompt name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "frozen": {
                    "description": "If the flag is true, the prompt is frozen and cannot be modified by users.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "prompt_text": {
                    "description": "Prompt text that is sent to a LLM on no-match default, placeholders are filled downstream. For example: \"Here is a conversation $conversation, a response is: \"",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Stored prompts that can be selected, for example default templates like \"conservative\" or \"chatty\", or user defined ones.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Settings for Generative Fallback.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "generative_safety_settings": {
        "block": {
          "attributes": {
            "default_banned_phrase_match_strategy": {
              "description": "Optional. Default phrase match strategy for banned phrases.\nSee [PhraseMatchStrategy](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/GenerativeSettings#phrasematchstrategy) for valid values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "banned_phrases": {
              "block": {
                "attributes": {
                  "language_code": {
                    "description": "Language code of the phrase.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "text": {
                    "description": "Text input which can be used for prompt or banned phrases.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Banned phrases for generated text.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Settings for Generative Safety.\nw",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "knowledge_connector_settings": {
        "block": {
          "attributes": {
            "agent": {
              "description": "Name of the virtual agent. Used for LLM prompt. Can be left empty.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "agent_identity": {
              "description": "Identity of the agent, e.g. \"virtual agent\", \"AI assistant\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "agent_scope": {
              "description": "Agent scope, e.g. \"Example company website\", \"internal Example company website for employees\", \"manual of car owner\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "business": {
              "description": "Name of the company, organization or other entity that the agent represents. Used for knowledge connector LLM prompt and for knowledge search.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "business_description": {
              "description": "Company description, used for LLM prompt, e.g. \"a family company selling freshly roasted coffee beans\".''",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disable_data_store_fallback": {
              "description": "Whether to disable fallback to Data Store search results (in case the LLM couldn't pick a proper answer). Per default the feature is enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Settings for knowledge connector.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "llm_model_settings": {
        "block": {
          "attributes": {
            "model": {
              "description": "The selected LLM model.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prompt_text": {
              "description": "The custom prompt to use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "LLM model settings.",
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

func GoogleDialogflowCxGenerativeSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxGenerativeSettings), &result)
	return &result
}
