package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityFirewallEndpoint = `{
  "block": {
    "attributes": {
      "associated_networks": {
        "computed": true,
        "description": "List of networks that are associated with this endpoint in the local zone.\nThis is a projection of the FirewallEndpointAssociations pointing at this\nendpoint. A network will only appear in this list after traffic routing is\nfully configured. Format: projects/{project}/global/networks/{name}.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "billing_project_id": {
        "description": "Project to bill on endpoint uptime usage.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time the firewall endpoint was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A map of key/value label pairs to assign to the resource.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location (zone) of the firewall endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the firewall endpoint resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The name of the parent this firewall endpoint belongs to.\nFormat: organizations/{organization_id}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Whether reconciling is in progress, recommended per https://google.aip.dev/128.",
        "description_kind": "plain",
        "type": "bool"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Time the firewall endpoint was updated in UTC.",
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

func GoogleNetworkSecurityFirewallEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityFirewallEndpoint), &result)
	return &result
}
