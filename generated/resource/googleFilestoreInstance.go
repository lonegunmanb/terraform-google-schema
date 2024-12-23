package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFilestoreInstance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "description": "Indicates whether the instance is protected against deletion.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deletion_protection_reason": {
        "description": "The reason for enabling deletion protection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "A description of the instance.",
        "description_kind": "plain",
        "optional": true,
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
      "etag": {
        "computed": true,
        "description": "Server-specified ETag for the instance resource to prevent\nsimultaneous updates from overwriting each other.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "description": "KMS key name used for data encryption.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Resource labels to represent user-provided metadata.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "computed": true,
        "description": "The name of the location of the instance. This can be a region for ENTERPRISE tier instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "description": "Either NFSv3, for using NFS version 3 as file sharing protocol,\nor NFSv4.1, for using NFS version 4.1 as file sharing protocol.\nNFSv4.1 can be used with HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE.\nThe default is NFSv3. Default value: \"NFS_V3\" Possible values: [\"NFS_V3\", \"NFS_V4_1\"]",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The service tier of the instance.\nPossible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone": {
        "computed": true,
        "deprecated": true,
        "description": "The name of the Filestore zone of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "file_shares": {
        "block": {
          "attributes": {
            "capacity_gb": {
              "description": "File share capacity in GiB. This must be at least 1024 GiB\nfor the standard tier, or 2560 GiB for the premium tier.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "name": {
              "description": "The name of the fileshare (16 characters or less)",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_backup": {
              "description": "The resource name of the backup, in the format\nprojects/{projectId}/locations/{locationId}/backups/{backupId},\nthat this file share has been restored from.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "nfs_export_options": {
              "block": {
                "attributes": {
                  "access_mode": {
                    "description": "Either READ_ONLY, for allowing only read requests on the exported directory,\nor READ_WRITE, for allowing both read and write requests. The default is READ_WRITE. Default value: \"READ_WRITE\" Possible values: [\"READ_ONLY\", \"READ_WRITE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "anon_gid": {
                    "description": "An integer representing the anonymous group id with a default value of 65534.\nAnon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned\nif this field is specified for other squashMode settings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "anon_uid": {
                    "description": "An integer representing the anonymous user id with a default value of 65534.\nAnon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned\nif this field is specified for other squashMode settings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ip_ranges": {
                    "description": "List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.\nOverlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.\nThe limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "squash_mode": {
                    "description": "Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,\nfor not allowing root access. The default is NO_ROOT_SQUASH. Default value: \"NO_ROOT_SQUASH\" Possible values: [\"NO_ROOT_SQUASH\", \"ROOT_SQUASH\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Nfs Export Options. There is a limit of 10 export options per file share.",
                "description_kind": "plain"
              },
              "max_items": 10,
              "nesting_mode": "list"
            }
          },
          "description": "File system shares on the instance. For this version, only a\nsingle file share is supported.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "networks": {
        "block": {
          "attributes": {
            "connect_mode": {
              "description": "The network connect mode of the Filestore instance.\nIf not provided, the connect mode defaults to\nDIRECT_PEERING. Default value: \"DIRECT_PEERING\" Possible values: [\"DIRECT_PEERING\", \"PRIVATE_SERVICE_ACCESS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_addresses": {
              "computed": true,
              "description": "A list of IPv4 or IPv6 addresses.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "modes": {
              "description": "IP versions for which the instance has\nIP addresses assigned. Possible values: [\"ADDRESS_MODE_UNSPECIFIED\", \"MODE_IPV4\", \"MODE_IPV6\"]",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "network": {
              "description": "The name of the GCE VPC network to which the\ninstance is connected.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "reserved_ip_range": {
              "computed": true,
              "description": "A /29 CIDR block that identifies the range of IP\naddresses reserved for this instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "VPC networks to which the instance is connected. For this version,\nonly a single network is supported.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "performance_config": {
        "block": {
          "block_types": {
            "fixed_iops": {
              "block": {
                "attributes": {
                  "max_iops": {
                    "description": "The number of IOPS to provision for the instance.\nmax_iops must be in multiple of 1000.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The instance will have a fixed provisioned IOPS value,\nwhich will remain constant regardless of instance\ncapacity.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "iops_per_tb": {
              "block": {
                "attributes": {
                  "max_iops_per_tb": {
                    "description": "The instance max IOPS will be calculated by multiplying\nthe capacity of the instance (TB) by max_iops_per_tb,\nand rounding to the nearest 1000. The instance max IOPS\nwill be changed dynamically based on the instance\ncapacity.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The instance provisioned IOPS will change dynamically\nbased on the capacity of the instance.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Performance configuration for the instance. If not provided,\nthe default performance settings will be used.",
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
  "version": 1
}`

func GoogleFilestoreInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFilestoreInstance), &result)
	return &result
}
