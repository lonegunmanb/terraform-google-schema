package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEdgecontainerVpnConnection = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "The canonical Cluster name to connect to. It is in the form of projects/{project}/locations/{location}/clusters/{cluster}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the VPN connection was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "details": {
        "computed": true,
        "description": "A nested object resource",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloud_router": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "cloud_vpns": [
                "list",
                [
                  "object",
                  {
                    "gateway": "string"
                  }
                ]
              ],
              "error": "string",
              "state": "string"
            }
          ]
        ]
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
      "enable_high_availability": {
        "computed": true,
        "description": "Whether this VPN connection has HA enabled on cluster side. If enabled, when creating VPN connection we will attempt to use 2 ANG floating IPs.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels associated with this resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Google Cloud Platform location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of VPN connection",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_gateway_ip": {
        "description": "NAT gateway IP, or WAN IP address. If a customer has multiple NAT IPs, the customer needs to configure NAT such that only one external IP maps to the GMEC Anthos cluster.\nThis is empty if NAT is not used.",
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
      "router": {
        "description": "The VPN connection Cloud Router name.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The time when the VPN connection was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc": {
        "description": "The network ID of VPC to connect to.",
        "description_kind": "plain",
        "optional": true,
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
      },
      "vpc_project": {
        "block": {
          "attributes": {
            "project_id": {
              "description": "The project of the VPC to connect to. If not specified, it is the same as the cluster project.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Project detail of the VPC network. Required if VPC is in a different project than the cluster project.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleEdgecontainerVpnConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEdgecontainerVpnConnection), &result)
	return &result
}
