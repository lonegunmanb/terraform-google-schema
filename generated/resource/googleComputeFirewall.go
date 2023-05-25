package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeFirewall = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_ranges": {
        "computed": true,
        "description": "If destination ranges are specified, the firewall will apply only to\ntraffic that has destination IP address in these ranges. These ranges\nmust be expressed in CIDR format. Only IPv4 is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "direction": {
        "computed": true,
        "description": "Direction of traffic to which this firewall applies; default is\nINGRESS. Note: For INGRESS traffic, it is NOT supported to specify\ndestinationRanges; For EGRESS traffic, it is NOT supported to specify\nsourceRanges OR sourceTags. Possible values: [\"INGRESS\", \"EGRESS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description": "Denotes whether the firewall rule is disabled, i.e not applied to the\nnetwork it is associated with. When set to true, the firewall rule is\nnot enforced and the network behaves as if it did not exist. If this\nis unspecified, the firewall rule will be enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_logging": {
        "computed": true,
        "deprecated": true,
        "description": "This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported to Stackdriver.",
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
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The name or self_link of the network to attach this firewall to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "description": "Priority for this rule. This is an integer between 0 and 65535, both\ninclusive. When not specified, the value assumed is 1000. Relative\npriorities determine precedence of conflicting rules. Lower value of\npriority implies higher precedence (eg, a rule with priority 0 has\nhigher precedence than a rule with priority 1). DENY rules take\nprecedence over ALLOW rules having equal priority.",
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_ranges": {
        "computed": true,
        "description": "If source ranges are specified, the firewall will apply only to\ntraffic that has source IP address in these ranges. These ranges must\nbe expressed in CIDR format. One or both of sourceRanges and\nsourceTags may be set. If both properties are set, the firewall will\napply to traffic that has source IP address within sourceRanges OR the\nsource IP that belongs to a tag listed in the sourceTags property. The\nconnection does not need to match both properties for the firewall to\napply. Only IPv4 is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source_service_accounts": {
        "description": "If source service accounts are specified, the firewall will apply only\nto traffic originating from an instance with a service account in this\nlist. Source service accounts cannot be used to control traffic to an\ninstance's external IP address because service accounts are associated\nwith an instance, not an IP address. sourceRanges can be set at the\nsame time as sourceServiceAccounts. If both are set, the firewall will\napply to traffic that has source IP address within sourceRanges OR the\nsource IP belongs to an instance with service account listed in\nsourceServiceAccount. The connection does not need to match both\nproperties for the firewall to apply. sourceServiceAccounts cannot be\nused at the same time as sourceTags or targetTags.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source_tags": {
        "description": "If source tags are specified, the firewall will apply only to traffic\nwith source IP that belongs to a tag listed in source tags. Source\ntags cannot be used to control traffic to an instance's external IP\naddress. Because tags are associated with an instance, not an IP\naddress. One or both of sourceRanges and sourceTags may be set. If\nboth properties are set, the firewall will apply to traffic that has\nsource IP address within sourceRanges OR the source IP that belongs to\na tag listed in the sourceTags property. The connection does not need\nto match both properties for the firewall to apply.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "target_service_accounts": {
        "description": "A list of service accounts indicating sets of instances located in the\nnetwork that may make network connections as specified in allowed[].\ntargetServiceAccounts cannot be used at the same time as targetTags or\nsourceTags. If neither targetServiceAccounts nor targetTags are\nspecified, the firewall rule applies to all instances on the specified\nnetwork.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "target_tags": {
        "description": "A list of instance tags indicating sets of instances located in the\nnetwork that may make network connections as specified in allowed[].\nIf no targetTags are specified, the firewall rule applies to all\ninstances on the specified network.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "allow": {
        "block": {
          "attributes": {
            "ports": {
              "description": "An optional list of ports to which this rule applies. This field\nis only applicable for UDP or TCP protocol. Each entry must be\neither an integer or a range. If not specified, this rule\napplies to connections through any port.\n\nExample inputs include: [\"22\"], [\"80\",\"443\"], and\n[\"12345-12349\"].",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "protocol": {
              "description": "The IP protocol to which this rule applies. The protocol type is\nrequired when creating a firewall rule. This value can either be\none of the following well known protocol strings (tcp, udp,\nicmp, esp, ah, sctp, ipip, all), or the IP protocol number.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The list of ALLOW rules specified by this firewall. Each rule\nspecifies a protocol and port-range tuple that describes a permitted\nconnection.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "deny": {
        "block": {
          "attributes": {
            "ports": {
              "description": "An optional list of ports to which this rule applies. This field\nis only applicable for UDP or TCP protocol. Each entry must be\neither an integer or a range. If not specified, this rule\napplies to connections through any port.\n\nExample inputs include: [\"22\"], [\"80\",\"443\"], and\n[\"12345-12349\"].",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "protocol": {
              "description": "The IP protocol to which this rule applies. The protocol type is\nrequired when creating a firewall rule. This value can either be\none of the following well known protocol strings (tcp, udp,\nicmp, esp, ah, sctp, ipip, all), or the IP protocol number.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The list of DENY rules specified by this firewall. Each rule specifies\na protocol and port-range tuple that describes a denied connection.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "log_config": {
        "block": {
          "attributes": {
            "metadata": {
              "description": "This field denotes whether to include or exclude metadata for firewall logs. Possible values: [\"EXCLUDE_ALL_METADATA\", \"INCLUDE_ALL_METADATA\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "This field denotes the logging options for a particular firewall rule.\nIf defined, logging is enabled, and logs will be exported to Cloud Logging.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleComputeFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeFirewall), &result)
	return &result
}
