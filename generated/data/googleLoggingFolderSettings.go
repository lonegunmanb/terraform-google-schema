package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingFolderSettings = `{
  "block": {
    "attributes": {
      "disable_default_sink": {
        "computed": true,
        "description": "If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.",
        "description_kind": "plain",
        "type": "bool"
      },
      "folder": {
        "description": "The folder for which to retrieve settings.",
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
      "kms_key_name": {
        "computed": true,
        "description": "The resource name for the configured Cloud KMS key.\n\t\t\t\tKMS key name format:\n\t\t\t\t\"projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]\"\n\t\t\t\tTo enable CMEK for the bucket, set this field to a valid kmsKeyName for which the associated service account has the required cloudkms.cryptoKeyEncrypterDecrypter roles assigned for the key.\n\t\t\t\tThe Cloud KMS key used by the bucket can be updated by changing the kmsKeyName to a new valid key name. Encryption operations that are in progress will be completed with the key that was in use when they started. Decryption operations will be completed using the key that was used at the time of encryption unless access to that key has been revoked.\n\t\t\t\tSee [Enabling CMEK for Logging Buckets](https://cloud.google.com/logging/docs/routing/managed-encryption-storage) for more information.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_service_account_id": {
        "computed": true,
        "description": "The service account associated with a project for which CMEK will apply.\n\t\t\t\tBefore enabling CMEK for a logging bucket, you must first assign the cloudkms.cryptoKeyEncrypterDecrypter role to the service account associated with the project for which CMEK will apply. Use [v2.getCmekSettings](https://cloud.google.com/logging/docs/reference/v2/rest/v2/TopLevel/getCmekSettings#google.logging.v2.ConfigServiceV2.GetCmekSettings) to obtain the service account ID.\n\t\t\t\tSee [Enabling CMEK for Logging Buckets](https://cloud.google.com/logging/docs/routing/managed-encryption-storage) for more information.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_service_account_id": {
        "computed": true,
        "description": "The service account for the given container. Sinks use this service account as their writerIdentity if no custom service account is provided.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the CMEK settings.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_location": {
        "computed": true,
        "description": "The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleLoggingFolderSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingFolderSettings), &result)
	return &result
}
