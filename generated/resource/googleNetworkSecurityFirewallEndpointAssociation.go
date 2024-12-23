package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityFirewallEndpointAssociation = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the firewall endpoint was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "description": "Whether the association is disabled. True indicates that traffic will not be intercepted.\n\n~\u003e **Note:** The API will reject the request if this value is set to true when creating the resource,\notherwise on an update the association can be disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "firewall_endpoint": {
        "description": "The URL of the firewall endpoint that is being associated.",
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
        "description": "The location (zone) of the firewall endpoint association.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the firewall endpoint association resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The URL of the network that is being associated.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The name of the parent this firewall endpoint association belongs to.\nFormat: projects/{project_id}.",
        "description_kind": "plain",
        "optional": true,
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
      "tls_inspection_policy": {
        "description": "The URL of the TlsInspectionPolicy that is being associated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func GoogleNetworkSecurityFirewallEndpointAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityFirewallEndpointAssociation), &result)
	return &result
}
