package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerGcpUserAccessBinding = `{
  "block": {
    "attributes": {
      "access_levels": {
        "description": "Optional. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: \"accessPolicies/9522/accessLevels/device_trusted\"",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "group_key": {
        "description": "Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See \"id\" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: \"01d520gv4vjcrht\"",
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
      "name": {
        "computed": true,
        "description": "Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by RFC 3986 Section 2.3). Should not be specified by the client during creation. Example: \"organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N\"",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_id": {
        "description": "Required. ID of the parent organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "scoped_access_settings": {
        "block": {
          "block_types": {
            "active_settings": {
              "block": {
                "attributes": {
                  "access_levels": {
                    "description": "Optional. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: \"accessPolicies/9522/accessLevels/device_trusted\"",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "session_settings": {
                    "block": {
                      "attributes": {
                        "max_inactivity": {
                          "description": "Optional. How long a user is allowed to take between actions before a new access token must be issued. Only set for Google Cloud apps.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "session_length": {
                          "description": "Optional. The session length. Setting this field to zero is equal to disabling session. Also can set infinite session by flipping the enabled bit to false below. If useOidcMaxAge is true, for OIDC apps, the session length will be the minimum of this field and OIDC max_age param.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "session_length_enabled": {
                          "description": "Optional. This field enables or disables Google Cloud session length. When false, all fields set above will be disregarded and the session length is basically infinite.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "session_reauth_method": {
                          "description": "Optional. The session challenges proposed to users when the Google Cloud session length is up. Possible values: [\"LOGIN\", \"SECURITY_KEY\", \"PASSWORD\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "use_oidc_max_age": {
                          "description": "Optional. Only useful for OIDC apps. When false, the OIDC max_age param, if passed in the authentication request will be ignored. When true, the re-auth period will be the minimum of the sessionLength field and the max_age OIDC param.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Optional. Session settings applied to user access on a given AccessScope.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Access settings for this scoped access settings. This field may be empty if dryRunSettings is set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "dry_run_settings": {
              "block": {
                "attributes": {
                  "access_levels": {
                    "description": "Optional. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: \"accessPolicies/9522/accessLevels/device_trusted\"",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Optional. Dry-run access settings for this scoped access settings. This field may be empty if activeSettings is set. Cannot contain session settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "scope": {
              "block": {
                "block_types": {
                  "client_scope": {
                    "block": {
                      "block_types": {
                        "restricted_client_application": {
                          "block": {
                            "attributes": {
                              "client_id": {
                                "description": "The OAuth client ID of the application.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "description": "The name of the application. Example: \"Cloud Console\"",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Optional. The application that is subject to this binding's scope. Only one of clientId or name should be specified.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Optional. Client scope for this access scope.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Application, etc. to which the access settings will be applied to. Implicitly, this is the scoped access settings key; as such, it must be unique and non-empty.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. A list of scoped access settings that set this binding's restrictions on a subset of applications.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "session_settings": {
        "block": {
          "attributes": {
            "max_inactivity": {
              "description": "Optional. How long a user is allowed to take between actions before a new access token must be issued. Only set for Google Cloud apps.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "session_length": {
              "description": "Optional. The session length. Setting this field to zero is equal to disabling session. Also can set infinite session by flipping the enabled bit to false below. If useOidcMaxAge is true, for OIDC apps, the session length will be the minimum of this field and OIDC max_age param.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "session_length_enabled": {
              "description": "Optional. This field enables or disables Google Cloud session length. When false, all fields set above will be disregarded and the session length is basically infinite.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "session_reauth_method": {
              "description": "Optional. The session challenges proposed to users when the Google Cloud session length is up. Possible values: [\"LOGIN\", \"SECURITY_KEY\", \"PASSWORD\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_oidc_max_age": {
              "description": "Optional. Only useful for OIDC apps. When false, the OIDC max_age param, if passed in the authentication request will be ignored. When true, the re-auth period will be the minimum of the sessionLength field and the max_age OIDC param.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Optional. The Google Cloud session length (GCSL) policy for the group key.",
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

func GoogleAccessContextManagerGcpUserAccessBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerGcpUserAccessBinding), &result)
	return &result
}
