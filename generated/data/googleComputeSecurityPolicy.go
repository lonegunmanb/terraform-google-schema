package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSecurityPolicy = `{
  "block": {
    "attributes": {
      "adaptive_protection_config": {
        "computed": true,
        "description": "Adaptive Protection Config of this security policy.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "layer_7_ddos_defense_config": [
                "list",
                [
                  "object",
                  {
                    "enable": "bool",
                    "rule_visibility": "string",
                    "threshold_configs": [
                      "list",
                      [
                        "object",
                        {
                          "auto_deploy_confidence_threshold": "number",
                          "auto_deploy_expiration_sec": "number",
                          "auto_deploy_impacted_baseline_threshold": "number",
                          "auto_deploy_load_threshold": "number",
                          "detection_absolute_qps": "number",
                          "detection_load_threshold": "number",
                          "detection_relative_to_baseline_qps": "number",
                          "name": "string",
                          "traffic_granularity_configs": [
                            "list",
                            [
                              "object",
                              {
                                "enable_each_unique_value": "bool",
                                "type": "string",
                                "value": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "advanced_options_config": {
        "computed": true,
        "description": "Advanced Options Config of this security policy.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "json_custom_config": [
                "list",
                [
                  "object",
                  {
                    "content_types": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "json_parsing": "string",
              "log_level": "string",
              "user_ip_request_headers": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "An optional description of this security policy. Max size is 2048.",
        "description_kind": "plain",
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
        "optional": true,
        "type": "string"
      },
      "project": {
        "description": "The project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recaptcha_options_config": {
        "computed": true,
        "description": "reCAPTCHA configuration options to be applied for the security policy.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "redirect_site_key": "string"
            }
          ]
        ]
      },
      "rule": {
        "computed": true,
        "description": "The set of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match \"*\"). If no rules are provided when creating a security policy, a default rule with action \"allow\" will be added.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "action": "string",
              "description": "string",
              "header_action": [
                "list",
                [
                  "object",
                  {
                    "request_headers_to_adds": [
                      "list",
                      [
                        "object",
                        {
                          "header_name": "string",
                          "header_value": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "match": [
                "list",
                [
                  "object",
                  {
                    "config": [
                      "list",
                      [
                        "object",
                        {
                          "src_ip_ranges": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "expr": [
                      "list",
                      [
                        "object",
                        {
                          "expression": "string"
                        }
                      ]
                    ],
                    "expr_options": [
                      "list",
                      [
                        "object",
                        {
                          "recaptcha_options": [
                            "list",
                            [
                              "object",
                              {
                                "action_token_site_keys": [
                                  "list",
                                  "string"
                                ],
                                "session_token_site_keys": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "versioned_expr": "string"
                  }
                ]
              ],
              "preconfigured_waf_config": [
                "list",
                [
                  "object",
                  {
                    "exclusion": [
                      "list",
                      [
                        "object",
                        {
                          "request_cookie": [
                            "list",
                            [
                              "object",
                              {
                                "operator": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "request_header": [
                            "list",
                            [
                              "object",
                              {
                                "operator": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "request_query_param": [
                            "list",
                            [
                              "object",
                              {
                                "operator": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "request_uri": [
                            "list",
                            [
                              "object",
                              {
                                "operator": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "target_rule_ids": [
                            "set",
                            "string"
                          ],
                          "target_rule_set": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "preview": "bool",
              "priority": "number",
              "rate_limit_options": [
                "list",
                [
                  "object",
                  {
                    "ban_duration_sec": "number",
                    "ban_threshold": [
                      "list",
                      [
                        "object",
                        {
                          "count": "number",
                          "interval_sec": "number"
                        }
                      ]
                    ],
                    "conform_action": "string",
                    "enforce_on_key": "string",
                    "enforce_on_key_name": "string",
                    "exceed_action": "string",
                    "exceed_redirect_options": [
                      "list",
                      [
                        "object",
                        {
                          "target": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "rate_limit_threshold": [
                      "list",
                      [
                        "object",
                        {
                          "count": "number",
                          "interval_sec": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "redirect_options": [
                "list",
                [
                  "object",
                  {
                    "target": "string",
                    "type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "self_link": {
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.",
        "description_kind": "plain",
        "type": "string"
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
