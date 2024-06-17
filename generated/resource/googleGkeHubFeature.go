package resource

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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "GCP labels for this Feature.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
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
    "block_types": {
      "fleet_default_member_config": {
        "block": {
          "block_types": {
            "configmanagement": {
              "block": {
                "attributes": {
                  "version": {
                    "description": "Version of ACM installed",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "config_sync": {
                    "block": {
                      "attributes": {
                        "prevent_drift": {
                          "description": "Set to true to enable the Config Sync admission webhook to prevent drifts. If set to 'false', disables the Config Sync admission webhook and does not prevent drifts.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "source_format": {
                          "description": "Specifies whether the Config Sync Repo is in hierarchical or unstructured mode",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "git": {
                          "block": {
                            "attributes": {
                              "gcp_service_account_email": {
                                "description": "The Google Cloud Service Account Email used for auth when secretType is gcpServiceAccount",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "https_proxy": {
                                "description": "URL for the HTTPS Proxy to be used when communicating with the Git repo",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "policy_dir": {
                                "description": "The path within the Git repository that represents the top level of the repo to sync",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "secret_type": {
                                "description": "Type of secret configured for access to the Git repo",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "sync_branch": {
                                "description": "The branch of the repository to sync from. Default: master",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sync_repo": {
                                "description": "The URL of the Git repository to use as the source of truth",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sync_rev": {
                                "description": "Git revision (tag or hash) to check out. Default HEAD",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sync_wait_secs": {
                                "description": "Period in seconds between consecutive syncs. Default: 15",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Git repo configuration for the cluster",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "oci": {
                          "block": {
                            "attributes": {
                              "gcp_service_account_email": {
                                "description": "The Google Cloud Service Account Email used for auth when secretType is gcpServiceAccount",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "policy_dir": {
                                "description": "The absolute path of the directory that contains the local resources. Default: the root directory of the image",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "secret_type": {
                                "description": "Type of secret configured for access to the Git repo",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "sync_repo": {
                                "description": "The OCI image repository URL for the package to sync from",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sync_wait_secs": {
                                "description": "Period in seconds between consecutive syncs. Default: 15",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "version": {
                                "deprecated": true,
                                "description": "Version of ACM installed",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "OCI repo configuration for the cluster",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "ConfigSync configuration for the cluster",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Config Management spec",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "mesh": {
              "block": {
                "attributes": {
                  "management": {
                    "description": "Whether to automatically manage Service Mesh Possible values: [\"MANAGEMENT_UNSPECIFIED\", \"MANAGEMENT_AUTOMATIC\", \"MANAGEMENT_MANUAL\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Service Mesh spec",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "policycontroller": {
              "block": {
                "attributes": {
                  "version": {
                    "computed": true,
                    "description": "Configures the version of Policy Controller",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "policy_controller_hub_config": {
                    "block": {
                      "attributes": {
                        "audit_interval_seconds": {
                          "description": "Interval for Policy Controller Audit scans (in seconds). When set to 0, this disables audit functionality altogether.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "constraint_violation_limit": {
                          "description": "The maximum number of audit violations to be stored in a constraint. If not set, the internal default of 20 will be used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "exemptable_namespaces": {
                          "description": "The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "install_spec": {
                          "description": "Configures the mode of the Policy Controller installation Possible values: [\"INSTALL_SPEC_UNSPECIFIED\", \"INSTALL_SPEC_NOT_INSTALLED\", \"INSTALL_SPEC_ENABLED\", \"INSTALL_SPEC_SUSPENDED\", \"INSTALL_SPEC_DETACHED\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "log_denies_enabled": {
                          "description": "Logs all denies and dry run failures.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "mutation_enabled": {
                          "description": "Enables the ability to mutate resources using Policy Controller.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "referential_rules_enabled": {
                          "description": "Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "deployment_configs": {
                          "block": {
                            "attributes": {
                              "component": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "pod_affinity": {
                                "computed": true,
                                "description": "Pod affinity configuration. Possible values: [\"AFFINITY_UNSPECIFIED\", \"NO_AFFINITY\", \"ANTI_AFFINITY\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "replica_count": {
                                "computed": true,
                                "description": "Pod replica count.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "container_resources": {
                                "block": {
                                  "block_types": {
                                    "limits": {
                                      "block": {
                                        "attributes": {
                                          "cpu": {
                                            "description": "CPU requirement expressed in Kubernetes resource units.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "memory": {
                                            "description": "Memory requirement expressed in Kubernetes resource units.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Limits describes the maximum amount of compute resources allowed for use by the running container.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "requests": {
                                      "block": {
                                        "attributes": {
                                          "cpu": {
                                            "description": "CPU requirement expressed in Kubernetes resource units.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "memory": {
                                            "description": "Memory requirement expressed in Kubernetes resource units.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Requests describes the amount of compute resources reserved for the container by the kube-scheduler.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Container resource requirements.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "pod_toleration": {
                                "block": {
                                  "attributes": {
                                    "effect": {
                                      "description": "Matches a taint effect.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "description": "Matches a taint key (not necessarily unique).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "operator": {
                                      "description": "Matches a taint operator.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "Matches a taint value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Pod tolerations of node taints.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Map of deployment configs to deployments (\"admission\", \"audit\", \"mutation\").",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "set"
                        },
                        "monitoring": {
                          "block": {
                            "attributes": {
                              "backends": {
                                "computed": true,
                                "description": "Specifies the list of backends Policy Controller will export to. An empty list would effectively disable metrics export. Possible values: [\"MONITORING_BACKEND_UNSPECIFIED\", \"PROMETHEUS\", \"CLOUD_MONITORING\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "Monitoring specifies the configuration of monitoring Policy Controller.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "policy_content": {
                          "block": {
                            "block_types": {
                              "bundles": {
                                "block": {
                                  "attributes": {
                                    "bundle": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "exempted_namespaces": {
                                      "description": "The set of namespaces to be exempted from the bundle.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "Configures which bundles to install and their corresponding install specs.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "set"
                              },
                              "template_library": {
                                "block": {
                                  "attributes": {
                                    "installation": {
                                      "description": "Configures the manner in which the template library is installed on the cluster. Possible values: [\"INSTALATION_UNSPECIFIED\", \"NOT_INSTALLED\", \"ALL\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Configures the installation of the Template Library.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the desired policy content on the cluster.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Configuration of Policy Controller",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Policy Controller spec",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. Fleet Default Membership Configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "block_types": {
            "clusterupgrade": {
              "block": {
                "attributes": {
                  "upstream_fleets": {
                    "description": "Specified if other fleet should be considered as a source of upgrades. Currently, at most one upstream fleet is allowed. The fleet name should be either fleet project number or id.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "gke_upgrade_overrides": {
                    "block": {
                      "block_types": {
                        "post_conditions": {
                          "block": {
                            "attributes": {
                              "soaking": {
                                "description": "Amount of time to \"soak\" after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Post conditions to override for the specified upgrade.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "upgrade": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the upgrade, e.g., \"k8s_control_plane\". It should be a valid upgrade name. It must not exceet 99 characters.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "version": {
                                "description": "Version of the upgrade, e.g., \"1.22.1-gke.100\". It should be a valid version. It must not exceet 99 characters.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Which upgrade to override.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Configuration overrides for individual upgrades.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "post_conditions": {
                    "block": {
                      "attributes": {
                        "soaking": {
                          "description": "Amount of time to \"soak\" after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Post conditions to override for the specified upgrade.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Clusterupgrade feature spec.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "fleetobservability": {
              "block": {
                "block_types": {
                  "logging_config": {
                    "block": {
                      "block_types": {
                        "default_config": {
                          "block": {
                            "attributes": {
                              "mode": {
                                "description": "Specified if fleet logging feature is enabled. Possible values: [\"MODE_UNSPECIFIED\", \"COPY\", \"MOVE\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Specified if applying the default routing config to logs not specified in other configs.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "fleet_scope_logs_config": {
                          "block": {
                            "attributes": {
                              "mode": {
                                "description": "Specified if fleet logging feature is enabled. Possible values: [\"MODE_UNSPECIFIED\", \"COPY\", \"MOVE\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Specified if applying the routing config to all logs for all fleet scopes.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specified if fleet logging feature is enabled for the entire fleet. If UNSPECIFIED, fleet logging feature is disabled for the entire fleet.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Fleet Observability feature spec.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "multiclusteringress": {
              "block": {
                "attributes": {
                  "config_membership": {
                    "description": "Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: 'projects/foo-proj/locations/global/memberships/bar'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Multicluster Ingress-specific spec.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.",
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

func GoogleGkeHubFeatureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubFeature), &result)
	return &result
}
