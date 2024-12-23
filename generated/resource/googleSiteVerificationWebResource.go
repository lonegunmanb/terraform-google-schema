package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSiteVerificationWebResource = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owners": {
        "computed": true,
        "description": "The email addresses of all direct, verified owners of this exact property. Indirect owners —\nfor example verified owners of the containing domain—are not included in this list.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "verification_method": {
        "description": "The verification method for the Site Verification system to use to verify\nthis site or domain. Possible values: [\"ANALYTICS\", \"DNS_CNAME\", \"DNS_TXT\", \"FILE\", \"META\", \"TAG_MANAGER\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "web_resource_id": {
        "computed": true,
        "description": "The string used to identify this web resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "site": {
        "block": {
          "attributes": {
            "identifier": {
              "description": "The site identifier. If the type is set to SITE, the identifier is a URL. If the type is\nset to INET_DOMAIN, the identifier is a domain name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "The type of resource to be verified. Possible values: [\"INET_DOMAIN\", \"SITE\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Container for the address and type of a site for which a verification token will be verified.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleSiteVerificationWebResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSiteVerificationWebResource), &result)
	return &result
}
