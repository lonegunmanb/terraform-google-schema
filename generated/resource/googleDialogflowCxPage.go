package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxPage = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "The human-readable name of the page, unique within the agent.",
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
        "description": "The language of the following fields in page:\n\nPage.entry_fulfillment.messages\nPage.entry_fulfillment.conditional_cases\nPage.event_handlers.trigger_fulfillment.messages\nPage.event_handlers.trigger_fulfillment.conditional_cases\nPage.form.parameters.fill_behavior.initial_prompt_fulfillment.messages\nPage.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases\nPage.form.parameters.fill_behavior.reprompt_event_handlers.messages\nPage.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases\nPage.transition_routes.trigger_fulfillment.messages\nPage.transition_routes.trigger_fulfillment.conditional_cases\nIf not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the page.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/pages/\u003cPage ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The flow to create a page for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transition_route_groups": {
        "description": "Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.\nIf multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -\u003e page's transition route group -\u003e flow's transition routes.\nIf multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.\nFormat:projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/transitionRouteGroups/\u003cTransitionRouteGroup ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "advanced_settings": {
        "block": {
          "block_types": {
            "dtmf_settings": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "If true, incoming audio is processed for DTMF (dual tone multi frequency) events. For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will detect the event (e.g. a \"3\" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "finish_digit": {
                    "description": "The digit that terminates a DTMF digit sequence.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_digits": {
                    "description": "Max length of DTMF digits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Define behaviors for DTMF (dual tone multi frequency). DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Exposed at the following levels:\n* Agent level\n* Flow level\n* Page level\n* Parameter level",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Hierarchical advanced settings for this page. The settings exposed at the lower level overrides the settings exposed at the higher level.\nHierarchy: Agent-\u003eFlow-\u003ePage-\u003eFulfillment/Parameter.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "entry_fulfillment": {
        "block": {
          "attributes": {
            "return_partial_responses": {
              "description": "Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tag": {
              "description": "The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "webhook": {
              "description": "The webhook to call. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "conditional_cases": {
              "block": {
                "attributes": {
                  "cases": {
                    "description": "A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.\nSee [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Conditional cases for this fulfillment.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "messages": {
              "block": {
                "attributes": {
                  "channel": {
                    "description": "The channel which the response is associated with. Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "payload": {
                    "description": "A custom, platform-specific payload.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "conversation_success": {
                    "block": {
                      "attributes": {
                        "metadata": {
                          "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Indicates that the conversation succeeded, i.e., the bot handled the issue that the customer talked to it about.\nDialogflow only uses this to determine which conversations should be counted as successful and doesn't process the metadata in this message in any way. Note that Dialogflow also considers conversations that get to the conversation end page as successful even if they don't return ConversationSuccess.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates that the conversation succeeded.\n* In a webhook response when you determine that you handled the customer issue.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "live_agent_handoff": {
                    "block": {
                      "attributes": {
                        "metadata": {
                          "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Indicates that the conversation should be handed off to a live agent.\nDialogflow only uses this to determine which conversations were handed off to a human agent for measurement purposes. What else to do with this signal is up to you and your handoff procedures.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates something went extremely wrong in the conversation.\n* In a webhook response when you determine that the customer issue can only be handled by a human.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "output_audio_text": {
                    "block": {
                      "attributes": {
                        "allow_playback_interruption": {
                          "computed": true,
                          "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "ssml": {
                          "description": "The SSML text to be synthesized. For more information, see SSML.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "text": {
                          "description": "The raw text to be synthesized.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "play_audio": {
                    "block": {
                      "attributes": {
                        "allow_playback_interruption": {
                          "computed": true,
                          "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "audio_uri": {
                          "description": "URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies an audio clip to be played by the client as part of the response.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "telephony_transfer_call": {
                    "block": {
                      "attributes": {
                        "phone_number": {
                          "description": "Transfer the call to a phone number in E.164 format.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents the signal that telles the client to transfer the phone call connected to the agent to a third-party endpoint.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "text": {
                    "block": {
                      "attributes": {
                        "allow_playback_interruption": {
                          "computed": true,
                          "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
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
                      "description": "The text response message.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The list of rich message responses to present to the user.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "set_parameter_actions": {
              "block": {
                "attributes": {
                  "parameter": {
                    "description": "Display name of the parameter.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The new JSON-encoded value of the parameter. A null value clears the parameter.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Set parameter values before executing the webhook.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The fulfillment to call when the session is entering the page.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "event_handlers": {
        "block": {
          "attributes": {
            "event": {
              "description": "The name of the event to handle.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The unique identifier of this event handler.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_flow": {
              "description": "The target flow to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_page": {
              "description": "The target page to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/pages/\u003cPage ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "trigger_fulfillment": {
              "block": {
                "attributes": {
                  "return_partial_responses": {
                    "description": "Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tag": {
                    "description": "The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "webhook": {
                    "description": "The webhook to call. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "conditional_cases": {
                    "block": {
                      "attributes": {
                        "cases": {
                          "description": "A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.\nSee [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Conditional cases for this fulfillment.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "messages": {
                    "block": {
                      "attributes": {
                        "channel": {
                          "description": "The channel which the response is associated with. Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "payload": {
                          "description": "A custom, platform-specific payload.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "conversation_success": {
                          "block": {
                            "attributes": {
                              "metadata": {
                                "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Indicates that the conversation succeeded, i.e., the bot handled the issue that the customer talked to it about.\nDialogflow only uses this to determine which conversations should be counted as successful and doesn't process the metadata in this message in any way. Note that Dialogflow also considers conversations that get to the conversation end page as successful even if they don't return ConversationSuccess.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates that the conversation succeeded.\n* In a webhook response when you determine that you handled the customer issue.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "live_agent_handoff": {
                          "block": {
                            "attributes": {
                              "metadata": {
                                "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Indicates that the conversation should be handed off to a live agent.\nDialogflow only uses this to determine which conversations were handed off to a human agent for measurement purposes. What else to do with this signal is up to you and your handoff procedures.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates something went extremely wrong in the conversation.\n* In a webhook response when you determine that the customer issue can only be handled by a human.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "output_audio_text": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "ssml": {
                                "description": "The SSML text to be synthesized. For more information, see SSML.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "text": {
                                "description": "The raw text to be synthesized.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "play_audio": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "audio_uri": {
                                "description": "URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies an audio clip to be played by the client as part of the response.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "telephony_transfer_call": {
                          "block": {
                            "attributes": {
                              "phone_number": {
                                "description": "Transfer the call to a phone number in E.164 format.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Represents the signal that telles the client to transfer the phone call connected to the agent to a third-party endpoint.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "text": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
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
                            "description": "The text response message.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The list of rich message responses to present to the user.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "set_parameter_actions": {
                    "block": {
                      "attributes": {
                        "parameter": {
                          "description": "Display name of the parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "The new JSON-encoded value of the parameter. A null value clears the parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Set parameter values before executing the webhook.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Handlers associated with the page to handle events such as webhook errors, no match or no input.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "form": {
        "block": {
          "block_types": {
            "parameters": {
              "block": {
                "attributes": {
                  "default_value": {
                    "description": "The default value of an optional parameter. If the parameter is required, the default value will be ignored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "display_name": {
                    "description": "The human-readable name of the parameter, unique within the form.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "entity_type": {
                    "description": "The entity type of the parameter.\nFormat: projects/-/locations/-/agents/-/entityTypes/\u003cSystem Entity Type ID\u003e for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/entityTypes/\u003cEntity Type ID\u003e for developer entity types.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_list": {
                    "description": "Indicates whether the parameter represents a list of values.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "redact": {
                    "description": "Indicates whether the parameter content should be redacted in log.\nIf redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "required": {
                    "description": "Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them.\nRequired parameters must be filled before form filling concludes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "advanced_settings": {
                    "block": {
                      "block_types": {
                        "dtmf_settings": {
                          "block": {
                            "attributes": {
                              "enabled": {
                                "description": "If true, incoming audio is processed for DTMF (dual tone multi frequency) events. For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will detect the event (e.g. a \"3\" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "finish_digit": {
                                "description": "The digit that terminates a DTMF digit sequence.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "max_digits": {
                                "description": "Max length of DTMF digits.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Define behaviors for DTMF (dual tone multi frequency). DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Exposed at the following levels:\n* Agent level\n* Flow level\n* Page level\n* Parameter level",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Hierarchical advanced settings for this parameter. The settings exposed at the lower level overrides the settings exposed at the higher level.\nHierarchy: Agent-\u003eFlow-\u003ePage-\u003eFulfillment/Parameter.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "fill_behavior": {
                    "block": {
                      "block_types": {
                        "initial_prompt_fulfillment": {
                          "block": {
                            "attributes": {
                              "return_partial_responses": {
                                "description": "Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "tag": {
                                "description": "The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "webhook": {
                                "description": "The webhook to call. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "conditional_cases": {
                                "block": {
                                  "attributes": {
                                    "cases": {
                                      "description": "A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.\nSee [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Conditional cases for this fulfillment.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "messages": {
                                "block": {
                                  "attributes": {
                                    "channel": {
                                      "description": "The channel which the response is associated with. Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "payload": {
                                      "description": "A custom, platform-specific payload.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "conversation_success": {
                                      "block": {
                                        "attributes": {
                                          "metadata": {
                                            "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Indicates that the conversation succeeded, i.e., the bot handled the issue that the customer talked to it about.\nDialogflow only uses this to determine which conversations should be counted as successful and doesn't process the metadata in this message in any way. Note that Dialogflow also considers conversations that get to the conversation end page as successful even if they don't return ConversationSuccess.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates that the conversation succeeded.\n* In a webhook response when you determine that you handled the customer issue.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "live_agent_handoff": {
                                      "block": {
                                        "attributes": {
                                          "metadata": {
                                            "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Indicates that the conversation should be handed off to a live agent.\nDialogflow only uses this to determine which conversations were handed off to a human agent for measurement purposes. What else to do with this signal is up to you and your handoff procedures.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates something went extremely wrong in the conversation.\n* In a webhook response when you determine that the customer issue can only be handled by a human.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "output_audio_text": {
                                      "block": {
                                        "attributes": {
                                          "allow_playback_interruption": {
                                            "computed": true,
                                            "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "ssml": {
                                            "description": "The SSML text to be synthesized. For more information, see SSML.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "text": {
                                            "description": "The raw text to be synthesized.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "play_audio": {
                                      "block": {
                                        "attributes": {
                                          "allow_playback_interruption": {
                                            "computed": true,
                                            "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "audio_uri": {
                                            "description": "URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Specifies an audio clip to be played by the client as part of the response.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "telephony_transfer_call": {
                                      "block": {
                                        "attributes": {
                                          "phone_number": {
                                            "description": "Transfer the call to a phone number in E.164 format.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Represents the signal that telles the client to transfer the phone call connected to the agent to a third-party endpoint.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "text": {
                                      "block": {
                                        "attributes": {
                                          "allow_playback_interruption": {
                                            "computed": true,
                                            "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
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
                                        "description": "The text response message.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The list of rich message responses to present to the user.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "set_parameter_actions": {
                                "block": {
                                  "attributes": {
                                    "parameter": {
                                      "description": "Display name of the parameter.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The new JSON-encoded value of the parameter. A null value clears the parameter.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Set parameter values before executing the webhook.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The fulfillment to provide the initial prompt that the agent can present to the user in order to fill the parameter.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "reprompt_event_handlers": {
                          "block": {
                            "attributes": {
                              "event": {
                                "description": "The name of the event to handle.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The unique identifier of this event handler.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "target_flow": {
                                "description": "The target flow to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "target_page": {
                                "description": "The target page to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/pages/\u003cPage ID\u003e.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "trigger_fulfillment": {
                                "block": {
                                  "attributes": {
                                    "return_partial_responses": {
                                      "description": "Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "tag": {
                                      "description": "The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "webhook": {
                                      "description": "The webhook to call. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "conditional_cases": {
                                      "block": {
                                        "attributes": {
                                          "cases": {
                                            "description": "A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.\nSee [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Conditional cases for this fulfillment.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "messages": {
                                      "block": {
                                        "attributes": {
                                          "channel": {
                                            "description": "The channel which the response is associated with. Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "description": "A custom, platform-specific payload.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "conversation_success": {
                                            "block": {
                                              "attributes": {
                                                "metadata": {
                                                  "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Indicates that the conversation succeeded, i.e., the bot handled the issue that the customer talked to it about.\nDialogflow only uses this to determine which conversations should be counted as successful and doesn't process the metadata in this message in any way. Note that Dialogflow also considers conversations that get to the conversation end page as successful even if they don't return ConversationSuccess.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates that the conversation succeeded.\n* In a webhook response when you determine that you handled the customer issue.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "live_agent_handoff": {
                                            "block": {
                                              "attributes": {
                                                "metadata": {
                                                  "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Indicates that the conversation should be handed off to a live agent.\nDialogflow only uses this to determine which conversations were handed off to a human agent for measurement purposes. What else to do with this signal is up to you and your handoff procedures.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates something went extremely wrong in the conversation.\n* In a webhook response when you determine that the customer issue can only be handled by a human.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "output_audio_text": {
                                            "block": {
                                              "attributes": {
                                                "allow_playback_interruption": {
                                                  "computed": true,
                                                  "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                                  "description_kind": "plain",
                                                  "type": "bool"
                                                },
                                                "ssml": {
                                                  "description": "The SSML text to be synthesized. For more information, see SSML.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "text": {
                                                  "description": "The raw text to be synthesized.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "play_audio": {
                                            "block": {
                                              "attributes": {
                                                "allow_playback_interruption": {
                                                  "computed": true,
                                                  "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                                  "description_kind": "plain",
                                                  "type": "bool"
                                                },
                                                "audio_uri": {
                                                  "description": "URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Specifies an audio clip to be played by the client as part of the response.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "telephony_transfer_call": {
                                            "block": {
                                              "attributes": {
                                                "phone_number": {
                                                  "description": "Transfer the call to a phone number in E.164 format.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Represents the signal that telles the client to transfer the phone call connected to the agent to a third-party endpoint.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "text": {
                                            "block": {
                                              "attributes": {
                                                "allow_playback_interruption": {
                                                  "computed": true,
                                                  "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                                  "description_kind": "plain",
                                                  "type": "bool"
                                                },
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
                                              "description": "The text response message.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "The list of rich message responses to present to the user.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "set_parameter_actions": {
                                      "block": {
                                        "attributes": {
                                          "parameter": {
                                            "description": "Display name of the parameter.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "description": "The new JSON-encoded value of the parameter. A null value clears the parameter.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Set parameter values before executing the webhook.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The handlers for parameter-level events, used to provide reprompt for the parameter or transition to a different page/flow. The supported events are:\n* sys.no-match-\u003cN\u003e, where N can be from 1 to 6\n* sys.no-match-default\n* sys.no-input-\u003cN\u003e, where N can be from 1 to 6\n* sys.no-input-default\n* sys.invalid-parameter\n[initialPromptFulfillment][initialPromptFulfillment] provides the first prompt for the parameter.\nIf the user's response does not fill the parameter, a no-match/no-input event will be triggered, and the fulfillment associated with the sys.no-match-1/sys.no-input-1 handler (if defined) will be called to provide a prompt. The sys.no-match-2/sys.no-input-2 handler (if defined) will respond to the next no-match/no-input event, and so on.\nA sys.no-match-default or sys.no-input-default handler will be used to handle all following no-match/no-input events after all numbered no-match/no-input handlers for the parameter are consumed.\nA sys.invalid-parameter handler can be defined to handle the case where the parameter values have been invalidated by webhook. For example, if the user's response fill the parameter, however the parameter was invalidated by webhook, the fulfillment associated with the sys.invalid-parameter handler (if defined) will be called to provide a prompt.\nIf the event handler for the corresponding event can't be found on the parameter, initialPromptFulfillment will be re-prompted.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Defines fill behavior for the parameter.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Parameters to collect from the user.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The form associated with the page, used for collecting parameters relevant to the page.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "knowledge_connector_settings": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether Knowledge Connector is enabled or not.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "target_flow": {
              "description": "The target flow to transition to. Format: projects/\u003cProjectID\u003e/locations/\u003cLocationID\u003e/agents/\u003cAgentID\u003e/flows/\u003cFlowID\u003e.\nThis field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_page": {
              "description": "The target page to transition to. Format: projects/\u003cProjectID\u003e/locations/\u003cLocationID\u003e/agents/\u003cAgentID\u003e/flows/\u003cFlowID\u003e/pages/\u003cPageID\u003e.\nThe page must be in the same host flow (the flow that owns this 'KnowledgeConnectorSettings').\nThis field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_store_connections": {
              "block": {
                "attributes": {
                  "data_store": {
                    "description": "The full name of the referenced data store. Formats: projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore} projects/{project}/locations/{location}/dataStores/{dataStore}",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_store_type": {
                    "description": "The type of the connected data store.\n* PUBLIC_WEB: A data store that contains public web content.\n* UNSTRUCTURED: A data store that contains unstructured private data.\n* STRUCTURED: A data store that contains structured data (for example FAQ). Possible values: [\"PUBLIC_WEB\", \"UNSTRUCTURED\", \"STRUCTURED\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "document_processing_mode": {
                    "description": "The document processing mode for the data store connection. Should only be set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is considered as DOCUMENTS, as this is the legacy mode.\n* DOCUMENTS: Documents are processed as documents.\n* CHUNKS: Documents are converted to chunks. Possible values: [\"DOCUMENTS\", \"CHUNKS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Optional. List of related data store connections.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "trigger_fulfillment": {
              "block": {
                "attributes": {
                  "enable_generative_fallback": {
                    "description": "If the flag is true, the agent will utilize LLM to generate a text response. If LLM generation fails, the defined responses in the fulfillment will be respected. This flag is only useful for fulfillments associated with no-match event handlers.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "return_partial_responses": {
                    "description": "Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tag": {
                    "description": "The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "webhook": {
                    "description": "The webhook to call. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "advanced_settings": {
                    "block": {
                      "block_types": {
                        "dtmf_settings": {
                          "block": {
                            "attributes": {
                              "enabled": {
                                "description": "If true, incoming audio is processed for DTMF (dual tone multi frequtectency) events. For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will de the event (e.g. a \"3\" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "endpointing_timeout_duration": {
                                "description": "Endpoint timeout setting for matching dtmf input to regex.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.500s\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "finish_digit": {
                                "description": "The digit that terminates a DTMF digit sequence.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "interdigit_timeout_duration": {
                                "description": "Interdigit timeout setting for matching dtmf input to regex.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.500s\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "max_digits": {
                                "description": "Max length of DTMF digits.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Define behaviors for DTMF (dual tone multi frequency). DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Exposed at the following levels:\n* Agent level\n* Flow level\n* Page level\n* Parameter level",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "logging_settings": {
                          "block": {
                            "attributes": {
                              "enable_consent_based_redaction": {
                                "description": "Enables consent-based end-user input redaction, if true, a pre-defined session parameter **$session.params.conversation-redaction** will be used to determine if the utterance should be redacted.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "enable_interaction_logging": {
                                "description": "Enables DF Interaction logging.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "enable_stackdriver_logging": {
                                "description": "Enables Google Cloud Logging.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Settings for logging. Settings for Dialogflow History, Contact Center messages, StackDriver logs, and speech logging. Exposed at the following levels:\n* Agent level",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "speech_settings": {
                          "block": {
                            "attributes": {
                              "endpointer_sensitivity": {
                                "description": "Sensitivity of the speech model that detects the end of speech. Scale from 0 to 100.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "models": {
                                "description": "Mapping from language to Speech-to-Text model. The mapped Speech-to-Text model will be selected for requests from its corresponding language. For more information, see [Speech models](https://cloud.google.com/dialogflow/cx/docs/concept/speech-models).\nAn object containing a list of **\"key\": value** pairs. Example: **{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }**.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "no_speech_timeout": {
                                "description": "Timeout before detecting no speech.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.500s\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "use_timeout_based_endpointing": {
                                "description": "Use timeout based endpointing, interpreting endpointer sensitivity as seconds of timeout value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Settings for speech to text detection. Exposed at the following levels:\n* Agent level\n* Flow level\n* Page level\n* Parameter level",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Hierarchical advanced settings for agent/flow/page/fulfillment/parameter. Settings exposed at lower level overrides the settings exposed at higher level. Overriding occurs at the sub-setting level. For example, the playbackInterruptionSettings at fulfillment level only overrides the playbackInterruptionSettings at the agent level, leaving other settings at the agent level unchanged.\nDTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel.\nHierarchy: Agent-\u003eFlow-\u003ePage-\u003eFulfillment/Parameter.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "conditional_cases": {
                    "block": {
                      "attributes": {
                        "cases": {
                          "description": "A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.\nSee [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Conditional cases for this fulfillment.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "messages": {
                    "block": {
                      "attributes": {
                        "channel": {
                          "description": "The channel which the response is associated with. Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "end_interaction": {
                          "computed": true,
                          "description": "This type has no fields.\nIndicates that interaction with the Dialogflow agent has ended. This message is generated by Dialogflow only and not supposed to be defined by the user.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            [
                              "object",
                              {}
                            ]
                          ]
                        },
                        "mixed_audio": {
                          "computed": true,
                          "description": "Represents an audio message that is composed of both segments synthesized from the Dialogflow agent prompts and ones hosted externally at the specified URIs. The external URIs are specified via playAudio. This message is generated by Dialogflow only and not supposed to be defined by the user.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            [
                              "object",
                              {
                                "segments": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "allow_playback_interruption": "bool",
                                      "audio": "string",
                                      "uri": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        },
                        "payload": {
                          "description": "Returns a response containing a custom, platform-specific payload.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "conversation_success": {
                          "block": {
                            "attributes": {
                              "metadata": {
                                "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Indicates that the conversation succeeded, i.e., the bot handled the issue that the customer talked to it about.\nDialogflow only uses this to determine which conversations should be counted as successful and doesn't process the metadata in this message in any way. Note that Dialogflow also considers conversations that get to the conversation end page as successful even if they don't return ConversationSuccess.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates that the conversation succeeded.\n* In a webhook response when you determine that you handled the customer issue.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "knowledge_info_card": {
                          "block": {
                            "description": "This type has no fields.\nRepresents info card response. If the response contains generative knowledge prediction, Dialogflow will return a payload with Infobot Messenger compatible info card.\nOtherwise, the info card response is skipped.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "live_agent_handoff": {
                          "block": {
                            "attributes": {
                              "metadata": {
                                "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Indicates that the conversation should be handed off to a live agent.\nDialogflow only uses this to determine which conversations were handed off to a human agent for measurement purposes. What else to do with this signal is up to you and your handoff procedures.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates something went extremely wrong in the conversation.\n* In a webhook response when you determine that the customer issue can only be handled by a human.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "output_audio_text": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "ssml": {
                                "description": "The SSML text to be synthesized. For more information, see SSML.\nThis field is part of a union field 'source': Only one of 'text' or 'ssml' may be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "text": {
                                "description": "The raw text to be synthesized.\nThis field is part of a union field 'source': Only one of 'text' or 'ssml' may be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "play_audio": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "audio_uri": {
                                "description": "URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies an audio clip to be played by the client as part of the response.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "telephony_transfer_call": {
                          "block": {
                            "attributes": {
                              "phone_number": {
                                "description": "Transfer the call to a phone number in E.164 format.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Represents the signal that telles the client to transfer the phone call connected to the agent to a third-party endpoint.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "text": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "text": {
                                "description": "A collection of text response variants. If multiple variants are defined, only one text response variant is returned at runtime.\nrequired: true",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "The text response message.\nThis field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The list of rich message responses to present to the user.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "set_parameter_actions": {
                    "block": {
                      "attributes": {
                        "parameter": {
                          "description": "Display name of the parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "The new JSON-encoded value of the parameter. A null value clears the parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Set parameter values before executing the webhook.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The fulfillment to be triggered.\nWhen the answers from the Knowledge Connector are selected by Dialogflow, you can utitlize the request scoped parameter $request.knowledge.answers (contains up to the 5 highest confidence answers) and $request.knowledge.questions (contains the corresponding questions) to construct the fulfillment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Knowledge connector configuration.",
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
      },
      "transition_routes": {
        "block": {
          "attributes": {
            "condition": {
              "description": "The condition to evaluate against form parameters or session parameters.\nAt least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "intent": {
              "description": "The unique identifier of an Intent.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/intents/\u003cIntent ID\u003e. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The unique identifier of this transition route.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_flow": {
              "description": "The target flow to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_page": {
              "description": "The target page to transition to.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/pages/\u003cPage ID\u003e.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "trigger_fulfillment": {
              "block": {
                "attributes": {
                  "return_partial_responses": {
                    "description": "Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tag": {
                    "description": "The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "webhook": {
                    "description": "The webhook to call. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "conditional_cases": {
                    "block": {
                      "attributes": {
                        "cases": {
                          "description": "A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.\nSee [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Conditional cases for this fulfillment.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "messages": {
                    "block": {
                      "attributes": {
                        "channel": {
                          "description": "The channel which the response is associated with. Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "payload": {
                          "description": "A custom, platform-specific payload.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "conversation_success": {
                          "block": {
                            "attributes": {
                              "metadata": {
                                "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Indicates that the conversation succeeded, i.e., the bot handled the issue that the customer talked to it about.\nDialogflow only uses this to determine which conversations should be counted as successful and doesn't process the metadata in this message in any way. Note that Dialogflow also considers conversations that get to the conversation end page as successful even if they don't return ConversationSuccess.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates that the conversation succeeded.\n* In a webhook response when you determine that you handled the customer issue.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "live_agent_handoff": {
                          "block": {
                            "attributes": {
                              "metadata": {
                                "description": "Custom metadata. Dialogflow doesn't impose any structure on this.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Indicates that the conversation should be handed off to a live agent.\nDialogflow only uses this to determine which conversations were handed off to a human agent for measurement purposes. What else to do with this signal is up to you and your handoff procedures.\nYou may set this, for example:\n* In the entryFulfillment of a Page if entering the page indicates something went extremely wrong in the conversation.\n* In a webhook response when you determine that the customer issue can only be handled by a human.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "output_audio_text": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "ssml": {
                                "description": "The SSML text to be synthesized. For more information, see SSML.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "text": {
                                "description": "The raw text to be synthesized.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "play_audio": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "audio_uri": {
                                "description": "URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies an audio clip to be played by the client as part of the response.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "telephony_transfer_call": {
                          "block": {
                            "attributes": {
                              "phone_number": {
                                "description": "Transfer the call to a phone number in E.164 format.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Represents the signal that telles the client to transfer the phone call connected to the agent to a third-party endpoint.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "text": {
                          "block": {
                            "attributes": {
                              "allow_playback_interruption": {
                                "computed": true,
                                "description": "Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
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
                            "description": "The text response message.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The list of rich message responses to present to the user.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "set_parameter_actions": {
                    "block": {
                      "attributes": {
                        "parameter": {
                          "description": "Display name of the parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "The new JSON-encoded value of the parameter. A null value clears the parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Set parameter values before executing the webhook.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The fulfillment to call when the condition is satisfied. At least one of triggerFulfillment and target must be specified. When both are defined, triggerFulfillment is executed first.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.\nWhen we are in a certain page, the TransitionRoutes are evalauted in the following order:\nTransitionRoutes defined in the page with intent specified.\nTransitionRoutes defined in the transition route groups with intent specified.\nTransitionRoutes defined in flow with intent specified.\nTransitionRoutes defined in the transition route groups with intent specified.\nTransitionRoutes defined in the page with only condition specified.\nTransitionRoutes defined in the transition route groups with only condition specified.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDialogflowCxPageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxPage), &result)
	return &result
}
