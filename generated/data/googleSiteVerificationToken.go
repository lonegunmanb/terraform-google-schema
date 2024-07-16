package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSiteVerificationToken = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identifier": {
        "description": "The site identifier. If the type is set to SITE, the identifier is a URL. If the type is\nset to INET_DOMAIN, the identifier is a domain name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "token": {
        "computed": true,
        "description": "The returned token for use in subsequent verification steps.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "The type of resource to be verified, either a domain or a web site. Possible values: [\"INET_DOMAIN\", \"SITE\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "verification_method": {
        "description": "The verification method for the Site Verification system to use to verify\nthis site or domain. Possible values: [\"ANALYTICS\", \"DNS_CNAME\", \"DNS_TXT\", \"FILE\", \"META\", \"TAG_MANAGER\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func GoogleSiteVerificationTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSiteVerificationToken), &result)
	return &result
}
