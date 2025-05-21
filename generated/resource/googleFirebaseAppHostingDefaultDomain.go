package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaseAppHostingDefaultDomain = `{
  "block": {
    "attributes": {
      "backend": {
        "description": "The ID of the Backend that this Domain is associated with",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time at which the domain was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "computed": true,
        "description": "Whether the domain is disabled. Defaults to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "domain_id": {
        "description": "Id of the domain. For default domain, it should be {{backend}}--{{project_id}}.{{location}}.hosted.app",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Server-computed checksum based on other values; may be sent\non update or delete to ensure operation is done on expected resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the Backend that this Domain is associated with",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the domain, e.g.\n'projects/{project}/locations/{locationId}/backends/{backendId}/domains/{domainId}'",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-assigned, unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time at which the domain was last updated.",
        "description_kind": "plain",
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

func GoogleFirebaseAppHostingDefaultDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaseAppHostingDefaultDomain), &result)
	return &result
}
