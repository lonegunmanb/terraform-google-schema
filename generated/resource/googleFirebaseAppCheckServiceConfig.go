package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaseAppCheckServiceConfig = `{
  "block": {
    "attributes": {
      "enforcement_mode": {
        "description": "The App Check enforcement mode for a service supported by App Check. Valid values are\n\n(Unset)\nFirebase App Check is not enforced for the service, nor are App Check metrics collected.\nThough the service is not protected by App Check in this mode, other applicable protections,\nsuch as user authorization, are still enforced. An unconfigured service is in this mode by default.\nThis is equivalent to OFF in the REST API. Deleting the Terraform resource will also switch the\nenforcement to OFF for this service.\n\nUNENFORCED\nFirebase App Check is not enforced for the service. App Check metrics are collected to help you\ndecide when to turn on enforcement for the service. Though the service is not protected by App Check\nin this mode, other applicable protections, such as user authorization, are still enforced.\n\nENFORCED\nFirebase App Check is enforced for the service. The service will reject any request that attempts to\naccess your project's resources if it does not have valid App Check token attached, with some exceptions\ndepending on the service; for example, some services will still allow requests bearing the developer's\nprivileged service account credentials without an App Check token. App Check metrics continue to be\ncollected to help you detect issues with your App Check integration and monitor the composition of your\ncallers. While the service is protected by App Check, other applicable protections, such as user\nauthorization, continue to be enforced at the same time.\n\nUse caution when choosing to enforce App Check on a Firebase service. If your users have not updated\nto an App Check capable version of your app, their apps will no longer be able to use your Firebase\nservices that are enforcing App Check. App Check metrics can help you decide whether to enforce App\nCheck on your Firebase services.\n\nIf your app has not launched yet, you should enable enforcement immediately, since there are no outdated\nclients in use. Possible values: [\"UNENFORCED\", \"ENFORCED\"]",
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
      "name": {
        "computed": true,
        "description": "The fully-qualified resource name of the service enforcement configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_id": {
        "description": "The identifier of the service to configure enforcement. Currently, the following service IDs are supported:\n  firebasestorage.googleapis.com (Cloud Storage for Firebase)\n  firebasedatabase.googleapis.com (Firebase Realtime Database)\n  firestore.googleapis.com (Cloud Firestore)\n  identitytoolkit.googleapis.com (Authentication)",
        "description_kind": "plain",
        "required": true,
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

func GoogleFirebaseAppCheckServiceConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaseAppCheckServiceConfig), &result)
	return &result
}
