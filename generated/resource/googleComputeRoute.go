package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRoute = `{
  "block": {
    "attributes": {
      "as_paths": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "as_lists": [
                "list",
                "number"
              ],
              "path_segment_type": "string"
            }
          ]
        ]
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property\nwhen you create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dest_range": {
        "description": "The destination range of outgoing packets that this route applies to.\nOnly IPv4 is supported.",
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
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035.  Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the\nlast character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The network that this route applies to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "next_hop_gateway": {
        "description": "URL to a gateway that should handle matching packets.\nCurrently, you can only specify the internet gateway, using a full or\npartial valid URL:\n* 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'\n* 'projects/project/global/gateways/default-internet-gateway'\n* 'global/gateways/default-internet-gateway'\n* The string 'default-internet-gateway'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_hub": {
        "computed": true,
        "description": "The hub network that should handle matching packets, which should conform to RFC1035.",
        "description_kind": "plain",
        "type": "string"
      },
      "next_hop_ilb": {
        "description": "The IP address or URL to a forwarding rule of type\nloadBalancingScheme=INTERNAL that should handle matching\npackets.\n\nWith the GA provider you can only specify the forwarding\nrule as a partial or full URL. For example, the following\nare all valid values:\n* 10.128.0.56\n* https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule\n* regions/region/forwardingRules/forwardingRule\n\nWhen the beta provider, you can also specify the IP address\nof a forwarding rule from the same VPC or any peered VPC.\n\nNote that this can only be used when the destinationRange is\na public (non-RFC 1918) IP CIDR range.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_instance": {
        "description": "URL to an instance that should handle matching packets.\nYou can specify this as a full or partial URL. For example:\n* 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance'\n* 'projects/project/zones/zone/instances/instance'\n* 'zones/zone/instances/instance'\n* Just the instance name, with the zone in 'next_hop_instance_zone'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_instance_zone": {
        "computed": true,
        "description": "The zone of the instance specified in next_hop_instance. Omit if next_hop_instance is specified as a URL.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_inter_region_cost": {
        "computed": true,
        "description": "Internal fixed region-to-region cost that Google Cloud calculates based on factors such as network performance, distance, and available bandwidth between regions.",
        "description_kind": "plain",
        "type": "string"
      },
      "next_hop_ip": {
        "computed": true,
        "description": "Network IP address of an instance that should handle matching packets.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_med": {
        "computed": true,
        "description": "Multi-Exit Discriminator, a BGP route metric that indicates the desirability of a particular route in a network.",
        "description_kind": "plain",
        "type": "string"
      },
      "next_hop_network": {
        "computed": true,
        "description": "URL to a Network that should handle matching packets.",
        "description_kind": "plain",
        "type": "string"
      },
      "next_hop_origin": {
        "computed": true,
        "description": "Indicates the origin of the route. Can be IGP (Interior Gateway Protocol), EGP (Exterior Gateway Protocol), or INCOMPLETE.",
        "description_kind": "plain",
        "type": "string"
      },
      "next_hop_peering": {
        "computed": true,
        "description": "The network peering name that should handle matching packets, which should conform to RFC1035.",
        "description_kind": "plain",
        "type": "string"
      },
      "next_hop_vpn_tunnel": {
        "description": "URL to a VpnTunnel that should handle matching packets.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description": "The priority of this route. Priority is used to break ties in cases\nwhere there is more than one matching route of equal prefix length.\n\nIn the case of two routes with equal prefix length, the one with the\nlowest-numbered priority value wins.\n\nDefault value is 1000. Valid range is 0 through 65535.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_status": {
        "computed": true,
        "description": "The status of the route, which can be one of the following values:\n- 'ACTIVE' for an active route\n- 'INACTIVE' for an inactive route",
        "description_kind": "plain",
        "type": "string"
      },
      "route_type": {
        "computed": true,
        "description": "The type of this route, which can be one of the following values:\n- 'TRANSIT' for a transit route that this router learned from another Cloud Router and will readvertise to one of its BGP peers\n- 'SUBNET' for a route from a subnet of the VPC\n- 'BGP' for a route learned from a BGP peer of this router\n- 'STATIC' for a static route",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description": "A list of instance tags to which this route applies.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "warnings": {
        "computed": true,
        "description": "If potential misconfigurations are detected for this route, this field will be populated with warning messages.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string",
              "data": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "value": "string"
                  }
                ]
              ],
              "message": "string"
            }
          ]
        ]
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

func GoogleComputeRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRoute), &result)
	return &result
}
