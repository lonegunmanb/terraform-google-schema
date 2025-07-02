package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerRegionalSecrets = `{
  "block": {
    "attributes": {
      "filter": {
        "description": "Filter string, adhering to the rules in List-operation filtering (https://cloud.google.com/secret-manager/docs/filtering).\nList only secrets matching the filter. If filter is empty, all regional secrets are listed from the specified location.",
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
      "secrets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "annotations": [
                "map",
                "string"
              ],
              "create_time": "string",
              "customer_managed_encryption": [
                "list",
                [
                  "object",
                  {
                    "kms_key_name": "string"
                  }
                ]
              ],
              "deletion_protection": "bool",
              "effective_annotations": [
                "map",
                "string"
              ],
              "effective_labels": [
                "map",
                "string"
              ],
              "expire_time": "string",
              "labels": [
                "map",
                "string"
              ],
              "location": "string",
              "name": "string",
              "project": "string",
              "rotation": [
                "list",
                [
                  "object",
                  {
                    "next_rotation_time": "string",
                    "rotation_period": "string"
                  }
                ]
              ],
              "secret_id": "string",
              "terraform_labels": [
                "map",
                "string"
              ],
              "topics": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "ttl": "string",
              "version_aliases": [
                "map",
                "string"
              ],
              "version_destroy_ttl": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSecretManagerRegionalSecretsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecretManagerRegionalSecrets), &result)
	return &result
}
