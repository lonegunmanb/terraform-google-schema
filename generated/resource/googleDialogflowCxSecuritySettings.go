package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxSecuritySettings = `{
  "block": {
    "attributes": {
      "deidentify_template": {
        "description": "[DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. If empty, Dialogflow replaces sensitive info with [redacted] text.\nNote: deidentifyTemplate must be located in the same region as the SecuritySettings.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/deidentifyTemplates/\u003cTemplate ID\u003e OR organizations/\u003cOrganization ID\u003e/locations/\u003cLocation ID\u003e/deidentifyTemplates/\u003cTemplate ID\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the security settings, unique within the location.",
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
      "inspect_template": {
        "description": "[DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config.\nNote: inspectTemplate must be located in the same region as the SecuritySettings.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/inspectTemplates/\u003cTemplate ID\u003e OR organizations/\u003cOrganization ID\u003e/locations/\u003cLocation ID\u003e/inspectTemplates/\u003cTemplate ID\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location these settings are located in. Settings can only be applied to an agent in the same location.\nSee [Available Regions](https://cloud.google.com/dialogflow/cx/docs/concept/region#avail) for a list of supported locations.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the settings.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/securitySettings/\u003cSecurity Settings ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "purge_data_types": {
        "description": "List of types of data to remove when retention settings triggers purge. Possible values: [\"DIALOGFLOW_HISTORY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "redaction_scope": {
        "description": "Defines what types of data to redact. If not set, defaults to not redacting any kind of data.\n* REDACT_DISK_STORAGE: On data to be written to disk or similar devices that are capable of holding data even if power is disconnected. This includes data that are temporarily saved on disk. Possible values: [\"REDACT_DISK_STORAGE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redaction_strategy": {
        "description": "Defines how we redact data. If not set, defaults to not redacting.\n* REDACT_WITH_SERVICE: Call redaction service to clean up the data to be persisted. Possible values: [\"REDACT_WITH_SERVICE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_strategy": {
        "description": "Defines how long we retain persisted data that contains sensitive info. Only one of 'retention_window_days' and 'retention_strategy' may be set.\n* REMOVE_AFTER_CONVERSATION: Removes data when the conversation ends. If there is no conversation explicitly established, a default conversation ends when the corresponding Dialogflow session ends. Possible values: [\"REMOVE_AFTER_CONVERSATION\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_window_days": {
        "description": "Retains the data for the specified number of days. User must set a value lower than Dialogflow's default 365d TTL (30 days for Agent Assist traffic), higher value will be ignored and use default. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use default TTL.\nOnly one of 'retention_window_days' and 'retention_strategy' may be set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "audio_export_settings": {
        "block": {
          "attributes": {
            "audio_export_pattern": {
              "description": "Filename pattern for exported audio.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "audio_format": {
              "description": "File format for exported audio file. Currently only in telephony recordings.\n* MULAW: G.711 mu-law PCM with 8kHz sample rate.\n* MP3: MP3 file format.\n* OGG: OGG Vorbis. Possible values: [\"MULAW\", \"MP3\", \"OGG\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_audio_redaction": {
              "description": "Enable audio redaction if it is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "gcs_bucket": {
              "description": "Cloud Storage bucket to export audio record to. Setting this field would grant the Storage Object Creator role to the Dialogflow Service Agent. API caller that tries to modify this field should have the permission of storage.buckets.setIamPolicy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Controls audio export settings for post-conversation analytics when ingesting audio to conversations.\nIf retention_strategy is set to REMOVE_AFTER_CONVERSATION or gcs_bucket is empty, audio export is disabled.\nIf audio export is enabled, audio is recorded and saved to gcs_bucket, subject to retention policy of gcs_bucket.\nThis setting won't effect audio input for implicit sessions via [Sessions.DetectIntent](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.sessions/detectIntent#google.cloud.dialogflow.cx.v3.Sessions.DetectIntent).",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "insights_export_settings": {
        "block": {
          "attributes": {
            "enable_insights_export": {
              "description": "If enabled, we will automatically exports conversations to Insights and Insights runs its analyzers.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Controls conversation exporting settings to Insights after conversation is completed.\nIf retentionStrategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.",
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

func GoogleDialogflowCxSecuritySettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxSecuritySettings), &result)
	return &result
}
