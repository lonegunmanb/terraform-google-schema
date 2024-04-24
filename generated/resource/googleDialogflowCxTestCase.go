package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxTestCase = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "When the test was created. A timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the test case, unique within the agent. Limit of 200 characters.",
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
      "last_test_result": {
        "computed": true,
        "description": "The latest test result.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "conversation_turns": [
                "list",
                [
                  "object",
                  {
                    "user_input": [
                      "list",
                      [
                        "object",
                        {
                          "enable_sentiment_analysis": "bool",
                          "injected_parameters": "string",
                          "input": [
                            "list",
                            [
                              "object",
                              {
                                "dtmf": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "digits": "string",
                                      "finish_digit": "string"
                                    }
                                  ]
                                ],
                                "event": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "event": "string"
                                    }
                                  ]
                                ],
                                "language_code": "string",
                                "text": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "text": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "is_webhook_enabled": "bool"
                        }
                      ]
                    ],
                    "virtual_agent_output": [
                      "list",
                      [
                        "object",
                        {
                          "current_page": [
                            "list",
                            [
                              "object",
                              {
                                "display_name": "string",
                                "name": "string"
                              }
                            ]
                          ],
                          "differences": [
                            "list",
                            [
                              "object",
                              {
                                "description": "string",
                                "type": "string"
                              }
                            ]
                          ],
                          "session_parameters": "string",
                          "status": [
                            "list",
                            [
                              "object",
                              {
                                "code": "number",
                                "details": "string",
                                "message": "string"
                              }
                            ]
                          ],
                          "text_responses": [
                            "list",
                            [
                              "object",
                              {
                                "text": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "triggered_intent": [
                            "list",
                            [
                              "object",
                              {
                                "display_name": "string",
                                "name": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "environment": "string",
              "name": "string",
              "test_result": "string",
              "test_time": "string"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the test case.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/testCases/\u003cTestCase ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "notes": {
        "description": "Additional freeform notes about the test case. Limit of 400 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent": {
        "description": "The agent to create the test case for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description": "Tags are short descriptions that users may apply to test cases for organizational and filtering purposes.\nEach tag should start with \"#\" and has a limit of 30 characters",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "test_case_conversation_turns": {
        "block": {
          "block_types": {
            "user_input": {
              "block": {
                "attributes": {
                  "enable_sentiment_analysis": {
                    "description": "Whether sentiment analysis is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "injected_parameters": {
                    "description": "Parameters that need to be injected into the conversation during intent detection.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_webhook_enabled": {
                    "description": "If webhooks should be allowed to trigger in response to the user utterance. Often if parameters are injected, webhooks should not be enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "input": {
                    "block": {
                      "attributes": {
                        "language_code": {
                          "description": "The language of the input. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes.\nNote that queries in the same session do not necessarily need to specify the same language.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "dtmf": {
                          "block": {
                            "attributes": {
                              "digits": {
                                "description": "The dtmf digits.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "finish_digit": {
                                "description": "The finish digit (if any).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The DTMF event to be handled.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "event": {
                          "block": {
                            "attributes": {
                              "event": {
                                "description": "Name of the event.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The event to be triggered.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "text": {
                          "block": {
                            "attributes": {
                              "text": {
                                "description": "The natural language text to be processed. Text length must not exceed 256 characters.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The natural language text to be processed.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "User input. Supports text input, event input, dtmf input in the test case.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The user input.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "virtual_agent_output": {
              "block": {
                "attributes": {
                  "session_parameters": {
                    "description": "The session parameters available to the bot at this point.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "current_page": {
                    "block": {
                      "attributes": {
                        "display_name": {
                          "computed": true,
                          "description": "The human-readable name of the page, unique within the flow.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "name": {
                          "description": "The unique identifier of the page.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/pages/\u003cPage ID\u003e.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The [Page](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows.pages#Page) on which the utterance was spoken.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "text_responses": {
                    "block": {
                      "attributes": {
                        "text": {
                          "description": "A collection of text responses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "The text responses from the agent for the turn.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "triggered_intent": {
                    "block": {
                      "attributes": {
                        "display_name": {
                          "computed": true,
                          "description": "The human-readable name of the intent, unique within the agent.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "name": {
                          "description": "The unique identifier of the intent.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/intents/\u003cIntent ID\u003e.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The [Intent](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.intents#Intent) that triggered the response.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The virtual agent output.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The conversation turns uttered when the test case was created, in chronological order. These include the canonical set of agent utterances that should occur when the agent is working properly.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "test_config": {
        "block": {
          "attributes": {
            "flow": {
              "description": "Flow name to start the test case with.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.\nOnly one of flow and page should be set to indicate the starting point of the test case. If neither is set, the test case will start with start page on the default start flow.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "page": {
              "description": "The page to start the test case with.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/pages/\u003cPage ID\u003e.\nOnly one of flow and page should be set to indicate the starting point of the test case. If neither is set, the test case will start with start page on the default start flow.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tracking_parameters": {
              "description": "Session parameters to be compared when calculating differences.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Config for the test case.",
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

func GoogleDialogflowCxTestCaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxTestCase), &result)
	return &result
}
