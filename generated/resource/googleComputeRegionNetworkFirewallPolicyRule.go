package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionNetworkFirewallPolicyRule = `{
  "block": {
    "attributes": {
      "action": {
        "description": "The Action to perform when the client connection triggers the rule. Valid actions are \"allow\", \"deny\", \"goto_next\" and \"apply_security_profile_group\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description for this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "direction": {
        "description": "The direction in which this rule applies. Possible values: [\"INGRESS\", \"EGRESS\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disabled": {
        "description": "Denotes whether the firewall policy rule is disabled.\nWhen set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist.\nIf this is unspecified, the firewall policy rule will be enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_logging": {
        "description": "Denotes whether to enable logging for a particular rule.\nIf logging is enabled, logs will be exported to the configured export destination in Stackdriver.\nLogs may be exported to BigQuery or Pub/Sub.\nNote: you cannot enable logging on \"goto_next\" rules.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "firewall_policy": {
        "description": "The firewall policy of the resource.",
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
      "kind": {
        "computed": true,
        "description": "Type of the resource. Always 'compute#firewallPolicyRule' for firewall policy rules",
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "description": "An integer indicating the priority of a rule in the list.\nThe priority must be a positive value between 0 and 2147483647.\nRules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The location of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_name": {
        "description": "An optional name for the rule. This field is not a unique identifier and can be updated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_tuple_count": {
        "computed": true,
        "description": "Calculation of the complexity of a single firewall policy rule.",
        "description_kind": "plain",
        "type": "number"
      },
      "security_profile_group": {
        "description": "A fully-qualified URL of a SecurityProfile resource instance.\nExample: https://networksecurity.googleapis.com/v1/projects/{project}/locations/{location}/securityProfileGroups/my-security-profile-group\nMust be specified if action = 'apply_security_profile_group' and cannot be specified for other actions.\n\nSecurity Profile Group and Firewall Policy Rule must be in the same scope.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_service_accounts": {
        "description": "A list of service accounts indicating the sets of instances that are applied with this rule.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tls_inspect": {
        "description": "Boolean flag indicating if the traffic should be TLS decrypted.\nCan be set only if action = 'apply_security_profile_group' and cannot be set for other actions.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "match": {
        "block": {
          "attributes": {
            "dest_address_groups": {
              "description": "Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dest_fqdns": {
              "description": "Fully Qualified Domain Name (FQDN) which should be matched against traffic destination. Maximum number of destination fqdn allowed is 100.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dest_ip_ranges": {
              "description": "CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dest_region_codes": {
              "description": "Region codes whose IP addresses will be used to match for destination of traffic. Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex.\"US\" Maximum number of dest region codes allowed is 5000.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dest_threat_intelligences": {
              "description": "Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic destination.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_address_groups": {
              "description": "Address groups which should be matched against the traffic source. Maximum number of source address groups is 10.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_fqdns": {
              "description": "Fully Qualified Domain Name (FQDN) which should be matched against traffic source. Maximum number of source fqdn allowed is 100.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_ip_ranges": {
              "description": "CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_region_codes": {
              "description": "Region codes whose IP addresses will be used to match for source of traffic. Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex.\"US\" Maximum number of source region codes allowed is 5000.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_threat_intelligences": {
              "description": "Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic source.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "layer4_configs": {
              "block": {
                "attributes": {
                  "ip_protocol": {
                    "description": "The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule.\nThis value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp), or the IP protocol number.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ports": {
                    "description": "An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port.\nExample inputs include: [\"22\"], [\"80\",\"443\"], and [\"12345-12349\"].",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Pairs of IP protocols and ports that the rule should match.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "src_secure_tags": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the secure tag, created with TagManager's TagValue API.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "State of the secure tag, either EFFECTIVE or INEFFECTIVE. A secure tag is INEFFECTIVE when it is deleted or its network is deleted.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "List of secure tag values, which should be matched at the source of the traffic. For INGRESS rule, if all the srcSecureTag are INEFFECTIVE, and there is no srcIpRange, this rule will be ignored. Maximum number of source tag values allowed is 256.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "target_secure_tags": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the secure tag, created with TagManager's TagValue API.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "State of the secure tag, either EFFECTIVE or INEFFECTIVE. A secure tag is INEFFECTIVE when it is deleted or its network is deleted.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "A list of secure tags that controls which instances the firewall rule applies to.\nIf targetSecureTag are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored.\ntargetSecureTag may not be set at the same time as targetServiceAccounts. If neither targetServiceAccounts nor targetSecureTag are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.",
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

func GoogleComputeRegionNetworkFirewallPolicyRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionNetworkFirewallPolicyRule), &result)
	return &result
}
