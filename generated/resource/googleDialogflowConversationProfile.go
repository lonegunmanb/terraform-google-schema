package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowConversationProfile = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "Required. Human readable name for this profile. Max length 1024 bytes.",
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
        "computed": true,
        "description": "Language code for the conversation profile. This should be a BCP-47 language tag.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "desc",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "name",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_settings": {
        "description": "Name of the CX SecuritySettings reference for the agent.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_zone": {
        "description": "The time zone of this conversational profile.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "automated_agent_config": {
        "block": {
          "attributes": {
            "agent": {
              "description": "ID of the Dialogflow agent environment to use.\nExpects the format \"projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agent/environments/\u003cEnvironmentID\u003e\"",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "session_ttl": {
              "description": "Configure lifetime of the Dialogflow session.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration for an automated agent to use with this profile",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "human_agent_assistant_config": {
        "block": {
          "block_types": {
            "end_user_suggestion_config": {
              "block": {
                "attributes": {
                  "disable_high_latency_features_sync_delivery": {
                    "description": "When disableHighLatencyFeaturesSyncDelivery is true and using the AnalyzeContent API, we will not deliver the responses from high latency features in the API response. The humanAgentAssistantConfig.notification_config must be configured and enableEventBasedSuggestion must be set to true to receive the responses from high latency features in Pub/Sub. High latency feature(s): KNOWLEDGE_ASSIST",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "generators": {
                    "description": "List of various generator resource names used in the conversation profile.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "group_suggestion_responses": {
                    "description": "If groupSuggestionResponses is false, and there are multiple featureConfigs in event based suggestion or StreamingAnalyzeContent, we will try to deliver suggestions to customers as soon as we get new suggestion. Different type of suggestions based on the same context will be in separate Pub/Sub event or StreamingAnalyzeContentResponse.\n\nIf groupSuggestionResponses set to true. All the suggestions to the same participant based on the same context will be grouped into a single Pub/Sub event or StreamingAnalyzeContentResponse.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "feature_configs": {
                    "block": {
                      "attributes": {
                        "disable_agent_query_logging": {
                          "description": "Disable the logging of search queries sent by human agents. It can prevent those queries from being stored at answer records.\nThis feature is only supported for types: KNOWLEDGE_SEARCH.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_conversation_augmented_query": {
                          "description": "Enable including conversation context during query answer generation.\nThis feature is only supported for types: KNOWLEDGE_SEARCH.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_event_based_suggestion": {
                          "description": "Automatically iterates all participants and tries to compile suggestions.\nThis feature is only supported for types: ARTICLE_SUGGESTION, FAQ, DIALOGFLOW_ASSIST, KNOWLEDGE_ASSIST.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_query_suggestion_only": {
                          "description": "Enable query suggestion only.\nThis feature is only supported for types: KNOWLEDGE_ASSIST",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_query_suggestion_when_no_answer": {
                          "description": "Enable query suggestion even if we can't find its answer. By default, queries are suggested only if we find its answer.\nThis feature is only supported for types: KNOWLEDGE_ASSIST.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "conversation_model_config": {
                          "block": {
                            "attributes": {
                              "baseline_model_version": {
                                "description": "Version of current baseline model. It will be ignored if model is set. Valid versions are: Article Suggestion baseline model: - 0.9 - 1.0 (default) Summarization baseline model: - 1.0",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "model": {
                                "description": "Conversation model resource name. Format: projects/\u003cProject ID\u003e/conversationModels/\u003cModel ID\u003e.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Configs of custom conversation model.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "conversation_process_config": {
                          "block": {
                            "attributes": {
                              "recent_sentences_count": {
                                "description": "Number of recent non-small-talk sentences to use as context for article and FAQ suggestion",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Config to process conversation.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "query_config": {
                          "block": {
                            "attributes": {
                              "confidence_threshold": {
                                "description": "Confidence threshold of query result.\nThis feature is only supported for types: ARTICLE_SUGGESTION, FAQ, SMART_REPLY, SMART_COMPOSE, KNOWLEDGE_SEARCH, KNOWLEDGE_ASSIST, ENTITY_EXTRACTION.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "max_results": {
                                "description": "Maximum number of results to return.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "context_filter_settings": {
                                "block": {
                                  "attributes": {
                                    "drop_handoff_messages": {
                                      "description": "If set to true, the last message from virtual agent (hand off message) and the message before it (trigger message of hand off) are dropped.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "drop_ivr_messages": {
                                      "description": "If set to true, all messages from ivr stage are dropped.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "drop_virtual_agent_messages": {
                                      "description": "If set to true, all messages from virtual agent are dropped.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Determines how recent conversation context is filtered when generating suggestions. If unspecified, no messages will be dropped.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "dialogflow_query_source": {
                                "block": {
                                  "attributes": {
                                    "agent": {
                                      "description": "he name of a Dialogflow virtual agent used for end user side intent detection and suggestion. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agent.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "human_agent_side_config": {
                                      "block": {
                                        "attributes": {
                                          "agent": {
                                            "description": "The name of a dialogflow virtual agent used for intent detection and suggestion triggered by human agent. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The Dialogflow assist configuration for human agent.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Query from Dialogflow agent.\nThis feature is supported for types: DIALOGFLOW_ASSIST.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "document_query_source": {
                                "block": {
                                  "attributes": {
                                    "documents": {
                                      "description": "Knowledge documents to query from. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/knowledgeBases/\u003cKnowledgeBase ID\u003e/documents/\u003cDocument ID\u003e.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "Query from knowledge base document.\nThis feature is supported for types: SMART_REPLY, SMART_COMPOSE.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "knowledge_base_query_source": {
                                "block": {
                                  "attributes": {
                                    "knowledge_bases": {
                                      "description": "Knowledge bases to query. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "Query from knowledgebase.\nThis feature is only supported for types: ARTICLE_SUGGESTION, FAQ.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "sections": {
                                "block": {
                                  "attributes": {
                                    "section_types": {
                                      "description": "The selected sections chosen to return when requesting a summary of a conversation\nIf not provided the default selection will be \"{SITUATION, ACTION, RESULT}\". Possible values: [\"SECTION_TYPE_UNSPECIFIED\", \"SITUATION\", \"ACTION\", \"RESOLUTION\", \"REASON_FOR_CANCELLATION\", \"CUSTOMER_SATISFACTION\", \"ENTITIES\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "he customized sections chosen to return when requesting a summary of a conversation.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Configs of query.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "suggestion_feature": {
                          "block": {
                            "attributes": {
                              "type": {
                                "description": "Type of Human Agent Assistant API feature to request.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The suggestion feature.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "suggestion_trigger_settings": {
                          "block": {
                            "attributes": {
                              "no_small_talk": {
                                "description": "Do not trigger if last utterance is small talk.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "only_end_user": {
                                "description": "Only trigger suggestion if participant role of last utterance is END_USER.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Settings of suggestion trigger.\nThis feature is only supported for types: ARTICLE_SUGGESTION, FAQ.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Configuration of different suggestion features. One feature can have only one config.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Configuration for agent assistance of end user participant.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "human_agent_suggestion_config": {
              "block": {
                "attributes": {
                  "disable_high_latency_features_sync_delivery": {
                    "description": "When disableHighLatencyFeaturesSyncDelivery is true and using the AnalyzeContent API, we will not deliver the responses from high latency features in the API response. The humanAgentAssistantConfig.notification_config must be configured and enableEventBasedSuggestion must be set to true to receive the responses from high latency features in Pub/Sub. High latency feature(s): KNOWLEDGE_ASSIST",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "generators": {
                    "description": "List of various generator resource names used in the conversation profile.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "group_suggestion_responses": {
                    "description": "If groupSuggestionResponses is false, and there are multiple featureConfigs in event based suggestion or StreamingAnalyzeContent, we will try to deliver suggestions to customers as soon as we get new suggestion. Different type of suggestions based on the same context will be in separate Pub/Sub event or StreamingAnalyzeContentResponse.\n\nIf groupSuggestionResponses set to true. All the suggestions to the same participant based on the same context will be grouped into a single Pub/Sub event or StreamingAnalyzeContentResponse.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "feature_configs": {
                    "block": {
                      "attributes": {
                        "disable_agent_query_logging": {
                          "description": "Disable the logging of search queries sent by human agents. It can prevent those queries from being stored at answer records.\nThis feature is only supported for types: KNOWLEDGE_SEARCH.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_conversation_augmented_query": {
                          "description": "Enable including conversation context during query answer generation.\nThis feature is only supported for types: KNOWLEDGE_SEARCH.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_event_based_suggestion": {
                          "description": "Automatically iterates all participants and tries to compile suggestions.\nThis feature is only supported for types: ARTICLE_SUGGESTION, FAQ, DIALOGFLOW_ASSIST, KNOWLEDGE_ASSIST.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_query_suggestion_only": {
                          "description": "Enable query suggestion only.\nThis feature is only supported for types: KNOWLEDGE_ASSIST",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_query_suggestion_when_no_answer": {
                          "description": "Enable query suggestion even if we can't find its answer. By default, queries are suggested only if we find its answer.\nThis feature is only supported for types: KNOWLEDGE_ASSIST.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "conversation_model_config": {
                          "block": {
                            "attributes": {
                              "baseline_model_version": {
                                "description": "Version of current baseline model. It will be ignored if model is set. Valid versions are: Article Suggestion baseline model: - 0.9 - 1.0 (default) Summarization baseline model: - 1.0",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "model": {
                                "description": "Conversation model resource name. Format: projects/\u003cProject ID\u003e/conversationModels/\u003cModel ID\u003e.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Configs of custom conversation model.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "conversation_process_config": {
                          "block": {
                            "attributes": {
                              "recent_sentences_count": {
                                "description": "Number of recent non-small-talk sentences to use as context for article and FAQ suggestion",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Config to process conversation.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "query_config": {
                          "block": {
                            "attributes": {
                              "confidence_threshold": {
                                "description": "Confidence threshold of query result.\nThis feature is only supported for types: ARTICLE_SUGGESTION, FAQ, SMART_REPLY, SMART_COMPOSE, KNOWLEDGE_SEARCH, KNOWLEDGE_ASSIST, ENTITY_EXTRACTION.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "max_results": {
                                "description": "Maximum number of results to return.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "context_filter_settings": {
                                "block": {
                                  "attributes": {
                                    "drop_handoff_messages": {
                                      "description": "If set to true, the last message from virtual agent (hand off message) and the message before it (trigger message of hand off) are dropped.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "drop_ivr_messages": {
                                      "description": "If set to true, all messages from ivr stage are dropped.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "drop_virtual_agent_messages": {
                                      "description": "If set to true, all messages from virtual agent are dropped.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Determines how recent conversation context is filtered when generating suggestions. If unspecified, no messages will be dropped.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "dialogflow_query_source": {
                                "block": {
                                  "attributes": {
                                    "agent": {
                                      "description": "he name of a Dialogflow virtual agent used for end user side intent detection and suggestion. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agent.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "human_agent_side_config": {
                                      "block": {
                                        "attributes": {
                                          "agent": {
                                            "description": "The name of a dialogflow virtual agent used for intent detection and suggestion triggered by human agent. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The Dialogflow assist configuration for human agent.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Query from Dialogflow agent.\nThis feature is supported for types: DIALOGFLOW_ASSIST.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "sections": {
                                "block": {
                                  "attributes": {
                                    "section_types": {
                                      "description": "The selected sections chosen to return when requesting a summary of a conversation\nIf not provided the default selection will be \"{SITUATION, ACTION, RESULT}\". Possible values: [\"SECTION_TYPE_UNSPECIFIED\", \"SITUATION\", \"ACTION\", \"RESOLUTION\", \"REASON_FOR_CANCELLATION\", \"CUSTOMER_SATISFACTION\", \"ENTITIES\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "he customized sections chosen to return when requesting a summary of a conversation.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Configs of query.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "suggestion_feature": {
                          "block": {
                            "attributes": {
                              "type": {
                                "description": "Type of Human Agent Assistant API feature to request.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The suggestion feature.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "suggestion_trigger_settings": {
                          "block": {
                            "attributes": {
                              "no_small_talk": {
                                "description": "Do not trigger if last utterance is small talk.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "only_end_user": {
                                "description": "Only trigger suggestion if participant role of last utterance is END_USER.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Settings of suggestion trigger.\nThis feature is only supported for types: ARTICLE_SUGGESTION, FAQ.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Configuration of different suggestion features. One feature can have only one config.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Configuration for agent assistance of human agent participant.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "message_analysis_config": {
              "block": {
                "attributes": {
                  "enable_entity_extraction": {
                    "description": "Enable entity extraction in conversation messages on agent assist stage.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_sentiment_analysis": {
                    "description": "Enable sentiment analysis in conversation messages on agent assist stage. Sentiment analysis inspects user input and identifies the prevailing subjective opinion, especially to determine a user's attitude as positive, negative, or neutral.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "desc",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "notification_config": {
              "block": {
                "attributes": {
                  "message_format": {
                    "description": "Format of the message Possible values: [\"MESSAGE_FORMAT_UNSPECIFIED\", \"PROTO\", \"JSON\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "topic": {
                    "description": "Name of the Pub/Sub topic to publish conversation events",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Pub/Sub topic on which to publish new agent assistant events.\nExpects the format \"projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/topics/\u003cTopic ID\u003e\"",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connecting to a live agent",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "human_agent_handoff_config": {
        "block": {
          "block_types": {
            "live_person_config": {
              "block": {
                "attributes": {
                  "account_number": {
                    "description": "Account number of the LivePerson account to connect.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Config for using LivePerson.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines the hand off to a live agent, typically on which external agent service provider to connect to a conversation.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "logging_config": {
        "block": {
          "attributes": {
            "enable_stackdriver_logging": {
              "description": "Whether to log conversation events",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Defines logging behavior for conversation lifecycle events.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "new_message_event_notification_config": {
        "block": {
          "attributes": {
            "message_format": {
              "description": "Format of the message Possible values: [\"MESSAGE_FORMAT_UNSPECIFIED\", \"PROTO\", \"JSON\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic": {
              "description": "Name of the Pub/Sub topic to publish conversation events",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Pub/Sub topic on which to publish new agent assistant events.\nExpects the format \"projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/topics/\u003cTopic ID\u003e\"",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "notification_config": {
        "block": {
          "attributes": {
            "message_format": {
              "description": "Format of the message Possible values: [\"MESSAGE_FORMAT_UNSPECIFIED\", \"PROTO\", \"JSON\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic": {
              "description": "Name of the Pub/Sub topic to publish conversation events",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Pub/Sub topic on which to publish new agent assistant events.\nExpects the format \"projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/topics/\u003cTopic ID\u003e\"",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "stt_config": {
        "block": {
          "attributes": {
            "audio_encoding": {
              "description": "Audio encoding of the audio content to process. Possible values: [\"AUDIO_ENCODING_UNSPECIFIED\", \"AUDIO_ENCODING_LINEAR_16\", \"AUDIO_ENCODING_FLAC\", \"AUDIO_ENCODING_MULAW\", \"AUDIO_ENCODING_AMR\", \"AUDIO_ENCODING_AMR_WB\", \"AUDIO_ENCODING_OGG_OPUS\", \"AUDIOENCODING_SPEEX_WITH_HEADER_BYTE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_word_info": {
              "description": "If true, Dialogflow returns SpeechWordInfo in StreamingRecognitionResult with information about the recognized speech words.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "language_code": {
              "computed": true,
              "description": "The language of the supplied audio.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model": {
              "description": "Which Speech model to select.\nLeave this field unspecified to use Agent Speech settings for model selection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sample_rate_hertz": {
              "description": "Sample rate (in Hertz) of the audio content sent in the query.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "speech_model_variant": {
              "description": "The speech model used in speech to text. Possible values: [\"SPEECH_MODEL_VARIANT_UNSPECIFIED\", \"USE_BEST_AVAILABLE\", \"USE_STANDARD\", \"USE_ENHANCED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_timeout_based_endpointing": {
              "description": "Use timeout based endpointing, interpreting endpointer sensitivy as seconds of timeout value.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Settings for speech transcription.",
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
      "tts_config": {
        "block": {
          "attributes": {
            "effects_profile_id": {
              "description": "An identifier which selects 'audio effects' profiles that are applied on (post synthesized) text to speech. Effects are applied on top of each other in the order they are given.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "pitch": {
              "description": "Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20 semitones from the original pitch. -20 means decrease 20 semitones from the original pitch.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "speaking_rate": {
              "description": "Speaking rate/speed, in the range [0.25, 4.0].",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "volume_gain_db": {
              "description": "Volume gain (in dB) of the normal native volume supported by the specific voice.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "voice": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The name of the voice.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssml_gender": {
                    "description": "The preferred gender of the voice. Possible values: [\"SSML_VOICE_GENDER_UNSPECIFIED\", \"SSML_VOICE_GENDER_MALE\", \"SSML_VOICE_GENDER_FEMALE\", \"SSML_VOICE_GENDER_NEUTRAL\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The desired voice of the synthesized audio.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for Text-to-Speech synthesization. If agent defines synthesization options as well, agent settings overrides the option here.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDialogflowConversationProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowConversationProfile), &result)
	return &result
}
