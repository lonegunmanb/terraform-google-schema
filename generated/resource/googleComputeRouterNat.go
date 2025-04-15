package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterNat = `{
  "block": {
    "attributes": {
      "auto_network_tier": {
        "computed": true,
        "description": "The network tier to use when automatically reserving NAT IP addresses.\nMust be one of: PREMIUM, STANDARD. If not specified, then the current\nproject-level default tier is used. Possible values: [\"PREMIUM\", \"STANDARD\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "drain_nat_ips": {
        "computed": true,
        "description": "A list of URLs of the IP resources to be drained. These IPs must be\nvalid static external IPs that have been assigned to the NAT.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "enable_dynamic_port_allocation": {
        "computed": true,
        "description": "Enable Dynamic Port Allocation.\nIf minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.\nIf minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.\nIf maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.\nIf maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.\n\nMutually exclusive with enableEndpointIndependentMapping.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_endpoint_independent_mapping": {
        "computed": true,
        "description": "Enable endpoint independent mapping.\nFor more information see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint_types": {
        "computed": true,
        "description": "Specifies the endpoint Types supported by the NAT Gateway.\nSupported values include:\n      'ENDPOINT_TYPE_VM', 'ENDPOINT_TYPE_SWG',\n      'ENDPOINT_TYPE_MANAGED_PROXY_LB'.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "icmp_idle_timeout_sec": {
        "description": "Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "initial_nat_ips": {
        "description": "Self-links of NAT IPs to be used as initial value for creation alongside a RouterNatAddress resource.\nConflicts with natIps and drainNatIps. Only valid if natIpAllocateOption is set to MANUAL_ONLY.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "max_ports_per_vm": {
        "description": "Maximum number of ports allocated to a VM from this NAT.\nThis field can only be set when enableDynamicPortAllocation is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_ports_per_vm": {
        "computed": true,
        "description": "Minimum number of ports allocated to a VM from this NAT. Defaults to 64 for static port allocation and 32 dynamic port allocation if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "Name of the NAT service. The name must be 1-63 characters long and\ncomply with RFC1035.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_ip_allocate_option": {
        "description": "How external IPs should be allocated for this NAT. Valid values are\n'AUTO_ONLY' for only allowing NAT IPs allocated by Google Cloud\nPlatform, or 'MANUAL_ONLY' for only user-allocated NAT IP addresses. Possible values: [\"MANUAL_ONLY\", \"AUTO_ONLY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_ips": {
        "computed": true,
        "description": "Self-links of NAT IPs. Only valid if natIpAllocateOption\nis set to MANUAL_ONLY.\nIf this field is used alongside with a count created list of address resources 'google_compute_address.foobar.*.self_link',\nthe access level resource for the address resource must have a 'lifecycle' block with 'create_before_destroy = true' so\nthe number of resources can be increased/decreased without triggering the 'resourceInUseByAnotherResource' error.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where the router and NAT reside.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router": {
        "description": "The name of the Cloud Router in which this NAT will be configured.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_subnetwork_ip_ranges_to_nat": {
        "description": "How NAT should be configured per Subnetwork.\nIf 'ALL_SUBNETWORKS_ALL_IP_RANGES', all of the\nIP ranges in every Subnetwork are allowed to Nat.\nIf 'ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES', all of the primary IP\nranges in every Subnetwork are allowed to Nat.\n'LIST_OF_SUBNETWORKS': A list of Subnetworks are allowed to Nat\n(specified in the field subnetwork below). Note that if this field\ncontains ALL_SUBNETWORKS_ALL_IP_RANGES or\nALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any\nother RouterNat section in any Router for this network in this region. Possible values: [\"ALL_SUBNETWORKS_ALL_IP_RANGES\", \"ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES\", \"LIST_OF_SUBNETWORKS\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tcp_established_idle_timeout_sec": {
        "description": "Timeout (in seconds) for TCP established connections.\nDefaults to 1200s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tcp_time_wait_timeout_sec": {
        "description": "Timeout (in seconds) for TCP connections that are in TIME_WAIT state.\nDefaults to 120s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tcp_transitory_idle_timeout_sec": {
        "description": "Timeout (in seconds) for TCP transitory connections.\nDefaults to 30s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "type": {
        "description": "Indicates whether this NAT is used for public or private IP translation.\nIf unspecified, it defaults to PUBLIC.\nIf 'PUBLIC' NAT used for public IP translation.\nIf 'PRIVATE' NAT used for private IP translation. Default value: \"PUBLIC\" Possible values: [\"PUBLIC\", \"PRIVATE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "udp_idle_timeout_sec": {
        "description": "Timeout (in seconds) for UDP connections. Defaults to 30s if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "log_config": {
        "block": {
          "attributes": {
            "enable": {
              "description": "Indicates whether or not to export logs.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "filter": {
              "description": "Specifies the desired filtering of logs on this NAT. Possible values: [\"ERRORS_ONLY\", \"TRANSLATIONS_ONLY\", \"ALL\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for logging on NAT",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rules": {
        "block": {
          "attributes": {
            "description": {
              "description": "An optional description of this rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "match": {
              "description": "CEL expression that specifies the match condition that egress traffic from a VM is evaluated against.\nIf it evaluates to true, the corresponding action is enforced.\n\nThe following examples are valid match expressions for public NAT:\n\n\"inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')\"\n\n\"destination.ip == '1.1.0.1' || destination.ip == '8.8.8.8'\"\n\nThe following example is a valid match expression for private NAT:\n\n\"nexthop.hub == 'https://networkconnectivity.googleapis.com/v1alpha1/projects/my-project/global/hub/hub-1'\"",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_number": {
              "description": "An integer uniquely identifying a rule in the list.\nThe rule number must be a positive value between 0 and 65000, and must be unique among rules within a NAT.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "source_nat_active_ips": {
                    "description": "A list of URLs of the IP resources used for this NAT rule.\nThese IP addresses must be valid static external IP addresses assigned to the project.\nThis field is used for public NAT.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "source_nat_active_ranges": {
                    "description": "A list of URLs of the subnetworks used as source ranges for this NAT Rule.\nThese subnetworks must have purpose set to PRIVATE_NAT.\nThis field is used for private NAT.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "source_nat_drain_ips": {
                    "description": "A list of URLs of the IP resources to be drained.\nThese IPs must be valid static external IPs that have been assigned to the NAT.\nThese IPs should be used for updating/patching a NAT rule only.\nThis field is used for public NAT.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "source_nat_drain_ranges": {
                    "description": "A list of URLs of subnetworks representing source ranges to be drained.\nThis is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.\nThis field is used for private NAT.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "The action to be enforced for traffic that matches this rule.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A list of rules associated with this NAT.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "subnetwork": {
        "block": {
          "attributes": {
            "name": {
              "description": "Self-link of subnetwork to NAT",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secondary_ip_range_names": {
              "description": "List of the secondary ranges of the subnetwork that are allowed\nto use NAT. This can be populated only if\n'LIST_OF_SECONDARY_IP_RANGES' is one of the values in\nsourceIpRangesToNat",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "source_ip_ranges_to_nat": {
              "description": "List of options for which source IPs in the subnetwork\nshould have NAT enabled. Supported values include:\n'ALL_IP_RANGES', 'LIST_OF_SECONDARY_IP_RANGES',\n'PRIMARY_IP_RANGE'.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "One or more subnetwork NAT configurations. Only used if\n'source_subnetwork_ip_ranges_to_nat' is set to 'LIST_OF_SUBNETWORKS'",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRouterNatSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterNat), &result)
	return &result
}
