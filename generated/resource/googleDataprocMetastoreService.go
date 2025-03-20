package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocMetastoreService = `{
  "block": {
    "attributes": {
      "artifact_gcs_uri": {
        "computed": true,
        "description": "A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time when the metastore service was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_type": {
        "description": "The database type that the Metastore service stores its data. Default value: \"MYSQL\" Possible values: [\"MYSQL\", \"SPANNER\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "description": "Indicates if the dataproc metastore should be protected against accidental deletions.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "endpoint_uri": {
        "computed": true,
        "description": "The URI of the endpoint used to access the metastore service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels for the metastore service.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the metastore service should reside.\nThe default value is 'global'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The relative resource name of the metastore service.",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:\n\n\"projects/{projectNumber}/global/networks/{network_id}\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The TCP port at which the metastore service is reached. Default: 9083.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_channel": {
        "description": "The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: \"STABLE\" Possible values: [\"CANARY\", \"STABLE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_id": {
        "description": "The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),\nand hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between\n3 and 63 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the metastore service.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "Additional information about the current state of the metastore service, if available.",
        "description_kind": "plain",
        "type": "string"
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
      "tier": {
        "computed": true,
        "description": "The tier of the service. Possible values: [\"DEVELOPER\", \"ENTERPRISE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "The globally unique resource identifier of the metastore service.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time when the metastore service was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_config": {
        "block": {
          "attributes": {
            "kms_key": {
              "description": "The fully qualified customer provided Cloud KMS key name to use for customer data encryption.\nUse the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Information used to configure the Dataproc Metastore service to encrypt\ncustomer data at rest.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "hive_metastore_config": {
        "block": {
          "attributes": {
            "config_overrides": {
              "computed": true,
              "description": "A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).\nThe mappings override system defaults (some keys cannot be overridden)",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "endpoint_protocol": {
              "description": "The protocol to use for the metastore service endpoint. If unspecified, defaults to 'THRIFT'. Default value: \"THRIFT\" Possible values: [\"THRIFT\", \"GRPC\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description": "The Hive metastore schema version.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "auxiliary_versions": {
              "block": {
                "attributes": {
                  "config_overrides": {
                    "description": "A mapping of Hive metastore configuration key-value pairs to apply to the auxiliary Hive metastore (configured in hive-site.xml) in addition to the primary version's overrides.\nIf keys are present in both the auxiliary version's overrides and the primary version's overrides, the value from the auxiliary version's overrides takes precedence.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "version": {
                    "description": "The Hive metastore version of the auxiliary service. It must be less than the primary Hive metastore service's version.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A mapping of Hive metastore version to the auxiliary version configuration.\nWhen specified, a secondary Hive metastore service is created along with the primary service.\nAll auxiliary versions must be less than the service's primary version.\nThe key is the auxiliary service name and it must match the regular expression a-z?.\nThis means that the first character must be a lowercase letter, and all the following characters must be hyphens, lowercase letters, or digits, except the last character, which cannot be a hyphen.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "kerberos_config": {
              "block": {
                "attributes": {
                  "krb5_config_gcs_uri": {
                    "description": "A Cloud Storage URI that specifies the path to a krb5.conf file. It is of the form gs://{bucket_name}/path/to/krb5.conf, although the file does not need to be named krb5.conf explicitly.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "principal": {
                    "description": "A Kerberos principal that exists in the both the keytab the KDC to authenticate as. A typical principal is of the form \"primary/instance@REALM\", but there is no exact format.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "keytab": {
                    "block": {
                      "attributes": {
                        "cloud_secret": {
                          "description": "The relative resource name of a Secret Manager secret version, in the following form:\n\n\"projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}\".",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A Kerberos keytab file that can be used to authenticate a service principal with a Kerberos Key Distribution Center (KDC).",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Information used to configure the Hive metastore service as a service principal in a Kerberos realm.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration information specific to running Hive metastore software as the metastore service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_window": {
        "block": {
          "attributes": {
            "day_of_week": {
              "description": "The day of week, when the window starts. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hour_of_day": {
              "description": "The hour of day (0-23) when the window starts.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "The one hour maintenance window of the metastore service.\nThis specifies when the service can be restarted for maintenance purposes in UTC time.\nMaintenance window is not needed for services with the 'SPANNER' database type.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "metadata_integration": {
        "block": {
          "block_types": {
            "data_catalog_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Defines whether the metastore metadata should be synced to Data Catalog. The default value is to disable syncing metastore metadata to Data Catalog.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "The integration config for the Data Catalog service.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The setting that defines how metastore metadata should be integrated with external services and systems.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "block_types": {
            "consumers": {
              "block": {
                "attributes": {
                  "endpoint_uri": {
                    "computed": true,
                    "description": "The URI of the endpoint used to access the metastore service.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subnetwork": {
                    "description": "The subnetwork of the customer project from which an IP address is reserved and used as the Dataproc Metastore service's endpoint.\nIt is accessible to hosts in the subnet and to all hosts in a subnet in the same region and same network.\nThere must be at least one IP address available in the subnet's primary range. The subnet is specified in the following form:\n'projects/{projectNumber}/regions/{region_id}/subnetworks/{subnetwork_id}",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The consumer-side network configuration for the Dataproc Metastore instance.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The configuration specifying the network settings for the Dataproc Metastore service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "scaling_config": {
        "block": {
          "attributes": {
            "instance_size": {
              "description": "Metastore instance sizes. Possible values: [\"EXTRA_SMALL\", \"SMALL\", \"MEDIUM\", \"LARGE\", \"EXTRA_LARGE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scaling_factor": {
              "description": "Scaling factor, in increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Represents the scaling configuration of a metastore service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "scheduled_backup": {
        "block": {
          "attributes": {
            "backup_location": {
              "description": "A Cloud Storage URI of a folder, in the format gs://\u003cbucket_name\u003e/\u003cpath_inside_bucket\u003e. A sub-folder \u003cbackup_folder\u003e containing backup files will be stored below it.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cron_schedule": {
              "description": "The scheduled interval in Cron format, see https://en.wikipedia.org/wiki/Cron The default is empty: scheduled backup is not enabled. Must be specified to enable scheduled backups.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description": "Defines whether the scheduled backup is enabled. The default value is false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "time_zone": {
              "computed": true,
              "description": "Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones), e.g. America/Los_Angeles or Africa/Abidjan. If left unspecified, the default is UTC.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The configuration of scheduled backup for the metastore service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "telemetry_config": {
        "block": {
          "attributes": {
            "log_format": {
              "description": "The output format of the Dataproc Metastore service's logs. Default value: \"JSON\" Possible values: [\"LEGACY\", \"JSON\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.",
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

func GoogleDataprocMetastoreServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocMetastoreService), &result)
	return &result
}
