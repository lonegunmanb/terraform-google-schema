package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSecurityPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "An optional description of this security policy. Max size is 2048.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource.",
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
        "description": "The name of the security policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "adaptive_protection_config": {
        "block": {
          "block_types": {
            "layer_7_ddos_defense_config": {
              "block": {
                "attributes": {
                  "enable": {
                    "description": "If set to true, enables CAAP for L7 DDoS detection.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "rule_visibility": {
                    "computed": true,
                    "description": "Rule visibility. Supported values include: \"STANDARD\", \"PREMIUM\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "threshold_configs": {
                    "block": {
                      "attributes": {
                        "auto_deploy_confidence_threshold": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "auto_deploy_expiration_sec": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "auto_deploy_impacted_baseline_threshold": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "auto_deploy_load_threshold": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "detection_absolute_qps": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "detection_load_threshold": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "detection_relative_to_baseline_qps": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "name": {
                          "description": "The name must be 1-63 characters long, and comply with RFC1035. The name must be unique within the security policy.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "traffic_granularity_configs": {
                          "block": {
                            "attributes": {
                              "enable_each_unique_value": {
                                "description": "If enabled, traffic matching each unique value for the specified type constitutes a separate traffic unit. It can only be set to true if value is empty.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "type": {
                                "description": "Type of this configuration.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "Requests that match this value constitute a granular traffic unit.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Configuration options for layer7 adaptive protection for various customizable thresholds.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Layer 7 DDoS Defense Config of this security policy",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Adaptive Protection Config of this security policy.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "advanced_options_config": {
        "block": {
          "attributes": {
            "json_parsing": {
              "computed": true,
              "description": "JSON body parsing. Supported values include: \"DISABLED\", \"STANDARD\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_level": {
              "computed": true,
              "description": "Logging level. Supported values include: \"NORMAL\", \"VERBOSE\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_ip_request_headers": {
              "description": "An optional list of case-insensitive request header names to use for resolving the callers client IP address.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "json_custom_config": {
              "block": {
                "attributes": {
                  "content_types": {
                    "description": "A list of custom Content-Type header values to apply the JSON parsing.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "Custom configuration to apply the JSON parsing. Only applicable when JSON parsing is set to STANDARD.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Advanced Options Config of this security policy.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "recaptcha_options_config": {
        "block": {
          "attributes": {
            "redirect_site_key": {
              "description": "A field to supply a reCAPTCHA site key to be used for all the rules using the redirect action with the type of GOOGLE_RECAPTCHA under the security policy. The specified site key needs to be created from the reCAPTCHA API. The user is responsible for the validity of the specified site key. If not specified, a Google-managed site key is used.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "reCAPTCHA configuration options to be applied for the security policy.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rule": {
        "block": {
          "attributes": {
            "action": {
              "description": "Action to take when match matches the request.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "description": "An optional description of this rule. Max size is 64.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preview": {
              "computed": true,
              "description": "When set to true, the action specified above is not enforced. Stackdriver logs for requests that trigger a preview action are annotated as such.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "priority": {
              "description": "An unique positive integer indicating the priority of evaluation for a rule. Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "header_action": {
              "block": {
                "block_types": {
                  "request_headers_to_adds": {
                    "block": {
                      "attributes": {
                        "header_name": {
                          "description": "The name of the header to set.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "header_value": {
                          "description": "The value to set the named header to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The list of request headers to add or overwrite if they're already present.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Additional actions that are performed on headers.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "match": {
              "block": {
                "attributes": {
                  "versioned_expr": {
                    "description": "Predefined rule expression. If this field is specified, config must also be specified. Available options:   SRC_IPS_V1: Must specify the corresponding src_ip_ranges field in config.",
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
                          "description": "Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation to match against inbound traffic. There is a limit of 10 IP ranges per rule. A value of '*' matches all IPs (can be used to override the default behavior).",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description": "The configuration options available when specifying versioned_expr. This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified.",
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
                  },
                  "expr_options": {
                    "block": {
                      "block_types": {
                        "recaptcha_options": {
                          "block": {
                            "attributes": {
                              "action_token_site_keys": {
                                "description": "A list of site keys to be used during the validation of reCAPTCHA action-tokens. The provided site keys need to be created from reCAPTCHA API under the same project where the security policy is created",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "session_token_site_keys": {
                                "description": "A list of site keys to be used during the validation of reCAPTCHA session-tokens. The provided site keys need to be created from reCAPTCHA API under the same project where the security policy is created.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "reCAPTCHA configuration options to be applied for the rule. If the rule does not evaluate reCAPTCHA tokens, this field has no effect.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The configuration options available when specifying a user defined CEVAL expression (i.e., 'expr').",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding action is enforced.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "preconfigured_waf_config": {
              "block": {
                "block_types": {
                  "exclusion": {
                    "block": {
                      "attributes": {
                        "target_rule_ids": {
                          "description": "A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion. If omitted, it refers to all the rule IDs under the WAF rule set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
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
                                "description": "You can specify an exact match or a partial match by using a field operator and a field value. Available options: EQUALS: The operator matches if the field value equals the specified value. STARTS_WITH: The operator matches if the field value starts with the specified value. ENDS_WITH: The operator matches if the field value ends with the specified value. CONTAINS: The operator matches if the field value contains the specified value. EQUALS_ANY: The operator matches if the field value is any value.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation. The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.",
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
                                "description": "You can specify an exact match or a partial match by using a field operator and a field value. Available options: EQUALS: The operator matches if the field value equals the specified value. STARTS_WITH: The operator matches if the field value starts with the specified value. ENDS_WITH: The operator matches if the field value ends with the specified value. CONTAINS: The operator matches if the field value contains the specified value. EQUALS_ANY: The operator matches if the field value is any value.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation. The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.",
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
                                "description": "You can specify an exact match or a partial match by using a field operator and a field value. Available options: EQUALS: The operator matches if the field value equals the specified value. STARTS_WITH: The operator matches if the field value starts with the specified value. ENDS_WITH: The operator matches if the field value ends with the specified value. CONTAINS: The operator matches if the field value contains the specified value. EQUALS_ANY: The operator matches if the field value is any value.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation. The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Request query parameter whose value will be excluded from inspection during preconfigured WAF evaluation.  Note that the parameter can be in the query string or in the POST body.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "request_uri": {
                          "block": {
                            "attributes": {
                              "operator": {
                                "description": "You can specify an exact match or a partial match by using a field operator and a field value. Available options: EQUALS: The operator matches if the field value equals the specified value. STARTS_WITH: The operator matches if the field value starts with the specified value. ENDS_WITH: The operator matches if the field value ends with the specified value. CONTAINS: The operator matches if the field value contains the specified value. EQUALS_ANY: The operator matches if the field value is any value.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation. The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Request URI from the request line to be excluded from inspection during preconfigured WAF evaluation. When specifying this field, the query or fragment part should be excluded.",
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
                "description": "Preconfigured WAF configuration to be applied for the rule. If the rule does not evaluate preconfigured WAF rules, i.e., if evaluatePreconfiguredWaf() is not used, this field will have no effect.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rate_limit_options": {
              "block": {
                "attributes": {
                  "ban_duration_sec": {
                    "description": "Can only be specified if the action for the rule is \"rate_based_ban\". If specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "conform_action": {
                    "description": "Action to take for requests that are under the configured rate limit threshold. Valid option is \"allow\" only.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "enforce_on_key": {
                    "description": "Determines the key to enforce the rateLimitThreshold on",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enforce_on_key_name": {
                    "description": "Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value. HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "exceed_action": {
                    "description": "Action to take for requests that are above the configured rate limit threshold, to either deny with a specified HTTP response code, or redirect to a different endpoint. Valid options are \"deny()\" where valid values for status are 403, 404, 429, and 502, and \"redirect\" where the redirect parameters come from exceedRedirectOptions below.",
                    "description_kind": "plain",
                    "required": true,
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
                          "required": true,
                          "type": "number"
                        },
                        "interval_sec": {
                          "description": "Interval over which the threshold is computed.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Can only be specified if the action for the rule is \"rate_based_ban\". If specified, the key will be banned for the configured 'banDurationSec' when the number of requests that exceed the 'rateLimitThreshold' also exceed this 'banThreshold'.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "enforce_on_key_configs": {
                    "block": {
                      "attributes": {
                        "enforce_on_key_name": {
                          "description": "Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value. HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enforce_on_key_type": {
                          "description": "Determines the key to enforce the rate_limit_threshold on",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Enforce On Key Config of this security policy",
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
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Parameters defining the redirect action that is used as the exceed action. Cannot be specified if the exceed action is not redirect.",
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
                          "required": true,
                          "type": "number"
                        },
                        "interval_sec": {
                          "description": "Interval over which the threshold is computed.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Threshold at which to begin ratelimiting.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Rate limit threshold for this security policy. Must be specified if the action is \"rate_based_ban\" or \"throttle\". Cannot be specified for any other actions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "redirect_options": {
              "block": {
                "attributes": {
                  "target": {
                    "description": "Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "Type of the redirect action. Available options: EXTERNAL_302: Must specify the corresponding target field in config. GOOGLE_RECAPTCHA: Cannot specify target field in config.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Parameters defining the redirect action. Cannot be specified for any other actions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The set of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match \"*\"). If no rules are provided when creating a security policy, a default rule with action \"allow\" will be added.",
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

func GoogleComputeSecurityPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSecurityPolicy), &result)
	return &result
}
