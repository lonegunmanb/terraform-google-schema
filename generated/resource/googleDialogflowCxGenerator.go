package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxGenerator = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "The human-readable name of the generator, unique within the agent.",
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
      "language_code": {
        "description": "The language to create generators for the following fields:\n* Generator.prompt_text.text\nIf not specified, the agent's default language is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the Generator.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/generators/\u003cGenerator ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a Generator for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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
          "description": "The LLM model settings.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "model_parameter": {
        "block": {
          "attributes": {
            "max_decode_steps": {
              "description": "The maximum number of tokens to generate.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "temperature": {
              "description": "The temperature used for sampling. Temperature sampling occurs after both topP and topK have been applied.\nValid range: [0.0, 1.0] Low temperature = less random. High temperature = more random.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "top_k": {
              "description": "If set, the sampling process in each step is limited to the topK tokens with highest probabilities.\nValid range: [1, 40] or 1000+. Small topK = less random. Large topK = more random.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "top_p": {
              "description": "If set, only the tokens comprising the top topP probability mass are considered.\nIf both topP and topK are set, topP will be used for further refining candidates selected with topK.\nValid range: (0.0, 1.0]. Small topP = less random. Large topP = more random.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Parameters passed to the LLM to configure its behavior.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "placeholders": {
        "block": {
          "attributes": {
            "id": {
              "description": "Unique ID used to map custom placeholder to parameters in fulfillment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "Custom placeholder value in the prompt text.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "List of custom placeholders in the prompt text.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "prompt_text": {
        "block": {
          "attributes": {
            "text": {
              "description": "Text input which can be used for prompt or banned phrases.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Prompt for the LLM model.",
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

func GoogleDialogflowCxGeneratorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxGenerator), &result)
	return &result
}
