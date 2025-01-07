package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxAgent = `{
  "block": {
    "attributes": {
      "avatar_uri": {
        "description": "The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_language_code": {
        "description": "The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)\nfor a list of the currently supported language codes. This field cannot be updated after creation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the agent, unique within the location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_spell_correction": {
        "description": "Indicates if automatic spell correction is enabled in detect intent requests.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_stackdriver_logging": {
        "deprecated": true,
        "description": "Determines whether this agent should log conversation queries.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The name of the location this agent is located in.\n\n~\u003e **Note:** The first time you are deploying an Agent in your project you must configure location settings.\n This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.\n Another options is to use global location so you don't need to manually configure location settings.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the agent.",
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
        "description": "Name of the SecuritySettings reference for the agent. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/securitySettings/\u003cSecurity Settings ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_flow": {
        "computed": true,
        "description": "Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "supported_language_codes": {
        "description": "The list of all languages supported by this agent (except for the default_language_code).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "time_zone": {
        "description": "The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,\nEurope/Paris.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
          "description": "Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.\nHierarchy: Agent-\u003eFlow-\u003ePage-\u003eFulfillment/Parameter.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "git_integration_settings": {
        "block": {
          "block_types": {
            "github_settings": {
              "block": {
                "attributes": {
                  "access_token": {
                    "description": "The access token used to authenticate the access to the GitHub repository.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "branches": {
                    "description": "A list of branches configured to be used from Dialogflow.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "display_name": {
                    "description": "The unique repository display name for the GitHub repository.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "repository_uri": {
                    "description": "The GitHub repository URI related to the agent.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tracking_branch": {
                    "description": "The branch of the GitHub repository tracked for this agent.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Settings of integration with GitHub.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Git integration settings for this agent.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "speech_to_text_settings": {
        "block": {
          "attributes": {
            "enable_speech_adaptation": {
              "description": "Whether to use speech adaptation for speech recognition.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Settings related to speech recognition.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "text_to_speech_settings": {
        "block": {
          "attributes": {
            "synthesize_speech_configs": {
              "description": "Configuration of how speech should be synthesized, mapping from [language](https://cloud.google.com/dialogflow/cx/docs/reference/language) to [SynthesizeSpeechConfig](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents#synthesizespeechconfig).\nThese settings affect:\n* The phone gateway synthesize configuration set via Agent.text_to_speech_settings.\n* How speech is synthesized when invoking session APIs. 'Agent.text_to_speech_settings' only applies if 'OutputAudioConfig.synthesize_speech_config' is not specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Settings related to speech synthesizing.",
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

func GoogleDialogflowCxAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxAgent), &result)
	return &result
}
