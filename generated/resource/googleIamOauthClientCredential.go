package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamOauthClientCredential = `{
  "block": {
    "attributes": {
      "client_secret": {
        "computed": true,
        "description": "The system-generated OAuth client secret.\n\nThe client secret must be stored securely. If the client secret is\nleaked, you must delete and re-create the client credential. To learn\nmore, see [OAuth client and credential security risks and\nmitigations](https://cloud.google.com/iam/docs/workforce-oauth-app#security)",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "description": "Whether the OauthClientCredential is disabled. You cannot use a\ndisabled OauthClientCredential.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "A user-specified display name of the OauthClientCredential.\n\nCannot exceed 32 characters.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Immutable. Identifier. The resource name of the OauthClientCredential.\n\nFormat:\n'projects/{project}/locations/{location}/oauthClients/{oauth_client}/credentials/{credential}'",
        "description_kind": "plain",
        "type": "string"
      },
      "oauth_client_credential_id": {
        "description": "Required. The ID to use for the OauthClientCredential, which becomes the\nfinal component of the resource name. This value should be 4-32 characters,\nand may contain the characters [a-z0-9-]. The prefix 'gcp-' is\nreserved for use by Google, and may not be specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "oauthclient": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
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

func GoogleIamOauthClientCredentialSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamOauthClientCredential), &result)
	return &result
}
