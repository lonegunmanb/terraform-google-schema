package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocCluster = `{
  "block": {
    "attributes": {
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "graceful_decommission_timeout": {
        "description": "The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a terraform apply",
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
      "labels": {
        "description": "The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.\n\t\t\t\t\n\t\t\t\t**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\n\t\t\t\tPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name of the cluster, unique within the project and zone.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region in which the cluster and associated nodes will be created in. Defaults to global.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "cluster_config": {
        "block": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description": " The name of the cloud storage bucket ultimately used to house the staging data for the cluster. If staging_bucket is specified, it will contain this value, otherwise it will be the auto generated name.",
              "description_kind": "plain",
              "type": "string"
            },
            "staging_bucket": {
              "description": "The Cloud Storage staging bucket used to stage files, such as Hadoop jars, between client machines and the cluster. Note: If you don't explicitly specify a staging_bucket then GCP will auto create / assign one for you. However, you are not guaranteed an auto generated bucket which is solely dedicated to your cluster; it may be shared with other clusters in the same region/zone also choosing to use the auto generation option.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "temp_bucket": {
              "computed": true,
              "description": "The Cloud Storage temp bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. Note: If you don't explicitly specify a temp_bucket then GCP will auto create / assign one for you.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "autoscaling_config": {
              "block": {
                "attributes": {
                  "policy_uri": {
                    "description": "The autoscaling policy used by the cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The autoscaling policy config associated with the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "auxiliary_node_groups": {
              "block": {
                "attributes": {
                  "node_group_id": {
                    "computed": true,
                    "description": "A node group ID. Generated if not specified. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of from 3 to 33 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "node_group": {
                    "block": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The Node group resource name.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "roles": {
                          "description": "Node group roles.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "node_group_config": {
                          "block": {
                            "attributes": {
                              "instance_names": {
                                "computed": true,
                                "description": "List of auxiliary node group instance names which have been assigned to the cluster.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "machine_type": {
                                "computed": true,
                                "description": "The name of a Google Compute Engine machine type to create for the master",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "min_cpu_platform": {
                                "computed": true,
                                "description": "The name of a minimum generation of CPU family for the auxiliary node group. If not specified, GCP will default to a predetermined computed value for each zone.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "num_instances": {
                                "computed": true,
                                "description": "Specifies the number of auxiliary nodes to create. If not specified, GCP will default to a predetermined computed value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "accelerators": {
                                "block": {
                                  "attributes": {
                                    "accelerator_count": {
                                      "description": "The number of the accelerator cards of this type exposed to this instance. Often restricted to one of 1, 2, 4, or 8.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "accelerator_type": {
                                      "description": "The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "set"
                              },
                              "disk_config": {
                                "block": {
                                  "attributes": {
                                    "boot_disk_size_gb": {
                                      "computed": true,
                                      "description": "Size of the primary disk attached to each node, specified in GB. The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "boot_disk_type": {
                                      "description": "The disk type of the primary disk attached to each node. Such as \"pd-ssd\" or \"pd-standard\". Defaults to \"pd-standard\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "local_ssd_interface": {
                                      "description": "Interface type of local SSDs (default is \"scsi\"). Valid values: \"scsi\" (Small Computer System Interface), \"nvme\" (Non-Volatile Memory Express).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "num_local_ssds": {
                                      "computed": true,
                                      "description": "The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "Disk Config",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The node group instance group configuration.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Node group configuration.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The node group settings.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "dataproc_metric_config": {
              "block": {
                "block_types": {
                  "metrics": {
                    "block": {
                      "attributes": {
                        "metric_overrides": {
                          "description": "Specify one or more [available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics) to collect.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "metric_source": {
                          "description": "A source for the collection of Dataproc OSS metrics (see [available OSS metrics] (https://cloud.google.com//dataproc/docs/guides/monitoring#available_oss_metrics)).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Metrics sources to enable.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The config for Dataproc metrics.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "encryption_config": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The Customer managed encryption keys settings for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "endpoint_config": {
              "block": {
                "attributes": {
                  "enable_http_port_access": {
                    "description": "The flag to enable http access to specific ports on the cluster from external sources (aka Component Gateway). Defaults to false.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "http_ports": {
                    "computed": true,
                    "description": "The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true.",
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "The config settings for port access on the cluster. Structure defined below.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gce_cluster_config": {
              "block": {
                "attributes": {
                  "internal_ip_only": {
                    "description": "By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. If set to true, all instances in the cluster will only have internal IP addresses. Note: Private Google Access (also known as privateIpGoogleAccess) must be enabled on the subnetwork that the cluster will be launched in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "metadata": {
                    "description": "A map of the Compute Engine metadata entries to add to all instances",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "network": {
                    "computed": true,
                    "description": "The name or self_link of the Google Compute Engine network to the cluster will be part of. Conflicts with subnetwork. If neither is specified, this defaults to the \"default\" network.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account": {
                    "description": "The service account to be used by the Node VMs. If not specified, the \"default\" service account is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_scopes": {
                    "computed": true,
                    "description": "The set of Google API scopes to be made available on all of the node VMs under the service_account specified. These can be either FQDNs, or scope aliases.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnetwork": {
                    "description": "The name or self_link of the Google Compute Engine subnetwork the cluster will be part of. Conflicts with network.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "description": "The list of instance tags applied to instances in the cluster. Tags are used to identify valid sources or targets for network firewalls.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "zone": {
                    "computed": true,
                    "description": "The GCP zone where your data is stored and used (i.e. where the master and the worker nodes will be created in). If region is set to 'global' (default) then zone is mandatory, otherwise GCP is able to make use of Auto Zone Placement to determine this automatically for you. Note: This setting additionally determines and restricts which computing resources are available for use with other configs such as cluster_config.master_config.machine_type and cluster_config.worker_config.machine_type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "node_group_affinity": {
                    "block": {
                      "attributes": {
                        "node_group_uri": {
                          "description": "The URI of a sole-tenant that the cluster will be created on.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Node Group Affinity for sole-tenant clusters.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "reservation_affinity": {
                    "block": {
                      "attributes": {
                        "consume_reservation_type": {
                          "description": "Type of reservation to consume.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key": {
                          "description": "Corresponds to the label key of reservation resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "values": {
                          "description": "Corresponds to the label values of reservation resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description": "Reservation Affinity for consuming Zonal reservation.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "shielded_instance_config": {
                    "block": {
                      "attributes": {
                        "enable_integrity_monitoring": {
                          "description": "Defines whether instances have integrity monitoring enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_secure_boot": {
                          "description": "Defines whether instances have Secure Boot enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_vtpm": {
                          "description": "Defines whether instances have the vTPM enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Shielded Instance Config for clusters using Compute Engine Shielded VMs.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Common config settings for resources of Google Compute Engine cluster instances, applicable to all instances in the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "initialization_action": {
              "block": {
                "attributes": {
                  "script": {
                    "description": "The script to be executed during initialization of the cluster. The script must be a GCS file with a gs:// prefix.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout_sec": {
                    "description": "The maximum duration (in seconds) which script is allowed to take to execute its action. GCP will default to a predetermined computed value if not set (currently 300).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Commands to execute on each node after config is completed. You can specify multiple versions of these.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "lifecycle_config": {
              "block": {
                "attributes": {
                  "auto_delete_time": {
                    "description": "The time when cluster will be auto-deleted. A timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "idle_delete_ttl": {
                    "description": "The duration to keep the cluster alive while idling (no jobs running). After this TTL, the cluster will be deleted. Valid range: [10m, 14d].",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "idle_start_time": {
                    "computed": true,
                    "description": "Time when the cluster became idle (most recent job finished) and became eligible for deletion due to idleness.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "The settings for auto deletion cluster schedule.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "master_config": {
              "block": {
                "attributes": {
                  "image_uri": {
                    "computed": true,
                    "description": "The URI for the image to use for this master",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_names": {
                    "computed": true,
                    "description": "List of master instance names which have been assigned to the cluster.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "machine_type": {
                    "computed": true,
                    "description": "The name of a Google Compute Engine machine type to create for the master",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_cpu_platform": {
                    "computed": true,
                    "description": "The name of a minimum generation of CPU family for the master. If not specified, GCP will default to a predetermined computed value for each zone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "num_instances": {
                    "computed": true,
                    "description": "Specifies the number of master nodes to create. If not specified, GCP will default to a predetermined computed value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "accelerators": {
                    "block": {
                      "attributes": {
                        "accelerator_count": {
                          "description": "The number of the accelerator cards of this type exposed to this instance. Often restricted to one of 1, 2, 4, or 8.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "accelerator_type": {
                          "description": "The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description": "Size of the primary disk attached to each node, specified in GB. The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description": "The disk type of the primary disk attached to each node. Such as \"pd-ssd\" or \"pd-standard\". Defaults to \"pd-standard\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_ssd_interface": {
                          "description": "Interface type of local SSDs (default is \"scsi\"). Valid values: \"scsi\" (Small Computer System Interface), \"nvme\" (Non-Volatile Memory Express).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
                          "description": "The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Disk Config",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The Compute Engine config settings for the cluster's master instance.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "metastore_config": {
              "block": {
                "attributes": {
                  "dataproc_metastore_service": {
                    "description": "Resource name of an existing Dataproc Metastore service.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specifies a Metastore configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "preemptible_worker_config": {
              "block": {
                "attributes": {
                  "instance_names": {
                    "computed": true,
                    "description": "List of preemptible instance names which have been assigned to the cluster.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "num_instances": {
                    "computed": true,
                    "description": "Specifies the number of preemptible nodes to create. Defaults to 0.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "preemptibility": {
                    "description": "Specifies the preemptibility of the secondary nodes. Defaults to PREEMPTIBLE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description": "Size of the primary disk attached to each preemptible worker node, specified in GB. The smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description": "The disk type of the primary disk attached to each preemptible worker node. Such as \"pd-ssd\" or \"pd-standard\". Defaults to \"pd-standard\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_ssd_interface": {
                          "description": "Interface type of local SSDs (default is \"scsi\"). Valid values: \"scsi\" (Small Computer System Interface), \"nvme\" (Non-Volatile Memory Express).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
                          "description": "The amount of local SSD disks that will be attached to each preemptible worker node. Defaults to 0.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Disk Config",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "instance_flexibility_policy": {
                    "block": {
                      "attributes": {
                        "instance_selection_results": {
                          "computed": true,
                          "description": "A list of instance selection results in the group.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            [
                              "object",
                              {
                                "machine_type": "string",
                                "vm_count": "number"
                              }
                            ]
                          ]
                        }
                      },
                      "block_types": {
                        "instance_selection_list": {
                          "block": {
                            "attributes": {
                              "machine_types": {
                                "computed": true,
                                "description": "Full machine-type names, e.g. \"n1-standard-16\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "rank": {
                                "computed": true,
                                "description": "Preference of this instance selection. Lower number means higher preference. Dataproc will first try to create a VM based on the machine-type with priority rank and fallback to next rank based on availability. Machine types and instance selections with the same priority have the same preference.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "List of instance selection options that the group will use when creating new VMs.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Instance flexibility Policy allowing a mixture of VM shapes and provisioning models.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The Google Compute Engine config settings for the additional (aka preemptible) instances in a cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "security_config": {
              "block": {
                "block_types": {
                  "kerberos_config": {
                    "block": {
                      "attributes": {
                        "cross_realm_trust_admin_server": {
                          "description": "The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_realm_trust_kdc": {
                          "description": "The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_realm_trust_realm": {
                          "description": "The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cross_realm_trust_shared_password_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster\nKerberos realm and the remote trusted realm, in a cross realm trust relationship.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enable_kerberos": {
                          "description": "Flag to indicate whether to Kerberize the cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "kdc_db_key_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_password_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "keystore_password_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing\nthe password to the user provided keystore. For the self-signed certificate, this password is generated\nby Dataproc",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "keystore_uri": {
                          "description": "The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kms_key_uri": {
                          "description": "The uri of the KMS key used to encrypt various sensitive files.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "realm": {
                          "description": "The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "root_principal_password_uri": {
                          "description": "The cloud Storage URI of a KMS encrypted file containing the root principal password.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "tgt_lifetime_hours": {
                          "description": "The lifetime of the ticket granting ticket, in hours.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "truststore_password_uri": {
                          "description": "The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "truststore_uri": {
                          "description": "The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Kerberos related configuration",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Security related configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "software_config": {
              "block": {
                "attributes": {
                  "image_version": {
                    "computed": true,
                    "description": "The Cloud Dataproc image version to use for the cluster - this controls the sets of software versions installed onto the nodes when you create clusters. If not specified, defaults to the latest version.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "optional_components": {
                    "description": "The set of optional components to activate on the cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "override_properties": {
                    "description": "A list of override and additional properties (key/value pairs) used to modify various aspects of the common configuration files used when creating a cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "properties": {
                    "computed": true,
                    "description": "A list of the properties used to set the daemon config files. This will include any values supplied by the user via cluster_config.software_config.override_properties",
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "The config settings for software inside the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "worker_config": {
              "block": {
                "attributes": {
                  "image_uri": {
                    "computed": true,
                    "description": "The URI for the image to use for this master/worker",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_names": {
                    "computed": true,
                    "description": "List of master/worker instance names which have been assigned to the cluster.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "machine_type": {
                    "computed": true,
                    "description": "The name of a Google Compute Engine machine type to create for the master/worker",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_cpu_platform": {
                    "computed": true,
                    "description": "The name of a minimum generation of CPU family for the master/worker. If not specified, GCP will default to a predetermined computed value for each zone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_num_instances": {
                    "computed": true,
                    "description": "The minimum number of primary worker instances to create.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "num_instances": {
                    "computed": true,
                    "description": "Specifies the number of worker nodes to create. If not specified, GCP will default to a predetermined computed value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "accelerators": {
                    "block": {
                      "attributes": {
                        "accelerator_count": {
                          "description": "The number of the accelerator cards of this type exposed to this instance. Often restricted to one of 1, 2, 4, or 8.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "accelerator_type": {
                          "description": "The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "disk_config": {
                    "block": {
                      "attributes": {
                        "boot_disk_size_gb": {
                          "computed": true,
                          "description": "Size of the primary disk attached to each node, specified in GB. The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "boot_disk_type": {
                          "description": "The disk type of the primary disk attached to each node. Such as \"pd-ssd\" or \"pd-standard\". Defaults to \"pd-standard\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_ssd_interface": {
                          "description": "Interface type of local SSDs (default is \"scsi\"). Valid values: \"scsi\" (Small Computer System Interface), \"nvme\" (Non-Volatile Memory Express).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "num_local_ssds": {
                          "computed": true,
                          "description": "The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Disk Config",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The Compute Engine config settings for the cluster's worker instances.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Allows you to configure various aspects of the cluster.",
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
      "virtual_cluster_config": {
        "block": {
          "attributes": {
            "staging_bucket": {
              "description": "A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "auxiliary_services_config": {
              "block": {
                "block_types": {
                  "metastore_config": {
                    "block": {
                      "attributes": {
                        "dataproc_metastore_service": {
                          "description": "The Hive Metastore configuration for this workload.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The Hive Metastore configuration for this workload.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "spark_history_server_config": {
                    "block": {
                      "attributes": {
                        "dataproc_cluster": {
                          "description": "Resource name of an existing Dataproc Cluster to act as a Spark History Server for the workload.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The Spark History Server configuration for the workload.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Auxiliary services configuration for a Cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "kubernetes_cluster_config": {
              "block": {
                "attributes": {
                  "kubernetes_namespace": {
                    "description": "A namespace within the Kubernetes cluster to deploy into. If this namespace does not exist, it is created. If it exists, Dataproc verifies that another Dataproc VirtualCluster is not installed into it. If not specified, the name of the Dataproc Cluster is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "gke_cluster_config": {
                    "block": {
                      "attributes": {
                        "gke_cluster_target": {
                          "description": "A target GKE cluster to deploy to. It must be in the same project and region as the Dataproc cluster (the GKE cluster can be zonal or regional). Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "node_pool_target": {
                          "block": {
                            "attributes": {
                              "node_pool": {
                                "description": "The target GKE node pool. Format: 'projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{nodePool}'",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "roles": {
                                "description": "The roles associated with the GKE node pool.",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "node_pool_config": {
                                "block": {
                                  "attributes": {
                                    "locations": {
                                      "description": "The list of Compute Engine zones where node pool nodes associated with a Dataproc on GKE virtual cluster will be located.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "autoscaling": {
                                      "block": {
                                        "attributes": {
                                          "max_node_count": {
                                            "description": "The maximum number of nodes in the node pool. Must be \u003e= minNodeCount, and must be \u003e 0.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "min_node_count": {
                                            "description": "The minimum number of nodes in the node pool. Must be \u003e= 0 and \u003c= maxNodeCount.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "description": "The autoscaler configuration for this node pool. The autoscaler is enabled only when a valid configuration is present.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "config": {
                                      "block": {
                                        "attributes": {
                                          "local_ssd_count": {
                                            "description": "The minimum number of nodes in the node pool. Must be \u003e= 0 and \u003c= maxNodeCount.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "machine_type": {
                                            "description": "The name of a Compute Engine machine type.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "min_cpu_platform": {
                                            "description": "Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or a newer CPU platform. Specify the friendly names of CPU platforms, such as \"Intel Haswell\" or \"Intel Sandy Bridge\".",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "preemptible": {
                                            "description": "Whether the nodes are created as preemptible VM instances. Preemptible nodes cannot be used in a node pool with the CONTROLLER role or in the DEFAULT node pool if the CONTROLLER role is not assigned (the DEFAULT node pool will assume the CONTROLLER role).",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "spot": {
                                            "description": "Spot flag for enabling Spot VM, which is a rebrand of the existing preemptible flag.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          }
                                        },
                                        "description": "The node pool configuration.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Input only. The configuration for the GKE node pool.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "GKE node pools where workloads will be scheduled. At least one node pool must be assigned the DEFAULT GkeNodePoolTarget.Role. If a GkeNodePoolTarget is not specified, Dataproc constructs a DEFAULT GkeNodePoolTarget.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The configuration for running the Dataproc cluster on GKE.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "kubernetes_software_config": {
                    "block": {
                      "attributes": {
                        "component_version": {
                          "description": "The components that should be installed in this Dataproc cluster. The key must be a string from the KubernetesComponent enumeration. The value is the version of the software to be installed.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "properties": {
                          "computed": true,
                          "description": "The properties to set on daemon config files. Property keys are specified in prefix:property format, for example spark:spark.kubernetes.container.image.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "The software configuration for this Dataproc cluster running on Kubernetes.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration for running the Dataproc cluster on Kubernetes.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying compute resources, for example, when creating a Dataproc-on-GKE cluster. Dataproc may set default values, and values may change when clusters are updated. Exactly one of config or virtualClusterConfig must be specified.",
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

func GoogleDataprocClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocCluster), &result)
	return &result
}
