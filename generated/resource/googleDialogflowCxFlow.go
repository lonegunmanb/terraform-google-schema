package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxFlow = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the flow.",
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
      "is_default_start_flow": {
        "description": "Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent. When you create an agent, the Default Start Flow is created automatically.\nThe Default Start Flow cannot be deleted; deleting the 'google_dialogflow_cx_flow' resource does nothing to the underlying GCP resources.\n\n~\u003e Avoid having multiple 'google_dialogflow_cx_flow' resources linked to the same agent with 'is_default_start_flow = true' because they will compete to control a single Default Start Flow resource in GCP.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "language_code": {
        "description": "The language of the following fields in flow:\nFlow.event_handlers.trigger_fulfillment.messages\nFlow.event_handlers.trigger_fulfillment.conditional_cases\nFlow.transition_routes.trigger_fulfillment.messages\nFlow.transition_routes.trigger_fulfillment.conditional_cases\nIf not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the flow.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a flow for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transition_route_groups": {
        "description": "A flow's transition route group serve two purposes:\nThey are responsible for matching the user's first utterances in the flow.\nThey are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route groups defined in the page have higher priority than those defined in the flow.\nFormat:projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e/transitionRouteGroups/\u003cTransitionRouteGroup ID\u003e.",
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
            "audio_export_gcs_destination": {
              "block": {
                "attributes": {
                  "uri": {
                    "description": "The Google Cloud Storage URI for the exported objects. Whether a full object name, or just a prefix, its usage depends on the Dialogflow operation.\nFormat: gs://bucket/object-name-or-prefix",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "If present, incoming audio is exported by Dialogflow to the configured Google Cloud Storage destination. Exposed at the following levels:\n* Agent level\n* Flow level",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
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
                    "description": "Timeout before detecting no speech.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
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
          "description": "Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at the higher level.\nHierarchy: Agent-\u003eFlow-\u003ePage-\u003eFulfillment/Parameter.",
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
          "description": "A flow's event handlers serve two purposes:\nThey are responsible for handling events (e.g. no match, webhook errors) in the flow.\nThey are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow.\nUnlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.",
          "description_kind": "plain"
        },
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
      "nlu_settings": {
        "block": {
          "attributes": {
            "classification_threshold": {
              "description": "To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.\nIf the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "model_training_mode": {
              "description": "Indicates NLU model training mode.\n* MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.\n* MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train. Possible values: [\"MODEL_TRAINING_MODE_AUTOMATIC\", \"MODEL_TRAINING_MODE_MANUAL\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_type": {
              "description": "Indicates the type of NLU model.\n* MODEL_TYPE_STANDARD: Use standard NLU model.\n* MODEL_TYPE_ADVANCED: Use advanced NLU model. Possible values: [\"MODEL_TYPE_STANDARD\", \"MODEL_TYPE_ADVANCED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "NLU related settings of the flow.",
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
          "description": "A flow's transition routes serve two purposes:\nThey are responsible for matching the user's first utterances in the flow.\nThey are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as the user saying \"help\" or \"can I talk to a human?\", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow.\n\nTransitionRoutes are evalauted in the following order:\n  TransitionRoutes with intent specified.\n  TransitionRoutes with only condition specified.\n  TransitionRoutes with intent specified are inherited by pages in the flow.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDialogflowCxFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxFlow), &result)
	return &result
}
