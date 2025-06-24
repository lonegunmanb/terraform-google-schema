package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApihubPluginInstance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Timestamp indicating when the plugin instance was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "disable": {
        "description": "The display name for this plugin instance. Max length is 255 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "The display name for this plugin instance. Max length is 255 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "error_message": {
        "computed": true,
        "description": "Error message describing the failure, if any, during Create, Delete or\nApplyConfig operation corresponding to the plugin instance.This field will\nonly be populated if the plugin instance is in the ERROR or FAILED state.",
        "description_kind": "plain",
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
        "description": "Identifier. The unique name of the plugin instance resource.\nFormat:\n'projects/{project}/locations/{location}/plugins/{plugin}/instances/{instance}'",
        "description_kind": "plain",
        "type": "string"
      },
      "plugin": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plugin_instance_id": {
        "description": "The ID to use for the plugin instance, which will become the final\ncomponent of the plugin instance's resource name. This field is optional.\n\n* If provided, the same will be used. The service will throw an error if\nthe specified id is already used by another plugin instance in the plugin\nresource.\n* If not provided, a system generated id will be used.\n\nThis value should be 4-63 characters, and valid characters\nare /a-z[0-9]-_/.",
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
        "description": "The current state of the plugin instance (e.g., enabled, disabled,\nprovisioning).\nPossible values:\nSTATE_UNSPECIFIED\nCREATING\nACTIVE\nAPPLYING_CONFIG\nERROR\nFAILED\nDELETING",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Timestamp indicating when the plugin instance was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "actions": {
        "block": {
          "attributes": {
            "action_id": {
              "description": "This should map to one of the action id specified\nin actions_config in the plugin.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hub_instance_action": {
              "computed": true,
              "description": "The execution status for the plugin instance.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "current_execution_state": "string",
                    "last_execution": [
                      "list",
                      [
                        "object",
                        {
                          "end_time": "string",
                          "error_message": "string",
                          "result": "string",
                          "start_time": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            },
            "schedule_cron_expression": {
              "computed": true,
              "description": "The schedule for this plugin instance action. This can only be set if the\nplugin supports API_HUB_SCHEDULE_TRIGGER mode for this action.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_time_zone": {
              "computed": true,
              "description": "The time zone for the schedule cron expression. If not provided, UTC will\nbe used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "The current state of the plugin action in the plugin instance.\nPossible values:\nSTATE_UNSPECIFIED\nENABLED\nDISABLED\nENABLING\nDISABLING\nERROR",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "curation_config": {
              "block": {
                "attributes": {
                  "curation_type": {
                    "computed": true,
                    "description": "Possible values:\nCURATION_TYPE_UNSPECIFIED\nDEFAULT_CURATION_FOR_API_METADATA\nCUSTOM_CURATION_FOR_API_METADATA",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_curation": {
                    "block": {
                      "attributes": {
                        "curation": {
                          "description": "The unique name of the curation resource. This will be the name of the\ncuration resource in the format:\n'projects/{project}/locations/{location}/curations/{curation}'",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Custom curation information for this plugin instance.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The curation information for this plugin instance.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The action status for the plugin instance.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "auth_config": {
        "block": {
          "attributes": {
            "auth_type": {
              "description": "Possible values:\nAUTH_TYPE_UNSPECIFIED\nNO_AUTH\nGOOGLE_SERVICE_ACCOUNT\nUSER_PASSWORD\nAPI_KEY\nOAUTH2_CLIENT_CREDENTIALS",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "api_key_config": {
              "block": {
                "attributes": {
                  "http_element_location": {
                    "description": "The location of the API key.\nThe default value is QUERY.\nPossible values:\nHTTP_ELEMENT_LOCATION_UNSPECIFIED\nQUERY\nHEADER\nPATH\nBODY\nCOOKIE",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The parameter name of the API key.\nE.g. If the API request is \"https://example.com/act?api_key=\",\n\"api_key\" would be the parameter name.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "api_key": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: 'projects/*/secrets/*/versions/*'.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Secret provides a reference to entries in Secret Manager.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Config for authentication with API key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "google_service_account_config": {
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
            },
            "oauth2_client_credentials_config": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The client identifier.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "client_secret": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: 'projects/*/secrets/*/versions/*'.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Secret provides a reference to entries in Secret Manager.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Parameters to support Oauth 2.0 client credentials grant authentication.\nSee https://tools.ietf.org/html/rfc6749#section-1.3.4 for more details.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "user_password_config": {
              "block": {
                "attributes": {
                  "username": {
                    "description": "Username.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "password": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: 'projects/*/secrets/*/versions/*'.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Secret provides a reference to entries in Secret Manager.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Parameters to support Username and Password Authentication.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "AuthConfig represents the authentication information.",
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

func GoogleApihubPluginInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApihubPluginInstance), &result)
	return &result
}
