package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIdentityPlatformConfig = `{
  "block": {
    "attributes": {
      "authorized_domains": {
        "computed": true,
        "description": "List of domains authorized for OAuth redirects.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "autodelete_anonymous_users": {
        "description": "Whether anonymous users will be auto-deleted after a period of 30 days",
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
      "name": {
        "computed": true,
        "description": "The name of the Config resource",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "blocking_functions": {
        "block": {
          "block_types": {
            "forward_inbound_credentials": {
              "block": {
                "attributes": {
                  "access_token": {
                    "description": "Whether to pass the user's OAuth identity provider's access token.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "id_token": {
                    "description": "Whether to pass the user's OIDC identity provider's ID token.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "refresh_token": {
                    "description": "Whether to pass the user's OAuth identity provider's refresh token.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "The user credentials to include in the JWT payload that is sent to the registered Blocking Functions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "triggers": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "function_uri": {
                    "description": "HTTP URI trigger for the Cloud Function.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "update_time": {
                    "computed": true,
                    "description": "When the trigger was changed.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Map of Trigger to event type. Key should be one of the supported event types: \"beforeCreate\", \"beforeSignIn\".",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description": "Configuration related to blocking functions.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "quota": {
        "block": {
          "block_types": {
            "sign_up_quota_config": {
              "block": {
                "attributes": {
                  "quota": {
                    "description": "A sign up APIs quota that customers can override temporarily.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "quota_duration": {
                    "description": "How long this quota will be active for. It is measurred in seconds, e.g., Example: \"9.615s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start_time": {
                    "description": "When this quota will take affect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Quota for the Signup endpoint, if overwritten. Signup quota is measured in sign ups per project per hour per IP.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration related to quotas.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sign_in": {
        "block": {
          "attributes": {
            "allow_duplicate_emails": {
              "description": "Whether to allow more than one account to have the same email.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hash_config": {
              "computed": true,
              "description": "Output only. Hash config information.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "algorithm": "string",
                    "memory_cost": "number",
                    "rounds": "number",
                    "salt_separator": "string",
                    "signer_key": "string"
                  }
                ]
              ]
            }
          },
          "block_types": {
            "anonymous": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether anonymous user auth is enabled for the project or not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration options related to authenticating an anonymous user.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "email": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether email auth is enabled for the project or not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "password_required": {
                    "description": "Whether a password is required for email auth or not. If true, both an email and\npassword must be provided to sign in. If false, a user may sign in via either\nemail/password or email link.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration options related to authenticating a user by their email address.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "phone_number": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether phone number auth is enabled for the project or not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "test_phone_numbers": {
                    "description": "A map of \u003ctest phone number, fake code\u003e that can be used for phone auth testing.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "Configuration options related to authenticated a user by their phone number.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration related to local sign in methods.",
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

func GoogleIdentityPlatformConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIdentityPlatformConfig), &result)
	return &result
}
