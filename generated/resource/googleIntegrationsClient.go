package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIntegrationsClient = `{
  "block": {
    "attributes": {
      "create_sample_integrations": {
        "description": "Indicates if sample integrations should be created along with provisioning.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_sample_workflows": {
        "deprecated": true,
        "description": "Indicates if sample workflow should be created along with provisioning.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Location in which client needs to be provisioned.",
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
      "provision_gmek": {
        "deprecated": true,
        "description": "Indicates provision with GMEK or CMEK.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "run_as_service_account": {
        "description": "User input run-as service account, if empty, will bring up a new default service account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "cloud_kms_config": {
        "block": {
          "attributes": {
            "key": {
              "description": "A Cloud KMS key is a named object containing one or more key versions, along\nwith metadata for the key. A key exists on exactly one key ring tied to a\nspecific location.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_version": {
              "description": "Each version of a key contains key material used for encryption or signing.\nA key's version is represented by an integer, starting at 1. To decrypt data\nor verify a signature, you must use the same key version that was used to\nencrypt or sign the data.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_location": {
              "description": "Location name of the key ring, e.g. \"us-west1\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kms_project_id": {
              "description": "The Google Cloud project id of the project where the kms key stored. If empty,\nthe kms key is stored at the same project as customer's project and ecrypted\nwith CMEK, otherwise, the kms key is stored in the tenant project and\nencrypted with GMEK.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_ring": {
              "description": "A key ring organizes keys in a specific Google Cloud location and allows you to\nmanage access control on groups of keys. A key ring's name does not need to be\nunique across a Google Cloud project, but must be unique within a given location.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Cloud KMS config for AuthModule to encrypt/decrypt credentials.",
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

func GoogleIntegrationsClientSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIntegrationsClient), &result)
	return &result
}
