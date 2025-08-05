package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDiscoveryEngineCmekConfig = `{
  "block": {
    "attributes": {
      "cmek_config_id": {
        "description": "The unique id of the cmek config.",
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
      "is_default": {
        "computed": true,
        "description": "The default CmekConfig for the Customer.",
        "description_kind": "plain",
        "type": "bool"
      },
      "kms_key": {
        "description": "KMS key resource name which will be used to encrypt resources\n'projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_version": {
        "computed": true,
        "description": "KMS key version resource name which will be used to encrypt resources\n'\u003ckms_key\u003e/cryptoKeyVersions/{keyVersion}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_rotation_timestamp_micros": {
        "computed": true,
        "description": "The timestamp of the last key rotation.",
        "description_kind": "plain",
        "type": "number"
      },
      "location": {
        "description": "The geographic location where the CMEK config should reside. The value can\nonly be one of \"us\" and \"eu\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique full resource name of the cmek config. Values are of the format\n'projects/{project}/locations/{location}/cmekConfigs/{cmek_config_id}'.\nThis field must be a UTF-8 encoded string with a length limit of 1024\ncharacters.",
        "description_kind": "plain",
        "type": "string"
      },
      "notebooklm_state": {
        "computed": true,
        "description": "Whether the NotebookLM Corpus is ready to be used.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "set_default": {
        "description": "Set the following CmekConfig as the default to be used for child resources\nif one is not specified. The default value is true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The state of the CmekConfig.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "single_region_keys": {
        "block": {
          "attributes": {
            "kms_key": {
              "description": "Single-regional kms key resource name which will be used to encrypt\nresources\n'projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Single-regional CMEKs that are required for some VAIS features.",
          "description_kind": "plain"
        },
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

func GoogleDiscoveryEngineCmekConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDiscoveryEngineCmekConfig), &result)
	return &result
}
