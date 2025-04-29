package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionSecurityPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "An optional description of this resource. Provide this property when you create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. This field is used internally during\nupdates of this resource.",
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
        "description": "Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035.\nSpecifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region in which the created Region Security Policy should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link_with_policy_id": {
        "computed": true,
        "description": "Server-defined URL for this resource with the resource id.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "The type indicates the intended use of the security policy.\n- CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers.\n- CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.\n- CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application.\nThis field can be set only at resource creation time. Possible values: [\"CLOUD_ARMOR\", \"CLOUD_ARMOR_EDGE\", \"CLOUD_ARMOR_NETWORK\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "ddos_protection_config": {
        "block": {
          "attributes": {
            "ddos_protection": {
              "description": "Google Cloud Armor offers the following options to help protect systems against DDoS attacks:\n- STANDARD: basic always-on protection for network load balancers, protocol forwarding, or VMs with public IP addresses.\n- ADVANCED: additional protections for Managed Protection Plus subscribers who use network load balancers, protocol forwarding, or VMs with public IP addresses.\n- ADVANCED_PREVIEW: flag to enable the security policy in preview mode. Possible values: [\"ADVANCED\", \"ADVANCED_PREVIEW\", \"STANDARD\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for Google Cloud Armor DDOS Proctection Config.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rules": {
        "block": {
          "attributes": {
            "action": {
              "description": "The Action to perform when the rule is matched. The following are the valid actions:\n\n* allow: allow access to target.\n\n* deny(STATUS): deny access to target, returns the HTTP response code specified. Valid values for STATUS are 403, 404, and 502.\n\n* rate_based_ban: limit client traffic to the configured threshold and ban the client if the traffic exceeds the threshold. Configure parameters for this action in RateLimitOptions. Requires rateLimitOptions to be set.\n\n* redirect: redirect to a different target. This can either be an internal reCAPTCHA redirect, or an external URL-based redirect via a 302 response. Parameters for this action can be configured via redirectOptions. This action is only supported in Global Security Policies of type CLOUD_ARMOR.\n\n* throttle: limit client traffic to the configured threshold. Configure parameters for this action in rateLimitOptions. Requires rateLimitOptions to be set for this.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "description": "An optional description of this resource. Provide this property when you create the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preview": {
              "description": "If set to true, the specified action is not enforced.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "priority": {
              "description": "An integer indicating the priority of a rule in the list.\nThe priority must be a positive value between 0 and 2147483647.\nRules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest priority.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "match": {
              "block": {
                "attributes": {
                  "versioned_expr": {
                    "description": "Preconfigured versioned expression. If this field is specified, config must also be specified.\nAvailable preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding srcIpRange field in config. Possible values: [\"SRC_IPS_V1\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "config": {
                    "block": {
                      "attributes": {
                        "src_ip_ranges": {
                          "description": "CIDR IP address range. Maximum number of srcIpRanges allowed is 10.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "The configuration options available when specifying versionedExpr.\nThis field must be specified if versionedExpr is specified and cannot be specified if versionedExpr is not specified.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "expr": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description": "Textual representation of an expression in Common Expression Language syntax. The application context of the containing message determines which well-known feature set of CEL is supported.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header. See [Sample expressions](https://cloud.google.com/armor/docs/configure-security-policies#sample-expressions) for examples.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A match condition that incoming traffic is evaluated against.\nIf it evaluates to true, the corresponding 'action' is enforced.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "network_match": {
              "block": {
                "attributes": {
                  "dest_ip_ranges": {
                    "description": "Destination IPv4/IPv6 addresses or CIDR prefixes, in standard text format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "dest_ports": {
                    "description": "Destination port numbers for TCP/UDP/SCTP. Each element can be a 16-bit unsigned decimal number (e.g. \"80\") or range (e.g. \"0-1023\").",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ip_protocols": {
                    "description": "IPv4 protocol / IPv6 next header (after extension headers). Each element can be an 8-bit unsigned decimal number (e.g. \"6\"), range (e.g. \"253-254\"), or one of the following protocol names: \"tcp\", \"udp\", \"icmp\", \"esp\", \"ah\", \"ipip\", or \"sctp\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "src_asns": {
                    "description": "BGP Autonomous System Number associated with the source IP address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "src_ip_ranges": {
                    "description": "Source IPv4/IPv6 addresses or CIDR prefixes, in standard text format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "src_ports": {
                    "description": "Source port numbers for TCP/UDP/SCTP. Each element can be a 16-bit unsigned decimal number (e.g. \"80\") or range (e.g. \"0-1023\").",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "src_region_codes": {
                    "description": "Two-letter ISO 3166-1 alpha-2 country code associated with the source IP address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "user_defined_fields": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the user-defined field, as given in the definition.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "values": {
                          "description": "Matching values of the field. Each element can be a 32-bit unsigned decimal or hexadecimal (starting with \"0x\") number (e.g. \"64\") or range (e.g. \"0x400-0x7ff\").",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "User-defined fields. Each element names a defined field and lists the matching values for that field.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "A match condition that incoming packets are evaluated against for CLOUD_ARMOR_NETWORK security policies. If it matches, the corresponding 'action' is enforced.\nThe match criteria for a rule consists of built-in match fields (like 'srcIpRanges') and potentially multiple user-defined match fields ('userDefinedFields').\nField values may be extracted directly from the packet or derived from it (e.g. 'srcRegionCodes'). Some fields may not be present in every packet (e.g. 'srcPorts'). A user-defined field is only present if the base header is found in the packet and the entire field is in bounds.\nEach match field may specify which values can match it, listing one or more ranges, prefixes, or exact values that are considered a match for the field. A field value must be present in order to match a specified match field. If no match values are specified for a match field, then any field value is considered to match it, and it's not required to be present. For strings specifying '*' is also equivalent to match all.\nFor a packet to match a rule, all specified match fields must match the corresponding field values derived from the packet.\nExample:\nnetworkMatch: srcIpRanges: - \"192.0.2.0/24\" - \"198.51.100.0/24\" userDefinedFields: - name: \"ipv4_fragment_offset\" values: - \"1-0x1fff\"\nThe above match condition matches packets with a source IP in 192.0.2.0/24 or 198.51.100.0/24 and a user-defined field named \"ipv4_fragment_offset\" with a value between 1 and 0x1fff inclusive",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "preconfigured_waf_config": {
              "block": {
                "block_types": {
                  "exclusion": {
                    "block": {
                      "attributes": {
                        "target_rule_ids": {
                          "description": "A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion.\nIf omitted, it refers to all the rule IDs under the WAF rule set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "target_rule_set": {
                          "description": "Target WAF rule set to apply the preconfigured WAF exclusion.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "request_cookie": {
                          "block": {
                            "attributes": {
                              "operator": {
                                "description": "You can specify an exact match or a partial match by using a field operator and a field value.\nAvailable options:\nEQUALS: The operator matches if the field value equals the specified value.\nSTARTS_WITH: The operator matches if the field value starts with the specified value.\nENDS_WITH: The operator matches if the field value ends with the specified value.\nCONTAINS: The operator matches if the field value contains the specified value.\nEQUALS_ANY: The operator matches if the field value is any value. Possible values: [\"CONTAINS\", \"ENDS_WITH\", \"EQUALS\", \"EQUALS_ANY\", \"STARTS_WITH\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation.\nThe field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Request cookie whose value will be excluded from inspection during preconfigured WAF evaluation.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "request_header": {
                          "block": {
                            "attributes": {
                              "operator": {
                                "description": "You can specify an exact match or a partial match by using a field operator and a field value.\nAvailable options:\nEQUALS: The operator matches if the field value equals the specified value.\nSTARTS_WITH: The operator matches if the field value starts with the specified value.\nENDS_WITH: The operator matches if the field value ends with the specified value.\nCONTAINS: The operator matches if the field value contains the specified value.\nEQUALS_ANY: The operator matches if the field value is any value. Possible values: [\"CONTAINS\", \"ENDS_WITH\", \"EQUALS\", \"EQUALS_ANY\", \"STARTS_WITH\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation.\nThe field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Request header whose value will be excluded from inspection during preconfigured WAF evaluation.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "request_query_param": {
                          "block": {
                            "attributes": {
                              "operator": {
                                "description": "You can specify an exact match or a partial match by using a field operator and a field value.\nAvailable options:\nEQUALS: The operator matches if the field value equals the specified value.\nSTARTS_WITH: The operator matches if the field value starts with the specified value.\nENDS_WITH: The operator matches if the field value ends with the specified value.\nCONTAINS: The operator matches if the field value contains the specified value.\nEQUALS_ANY: The operator matches if the field value is any value. Possible values: [\"CONTAINS\", \"ENDS_WITH\", \"EQUALS\", \"EQUALS_ANY\", \"STARTS_WITH\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation.\nThe field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Request query parameter whose value will be excluded from inspection during preconfigured WAF evaluation.\nNote that the parameter can be in the query string or in the POST body.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "request_uri": {
                          "block": {
                            "attributes": {
                              "operator": {
                                "description": "You can specify an exact match or a partial match by using a field operator and a field value.\nAvailable options:\nEQUALS: The operator matches if the field value equals the specified value.\nSTARTS_WITH: The operator matches if the field value starts with the specified value.\nENDS_WITH: The operator matches if the field value ends with the specified value.\nCONTAINS: The operator matches if the field value contains the specified value.\nEQUALS_ANY: The operator matches if the field value is any value. Possible values: [\"CONTAINS\", \"ENDS_WITH\", \"EQUALS\", \"EQUALS_ANY\", \"STARTS_WITH\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation.\nThe field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Request URI from the request line to be excluded from inspection during preconfigured WAF evaluation.\nWhen specifying this field, the query or fragment part should be excluded.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "An exclusion to apply during preconfigured WAF evaluation.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Preconfigured WAF configuration to be applied for the rule.\nIf the rule does not evaluate preconfigured WAF rules, i.e., if evaluatePreconfiguredWaf() is not used, this field will have no effect.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rate_limit_options": {
              "block": {
                "attributes": {
                  "ban_duration_sec": {
                    "description": "Can only be specified if the action for the rule is \"rate_based_ban\".\nIf specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "conform_action": {
                    "description": "Action to take for requests that are under the configured rate limit threshold.\nValid option is \"allow\" only.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enforce_on_key": {
                    "description": "Determines the key to enforce the rateLimitThreshold on. Possible values are:\n* ALL: A single rate limit threshold is applied to all the requests matching this rule. This is the default value if \"enforceOnKey\" is not configured.\n* IP: The source IP address of the request is the key. Each IP has this limit enforced separately.\n* HTTP_HEADER: The value of the HTTP header whose name is configured under \"enforceOnKeyName\". The key value is truncated to the first 128 bytes of the header value. If no such header is present in the request, the key type defaults to ALL.\n* XFF_IP: The first IP address (i.e. the originating client IP address) specified in the list of IPs under X-Forwarded-For HTTP header. If no such header is present or the value is not a valid IP, the key defaults to the source IP address of the request i.e. key type IP.\n* HTTP_COOKIE: The value of the HTTP cookie whose name is configured under \"enforceOnKeyName\". The key value is truncated to the first 128 bytes of the cookie value. If no such cookie is present in the request, the key type defaults to ALL.\n* HTTP_PATH: The URL path of the HTTP request. The key value is truncated to the first 128 bytes.\n* SNI: Server name indication in the TLS session of the HTTPS request. The key value is truncated to the first 128 bytes. The key type defaults to ALL on a HTTP session.\n* REGION_CODE: The country/region from which the request originates.\n* TLS_JA3_FINGERPRINT: JA3 TLS/SSL fingerprint if the client connects using HTTPS, HTTP/2 or HTTP/3. If not available, the key type defaults to ALL.\n* USER_IP: The IP address of the originating client, which is resolved based on \"userIpRequestHeaders\" configured with the security policy. If there is no \"userIpRequestHeaders\" configuration or an IP address cannot be resolved from it, the key type defaults to IP. Possible values: [\"ALL\", \"IP\", \"HTTP_HEADER\", \"XFF_IP\", \"HTTP_COOKIE\", \"HTTP_PATH\", \"SNI\", \"REGION_CODE\", \"TLS_JA3_FINGERPRINT\", \"USER_IP\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enforce_on_key_name": {
                    "description": "Rate limit key name applicable only for the following key types:\nHTTP_HEADER -- Name of the HTTP header whose value is taken as the key value.\nHTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "exceed_action": {
                    "description": "Action to take for requests that are above the configured rate limit threshold, to deny with a specified HTTP response code.\nValid options are deny(STATUS), where valid values for STATUS are 403, 404, 429, and 502.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "ban_threshold": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description": "Number of HTTP(S) requests for calculating the threshold.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_sec": {
                          "description": "Interval over which the threshold is computed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Can only be specified if the action for the rule is \"rate_based_ban\".\nIf specified, the key will be banned for the configured 'banDurationSec' when the number of requests that exceed the 'rateLimitThreshold' also exceed this 'banThreshold'.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "enforce_on_key_configs": {
                    "block": {
                      "attributes": {
                        "enforce_on_key_name": {
                          "description": "Rate limit key name applicable only for the following key types:\nHTTP_HEADER -- Name of the HTTP header whose value is taken as the key value.\nHTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enforce_on_key_type": {
                          "description": "Determines the key to enforce the rateLimitThreshold on. Possible values are:\n* ALL: A single rate limit threshold is applied to all the requests matching this rule. This is the default value if \"enforceOnKeyConfigs\" is not configured.\n* IP: The source IP address of the request is the key. Each IP has this limit enforced separately.\n* HTTP_HEADER: The value of the HTTP header whose name is configured under \"enforceOnKeyName\". The key value is truncated to the first 128 bytes of the header value. If no such header is present in the request, the key type defaults to ALL.\n* XFF_IP: The first IP address (i.e. the originating client IP address) specified in the list of IPs under X-Forwarded-For HTTP header. If no such header is present or the value is not a valid IP, the key defaults to the source IP address of the request i.e. key type IP.\n* HTTP_COOKIE: The value of the HTTP cookie whose name is configured under \"enforceOnKeyName\". The key value is truncated to the first 128 bytes of the cookie value. If no such cookie is present in the request, the key type defaults to ALL.\n* HTTP_PATH: The URL path of the HTTP request. The key value is truncated to the first 128 bytes.\n* SNI: Server name indication in the TLS session of the HTTPS request. The key value is truncated to the first 128 bytes. The key type defaults to ALL on a HTTP session.\n* REGION_CODE: The country/region from which the request originates.\n* TLS_JA3_FINGERPRINT: JA3 TLS/SSL fingerprint if the client connects using HTTPS, HTTP/2 or HTTP/3. If not available, the key type defaults to ALL.\n* USER_IP: The IP address of the originating client, which is resolved based on \"userIpRequestHeaders\" configured with the security policy. If there is no \"userIpRequestHeaders\" configuration or an IP address cannot be resolved from it, the key type defaults to IP. Possible values: [\"ALL\", \"IP\", \"HTTP_HEADER\", \"XFF_IP\", \"HTTP_COOKIE\", \"HTTP_PATH\", \"SNI\", \"REGION_CODE\", \"TLS_JA3_FINGERPRINT\", \"USER_IP\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "If specified, any combination of values of enforceOnKeyType/enforceOnKeyName is treated as the key on which ratelimit threshold/action is enforced.\nYou can specify up to 3 enforceOnKeyConfigs.\nIf enforceOnKeyConfigs is specified, enforceOnKey must not be specified.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "rate_limit_threshold": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description": "Number of HTTP(S) requests for calculating the threshold.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_sec": {
                          "description": "Interval over which the threshold is computed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Threshold at which to begin ratelimiting.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Must be specified if the action is \"rate_based_ban\" or \"throttle\". Cannot be specified for any other actions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The set of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match \"*\"). If no rules are provided when creating a security policy, a default rule with action \"allow\" will be added.",
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
      },
      "user_defined_fields": {
        "block": {
          "attributes": {
            "base": {
              "description": "The base relative to which 'offset' is measured. Possible values are:\n- IPV4: Points to the beginning of the IPv4 header.\n- IPV6: Points to the beginning of the IPv6 header.\n- TCP: Points to the beginning of the TCP header, skipping over any IPv4 options or IPv6 extension headers. Not present for non-first fragments.\n- UDP: Points to the beginning of the UDP header, skipping over any IPv4 options or IPv6 extension headers. Not present for non-first fragments. Possible values: [\"IPV4\", \"IPV6\", \"TCP\", \"UDP\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "mask": {
              "description": "If specified, apply this mask (bitwise AND) to the field to ignore bits before matching.\nEncoded as a hexadecimal number (starting with \"0x\").\nThe last byte of the field (in network byte order) corresponds to the least significant byte of the mask.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "The name of this field. Must be unique within the policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "offset": {
              "description": "Offset of the first byte of the field (in network byte order) relative to 'base'.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "size": {
              "description": "Size of the field in bytes. Valid values: 1-4.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies.\nA user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits.\nRules may then specify matching values for these fields.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRegionSecurityPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionSecurityPolicy), &result)
	return &result
}
