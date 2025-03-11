package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamOauthClient = `{
  "block": {
    "attributes": {
      "allowed_grant_types": {
        "description": "Required. The list of OAuth grant types is allowed for the OauthClient.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "allowed_redirect_uris": {
        "description": "Required. The list of redirect uris that is allowed to redirect back\nwhen authorization process is completed.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "allowed_scopes": {
        "description": "Required. The list of scopes that the OauthClient is allowed to request during\nOAuth flows.\n\nThe following scopes are supported:\n\n* 'https://www.googleapis.com/auth/cloud-platform': See, edit, configure,\nand delete your Google Cloud data and see the email address for your Google\nAccount.\n* 'openid': The OAuth client can associate you with your personal\ninformation on Google Cloud.\n* 'email': The OAuth client can read a federated identity's email address.\n* 'groups': The OAuth client can read a federated identity's groups.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "client_id": {
        "computed": true,
        "description": "Output only. The system-generated OauthClient id.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_type": {
        "description": "Immutable. The type of OauthClient. Either public or private.\nFor private clients, the client secret can be managed using the dedicated\nOauthClientCredential resource.\nPossible values:\nCLIENT_TYPE_UNSPECIFIED\nPUBLIC_CLIENT\nCONFIDENTIAL_CLIENT",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "A user-specified description of the OauthClient.\n\nCannot exceed 256 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description": "Whether the OauthClient is disabled. You cannot use a disabled OAuth\nclient.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "A user-specified display name of the OauthClient.\n\nCannot exceed 32 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "expire_time": {
        "computed": true,
        "description": "Time after which the OauthClient will be permanently purged and cannot\nbe recovered.",
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
        "description": "Immutable. Identifier. The resource name of the OauthClient.\n\nFormat:'projects/{project}/locations/{location}/oauthClients/{oauth_client}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "oauth_client_id": {
        "description": "Required. The ID to use for the OauthClient, which becomes the final component of\nthe resource name. This value should be a string of 6 to 63 lowercase\nletters, digits, or hyphens. It must start with a letter, and cannot have a\ntrailing hyphen. The prefix 'gcp-' is reserved for use by Google, and may\nnot be specified.",
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
        "description": "The state of the OauthClient.\nPossible values:\nSTATE_UNSPECIFIED\nACTIVE\nDELETED",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleIamOauthClientSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamOauthClient), &result)
	return &result
}
