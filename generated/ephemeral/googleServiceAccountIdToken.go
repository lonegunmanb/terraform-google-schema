package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountIdToken = `{
  "block": {
    "attributes": {
      "delegates": {
        "description": "Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  Used only when using impersonation mode.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id_token": {
        "computed": true,
        "description": "The ` + "`" + `id_token` + "`" + ` representing the new generated identity.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "include_email": {
        "description": "Include the verified email in the claim. Used only when using impersonation mode.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "target_audience": {
        "description": "The audience claim for the ` + "`" + `id_token` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_service_account": {
        "description": "The email of the service account being impersonated.  Used only when using impersonation mode.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountIdTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountIdToken), &result)
	return &result
}
