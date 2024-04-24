package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerSecrets = `{
  "block": {
    "attributes": {
      "filter": {
        "description": "Filter string, adhering to the rules in List-operation filtering (https://cloud.google.com/secret-manager/docs/filtering).\nList only secrets matching the filter. If filter is empty, all secrets are listed.",
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
              "name": "string",
              "project": "string",
              "replication": [
                "list",
                [
                  "object",
                  {
                    "auto": [
                      "list",
                      [
                        "object",
                        {
                          "customer_managed_encryption": [
                            "list",
                            [
                              "object",
                              {
                                "kms_key_name": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "user_managed": [
                      "list",
                      [
                        "object",
                        {
                          "replicas": [
                            "list",
                            [
                              "object",
                              {
                                "customer_managed_encryption": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "kms_key_name": "string"
                                    }
                                  ]
                                ],
                                "location": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
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
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSecretManagerSecretsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecretManagerSecrets), &result)
	return &result
}
