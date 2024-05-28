package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSecurityPolicyRule = `{
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
      "id": {
        "computed": true,
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
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_policy": {
        "description": "The name of the security policy this rule belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
                "description": "User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header.",
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
                          "description": "You can specify an exact match or a partial match by using a field operator and a field value.\nAvailable options:\nEQUALS: The operator matches if the field value equals the specified value.\nSTARTS_WITH: The operator matches if the field value starts with the specified value.\nENDS_WITH: The operator matches if the field value ends with the specified value.\nCONTAINS: The operator matches if the field value contains the specified value.\nEQUALS_ANY: The operator matches if the field value is any value.",
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
                          "description": "You can specify an exact match or a partial match by using a field operator and a field value.\nAvailable options:\nEQUALS: The operator matches if the field value equals the specified value.\nSTARTS_WITH: The operator matches if the field value starts with the specified value.\nENDS_WITH: The operator matches if the field value ends with the specified value.\nCONTAINS: The operator matches if the field value contains the specified value.\nEQUALS_ANY: The operator matches if the field value is any value.",
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
                          "description": "You can specify an exact match or a partial match by using a field operator and a field value.\nAvailable options:\nEQUALS: The operator matches if the field value equals the specified value.\nSTARTS_WITH: The operator matches if the field value starts with the specified value.\nENDS_WITH: The operator matches if the field value ends with the specified value.\nCONTAINS: The operator matches if the field value contains the specified value.\nEQUALS_ANY: The operator matches if the field value is any value.",
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
                          "description": "You can specify an exact match or a partial match by using a field operator and a field value.\nAvailable options:\nEQUALS: The operator matches if the field value equals the specified value.\nSTARTS_WITH: The operator matches if the field value starts with the specified value.\nENDS_WITH: The operator matches if the field value ends with the specified value.\nCONTAINS: The operator matches if the field value contains the specified value.\nEQUALS_ANY: The operator matches if the field value is any value.",
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
              "description": "Action to take for requests that are above the configured rate limit threshold, to either deny with a specified HTTP response code, or redirect to a different endpoint.\nValid options are deny(STATUS), where valid values for STATUS are 403, 404, 429, and 502.",
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
            "exceed_redirect_options": {
              "block": {
                "attributes": {
                  "target": {
                    "description": "Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "Type of the redirect action.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Parameters defining the redirect action that is used as the exceed action. Cannot be specified if the exceed action is not redirect. This field is only supported in Global Security Policies of type CLOUD_ARMOR.",
                "description_kind": "plain"
              },
              "max_items": 1,
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

func GoogleComputeSecurityPolicyRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSecurityPolicyRule), &result)
	return &result
}
