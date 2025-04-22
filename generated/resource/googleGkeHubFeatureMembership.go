package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubFeatureMembership = `{
  "block": {
    "attributes": {
      "feature": {
        "description": "The name of the feature",
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
      "location": {
        "description": "The location of the feature",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "membership": {
        "description": "The name of the membership",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "membership_location": {
        "description": "The location of the membership",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project of the feature",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "configmanagement": {
        "block": {
          "attributes": {
            "management": {
              "computed": true,
              "description": "Set this field to MANAGEMENT_AUTOMATIC to enable Config Sync auto-upgrades, and set this field to MANAGEMENT_MANUAL or MANAGEMENT_UNSPECIFIED to disable Config Sync auto-upgrades.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "Optional. Version of ACM to install. Defaults to the latest version.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "binauthz": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether binauthz is enabled in this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "**DEPRECATED** Binauthz configuration for the cluster. This field will be ignored and should not be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "config_sync": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Enables the installation of ConfigSync. If set to true, ConfigSync resources will be created and the other ConfigSync fields will be applied if exist. If set to false, all other ConfigSync fields will be ignored, ConfigSync resources will be deleted. If omitted, ConfigSync resources will be managed depends on the presence of the git or oci field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "metrics_gcp_service_account_email": {
                    "description": "Deprecated: If Workload Identity Federation for GKE is enabled, Google Cloud Service Account is no longer needed for exporting Config Sync metrics: https://cloud.google.com/kubernetes-engine/enterprise/config-sync/docs/how-to/monitor-config-sync-cloud-monitoring#custom-monitoring.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prevent_drift": {
                    "computed": true,
                    "description": "Set to true to enable the Config Sync admission webhook to prevent drifts. If set to ` + "`" + `false` + "`" + `, disables the Config Sync admission webhook and does not prevent drifts.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_format": {
                    "description": "Specifies whether the Config Sync Repo is in \"hierarchical\" or \"unstructured\" mode.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "stop_syncing": {
                    "description": "Set to true to stop syncing configs for a single cluster. Default: false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "deployment_overrides": {
                    "block": {
                      "attributes": {
                        "deployment_name": {
                          "description": "The name of the Deployment.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "deployment_namespace": {
                          "description": "The namespace of the Deployment.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "containers": {
                          "block": {
                            "attributes": {
                              "container_name": {
                                "description": "The name of the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "cpu_limit": {
                                "description": "The CPU limit of the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "cpu_request": {
                                "description": "The CPU request of the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "memory_limit": {
                                "description": "The memory limit of the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "memory_request": {
                                "description": "The memory request of the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The override configurations for the containers in the Deployment.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The override configurations for the Config Sync Deployments.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "git": {
                    "block": {
                      "attributes": {
                        "gcp_service_account_email": {
                          "description": "The GCP Service Account Email used for auth when secretType is gcpServiceAccount.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "https_proxy": {
                          "description": "URL for the HTTPS proxy to be used when communicating with the Git repo.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "policy_dir": {
                          "description": "The path within the Git repository that represents the top level of the repo to sync. Default: the root directory of the repository.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_type": {
                          "description": "Type of secret configured for access to the Git repo. Must be one of ssh, cookiefile, gcenode, token, gcpserviceaccount or none. The validation of this is case-sensitive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_branch": {
                          "description": "The branch of the repository to sync from. Default: master.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_repo": {
                          "description": "The URL of the Git repository to use as the source of truth.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_rev": {
                          "description": "Git revision (tag or hash) to check out. Default HEAD.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_wait_secs": {
                          "description": "Period in seconds between consecutive syncs. Default: 15.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "oci": {
                    "block": {
                      "attributes": {
                        "gcp_service_account_email": {
                          "description": "The GCP Service Account Email used for auth when secret_type is gcpserviceaccount. ",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "policy_dir": {
                          "description": "The absolute path of the directory that contains the local resources. Default: the root directory of the image.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_type": {
                          "description": "Type of secret configured for access to the OCI Image. Must be one of gcenode, gcpserviceaccount or none. The validation of this is case-sensitive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_repo": {
                          "description": "The OCI image repository URL for the package to sync from. e.g. LOCATION-docker.pkg.dev/PROJECT_ID/REPOSITORY_NAME/PACKAGE_NAME.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_wait_secs": {
                          "description": "Period in seconds(int64 format) between consecutive syncs. Default: 15.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Config Sync configuration for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "hierarchy_controller": {
              "block": {
                "attributes": {
                  "enable_hierarchical_resource_quota": {
                    "description": "Whether hierarchical resource quota is enabled in this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_pod_tree_labels": {
                    "description": "Whether pod tree labels are enabled in this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enabled": {
                    "description": "**DEPRECATED** Configuring Hierarchy Controller through the configmanagement feature is no longer recommended. Use https://github.com/kubernetes-sigs/hierarchical-namespaces instead.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Hierarchy Controller configuration for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "policy_controller": {
              "block": {
                "attributes": {
                  "audit_interval_seconds": {
                    "description": "Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description": "Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
                  "log_denies_enabled": {
                    "description": "Logs all denies and dry run failures.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "mutation_enabled": {
                    "description": "Enable or disable mutation in policy controller. If true, mutation CRDs, webhook and controller deployment will be deployed to the cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "referential_rules_enabled": {
                    "description": "Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "template_library_installed": {
                    "description": "Installs the default template library along with Policy Controller.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "monitoring": {
                    "block": {
                      "attributes": {
                        "backends": {
                          "computed": true,
                          "description": " Specifies the list of backends Policy Controller will export to. Specifying an empty value ` + "`" + `[]` + "`" + ` disables metrics export.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Specifies the backends Policy Controller should export metrics to. For example, to specify metrics should be exported to Cloud Monitoring and Prometheus, specify backends: [\"cloudmonitoring\", \"prometheus\"]. Default: [\"cloudmonitoring\", \"prometheus\"]",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "**DEPRECATED** Configuring Policy Controller through the configmanagement feature is no longer recommended. Use the policycontroller feature instead.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Config Management-specific spec.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mesh": {
        "block": {
          "attributes": {
            "control_plane": {
              "deprecated": true,
              "description": "**DEPRECATED** Whether to automatically manage Service Mesh control planes. Possible values: CONTROL_PLANE_MANAGEMENT_UNSPECIFIED, AUTOMATIC, MANUAL",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "management": {
              "description": "Whether to automatically manage Service Mesh. Possible values: MANAGEMENT_UNSPECIFIED, MANAGEMENT_AUTOMATIC, MANAGEMENT_MANUAL",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Manage Mesh Features",
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
              "description": "Optional. Version of Policy Controller to install. Defaults to the latest version.",
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
                    "description": "Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.",
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
                    "description": "Configures the mode of the Policy Controller installation. Possible values: INSTALL_SPEC_UNSPECIFIED, INSTALL_SPEC_NOT_INSTALLED, INSTALL_SPEC_ENABLED, INSTALL_SPEC_SUSPENDED, INSTALL_SPEC_DETACHED",
                    "description_kind": "plain",
                    "optional": true,
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
                        "component_name": {
                          "description": "The name for the key in the map for which this object is mapped to in the API",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "pod_affinity": {
                          "description": "Pod affinity configuration. Possible values: AFFINITY_UNSPECIFIED, NO_AFFINITY, ANTI_AFFINITY",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "replica_count": {
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
                        "pod_tolerations": {
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
                          "description": " Specifies the list of backends Policy Controller will export to. Specifying an empty value ` + "`" + `[]` + "`" + ` disables metrics export.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Specifies the backends Policy Controller should export metrics to. For example, to specify metrics should be exported to Cloud Monitoring and Prometheus, specify backends: [\"cloudmonitoring\", \"prometheus\"]. Default: [\"cloudmonitoring\", \"prometheus\"]",
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
                              "bundle_name": {
                                "description": "The name for the key in the map for which this object is mapped to in the API",
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
                            "description": "map of bundle name to BundleInstallSpec. The bundle name maps to the ` + "`" + `bundleName` + "`" + ` key in the ` + "`" + `policycontroller.gke.io/constraintData` + "`" + ` annotation on a constraint.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "set"
                        },
                        "template_library": {
                          "block": {
                            "attributes": {
                              "installation": {
                                "description": "Configures the manner in which the template library is installed on the cluster. Possible values: INSTALLATION_UNSPECIFIED, NOT_INSTALLED, ALL",
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
                "description": "Policy Controller configuration for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Policy Controller-specific spec.",
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

func GoogleGkeHubFeatureMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubFeatureMembership), &result)
	return &result
}
