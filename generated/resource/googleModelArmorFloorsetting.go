package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleModelArmorFloorsetting = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "[Output only] Create timestamp",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_floor_setting_enforcement": {
        "description": "Floor Settings enforcement status.",
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
      "integrated_services": {
        "description": "List of integrated services for which the floor setting is applicable.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "Will be any one of these:\n\n* 'projects/{project}'\n* 'folders/{folder}'\n* 'organizations/{organizationId}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "[Output only] Update timestamp",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "ai_platform_floor_setting": {
        "block": {
          "attributes": {
            "enable_cloud_logging": {
              "description": "If true, log Model Armor filter results to Cloud Logging.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "inspect_and_block": {
              "description": "If true, Model Armor filters will be run in inspect and block mode.\nRequests that trip Model Armor filters will be blocked.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "inspect_only": {
              "description": "If true, Model Armor filters will be run in inspect only mode. No action\nwill be taken on the request.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "AI Platform floor setting.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "filter_config": {
        "block": {
          "block_types": {
            "malicious_uri_filter_settings": {
              "block": {
                "attributes": {
                  "filter_enforcement": {
                    "description": "Tells whether the Malicious URI filter is enabled or disabled.\nPossible values:\nENABLED\nDISABLED",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Malicious URI filter settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pi_and_jailbreak_filter_settings": {
              "block": {
                "attributes": {
                  "confidence_level": {
                    "description": "Possible values:\nLOW_AND_ABOVE\nMEDIUM_AND_ABOVE\nHIGH",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "filter_enforcement": {
                    "description": "Tells whether Prompt injection and Jailbreak filter is enabled or\ndisabled.\nPossible values:\nENABLED\nDISABLED",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Prompt injection and Jailbreak Filter settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rai_settings": {
              "block": {
                "block_types": {
                  "rai_filters": {
                    "block": {
                      "attributes": {
                        "confidence_level": {
                          "description": "Possible values:\nLOW_AND_ABOVE\nMEDIUM_AND_ABOVE\nHIGH",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "filter_type": {
                          "description": "Possible values:\nSEXUALLY_EXPLICIT\nHATE_SPEECH\nHARASSMENT\nDANGEROUS",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "List of Responsible AI filters enabled for template.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Responsible AI Filter settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sdp_settings": {
              "block": {
                "block_types": {
                  "advanced_config": {
                    "block": {
                      "attributes": {
                        "deidentify_template": {
                          "description": "Optional Sensitive Data Protection Deidentify template resource name.\n\nIf provided then DeidentifyContent action is performed during Sanitization\nusing this template and inspect template. The De-identified data will\nbe returned in SdpDeidentifyResult.\nNote that all info-types present in the deidentify template must be present\nin inspect template.\n\ne.g.\n'projects/{project}/locations/{location}/deidentifyTemplates/{deidentify_template}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "inspect_template": {
                          "description": "Sensitive Data Protection inspect template resource name\n\nIf only inspect template is provided (de-identify template not provided),\nthen Sensitive Data Protection InspectContent action is performed during\nSanitization. All Sensitive Data Protection findings identified during\ninspection will be returned as SdpFinding in SdpInsepctionResult.\n\ne.g:-\n'projects/{project}/locations/{location}/inspectTemplates/{inspect_template}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Sensitive Data Protection Advanced configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "basic_config": {
                    "block": {
                      "attributes": {
                        "filter_enforcement": {
                          "description": "Tells whether the Sensitive Data Protection basic config is enabled or\ndisabled.\nPossible values:\nENABLED\nDISABLED",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Sensitive Data Protection basic configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Sensitive Data Protection settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Filters configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "floor_setting_metadata": {
        "block": {
          "block_types": {
            "multi_language_detection": {
              "block": {
                "attributes": {
                  "enable_multi_language_detection": {
                    "description": "If true, multi language detection will be enabled.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Metadata for multi language detection.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Metadata to enable multi language detection via floor setting.",
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

func GoogleModelArmorFloorsettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleModelArmorFloorsetting), &result)
	return &result
}
