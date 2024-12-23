package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountAccessToken = `{
  "block": {
    "attributes": {
      "access_token": {
        "computed": true,
        "description": "The ` + "`" + `access_token` + "`" + ` representing the new generated identity.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "delegates": {
        "description": "Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. ` + "`" + `['projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com']` + "`" + `)",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "lifetime": {
        "computed": true,
        "description": "Lifetime of the impersonated token (defaults to its max: ` + "`" + `3600s` + "`" + `)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scopes": {
        "description": "The scopes the new credential should have (e.g. ` + "`" + `['cloud-platform']` + "`" + `)",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "target_service_account": {
        "description": "The service account to impersonate (e.g. ` + "`" + `service_B@your-project-id.iam.gserviceaccount.com` + "`" + `)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountAccessTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountAccessToken), &result)
	return &result
}
