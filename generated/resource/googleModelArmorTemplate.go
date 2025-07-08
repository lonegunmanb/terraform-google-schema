package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleModelArmorTemplate = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Create time stamp",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels as key value pairs\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
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
        "description": "Identifier. name of resource",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_id": {
        "description": "Id of the requesting object\nIf auto-generating Id server-side, remove this field and\ntemplate_id from the method_signature of Create RPC",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Update time stamp",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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
                          "description": "Optional Sensitive Data Protection Deidentify template resource name.\nIf provided then DeidentifyContent action is performed during Sanitization\nusing this template and inspect template. The De-identified data will\nbe returned in SdpDeidentifyResult.\nNote that all info-types present in the deidentify template must be present\nin inspect template.\ne.g.\n'projects/{project}/locations/{location}/deidentifyTemplates/{deidentify_template}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "inspect_template": {
                          "description": "Sensitive Data Protection inspect template resource name\nIf only inspect template is provided (de-identify template not provided),\nthen Sensitive Data Protection InspectContent action is performed during\nSanitization. All Sensitive Data Protection findings identified during\ninspection will be returned as SdpFinding in SdpInsepctionResult.\ne.g:-\n'projects/{project}/locations/{location}/inspectTemplates/{inspect_template}'",
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
      "template_metadata": {
        "block": {
          "attributes": {
            "custom_llm_response_safety_error_code": {
              "description": "Indicates the custom error code set by the user to be returned to the end\nuser if the LLM response trips Model Armor filters.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "custom_llm_response_safety_error_message": {
              "description": "Indicates the custom error message set by the user to be returned to the\nend user if the LLM response trips Model Armor filters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "custom_prompt_safety_error_code": {
              "description": "Indicates the custom error code set by the user to be returned to the end\nuser by the service extension if the prompt trips Model Armor filters.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "custom_prompt_safety_error_message": {
              "description": "Indicates the custom error message set by the user to be returned to the\nend user if the prompt trips Model Armor filters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enforcement_type": {
              "description": "Possible values:\nINSPECT_ONLY\nINSPECT_AND_BLOCK",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ignore_partial_invocation_failures": {
              "description": "If true, partial detector failures should be ignored.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "log_sanitize_operations": {
              "description": "If true, log sanitize operations.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "log_template_operations": {
              "description": "If true, log template crud operations.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
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
                "description": "Metadata to enable multi language detection via template.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Message describing TemplateMetadata",
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

func GoogleModelArmorTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleModelArmorTemplate), &result)
	return &result
}
