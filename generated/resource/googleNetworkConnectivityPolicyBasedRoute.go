package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkConnectivityPolicyBasedRoute = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time when the policy-based route was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
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
      "kind": {
        "computed": true,
        "description": "Type of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name of the policy based route.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "Fully-qualified URL of the network that this route applies to, for example: projects/my-project/global/networks/my-network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "next_hop_ilb_ip": {
        "description": "The IP address of a global-access-enabled L4 ILB that is the next hop for matching packets.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "next_hop_other_routes": {
        "description": "Other routes that will be referenced to determine the next hop of the packet. Possible values: [\"DEFAULT_ROUTING\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description": "The priority of this policy-based route. Priority is used to break ties in cases where there are more than one matching policy-based routes found. In cases where multiple policy-based routes are matched, the one with the lowest-numbered priority value wins. The default value is 1000. The priority value must be from 1 to 65535, inclusive.",
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
        "description": "Time when the policy-based route was created.",
        "description_kind": "plain",
        "type": "string"
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
                "map",
                "string"
              ],
              "warning_message": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "dest_range": {
              "description": "The destination IP range of outgoing packets that this policy-based route applies to. Default is \"0.0.0.0/0\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_protocol": {
              "description": "The IP protocol that this policy-based route applies to. Valid values are 'TCP', 'UDP', and 'ALL'. Default is 'ALL'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol_version": {
              "description": "Internet protocol versions this policy-based route applies to. Possible values: [\"IPV4\", \"IPV6\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "src_range": {
              "description": "The source IP range of outgoing packets that this policy-based route applies to. Default is \"0.0.0.0/0\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The filter to match L4 traffic.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "interconnect_attachment": {
        "block": {
          "attributes": {
            "region": {
              "description": "Cloud region to install this policy-based route on for Interconnect attachments. Use 'all' to install it on all Interconnect attachments.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The interconnect attachments that this policy-based route applies to.",
          "description_kind": "plain"
        },
        "max_items": 1,
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
      "virtual_machine": {
        "block": {
          "attributes": {
            "tags": {
              "description": "A list of VM instance tags that this policy-based route applies to. VM instances that have ANY of tags specified here will install this PBR.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "VM instances to which this policy-based route applies to.",
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

func GoogleNetworkConnectivityPolicyBasedRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkConnectivityPolicyBasedRoute), &result)
	return &result
}
