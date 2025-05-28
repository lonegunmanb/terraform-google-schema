package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionNetworkFirewallPolicyWithRules = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of the resource. This field is used internally during updates of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "User-provided name of the Network firewall policy.\nThe name should be unique in the project in which the firewall policy is created.\nThe name must be 1-63 characters long, and comply with RFC1035. Specifically,\nthe name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?\nwhich means the first character must be a lowercase letter, and all following characters must be a dash,\nlowercase letter, or digit, except the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_firewall_policy_id": {
        "computed": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "predefined_rules": {
        "computed": true,
        "description": "A list of firewall policy pre-defined rules.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action": "string",
              "description": "string",
              "direction": "string",
              "disabled": "bool",
              "enable_logging": "bool",
              "match": [
                "list",
                [
                  "object",
                  {
                    "dest_address_groups": [
                      "list",
                      "string"
                    ],
                    "dest_fqdns": [
                      "list",
                      "string"
                    ],
                    "dest_ip_ranges": [
                      "list",
                      "string"
                    ],
                    "dest_region_codes": [
                      "list",
                      "string"
                    ],
                    "dest_threat_intelligences": [
                      "list",
                      "string"
                    ],
                    "layer4_config": [
                      "list",
                      [
                        "object",
                        {
                          "ip_protocol": "string",
                          "ports": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "src_address_groups": [
                      "list",
                      "string"
                    ],
                    "src_fqdns": [
                      "list",
                      "string"
                    ],
                    "src_ip_ranges": [
                      "list",
                      "string"
                    ],
                    "src_region_codes": [
                      "list",
                      "string"
                    ],
                    "src_secure_tag": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "state": "string"
                        }
                      ]
                    ],
                    "src_threat_intelligences": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "priority": "number",
              "rule_name": "string",
              "security_profile_group": "string",
              "target_secure_tag": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "state": "string"
                  }
                ]
              ],
              "target_service_accounts": [
                "list",
                "string"
              ],
              "tls_inspect": "bool"
            }
          ]
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
        "description": "The region of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_tuple_count": {
        "computed": true,
        "description": "Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.",
        "description_kind": "plain",
        "type": "number"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link_with_id": {
        "computed": true,
        "description": "Server-defined URL for this resource with the resource id.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "rule": {
        "block": {
          "attributes": {
            "action": {
              "description": "The Action to perform when the client connection triggers the rule. Can currently be either\n\"allow\", \"deny\", \"apply_security_profile_group\" or \"goto_next\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "description": "A description of the rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "direction": {
              "description": "The direction in which this rule applies. If unspecified an INGRESS rule is created. Possible values: [\"INGRESS\", \"EGRESS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disabled": {
              "description": "Denotes whether the firewall policy rule is disabled. When set to true,\nthe firewall policy rule is not enforced and traffic behaves as if it did\nnot exist. If this is unspecified, the firewall policy rule will be\nenabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_logging": {
              "description": "Denotes whether to enable logging for a particular rule.\nIf logging is enabled, logs will be exported to the\nconfigured export destination in Stackdriver.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "priority": {
              "description": "An integer indicating the priority of a rule in the list. The priority must be a value\nbetween 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the\nhighest priority and 2147483647 is the lowest priority.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "rule_name": {
              "description": "An optional name for the rule. This field is not a unique identifier\nand can be updated.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_profile_group": {
              "description": "A fully-qualified URL of a SecurityProfile resource instance.\nExample:\nhttps://networksecurity.googleapis.com/v1/projects/{project}/locations/{location}/securityProfileGroups/my-security-profile-group\nMust be specified if action is 'apply_security_profile_group'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_service_accounts": {
              "description": "A list of service accounts indicating the sets of\ninstances that are applied with this rule.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tls_inspect": {
              "description": "Boolean flag indicating if the traffic should be TLS decrypted.\nIt can be set only if action = 'apply_security_profile_group' and cannot be set for other actions.",
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
                    "description": "Address groups which should be matched against the traffic destination.\nMaximum number of destination address groups is 10.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "dest_fqdns": {
                    "description": "Fully Qualified Domain Name (FQDN) which should be matched against\ntraffic destination. Maximum number of destination fqdn allowed is 100.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "dest_ip_ranges": {
                    "description": "Destination IP address range in CIDR format. Required for\nEGRESS rules.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "dest_region_codes": {
                    "description": "Region codes whose IP addresses will be used to match for destination\nof traffic. Should be specified as 2 letter country code defined as per\nISO 3166 alpha-2 country codes. ex.\"US\"\nMaximum number of destination region codes allowed is 5000.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "dest_threat_intelligences": {
                    "description": "Names of Network Threat Intelligence lists.\nThe IPs in these lists will be matched against traffic destination.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "src_address_groups": {
                    "description": "Address groups which should be matched against the traffic source.\nMaximum number of source address groups is 10.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "src_fqdns": {
                    "description": "Fully Qualified Domain Name (FQDN) which should be matched against\ntraffic source. Maximum number of source fqdn allowed is 100.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "src_ip_ranges": {
                    "description": "Source IP address range in CIDR format. Required for\nINGRESS rules.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "src_region_codes": {
                    "description": "Region codes whose IP addresses will be used to match for source\nof traffic. Should be specified as 2 letter country code defined as per\nISO 3166 alpha-2 country codes. ex.\"US\"\nMaximum number of source region codes allowed is 5000.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "src_threat_intelligences": {
                    "description": "Names of Network Threat Intelligence lists.\nThe IPs in these lists will be matched against traffic source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "layer4_config": {
                    "block": {
                      "attributes": {
                        "ip_protocol": {
                          "description": "The IP protocol to which this rule applies. The protocol\ntype is required when creating a firewall rule.\nThis value can either be one of the following well\nknown protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp),\nor the IP protocol number.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "ports": {
                          "description": "An optional list of ports to which this rule applies. This field\nis only applicable for UDP or TCP protocol. Each entry must be\neither an integer or a range. If not specified, this rule\napplies to connections through any port.\nExample inputs include: [\"22\"], [\"80\",\"443\"], and\n[\"12345-12349\"].",
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
                  "src_secure_tag": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the secure tag, created with TagManager's TagValue API.\n@pattern tagValues/[0-9]+",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "state": {
                          "computed": true,
                          "description": "[Output Only] State of the secure tag, either 'EFFECTIVE' or\n'INEFFECTIVE'. A secure tag is 'INEFFECTIVE' when it is deleted\nor its network is deleted.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "description": "List of secure tag values, which should be matched at the source\nof the traffic.\nFor INGRESS rule, if all the \u003ccode\u003esrcSecureTag\u003c/code\u003e are INEFFECTIVE,\nand there is no \u003ccode\u003esrcIpRange\u003c/code\u003e, this rule will be ignored.\nMaximum number of source tag values allowed is 256.",
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
            "target_secure_tag": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the secure tag, created with TagManager's TagValue API.\n@pattern tagValues/[0-9]+",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "[Output Only] State of the secure tag, either 'EFFECTIVE' or\n'INEFFECTIVE'. A secure tag is 'INEFFECTIVE' when it is deleted\nor its network is deleted.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "A list of secure tags that controls which instances the firewall rule\napplies to. If \u003ccode\u003etargetSecureTag\u003c/code\u003e are specified, then the\nfirewall rule applies only to instances in the VPC network that have one\nof those EFFECTIVE secure tags, if all the target_secure_tag are in\nINEFFECTIVE state, then this rule will be ignored.\n\u003ccode\u003etargetSecureTag\u003c/code\u003e may not be set at the same time as\n\u003ccode\u003etargetServiceAccounts\u003c/code\u003e.\nIf neither \u003ccode\u003etargetServiceAccounts\u003c/code\u003e nor\n\u003ccode\u003etargetSecureTag\u003c/code\u003e are specified, the firewall rule applies\nto all instances on the specified network.\nMaximum number of target label tags allowed is 256.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "A list of firewall policy rules.",
          "description_kind": "plain"
        },
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

func GoogleComputeRegionNetworkFirewallPolicyWithRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionNetworkFirewallPolicyWithRules), &result)
	return &result
}
