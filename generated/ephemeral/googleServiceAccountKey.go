package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountKey = `{
  "block": {
    "attributes": {
      "key_algorithm": {
        "computed": true,
        "description": "The algorithm used to generate the key.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the service account key. This must have format ` + "`" + `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{KEYID}` + "`" + `, where ` + "`" + `{ACCOUNT}` + "`" + ` is the email address or unique id of the service account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_key": {
        "computed": true,
        "description": "The public key, base64 encoded.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_type": {
        "description": "The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Get an ephemeral service account public key.",
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountKey), &result)
	return &result
}
