package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCertificateManagerCertificates = `{
  "block": {
    "attributes": {
      "certificates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "effective_labels": [
                "map",
                "string"
              ],
              "labels": [
                "map",
                "string"
              ],
              "location": "string",
              "managed": [
                "list",
                [
                  "object",
                  {
                    "authorization_attempt_info": [
                      "list",
                      [
                        "object",
                        {
                          "details": "string",
                          "domain": "string",
                          "failure_reason": "string",
                          "state": "string"
                        }
                      ]
                    ],
                    "dns_authorizations": [
                      "list",
                      "string"
                    ],
                    "domains": [
                      "list",
                      "string"
                    ],
                    "issuance_config": "string",
                    "provisioning_issue": [
                      "list",
                      [
                        "object",
                        {
                          "details": "string",
                          "reason": "string"
                        }
                      ]
                    ],
                    "state": "string"
                  }
                ]
              ],
              "name": "string",
              "project": "string",
              "san_dnsnames": [
                "list",
                "string"
              ],
              "scope": "string",
              "terraform_labels": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "filter": {
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
      "region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCertificateManagerCertificatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCertificateManagerCertificates), &result)
	return &result
}
