package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubFeature = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. When the Feature resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Output only. When the Feature resource was deleted.",
        "description_kind": "plain",
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
      "fleet_default_member_config": {
        "computed": true,
        "description": "Optional. Fleet Default Membership Configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "configmanagement": [
                "list",
                [
                  "object",
                  {
                    "config_sync": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool",
                          "git": [
                            "list",
                            [
                              "object",
                              {
                                "gcp_service_account_email": "string",
                                "https_proxy": "string",
                                "policy_dir": "string",
                                "secret_type": "string",
                                "sync_branch": "string",
                                "sync_repo": "string",
                                "sync_rev": "string",
                                "sync_wait_secs": "string"
                              }
                            ]
                          ],
                          "metrics_gcp_service_account_email": "string",
                          "oci": [
                            "list",
                            [
                              "object",
                              {
                                "gcp_service_account_email": "string",
                                "policy_dir": "string",
                                "secret_type": "string",
                                "sync_repo": "string",
                                "sync_wait_secs": "string",
                                "version": "string"
                              }
                            ]
                          ],
                          "prevent_drift": "bool",
                          "source_format": "string"
                        }
                      ]
                    ],
                    "management": "string",
                    "version": "string"
                  }
                ]
              ],
              "mesh": [
                "list",
                [
                  "object",
                  {
                    "management": "string"
                  }
                ]
              ],
              "policycontroller": [
                "list",
                [
                  "object",
                  {
                    "policy_controller_hub_config": [
                      "list",
                      [
                        "object",
                        {
                          "audit_interval_seconds": "number",
                          "constraint_violation_limit": "number",
                          "deployment_configs": [
                            "set",
                            [
                              "object",
                              {
                                "component": "string",
                                "container_resources": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "limits": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "cpu": "string",
                                            "memory": "string"
                                          }
                                        ]
                                      ],
                                      "requests": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "cpu": "string",
                                            "memory": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "pod_affinity": "string",
                                "pod_toleration": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "effect": "string",
                                      "key": "string",
                                      "operator": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "replica_count": "number"
                              }
                            ]
                          ],
                          "exemptable_namespaces": [
                            "list",
                            "string"
                          ],
                          "install_spec": "string",
                          "log_denies_enabled": "bool",
                          "monitoring": [
                            "list",
                            [
                              "object",
                              {
                                "backends": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "mutation_enabled": "bool",
                          "policy_content": [
                            "list",
                            [
                              "object",
                              {
                                "bundles": [
                                  "set",
                                  [
                                    "object",
                                    {
                                      "bundle": "string",
                                      "exempted_namespaces": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "template_library": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "installation": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "referential_rules_enabled": "bool"
                        }
                      ]
                    ],
                    "version": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "GCP labels for this Feature.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The full, unique name of this Feature resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_state": {
        "computed": true,
        "description": "State of the Feature resource itself.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "has_resources": "bool",
              "state": "string"
            }
          ]
        ]
      },
      "spec": {
        "computed": true,
        "description": "Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "clusterupgrade": [
                "list",
                [
                  "object",
                  {
                    "gke_upgrade_overrides": [
                      "list",
                      [
                        "object",
                        {
                          "post_conditions": [
                            "list",
                            [
                              "object",
                              {
                                "soaking": "string"
                              }
                            ]
                          ],
                          "upgrade": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "version": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "post_conditions": [
                      "list",
                      [
                        "object",
                        {
                          "soaking": "string"
                        }
                      ]
                    ],
                    "upstream_fleets": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "fleetobservability": [
                "list",
                [
                  "object",
                  {
                    "logging_config": [
                      "list",
                      [
                        "object",
                        {
                          "default_config": [
                            "list",
                            [
                              "object",
                              {
                                "mode": "string"
                              }
                            ]
                          ],
                          "fleet_scope_logs_config": [
                            "list",
                            [
                              "object",
                              {
                                "mode": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "multiclusteringress": [
                "list",
                [
                  "object",
                  {
                    "config_membership": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "Output only. The Hub-wide Feature state",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "state": [
                "list",
                [
                  "object",
                  {
                    "code": "string",
                    "description": "string",
                    "update_time": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
        "description": "Output only. When the Feature resource was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleGkeHubFeatureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubFeature), &result)
	return &result
}
