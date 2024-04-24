package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaseAppCheckDebugToken = `{
  "block": {
    "attributes": {
      "app_id": {
        "description": "The ID of a\n[Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id),\n[Apple App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id),\nor [Android App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "debug_token_id": {
        "computed": true,
        "description": "The last segment of the resource name of the debug token.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "A human readable display name used to identify this debug token.",
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "token": {
        "description": "The secret token itself. Must be provided during creation, and must be a UUID4,\ncase insensitive. You may use a method of your choice such as random/random_uuid\nto generate the token.\n\nThis field is immutable once set, and cannot be updated. You can, however, delete\nthis debug token to revoke it.\n\nFor security reasons, this field will never be populated in any response.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
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

func GoogleFirebaseAppCheckDebugTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaseAppCheckDebugToken), &result)
	return &result
}
