package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSubnetwork = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource. This field can be set only at resource\ncreation time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_flow_logs": {
        "computed": true,
        "deprecated": true,
        "description": "Whether to enable flow logging for this subnetwork. If this field is not explicitly set,\nit will not appear in get listings. If not set the default behavior is determined by the\norg policy, if there is no org policy specified, then it will default to disabled.\nThis field isn't supported if the subnet purpose field is set to REGIONAL_MANAGED_PROXY.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "external_ipv6_prefix": {
        "computed": true,
        "description": "The range of external IPv6 addresses that are owned by this subnetwork.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "deprecated": true,
        "description": "Fingerprint of this resource. This field is used internally during updates of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_address": {
        "computed": true,
        "description": "The gateway address for default routes to reach destination addresses\noutside this subnetwork.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internal_ipv6_prefix": {
        "computed": true,
        "description": "The internal IPv6 address range that is assigned to this subnetwork.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_cidr_range": {
        "computed": true,
        "description": "The range of internal addresses that are owned by this subnetwork.\nProvide this property when you create the subnetwork. For example,\n10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and\nnon-overlapping within a network. Only IPv4 is supported.\nField is optional when 'reserved_internal_range' is defined, otherwise required.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_collection": {
        "description": "Resource reference of a PublicDelegatedPrefix. The PDP must be a sub-PDP\nin EXTERNAL_IPV6_SUBNETWORK_CREATION mode.\nUse one of the following formats to specify a sub-PDP when creating an\nIPv6 NetLB forwarding rule using BYOIP:\nFull resource URL, as in:\n  * 'https://www.googleapis.com/compute/v1/projects/{{projectId}}/regions/{{region}}/publicDelegatedPrefixes/{{sub-pdp-name}}'\nPartial URL, as in:\n  * 'projects/{{projectId}}/regions/region/publicDelegatedPrefixes/{{sub-pdp-name}}'\n  * 'regions/{{region}}/publicDelegatedPrefixes/{{sub-pdp-name}}'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_access_type": {
        "description": "The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation\nor the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet\ncannot enable direct path. Possible values: [\"EXTERNAL\", \"INTERNAL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_range": {
        "computed": true,
        "description": "The range of internal IPv6 addresses that are owned by this subnetwork.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv6_gce_endpoint": {
        "computed": true,
        "description": "Possible endpoints of this subnetwork. It can be one of the following:\n  * VM_ONLY: The subnetwork can be used for creating instances and IPv6 addresses with VM endpoint type. Such a subnetwork\n  gets external IPv6 ranges from a public delegated prefix and cannot be used to create NetLb.\n  * VM_AND_FR: The subnetwork can be used for creating both VM instances and Forwarding Rules. It can also be used to reserve\n  IPv6 addresses with both VM and FR endpoint types. Such a subnetwork gets its IPv6 range from Google IP Pool directly.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the resource, provided by the client when initially\ncreating the resource. The name must be 1-63 characters long, and\ncomply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which\nmeans the first character must be a lowercase letter, and all\nfollowing characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The network this subnet belongs to.\nOnly networks that are in the distributed mode can have subnetworks.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_ip_google_access": {
        "computed": true,
        "description": "When enabled, VMs in this subnetwork without external IP addresses can\naccess Google APIs and services by using Private Google Access.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "private_ipv6_google_access": {
        "computed": true,
        "description": "The private IPv6 google access type for the VMs in this subnet.",
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
      "purpose": {
        "computed": true,
        "description": "The purpose of the resource. This field can be either 'PRIVATE', 'REGIONAL_MANAGED_PROXY', 'GLOBAL_MANAGED_PROXY', 'PRIVATE_SERVICE_CONNECT', 'PEER_MIGRATION' or 'PRIVATE_NAT'([Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html)).\nA subnet with purpose set to 'REGIONAL_MANAGED_PROXY' is a user-created subnetwork that is reserved for regional Envoy-based load balancers.\nA subnetwork in a given region with purpose set to 'GLOBAL_MANAGED_PROXY' is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.\nA subnetwork with purpose set to 'PRIVATE_SERVICE_CONNECT' reserves the subnet for hosting a Private Service Connect published service.\nA subnetwork with purpose set to 'PEER_MIGRATION' is a user created subnetwork that is reserved for migrating resources from one peered network to another.\nA subnetwork with purpose set to 'PRIVATE_NAT' is used as source range for Private NAT gateways.\nNote that 'REGIONAL_MANAGED_PROXY' is the preferred setting for all regional Envoy load balancers.\nIf unspecified, the purpose defaults to 'PRIVATE'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The GCP region for this subnetwork.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_internal_range": {
        "description": "The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com'\nE.g. 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role": {
        "description": "The role of subnetwork.\nCurrently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'.\nThe value can be set to 'ACTIVE' or 'BACKUP'.\nAn 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers in a region.\nA 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible values: [\"ACTIVE\", \"BACKUP\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "send_secondary_ip_range_if_empty": {
        "description": "Controls the removal behavior of secondary_ip_range.\nWhen false, removing secondary_ip_range from config will not produce a diff as\nthe provider will default to the API's value.\nWhen true, the provider will treat removing secondary_ip_range as sending an\nempty list of secondary IP ranges to the API.\nDefaults to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "stack_type": {
        "computed": true,
        "description": "The stack type for this subnet to identify whether the IPv6 feature is enabled or not.\nIf not specified IPV4_ONLY will be used. Possible values: [\"IPV4_ONLY\", \"IPV4_IPV6\", \"IPV6_ONLY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "'The state of the subnetwork, which can be one of the following values:\n READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose\n set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained.\n A subnetwork that is draining cannot be used or modified until it reaches a status of READY'",
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork_id": {
        "computed": true,
        "description": "The unique identifier number for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "log_config": {
        "block": {
          "attributes": {
            "aggregation_interval": {
              "description": "Can only be specified if VPC flow logging for this subnetwork is enabled.\nToggles the aggregation interval for collecting flow logs. Increasing the\ninterval time will reduce the amount of generated flow logs for long\nlasting connections. Default is an interval of 5 seconds per connection. Default value: \"INTERVAL_5_SEC\" Possible values: [\"INTERVAL_5_SEC\", \"INTERVAL_30_SEC\", \"INTERVAL_1_MIN\", \"INTERVAL_5_MIN\", \"INTERVAL_10_MIN\", \"INTERVAL_15_MIN\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filter_expr": {
              "description": "Export filter used to define which VPC flow logs should be logged, as as CEL expression. See\nhttps://cloud.google.com/vpc/docs/flow-logs#filtering for details on how to format this field.\nThe default value is 'true', which evaluates to include everything.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flow_sampling": {
              "description": "Can only be specified if VPC flow logging for this subnetwork is enabled.\nThe value of the field must be in [0, 1]. Set the sampling rate of VPC\nflow logs within the subnetwork where 1.0 means all collected logs are\nreported and 0.0 means no logs are reported. Default is 0.5 which means\nhalf of all collected logs are reported.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metadata": {
              "description": "Can only be specified if VPC flow logging for this subnetwork is enabled.\nConfigures whether metadata fields should be added to the reported VPC\nflow logs. Default value: \"INCLUDE_ALL_METADATA\" Possible values: [\"EXCLUDE_ALL_METADATA\", \"INCLUDE_ALL_METADATA\", \"CUSTOM_METADATA\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata_fields": {
              "description": "List of metadata fields that should be added to reported logs.\nCan only be specified if VPC flow logs for this subnetwork is enabled and \"metadata\" is set to CUSTOM_METADATA.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "This field denotes the VPC flow logging options for this subnetwork. If\nlogging is enabled, logs are exported to Cloud Logging. Flow logging\nisn't supported if the subnet 'purpose' field is set to subnetwork is\n'REGIONAL_MANAGED_PROXY' or 'GLOBAL_MANAGED_PROXY'.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "params": {
        "block": {
          "attributes": {
            "resource_manager_tags": {
              "description": "Resource manager tags to be bound to the subnetwork. Tag keys and values have the\nsame definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},\nand values are in the format tagValues/456. The field is ignored when empty.\nThe field is immutable and causes resource replacement when mutated. This field is only\nset at create time and modifying this field after creation will trigger recreation.\nTo apply tags to an existing resource, see the google_tags_tag_binding resource.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description": "Additional params passed with the request, but not persisted as part of resource payload",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "secondary_ip_range": {
        "block": {
          "attributes": {
            "ip_cidr_range": {
              "computed": true,
              "description": "The range of IP addresses belonging to this subnetwork secondary\nrange. Provide this property when you create the subnetwork.\nRanges must be unique and non-overlapping with all primary and\nsecondary IP ranges within a network. Only IPv4 is supported.\nField is optional when 'reserved_internal_range' is defined, otherwise required.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "range_name": {
              "description": "The name associated with this subnetwork secondary range, used\nwhen adding an alias IP range to a VM instance. The name must\nbe 1-63 characters long, and comply with RFC1035. The name\nmust be unique within the subnetwork.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "reserved_internal_range": {
              "description": "The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com'\nE.g. 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "An array of configurations for secondary IP ranges for VM instances\ncontained in this subnetwork. The primary IP of such VM must belong\nto the primary ipCidrRange of the subnetwork. The alias IPs may belong\nto either primary or secondary ranges.\n\n**Note**: This field uses [attr-as-block mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) to avoid\nbreaking users during the 0.12 upgrade. To explicitly send a list of zero objects,\nset 'send_secondary_ip_range_if_empty = true'",
          "description_kind": "plain"
        },
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeSubnetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSubnetwork), &result)
	return &result
}
