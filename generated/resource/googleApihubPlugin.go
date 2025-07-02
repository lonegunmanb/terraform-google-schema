package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApihubPlugin = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Timestamp indicating when the plugin was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The plugin description. Max length is 2000 characters (Unicode code\npoints).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the plugin. Max length is 50 characters (Unicode code\npoints).",
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
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the plugin.\nFormat: 'projects/{project}/locations/{location}/plugins/{plugin}'",
        "description_kind": "plain",
        "type": "string"
      },
      "ownership_type": {
        "computed": true,
        "description": "The type of the plugin, indicating whether it is 'SYSTEM_OWNED' or\n'USER_OWNED'.\nPossible values:\nOWNERSHIP_TYPE_UNSPECIFIED\nSYSTEM_OWNED\nUSER_OWNED",
        "description_kind": "plain",
        "type": "string"
      },
      "plugin_category": {
        "description": "Possible values:\nPLUGIN_CATEGORY_UNSPECIFIED\nAPI_GATEWAY\nAPI_PRODUCER",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plugin_id": {
        "description": "The ID to use for the Plugin resource, which will become the final\ncomponent of the Plugin's resource name. This field is optional.\n\n* If provided, the same will be used. The service will throw an error if\nthe specified id is already used by another Plugin resource in the API hub\ninstance.\n* If not provided, a system generated id will be used.\n\nThis value should be 4-63 characters, overall resource name which will be\nof format\n'projects/{project}/locations/{location}/plugins/{plugin}',\nits length is limited to 1000 characters and valid characters are\n/a-z[0-9]-_/.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Represents the state of the plugin.\nNote this field will not be set for plugins developed via plugin\nframework as the state will be managed at plugin instance level.\nPossible values:\nSTATE_UNSPECIFIED\nENABLED\nDISABLED",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Timestamp indicating when the plugin was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "actions_config": {
        "block": {
          "attributes": {
            "description": {
              "description": "The description of the operation performed by the action.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "display_name": {
              "description": "The display name of the action.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "description": "The id of the action.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "trigger_mode": {
              "description": "The trigger mode supported by the action.\nPossible values:\nTRIGGER_MODE_UNSPECIFIED\nAPI_HUB_ON_DEMAND_TRIGGER\nAPI_HUB_SCHEDULE_TRIGGER\nNON_API_HUB_MANAGED",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The configuration of actions supported by the plugin.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "config_template": {
        "block": {
          "block_types": {
            "additional_config_template": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "Description.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "id": {
                    "description": "ID of the config variable. Must be unique within the configuration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "required": {
                    "description": "Flag represents that this 'ConfigVariable' must be provided for a\nPluginInstance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "validation_regex": {
                    "description": "Regular expression in RE2 syntax used for validating the 'value' of a\n'ConfigVariable'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value_type": {
                    "description": "Type of the parameter: string, int, bool etc.\nPossible values:\nVALUE_TYPE_UNSPECIFIED\nSTRING\nINT\nBOOL\nSECRET\nENUM\nMULTI_SELECT\nMULTI_STRING\nMULTI_INT",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "enum_options": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description": "Description of the option.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "display_name": {
                          "description": "Display name of the option.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "id": {
                          "description": "Id of the option.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Enum options. To be populated if 'ValueType' is 'ENUM'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "multi_select_options": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description": "Description of the option.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "display_name": {
                          "description": "Display name of the option.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "id": {
                          "description": "Id of the option.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Multi select options. To be populated if 'ValueType' is 'MULTI_SELECT'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The list of additional configuration variables for the plugin's\nconfiguration.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "auth_config_template": {
              "block": {
                "attributes": {
                  "supported_auth_types": {
                    "description": "The list of authentication types supported by the plugin.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "service_account": {
                    "block": {
                      "attributes": {
                        "service_account": {
                          "description": "The service account to be used for authenticating request.\n\nThe 'iam.serviceAccounts.getAccessToken' permission should be granted on\nthis service account to the impersonator service account.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Config for Google service account authentication.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "AuthConfigTemplate represents the authentication template for a plugin.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "ConfigTemplate represents the configuration template for a plugin.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "documentation": {
        "block": {
          "attributes": {
            "external_uri": {
              "description": "The uri of the externally hosted documentation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Documentation details.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "hosting_service": {
        "block": {
          "attributes": {
            "service_uri": {
              "description": "The URI of the service implemented by the plugin developer, used to\ninvoke the plugin's functionality. This information is only required for\nuser defined plugins.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The information related to the service implemented by the plugin\ndeveloper, used to invoke the plugin's functionality.",
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

func GoogleApihubPluginSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApihubPlugin), &result)
	return &result
}
