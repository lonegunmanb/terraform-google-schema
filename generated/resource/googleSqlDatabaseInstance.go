package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlDatabaseInstance = `{
  "block": {
    "attributes": {
      "available_maintenance_versions": {
        "computed": true,
        "description": "Available Maintenance versions.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "connection_name": {
        "computed": true,
        "description": "The connection name of the instance to be used in connection strings. For example, when connecting with Cloud SQL Proxy.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_version": {
        "description": "The MySQL, PostgreSQL or SQL Server (beta) version to use. Supported values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, MYSQL_8_4, POSTGRES_9_6, POSTGRES_10, POSTGRES_11, POSTGRES_12, POSTGRES_13, POSTGRES_14, POSTGRES_15, POSTGRES_16, POSTGRES_17, SQLSERVER_2017_STANDARD, SQLSERVER_2017_ENTERPRISE, SQLSERVER_2017_EXPRESS, SQLSERVER_2017_WEB. Database Version Policies includes an up-to-date reference of supported versions.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deletion_protection": {
        "description": "Used to block Terraform from deleting a SQL Instance. Defaults to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dns_name": {
        "computed": true,
        "description": "The dns name of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "first_ip_address": {
        "computed": true,
        "description": "The first IPv4 address of any type assigned. This is to support accessing the first address in the list in a terraform output when the resource is configured with a count.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description": "The type of the instance. The valid values are:- 'SQL_INSTANCE_TYPE_UNSPECIFIED', 'CLOUD_SQL_INSTANCE', 'ON_PREMISES_INSTANCE' and 'READ_REPLICA_INSTANCE'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_address": "string",
              "time_to_retire": "string",
              "type": "string"
            }
          ]
        ]
      },
      "maintenance_version": {
        "computed": true,
        "description": "Maintenance version.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_instance_name": {
        "computed": true,
        "description": "The name of the instance that will act as the master in the replication setup. Note, this requires the master to have binary_log_enabled set, as well as existing backups.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the instance. If the name is left blank, Terraform will randomly generate one when the instance is first created. This is done because after a name is used, it cannot be reused for up to one week.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "IPv4 address assigned. This is a workaround for an issue fixed in Terraform 0.12 but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psc_service_attachment_link": {
        "computed": true,
        "description": "The link to service attachment of PSC instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip_address": {
        "computed": true,
        "description": "IPv4 address assigned. This is a workaround for an issue fixed in Terraform 0.12 but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region the instance will sit in. Note, Cloud SQL is not available in all regions. A valid region must be provided to use this resource. If a region is not provided in the resource definition, the provider region will be used instead, but this will be an apply-time error for instances if the provider region is not supported with Cloud SQL. If you choose not to provide the region argument for this resource, make sure you understand this.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica_names": {
        "computed": true,
        "description": "The replicas of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "root_password": {
        "description": "Initial root password. Required for MS SQL Server.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_ca_cert": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "list",
          [
            "object",
            {
              "cert": "string",
              "common_name": "string",
              "create_time": "string",
              "expiration_time": "string",
              "sha1_fingerprint": "string"
            }
          ]
        ]
      },
      "service_account_email_address": {
        "computed": true,
        "description": "The service account email address assigned to the instance.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "clone": {
        "block": {
          "attributes": {
            "allocated_ip_range": {
              "description": "The name of the allocated ip range for the private ip CloudSQL instance. For example: \"google-managed-services-default\". If set, the cloned instance ip will be created in the allocated range. The range name must comply with [RFC 1035](https://tools.ietf.org/html/rfc1035). Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "database_names": {
              "description": "(SQL Server only, use with point_in_time) clone only the specified databases from the source instance. Clone all databases if empty.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "point_in_time": {
              "description": "The timestamp of the point in time that should be restored.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preferred_zone": {
              "description": "(Point-in-time recovery for PostgreSQL only) Clone to an instance in the specified zone. If no zone is specified, clone to the same zone as the source instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_instance_name": {
              "description": "The name of the instance from which the point in time should be restored.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration for creating a new instance as a clone of another instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "replica_configuration": {
        "block": {
          "attributes": {
            "ca_certificate": {
              "description": "PEM representation of the trusted CA's x509 certificate.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cascadable_replica": {
              "description": "Specifies if a SQL Server replica is a cascadable replica. A cascadable replica is a SQL Server cross region replica that supports replica(s) under it.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "client_certificate": {
              "description": "PEM representation of the replica's x509 certificate.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_key": {
              "description": "PEM representation of the replica's private key. The corresponding public key in encoded in the client_certificate.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connect_retry_interval": {
              "description": "The number of seconds between connect retries. MySQL's default is 60 seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dump_file_path": {
              "description": "Path to a SQL file in Google Cloud Storage from which replica instances are created. Format is gs://bucket/filename.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failover_target": {
              "description": "Specifies if the replica is the failover target. If the field is set to true the replica will be designated as a failover replica. If the master instance fails, the replica instance will be promoted as the new master instance. Not supported for Postgres",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "master_heartbeat_period": {
              "description": "Time in ms between replication heartbeats.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "password": {
              "description": "Password for the replication connection.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "ssl_cipher": {
              "description": "Permissible ciphers for use in SSL encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "description": "Username for replication connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "verify_server_certificate": {
              "description": "True if the master's common name value is checked during the SSL handshake.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "The configuration for replication.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "restore_backup_context": {
        "block": {
          "attributes": {
            "backup_run_id": {
              "description": "The ID of the backup run to restore from.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "instance_id": {
              "description": "The ID of the instance that the backup was taken from.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project": {
              "description": "The full project ID of the source instance.",
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
      "settings": {
        "block": {
          "attributes": {
            "activation_policy": {
              "description": "This specifies when the instance should be active. Can be either ALWAYS, NEVER or ON_DEMAND.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "availability_type": {
              "description": "The availability type of the Cloud SQL instance, high availability\n(REGIONAL) or single zone (ZONAL). For all instances, ensure that\nsettings.backup_configuration.enabled is set to true.\nFor MySQL instances, ensure that settings.backup_configuration.binary_log_enabled is set to true.\nFor Postgres instances, ensure that settings.backup_configuration.point_in_time_recovery_enabled\nis set to true. Defaults to ZONAL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "collation": {
              "description": "The name of server instance collation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connector_enforcement": {
              "computed": true,
              "description": "Enables the enforcement of Cloud SQL Auth Proxy or Cloud SQL connectors for all the connections. If enabled, all the direct connections are rejected.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "deletion_protection_enabled": {
              "description": "Configuration to protect against accidental instance deletion.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disk_autoresize": {
              "description": "Enables auto-resizing of the storage size. Defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disk_autoresize_limit": {
              "description": "The maximum size, in GB, to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_size": {
              "computed": true,
              "description": "The size of data disk, in GB. Size of a running instance cannot be reduced but can be increased. The minimum value is 10GB.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_type": {
              "description": "The type of data disk: PD_SSD or PD_HDD. Defaults to PD_SSD.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "edition": {
              "computed": true,
              "description": "The edition of the instance, can be ENTERPRISE or ENTERPRISE_PLUS.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_dataplex_integration": {
              "description": "Enables Dataplex Integration.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_google_ml_integration": {
              "description": "Enables Vertex AI Integration.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pricing_plan": {
              "description": "Pricing plan for this instance, can only be PER_USE.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tier": {
              "description": "The machine type to use. See tiers for more details and supported versions. Postgres supports only shared-core machine types, and custom machine types such as db-custom-2-13312. See the Custom Machine Type Documentation to learn about specifying custom machine types.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "time_zone": {
              "description": "The time_zone to be used by the database engine (supported only for SQL Server), in SQL Server timezone format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_labels": {
              "computed": true,
              "description": "A set of key/value user label pairs to assign to the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "version": {
              "computed": true,
              "description": "Used to make sure changes to the settings block are atomic.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "block_types": {
            "active_directory_config": {
              "block": {
                "attributes": {
                  "domain": {
                    "description": "Domain name of the Active Directory for SQL Server (e.g., mydomain.com).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "advanced_machine_features": {
              "block": {
                "attributes": {
                  "threads_per_core": {
                    "description": "The number of threads per physical core. Can be 1 or 2.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "backup_configuration": {
              "block": {
                "attributes": {
                  "binary_log_enabled": {
                    "description": "True if binary logging is enabled. If settings.backup_configuration.enabled is false, this must be as well. Can only be used with MySQL.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enabled": {
                    "description": "True if backup configuration is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "location": {
                    "description": "Location of the backup configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "point_in_time_recovery_enabled": {
                    "description": "True if Point-in-time recovery is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "start_time": {
                    "computed": true,
                    "description": "HH:MM format time indicating when backup configuration starts.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transaction_log_retention_days": {
                    "computed": true,
                    "description": "The number of days of transaction logs we retain for point in time restore, from 1-7. (For PostgreSQL Enterprise Plus instances, from 1 to 35.)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "backup_retention_settings": {
                    "block": {
                      "attributes": {
                        "retained_backups": {
                          "description": "Number of backups to retain.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "retention_unit": {
                          "description": "The unit that 'retainedBackups' represents. Defaults to COUNT",
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
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "data_cache_config": {
              "block": {
                "attributes": {
                  "data_cache_enabled": {
                    "description": "Whether data cache is enabled for the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Data cache configurations.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "database_flags": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the flag.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value of the flag.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "deny_maintenance_period": {
              "block": {
                "attributes": {
                  "end_date": {
                    "description": "End date before which maintenance will not take place. The date is in format yyyy-mm-dd i.e., 2020-11-01, or mm-dd, i.e., 11-01",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start_date": {
                    "description": "Start date after which maintenance will not take place. The date is in format yyyy-mm-dd i.e., 2020-11-01, or mm-dd, i.e., 11-01",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "time": {
                    "description": "Time in UTC when the \"deny maintenance period\" starts on start_date and ends on end_date. The time is in format: HH:mm:SS, i.e., 00:00:00",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "insights_config": {
              "block": {
                "attributes": {
                  "query_insights_enabled": {
                    "description": "True if Query Insights feature is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "query_plans_per_minute": {
                    "computed": true,
                    "description": "Number of query execution plans captured by Insights per minute for all queries combined. Between 0 and 20. Default to 5.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "query_string_length": {
                    "description": "Maximum query length stored in bytes. Between 256 and 4500. Default to 1024.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "record_application_tags": {
                    "description": "True if Query Insights will record application tags from query when enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "record_client_address": {
                    "description": "True if Query Insights will record client address when enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration of Query Insights.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ip_configuration": {
              "block": {
                "attributes": {
                  "allocated_ip_range": {
                    "description": "The name of the allocated ip range for the private ip CloudSQL instance. For example: \"google-managed-services-default\". If set, the instance ip will be created in the allocated range. The range name must comply with RFC 1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_private_path_for_google_cloud_services": {
                    "description": "Whether Google Cloud services such as BigQuery are allowed to access data in this Cloud SQL instance over a private IP connection. SQLSERVER database type is not supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ipv4_enabled": {
                    "description": "Whether this Cloud SQL instance should be assigned a public IPV4 address. At least ipv4_enabled must be enabled or a private_network must be configured.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "private_network": {
                    "description": "The VPC network from which the Cloud SQL instance is accessible for private IP. For example, projects/myProject/global/networks/default. Specifying a network enables private IP. At least ipv4_enabled must be enabled or a private_network must be configured. This setting can be updated, but it cannot be removed after it is set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_ca_mode": {
                    "computed": true,
                    "description": "Specify how the server certificate's Certificate Authority is hosted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description": "Specify how SSL connection should be enforced in DB connections.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "authorized_networks": {
                    "block": {
                      "attributes": {
                        "expiration_time": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "psc_config": {
                    "block": {
                      "attributes": {
                        "allowed_consumer_projects": {
                          "description": "List of consumer projects that are allow-listed for PSC connections to this instance. This instance can be connected to with PSC from any network in these projects. Each consumer project in this list may be represented by a project number (numeric) or by a project id (alphanumeric).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "psc_enabled": {
                          "description": "Whether PSC connectivity is enabled for this instance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "psc_auto_connections": {
                          "block": {
                            "attributes": {
                              "consumer_network": {
                                "description": "The consumer network of this consumer endpoint. This must be a resource path that includes both the host project and the network name. The consumer host project of this network might be different from the consumer service project.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "consumer_service_project_id": {
                                "description": "The project ID of consumer service project of this consumer endpoint.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A comma-separated list of networks or a comma-separated list of network-project pairs. Each project in this list is represented by a project number (numeric) or by a project ID (alphanumeric). This allows Private Service Connect connections to be created automatically for the specified networks.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "PSC settings for a Cloud SQL instance.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "location_preference": {
              "block": {
                "attributes": {
                  "follow_gae_application": {
                    "description": "A Google App Engine application whose zone to remain in. Must be in the same region as this instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secondary_zone": {
                    "description": "The preferred Compute Engine zone for the secondary/failover",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "zone": {
                    "description": "The preferred compute engine zone.",
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
            "maintenance_window": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Day of week (1-7), starting on Monday",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "hour": {
                    "description": "Hour of day (0-23), ignored if day not set",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "update_track": {
                    "description": "Receive updates after one week (canary) or after two weeks (stable) or after five weeks (week5) of notification.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Declares a one-hour maintenance window when an Instance can automatically restart to apply updates. The maintenance window is specified in UTC time.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "password_validation_policy": {
              "block": {
                "attributes": {
                  "complexity": {
                    "description": "Password complexity.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disallow_username_substring": {
                    "description": "Disallow username as a part of the password.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_password_policy": {
                    "description": "Whether the password policy is enabled or not.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "min_length": {
                    "description": "Minimum number of characters allowed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "password_change_interval": {
                    "description": "Minimum interval after which the password can be changed. This flag is only supported for PostgresSQL.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reuse_interval": {
                    "description": "Number of previous passwords that cannot be reused.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sql_server_audit_config": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description": "The name of the destination bucket (e.g., gs://mybucket).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "retention_interval": {
                    "description": "How long to keep generated audit files. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\"..",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "upload_interval": {
                    "description": "How often to upload generated audit files. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
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
          "description": "The settings to use for the database. The configuration is detailed below.",
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

func GoogleSqlDatabaseInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlDatabaseInstance), &result)
	return &result
}
