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
