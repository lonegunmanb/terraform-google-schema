package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerNodePool = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "The cluster to create the node pool for. Cluster must be present in location provided for zonal clusters.",
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
      "initial_node_count": {
        "computed": true,
        "description": "The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "instance_group_urls": {
        "computed": true,
        "description": "The resource URLs of the managed instance groups associated with this node pool.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "location": {
        "computed": true,
        "description": "The location (region or zone) of the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_instance_group_urls": {
        "computed": true,
        "description": "List of instance group URLs which have been assigned to this node pool.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "max_pods_per_node": {
        "computed": true,
        "description": "The maximum number of pods per node in this node pool. Note that this does not work on node pools which are \"route-based\" - that is, node pools belonging to clusters that do not have IP Aliasing enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The name of the node pool. If left blank, Terraform will auto-generate a unique name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description": "Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_count": {
        "computed": true,
        "description": "The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "node_locations": {
        "computed": true,
        "description": "The list of zones in which the node pool's nodes should be located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "operation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which to create the node pool. If blank, the provider-configured project will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The Kubernetes version for the nodes in this pool. Note that if this field and auto_upgrade are both specified, they will fight each other for what the node version should be, so setting both is highly discouraged. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "autoscaling": {
        "block": {
          "attributes": {
            "location_policy": {
              "computed": true,
              "description": "Location policy specifies the algorithm used when scaling-up the node pool. \"BALANCED\" - Is a best effort policy that aims to balance the sizes of available zones. \"ANY\" - Instructs the cluster autoscaler to prioritize utilization of unused reservations, and reduces preemption risk for Spot VMs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_node_count": {
              "description": "Maximum number of nodes per zone in the node pool. Must be \u003e= min_node_count. Cannot be used with total limits.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_node_count": {
              "description": "Minimum number of nodes per zone in the node pool. Must be \u003e=0 and \u003c= max_node_count. Cannot be used with total limits.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "total_max_node_count": {
              "description": "Maximum number of all nodes in the node pool. Must be \u003e= total_min_node_count. Cannot be used with per zone limits.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "total_min_node_count": {
              "description": "Minimum number of all nodes in the node pool. Must be \u003e=0 and \u003c= total_max_node_count. Cannot be used with per zone limits.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "management": {
        "block": {
          "attributes": {
            "auto_repair": {
              "description": "Whether the nodes will be automatically repaired. Enabled by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "auto_upgrade": {
              "description": "Whether the nodes will be automatically upgraded. Enabled by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Node management configuration, wherein auto-repair and auto-upgrade is configured.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "attributes": {
            "create_pod_range": {
              "description": "Whether to create a new range for pod IPs in this node pool. Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_private_nodes": {
              "computed": true,
              "description": "Whether nodes have internal IP addresses only.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pod_ipv4_cidr_block": {
              "computed": true,
              "description": "The IP address range for pod IPs in this node pool. Only applicable if create_pod_range is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pod_range": {
              "computed": true,
              "description": "The ID of the secondary range for pod IPs. If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "additional_node_network_configs": {
              "block": {
                "attributes": {
                  "network": {
                    "description": "Name of the VPC where the additional interface belongs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnetwork": {
                    "description": "Name of the subnetwork where the additional interface belongs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "We specify the additional node networks for this node pool using this list. Each node network corresponds to an additional interface",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "additional_pod_network_configs": {
              "block": {
                "attributes": {
                  "max_pods_per_node": {
                    "computed": true,
                    "description": "The maximum number of pods per node which use this pod network.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "secondary_pod_range": {
                    "description": "The name of the secondary range on the subnet which provides IP address for this pod range.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnetwork": {
                    "description": "Name of the subnetwork where the additional pod network belongs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "We specify the additional pod networks for this node pool using this list. Each pod network corresponds to an additional alias IP range for the node",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "network_performance_config": {
              "block": {
                "attributes": {
                  "total_egress_bandwidth_tier": {
                    "description": "Specifies the total network bandwidth tier for the NodePool.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Network bandwidth tier configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pod_cidr_overprovision_config": {
              "block": {
                "attributes": {
                  "disabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for node-pool level pod cidr overprovision. If not set, the cluster level setting will be inherited",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_config": {
        "block": {
          "attributes": {
            "boot_disk_kms_key": {
              "description": "The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_size_gb": {
              "computed": true,
              "description": "Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_type": {
              "computed": true,
              "description": "Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "effective_taints": {
              "computed": true,
              "description": "List of kubernetes taints applied to each node.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "effect": "string",
                    "key": "string",
                    "value": "string"
                  }
                ]
              ]
            },
            "enable_confidential_storage": {
              "description": "If enabled boot disks are configured with confidential mode.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "image_type": {
              "computed": true,
              "description": "The image type to use for this node. Note that for a given image type, the latest version of it will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "computed": true,
              "description": "The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "local_ssd_count": {
              "computed": true,
              "description": "The number of local SSD disks to be attached to the node.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "logging_variant": {
              "computed": true,
              "description": "Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "machine_type": {
              "computed": true,
              "description": "The name of a Google Compute Engine machine type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description": "The metadata key/value pairs assigned to instances in the cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "min_cpu_platform": {
              "computed": true,
              "description": "Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_group": {
              "description": "Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on sole tenant nodes.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "oauth_scopes": {
              "computed": true,
              "description": "The set of Google API scopes to be made available on all of the node VMs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "preemptible": {
              "description": "Whether the nodes are created as preemptible VM instances.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "resource_labels": {
              "description": "The GCE resource labels (a map of key/value pairs) to be applied to the node pool.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "resource_manager_tags": {
              "description": "A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT \u0026 PATCH) when empty.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "service_account": {
              "computed": true,
              "description": "The Google Cloud Platform Service Account to be used by the node VMs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "spot": {
              "description": "Whether the nodes are created as spot VM instances.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "storage_pools": {
              "description": "The list of Storage Pools where boot disks are provisioned.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tags": {
              "description": "The list of instance tags applied to all nodes.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "advanced_machine_features": {
              "block": {
                "attributes": {
                  "enable_nested_virtualization": {
                    "description": "Whether the node should have nested virtualization enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "threads_per_core": {
                    "description": "The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Specifies options for controlling advanced machine features.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "confidential_nodes": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether Confidential Nodes feature is enabled for all nodes in this pool.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after pool creation without deleting and recreating the entire pool.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "containerd_config": {
              "block": {
                "block_types": {
                  "private_registry_access_config": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description": "Whether or not private registries are configured.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "certificate_authority_domain_config": {
                          "block": {
                            "attributes": {
                              "fqdns": {
                                "description": "List of fully-qualified-domain-names. IPv4s and port specification are supported.",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "gcp_secret_manager_certificate_config": {
                                "block": {
                                  "attributes": {
                                    "secret_uri": {
                                      "description": "URI for the secret that hosts a certificate. Must be in the format 'projects/PROJECT_NUM/secrets/SECRET_NAME/versions/VERSION_OR_LATEST'.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Parameters for configuring a certificate hosted in GCP SecretManager.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Parameters for configuring CA certificate and domains.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Parameters for private container registries configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Parameters for containerd configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ephemeral_storage_local_ssd_config": {
              "block": {
                "attributes": {
                  "local_ssd_count": {
                    "description": "Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "fast_socket": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether or not NCCL Fast Socket is enabled",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Enable or disable NCCL Fast Socket in the node pool.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcfs_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether or not GCFS is enabled",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "GCFS configuration for this node.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "guest_accelerator": {
              "block": {
                "attributes": {
                  "count": {
                    "description": "The number of the accelerator cards exposed to an instance.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "gpu_partition_size": {
                    "description": "Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "The accelerator type resource name.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "gpu_driver_installation_config": {
                    "block": {
                      "attributes": {
                        "gpu_driver_version": {
                          "description": "Mode for how the GPU driver is installed.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Configuration for auto installation of GPU driver.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "gpu_sharing_config": {
                    "block": {
                      "attributes": {
                        "gpu_sharing_strategy": {
                          "description": "The type of GPU sharing strategy to enable on the GPU node. Possible values are described in the API package (https://pkg.go.dev/google.golang.org/api/container/v1#GPUSharingConfig)",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "max_shared_clients_per_gpu": {
                          "description": "The maximum number of containers that can share a GPU.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Configuration for GPU sharing.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of the type and count of accelerator cards attached to the instance.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "gvnic": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether or not gvnic is enabled",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Enable or disable gvnic in the node pool.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "host_maintenance_policy": {
              "block": {
                "attributes": {
                  "maintenance_interval": {
                    "description": ".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The maintenance policy for the hosts on which the GKE VMs run on.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "kubelet_config": {
              "block": {
                "attributes": {
                  "cpu_cfs_quota": {
                    "description": "Enable CPU CFS quota enforcement for containers that specify CPU limits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "cpu_cfs_quota_period": {
                    "description": "Set the CPU CFS quota period value 'cpu.cfs_period_us'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cpu_manager_policy": {
                    "description": "Control the CPU management policy on the node.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "insecure_kubelet_readonly_port_enabled": {
                    "computed": true,
                    "description": "Controls whether the kubelet read-only port is enabled. It is strongly recommended to set this to ` + "`" + `FALSE` + "`" + `. Possible values: ` + "`" + `TRUE` + "`" + `, ` + "`" + `FALSE` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pod_pids_limit": {
                    "description": "Controls the maximum number of processes allowed to run in a pod.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Node kubelet configs.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "linux_node_config": {
              "block": {
                "attributes": {
                  "cgroup_mode": {
                    "computed": true,
                    "description": "cgroupMode specifies the cgroup mode to be used on the node.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sysctls": {
                    "description": "The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "hugepages_config": {
                    "block": {
                      "attributes": {
                        "hugepage_size_1g": {
                          "description": "Amount of 1G hugepages.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "hugepage_size_2m": {
                          "description": "Amount of 2M hugepages.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Amounts for 2M and 1G hugepages.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Parameters that can be configured on Linux nodes.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "local_nvme_ssd_block_config": {
              "block": {
                "attributes": {
                  "local_ssd_count": {
                    "description": "Number of raw-block local NVMe SSD disks to be attached to the node. Each local SSD is 375 GB in size.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Parameters for raw-block local NVMe SSDs.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "reservation_affinity": {
              "block": {
                "attributes": {
                  "consume_reservation_type": {
                    "description": "Corresponds to the type of reservation consumption.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "The label key of a reservation resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "description": "The label values of the reservation resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "The reservation affinity configuration for the node pool.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "secondary_boot_disks": {
              "block": {
                "attributes": {
                  "disk_image": {
                    "description": "Disk image to create the secondary boot disk from",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "mode": {
                    "description": "Mode for how the secondary boot disk is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Secondary boot disks for preloading data or container images.",
                "description_kind": "plain"
              },
              "max_items": 127,
              "nesting_mode": "list"
            },
            "shielded_instance_config": {
              "block": {
                "attributes": {
                  "enable_integrity_monitoring": {
                    "description": "Defines whether the instance has integrity monitoring enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_secure_boot": {
                    "description": "Defines whether the instance has Secure Boot enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Shielded Instance options.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sole_tenant_config": {
              "block": {
                "block_types": {
                  "node_affinity": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description": ".",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operator": {
                          "description": ".",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description": ".",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": ".",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description": "Node affinity options for sole tenant node pools.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "taint": {
              "block": {
                "attributes": {
                  "effect": {
                    "description": "Effect for taint.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Key for taint.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value for taint.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "List of Kubernetes taints to be applied to each node.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "workload_metadata_config": {
              "block": {
                "attributes": {
                  "mode": {
                    "description": "Mode is the configuration for how to expose metadata to workloads running on the node.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The workload metadata configuration for this node.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The configuration of the nodepool",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "placement_policy": {
        "block": {
          "attributes": {
            "policy_name": {
              "description": "If set, refers to the name of a custom resource policy supplied by the user. The resource policy must be in the same project and region as the node pool. If not found, InvalidArgument error is returned.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tpu_topology": {
              "description": "TPU placement topology for pod slice node pool. https://cloud.google.com/tpu/docs/types-topologies#tpu_topologies",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Type defines the type of placement policy",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Specifies the node placement policy",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "queued_provisioning": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether nodes in this node pool are obtainable solely through the ProvisioningRequest API",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Specifies the configuration of queued provisioning",
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
      },
      "upgrade_settings": {
        "block": {
          "attributes": {
            "max_surge": {
              "computed": true,
              "description": "The number of additional nodes that can be added to the node pool during an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unavailable": {
              "computed": true,
              "description": "The number of nodes that can be simultaneously unavailable during an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "strategy": {
              "description": "Update strategy for the given nodepool.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "blue_green_settings": {
              "block": {
                "attributes": {
                  "node_pool_soak_duration": {
                    "computed": true,
                    "description": "Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "standard_rollout_policy": {
                    "block": {
                      "attributes": {
                        "batch_node_count": {
                          "computed": true,
                          "description": "Number of blue nodes to drain in a batch.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "batch_percentage": {
                          "computed": true,
                          "description": "Percentage of the blue pool nodes to drain in a batch.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "batch_soak_duration": {
                          "computed": true,
                          "description": "Soak time after each batch gets drained.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Standard rollout policy is the default policy for blue-green.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Settings for BlueGreen node pool upgrade.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleContainerNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerNodePool), &result)
	return &result
}
