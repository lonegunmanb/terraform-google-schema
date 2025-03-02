package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountJwt = `{
  "block": {
    "attributes": {
      "delegates": {
        "description": "Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "expires_in": {
        "description": "Number of seconds until the JWT expires. If set and non-zero an ` + "`" + `exp` + "`" + ` claim will be added to the payload derived from the current timestamp plus expires_in seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "jwt": {
        "computed": true,
        "description": "The signed JWT containing the JWT Claims Set from the ` + "`" + `payload` + "`" + `.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "payload": {
        "description": "A JSON-encoded JWT claims set that will be included in the signed JWT.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_service_account": {
        "description": "The email of the service account that will sign the JWT.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Produces an arbitrary self-signed JWT for service accounts.",
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountJwtSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountJwt), &result)
	return &result
}
