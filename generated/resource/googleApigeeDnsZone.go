package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeDnsZone = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Description for the zone.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dns_zone_id": {
        "description": "ID of the dns zone.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain": {
        "description": "Doamin for the zone.",
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
      "name": {
        "computed": true,
        "description": "Name of the Dns Zone in the following format:\norganizations/{organization}/dnsZones/{dnsZone}.",
        "description_kind": "plain",
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee instance,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "peering_config": {
        "block": {
          "attributes": {
            "target_network_id": {
              "description": "The name of the producer VPC network.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_project_id": {
              "description": "The ID of the project that contains the producer VPC network.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Peering zone config",
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

func GoogleApigeeDnsZoneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeDnsZone), &result)
	return &result
}
