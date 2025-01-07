package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComposerEnvironment = `{
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: [a-z]([-a-z0-9]*[a-z0-9])?. Label values must be between 0 and 63 characters long and must conform to the regular expression ([a-z]([-a-z0-9]*[a-z0-9])?)?. No more than 64 labels can be associated with a given environment. Both keys and values must be \u003c= 128 bytes in size.\n\n\t\t\t\t**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\n\t\t\t\tPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Name of the environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The location or Compute Engine region for the environment.",
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
      "config": {
        "block": {
          "attributes": {
            "airflow_uri": {
              "computed": true,
              "description": "The URI of the Apache Airflow Web UI hosted within this environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "dag_gcs_prefix": {
              "computed": true,
              "description": "The Cloud Storage prefix of the DAGs for this environment. Although Cloud Storage objects reside in a flat namespace, a hierarchical file tree can be simulated using '/'-delimited object name prefixes. DAG objects for this environment reside in a simulated directory with this prefix.",
              "description_kind": "plain",
              "type": "string"
            },
            "enable_private_builds_only": {
              "computed": true,
              "description": "Optional. If true, builds performed during operations that install Python packages have only private connectivity to Google services. If false, the builds also have access to the internet.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_private_environment": {
              "computed": true,
              "description": "Optional. If true, a private Composer environment will be created.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "environment_size": {
              "computed": true,
              "description": "The size of the Cloud Composer environment. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gke_cluster": {
              "computed": true,
              "description": "The Kubernetes Engine cluster used to run this environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "node_count": {
              "computed": true,
              "description": "The number of nodes in the Kubernetes Engine cluster that will be used to run this environment. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "resilience_mode": {
              "computed": true,
              "description": "Whether high resilience is enabled or not. This field is supported for Cloud Composer environments in versions composer-2.1.15-airflow-*.*.* and newer.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_retention_config": {
              "block": {
                "block_types": {
                  "airflow_metadata_retention_config": {
                    "block": {
                      "attributes": {
                        "retention_days": {
                          "computed": true,
                          "description": "How many days data should be retained for. This field is supported for Cloud Composer environments in composer 3 and newer.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "retention_mode": {
                          "computed": true,
                          "description": "Whether database retention is enabled or not. This field is supported for Cloud Composer environments in composer 3 and newer.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Optional. The configuration setting for database retention.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "task_logs_retention_config": {
                    "block": {
                      "attributes": {
                        "storage_mode": {
                          "description": "Whether logs in cloud logging only is enabled or not. This field is supported for Cloud Composer environments in versions composer-2.0.32-airflow-2.1.4 and newer.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Optional. The configuration setting for Task Logs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration setting for Airflow data retention mechanism. This field is supported for Cloud Composer environments in versions composer-2.0.32-airflow-2.1.4. or newer",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "database_config": {
              "block": {
                "attributes": {
                  "machine_type": {
                    "description": "Optional. Cloud SQL machine type used by Airflow database. It has to be one of: db-n1-standard-2, db-n1-standard-4, db-n1-standard-8 or db-n1-standard-16. If not specified, db-n1-standard-2 will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "zone": {
                    "description": "Optional. Cloud SQL database preferred zone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The configuration of Cloud SQL instance that is used by the Apache Airflow software. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "encryption_config": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "Optional. Customer-managed Encryption Key available through Google's Key Management Service. Cannot be updated.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The encryption options for the Composer environment and its dependencies.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "maintenance_window": {
              "block": {
                "attributes": {
                  "end_time": {
                    "description": "Maintenance window end time. It is used only to calculate the duration of the maintenance window. The value for end-time must be in the future, relative to 'start_time'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "recurrence": {
                    "description": "Maintenance window recurrence. Format is a subset of RFC-5545 (https://tools.ietf.org/html/rfc5545) 'RRULE'. The only allowed values for 'FREQ' field are 'FREQ=DAILY' and 'FREQ=WEEKLY;BYDAY=...'. Example values: 'FREQ=WEEKLY;BYDAY=TU,WE', 'FREQ=DAILY'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start_time": {
                    "description": "Start time of the first recurrence of the maintenance window.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The configuration for Cloud Composer maintenance window.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "master_authorized_networks_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether or not master authorized networks is enabled.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "cidr_blocks": {
                    "block": {
                      "attributes": {
                        "cidr_block": {
                          "description": "cidr_block must be specified in CIDR notation.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "display_name": {
                          "description": "display_name is a field for users to identify CIDR blocks.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "cidr_blocks define up to 50 external networks that could access Kubernetes master through HTTPS.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description": "Configuration options for the master authorized networks feature. Enabled master authorized networks will disallow all external traffic to access Kubernetes master through HTTPS except traffic from the given CIDR blocks, Google Compute Engine Public IPs and Google Prod IPs.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "node_config": {
              "block": {
                "attributes": {
                  "composer_internal_ipv4_cidr_block": {
                    "computed": true,
                    "description": "IPv4 cidr range that will be used by Composer internal components.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "composer_network_attachment": {
                    "computed": true,
                    "description": "PSC (Private Service Connect) Network entry point. Customers can pre-create the Network Attachment and point Cloud Composer environment to use. It is possible to share network attachment among many environments, provided enough IP addresses are available.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_size_gb": {
                    "computed": true,
                    "description": "The disk size in GB used for node VMs. Minimum size is 20GB. If unspecified, defaults to 100GB. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "enable_ip_masq_agent": {
                    "computed": true,
                    "description": "Deploys 'ip-masq-agent' daemon set in the GKE cluster and defines nonMasqueradeCIDRs equals to pod IP range so IP masquerading is used for all destination addresses, except between pods traffic. See: https://cloud.google.com/kubernetes-engine/docs/how-to/ip-masquerade-agent",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "machine_type": {
                    "computed": true,
                    "description": "The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: \"projects/{project}/zones/{zone}/machineTypes/{machineType}\". Must belong to the enclosing environment's project and region/zone. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network": {
                    "computed": true,
                    "description": "The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: \"projects/{project}/zones/{zone}/machineTypes/{machineType}\". Must belong to the enclosing environment's project and region/zone. The network must belong to the environment's project. If unspecified, the \"default\" network ID in the environment's project is used. If a Custom Subnet Network is provided, subnetwork must also be provided.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "computed": true,
                    "description": "The set of Google API scopes to be made available on all node VMs. Cannot be updated. If empty, defaults to [\"https://www.googleapis.com/auth/cloud-platform\"]. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "service_account": {
                    "computed": true,
                    "description": "The Google Cloud Platform Service Account to be used by the node VMs. If a service account is not specified, the \"default\" Compute Engine service account is used. Cannot be updated. If given, note that the service account must have roles/composer.worker for any GCP resources created under the Cloud Composer Environment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnetwork": {
                    "computed": true,
                    "description": "The Compute Engine subnetwork to be used for machine communications, specified as a self-link, relative resource name (e.g. \"projects/{project}/regions/{region}/subnetworks/{subnetwork}\"), or by name. If subnetwork is provided, network must also be provided and the subnetwork must belong to the enclosing environment's project and region.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "description": "The list of instance tags applied to all node VMs. Tags are used to identify valid sources or targets for network firewalls. Each tag within the list must comply with RFC1035. Cannot be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "zone": {
                    "computed": true,
                    "description": "The Compute Engine zone in which to deploy the VMs running the Apache Airflow software, specified as the zone name or relative resource name (e.g. \"projects/{project}/zones/{zone}\"). Must belong to the enclosing environment's project and region. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "ip_allocation_policy": {
                    "block": {
                      "attributes": {
                        "cluster_ipv4_cidr_block": {
                          "description": "The IP address range used to allocate IP addresses to pods in the cluster. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. Specify either cluster_secondary_range_name or cluster_ipv4_cidr_block but not both.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cluster_secondary_range_name": {
                          "description": "The name of the cluster's secondary range used to allocate IP addresses to pods. Specify either cluster_secondary_range_name or cluster_ipv4_cidr_block but not both. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "services_ipv4_cidr_block": {
                          "description": "The IP address range used to allocate IP addresses in this cluster. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. Specify either services_secondary_range_name or services_ipv4_cidr_block but not both.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "services_secondary_range_name": {
                          "description": "The name of the services' secondary range used to allocate IP addresses to the cluster. Specify either services_secondary_range_name or services_ipv4_cidr_block but not both. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "use_ip_aliases": {
                          "description": "Whether or not to enable Alias IPs in the GKE cluster. If true, a VPC-native cluster is created. Defaults to true if the ip_allocation_policy block is present in config. This field is only supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*. Environments in newer versions always use VPC-native GKE clusters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Configuration for controlling how IPs are allocated in the GKE cluster. Cannot be updated.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration used for the Kubernetes Engine cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "private_environment_config": {
              "block": {
                "attributes": {
                  "cloud_composer_connection_subnetwork": {
                    "computed": true,
                    "description": "When specified, the environment will use Private Service Connect instead of VPC peerings to connect to Cloud SQL in the Tenant Project, and the PSC endpoint in the Customer Project will use an IP address from this subnetwork. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cloud_composer_network_ipv4_cidr_block": {
                    "computed": true,
                    "description": "The CIDR block from which IP range for Cloud Composer Network in tenant project will be reserved. Needs to be disjoint from private_cluster_config.master_ipv4_cidr_block and cloud_sql_ipv4_cidr_block. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cloud_sql_ipv4_cidr_block": {
                    "computed": true,
                    "description": "The CIDR block from which IP range in tenant project will be reserved for Cloud SQL. Needs to be disjoint from web_server_ipv4_cidr_block.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "connection_type": {
                    "computed": true,
                    "description": "Mode of internal communication within the Composer environment. Must be one of \"VPC_PEERING\" or \"PRIVATE_SERVICE_CONNECT\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_private_endpoint": {
                    "description": "If true, access to the public endpoint of the GKE cluster is denied. If this field is set to true, ip_allocation_policy.use_ip_aliases must be set to true for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_privately_used_public_ips": {
                    "computed": true,
                    "description": "When enabled, IPs from public (non-RFC1918) ranges can be used for ip_allocation_policy.cluster_ipv4_cidr_block and ip_allocation_policy.service_ipv4_cidr_block.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "master_ipv4_cidr_block": {
                    "computed": true,
                    "description": "The IP range in CIDR notation to use for the hosted master network. This range is used for assigning internal IP addresses to the cluster master or set of masters and to the internal load balancer virtual IP. This range must not overlap with any other ranges in use within the cluster's network. If left blank, the default value of '172.16.0.0/28' is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "web_server_ipv4_cidr_block": {
                    "computed": true,
                    "description": "The CIDR block from which IP range for web server will be reserved. Needs to be disjoint from master_ipv4_cidr_block and cloud_sql_ipv4_cidr_block. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The configuration used for the Private IP Cloud Composer environment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "recovery_config": {
              "block": {
                "block_types": {
                  "scheduled_snapshots_config": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description": "When enabled, Cloud Composer periodically saves snapshots of your environment to a Cloud Storage bucket.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "snapshot_creation_schedule": {
                          "description": "Snapshot schedule, in the unix-cron format.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "snapshot_location": {
                          "description": "the URI of a bucket folder where to save the snapshot.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "time_zone": {
                          "description": "A time zone for the schedule. This value is a time offset and does not take into account daylight saving time changes. Valid values are from UTC-12 to UTC+12. Examples: UTC, UTC-01, UTC+03.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The configuration settings for scheduled snapshots.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The recovery configuration settings for the Cloud Composer environment",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "software_config": {
              "block": {
                "attributes": {
                  "airflow_config_overrides": {
                    "description": "Apache Airflow configuration properties to override. Property keys contain the section and property names, separated by a hyphen, for example \"core-dags_are_paused_at_creation\". Section names must not contain hyphens (\"-\"), opening square brackets (\"[\"), or closing square brackets (\"]\"). The property name must not be empty and cannot contain \"=\" or \";\". Section and property names cannot contain characters: \".\" Apache Airflow configuration property names must be written in snake_case. Property values can contain any character, and can be written in any lower/upper case format. Certain Apache Airflow configuration property values are blacklisted, and cannot be overridden.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "env_variables": {
                    "description": "Additional environment variables to provide to the Apache Airflow scheduler, worker, and webserver processes. Environment variable names must match the regular expression [a-zA-Z_][a-zA-Z0-9_]*. They cannot specify Apache Airflow software configuration overrides (they cannot match the regular expression AIRFLOW__[A-Z0-9_]+__[A-Z0-9_]+), and they cannot match any of the following reserved names: AIRFLOW_HOME C_FORCE_ROOT CONTAINER_NAME DAGS_FOLDER GCP_PROJECT GCS_BUCKET GKE_CLUSTER_NAME SQL_DATABASE SQL_INSTANCE SQL_PASSWORD SQL_PROJECT SQL_REGION SQL_USER.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "image_version": {
                    "computed": true,
                    "description": "The version of the software running in the environment. This encapsulates both the version of Cloud Composer functionality and the version of Apache Airflow. It must match the regular expression composer-([0-9]+(\\.[0-9]+\\.[0-9]+(-preview\\.[0-9]+)?)?|latest)-airflow-([0-9]+(\\.[0-9]+(\\.[0-9]+)?)?). The Cloud Composer portion of the image version is a full semantic version, or an alias in the form of major version number or 'latest'. The Apache Airflow portion of the image version is a full semantic version that points to one of the supported Apache Airflow versions, or an alias in the form of only major or major.minor versions specified. See documentation for more details and version list.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pypi_packages": {
                    "description": "Custom Python Package Index (PyPI) packages to be installed in the environment. Keys refer to the lowercase package name (e.g. \"numpy\"). Values are the lowercase extras and version specifier (e.g. \"==1.12.0\", \"[devel,gcp_api]\", \"[devel]\u003e=1.8.2, \u003c1.9.2\"). To specify a package without pinning it to a version specifier, use the empty string as the value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "python_version": {
                    "computed": true,
                    "description": "The major version of Python used to run the Apache Airflow scheduler, worker, and webserver processes. Can be set to '2' or '3'. If not specified, the default is '2'. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*. Environments in newer versions always use Python major version 3.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scheduler_count": {
                    "computed": true,
                    "description": "The number of schedulers for Airflow. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-2.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "web_server_plugins_mode": {
                    "computed": true,
                    "description": "Should be either 'ENABLED' or 'DISABLED'. Defaults to 'ENABLED'. Used in Composer 3.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "cloud_data_lineage_integration": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description": "Whether or not Cloud Data Lineage integration is enabled.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "description": "The configuration for Cloud Data Lineage integration. Supported for Cloud Composer environments in versions composer-2.1.2-airflow-*.*.* and newer",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration settings for software inside the environment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "web_server_config": {
              "block": {
                "attributes": {
                  "machine_type": {
                    "description": "Optional. Machine type on which Airflow web server is running. It has to be one of: composer-n1-webserver-2, composer-n1-webserver-4 or composer-n1-webserver-8. If not specified, composer-n1-webserver-2 will be used. Value custom is returned only in response, if Airflow web server parameters were manually changed to a non-standard values.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The configuration settings for the Airflow web server App Engine instance. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "web_server_network_access_control": {
              "block": {
                "block_types": {
                  "allowed_ip_range": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description": "A description of this ip range.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "IP address or range, defined using CIDR notation, of requests that this rule applies to. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. IP range prefixes should be properly truncated. For example, 1.2.3.4/24 should be truncated to 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 should be truncated to 2001:db8::/32.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A collection of allowed IP ranges with descriptions.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description": "Network-level access control policy for the Airflow web server.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "workloads_config": {
              "block": {
                "block_types": {
                  "dag_processor": {
                    "block": {
                      "attributes": {
                        "count": {
                          "computed": true,
                          "description": "Number of DAG processors.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "cpu": {
                          "computed": true,
                          "description": "CPU request and limit for DAG processor.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "memory_gb": {
                          "computed": true,
                          "description": "Memory (GB) request and limit for DAG processor.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "storage_gb": {
                          "computed": true,
                          "description": "Storage (GB) request and limit for DAG processor.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Configuration for resources used by DAG processor.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "scheduler": {
                    "block": {
                      "attributes": {
                        "count": {
                          "computed": true,
                          "description": "The number of schedulers.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "cpu": {
                          "computed": true,
                          "description": "CPU request and limit for a single Airflow scheduler replica",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "memory_gb": {
                          "computed": true,
                          "description": "Memory (GB) request and limit for a single Airflow scheduler replica.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "storage_gb": {
                          "computed": true,
                          "description": "Storage (GB) request and limit for a single Airflow scheduler replica.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Configuration for resources used by Airflow schedulers.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "triggerer": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description": "The number of triggerers.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "cpu": {
                          "description": "CPU request and limit for a single Airflow triggerer replica.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "memory_gb": {
                          "description": "Memory (GB) request and limit for a single Airflow triggerer replica.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Configuration for resources used by Airflow triggerers.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "web_server": {
                    "block": {
                      "attributes": {
                        "cpu": {
                          "computed": true,
                          "description": "CPU request and limit for Airflow web server.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "memory_gb": {
                          "computed": true,
                          "description": "Memory (GB) request and limit for Airflow web server.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "storage_gb": {
                          "computed": true,
                          "description": "Storage (GB) request and limit for Airflow web server.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Configuration for resources used by Airflow web server.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "worker": {
                    "block": {
                      "attributes": {
                        "cpu": {
                          "computed": true,
                          "description": "CPU request and limit for a single Airflow worker replica.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_count": {
                          "computed": true,
                          "description": "Maximum number of workers for autoscaling.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "memory_gb": {
                          "computed": true,
                          "description": "Memory (GB) request and limit for a single Airflow worker replica.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min_count": {
                          "computed": true,
                          "description": "Minimum number of workers for autoscaling.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "storage_gb": {
                          "computed": true,
                          "description": "Storage (GB) request and limit for a single Airflow worker replica.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Configuration for resources used by Airflow workers.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The workloads configuration settings for the GKE cluster associated with the Cloud Composer environment. Supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration parameters for this environment.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "storage_config": {
        "block": {
          "attributes": {
            "bucket": {
              "description": "Optional. Name of an existing Cloud Storage bucket to be used by the environment.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration options for storage used by Composer environment.",
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

func GoogleComposerEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComposerEnvironment), &result)
	return &result
}
