package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetappVolume = `{
  "block": {
    "attributes": {
      "active_directory": {
        "computed": true,
        "description": "Reports the resource name of the Active Directory policy being used. Inherited from storage pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_gib": {
        "description": "Capacity of the volume (in GiB).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Create time of the volume. A timestamp in RFC3339 UTC \"Zulu\" format. Examples: \"2023-06-22T09:13:01.617Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_policy": {
        "description": "Policy to determine if the volume should be deleted forcefully.\nVolumes may have nested snapshot resources. Deleting such a volume will fail.\nSetting this parameter to FORCE will delete volumes including nested snapshots.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
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
      "encryption_type": {
        "computed": true,
        "description": "Reports the data-at-rest encryption type of the volume. Inherited from storage pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "has_replication": {
        "computed": true,
        "description": "Indicates whether the volume is part of a volume replication relationship.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kerberos_enabled": {
        "description": "Flag indicating if the volume is a kerberos volume or not, export policy rules control kerberos security modes (krb5, krb5i, krb5p).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_config": {
        "computed": true,
        "description": "Reports the CMEK policy resurce name being used for volume encryption. Inherited from storage pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels as key value pairs. Example: '{ \"owner\": \"Bob\", \"department\": \"finance\", \"purpose\": \"testing\" }'.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "ldap_enabled": {
        "computed": true,
        "description": "Flag indicating if the volume is NFS LDAP enabled or not. Inherited from storage pool.",
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "description": "Name of the pool location. Usually a region name, expect for some STANDARD service level pools which require a zone name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mount_options": {
        "computed": true,
        "description": "Reports mount instructions for this volume.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "export": "string",
              "export_full": "string",
              "instructions": "string",
              "protocol": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "The name of the volume. Needs to be unique per location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "VPC network name with format: 'projects/{{project}}/global/networks/{{network}}'. Inherited from storage pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocols": {
        "description": "The protocol of the volume. Allowed combinations are '['NFSV3']', '['NFSV4']', '['SMB']', '['NFSV3', 'NFSV4']', '['SMB', 'NFSV3']' and '['SMB', 'NFSV4']'. Possible values: [\"NFSV3\", \"NFSV4\", \"SMB\"]",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "psa_range": {
        "computed": true,
        "description": "Name of the Private Service Access allocated range. Inherited from storage pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "restricted_actions": {
        "description": "List of actions that are restricted on this volume. Possible values: [\"DELETE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "security_style": {
        "computed": true,
        "description": "Security Style of the Volume. Use UNIX to use UNIX or NFSV4 ACLs for file permissions.\nUse NTFS to use NTFS ACLs for file permissions. Can only be set for volumes which use SMB together with NFS as protocol. Possible values: [\"NTFS\", \"UNIX\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_level": {
        "computed": true,
        "description": "Service level of the volume. Inherited from storage pool. Supported values are : PREMIUM, EXTERME, STANDARD, FLEX.",
        "description_kind": "plain",
        "type": "string"
      },
      "share_name": {
        "description": "Share name (SMB) or export path (NFS) of the volume. Needs to be unique per location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "smb_settings": {
        "description": "Settings for volumes with SMB access. Possible values: [\"ENCRYPT_DATA\", \"BROWSABLE\", \"CHANGE_NOTIFY\", \"NON_BROWSABLE\", \"OPLOCKS\", \"SHOW_SNAPSHOT\", \"SHOW_PREVIOUS_VERSIONS\", \"ACCESS_BASED_ENUMERATION\", \"CONTINUOUSLY_AVAILABLE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "snapshot_directory": {
        "description": "If enabled, a NFS volume will contain a read-only .snapshot directory which provides access to each of the volume's snapshots. Will enable \"Previous Versions\" support for SMB.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "State of the volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_details": {
        "computed": true,
        "description": "State details of the volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_pool": {
        "description": "Name of the storage pool to create the volume in. Pool needs enough spare capacity to accomodate the volume.",
        "description_kind": "plain",
        "required": true,
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
      "unix_permissions": {
        "computed": true,
        "description": "Unix permission the mount point will be created with. Default is 0770. Applicable for UNIX security style volumes only.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "used_gib": {
        "computed": true,
        "description": "Used capacity of the volume (in GiB). This is computed periodically and it does not represent the realtime usage.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "export_policy": {
        "block": {
          "block_types": {
            "rules": {
              "block": {
                "attributes": {
                  "access_type": {
                    "description": "Defines the access type for clients matching the 'allowedClients' specification. Possible values: [\"READ_ONLY\", \"READ_WRITE\", \"READ_NONE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "allowed_clients": {
                    "description": "Defines the client ingress specification (allowed clients) as a comma seperated list with IPv4 CIDRs or IPv4 host addresses.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "has_root_access": {
                    "description": "If enabled, the root user (UID = 0) of the specified clients doesn't get mapped to nobody (UID = 65534). This is also known as no_root_squash.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kerberos5_read_only": {
                    "description": "If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'authentication' kerberos security mode.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "kerberos5_read_write": {
                    "description": "If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'authentication' kerberos security mode. The 'kerberos5ReadOnly' value is ignored if this is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "kerberos5i_read_only": {
                    "description": "If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'integrity' kerberos security mode.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "kerberos5i_read_write": {
                    "description": "If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'integrity' kerberos security mode. The 'kerberos5iReadOnly' value is ignored if this is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "kerberos5p_read_only": {
                    "description": "If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'privacy' kerberos security mode.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "kerberos5p_read_write": {
                    "description": "If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'privacy' kerberos security mode. The 'kerberos5pReadOnly' value is ignored if this is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "nfsv3": {
                    "description": "Enable to apply the export rule to NFSV3 clients.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "nfsv4": {
                    "description": "Enable to apply the export rule to NFSV4.1 clients.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Export rules (up to 5) control NFS volume access.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Export policy of the volume for NFSV3 and/or NFSV4.1 access.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "restore_parameters": {
        "block": {
          "attributes": {
            "source_backup": {
              "description": "Full name of the snapshot to use for creating this volume.\n'source_snapshot' and 'source_backup' cannot be used simultaneously.\nFormat: 'projects/{{project}}/locations/{{location}}/backupVaults/{{backupVaultId}}/backups/{{backup}}'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_snapshot": {
              "description": "Full name of the snapshot to use for creating this volume.\n'source_snapshot' and 'source_backup' cannot be used simultaneously.\nFormat: 'projects/{{project}}/locations/{{location}}/volumes/{{volume}}/snapshots/{{snapshot}}'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Used to create this volume from a snapshot (= cloning) or an backup.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "snapshot_policy": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Enables automated snapshot creation according to defined schedule. Default is false.\nTo disable automatic snapshot creation you have to remove the whole snapshot_policy block.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "daily_schedule": {
              "block": {
                "attributes": {
                  "hour": {
                    "description": "Set the hour to create the snapshot (0-23), defaults to midnight (0).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minute": {
                    "description": "Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "snapshots_to_keep": {
                    "description": "The maximum number of snapshots to keep for the daily schedule.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Daily schedule policy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "hourly_schedule": {
              "block": {
                "attributes": {
                  "minute": {
                    "description": "Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "snapshots_to_keep": {
                    "description": "The maximum number of snapshots to keep for the hourly schedule.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Hourly schedule policy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "monthly_schedule": {
              "block": {
                "attributes": {
                  "days_of_month": {
                    "description": "Set the day or days of the month to make a snapshot (1-31). Accepts a comma separated number of days. Defaults to '1'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "hour": {
                    "description": "Set the hour to create the snapshot (0-23), defaults to midnight (0).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minute": {
                    "description": "Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "snapshots_to_keep": {
                    "description": "The maximum number of snapshots to keep for the monthly schedule",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Monthly schedule policy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "weekly_schedule": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Set the day or days of the week to make a snapshot. Accepts a comma separated days of the week. Defaults to 'Sunday'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "hour": {
                    "description": "Set the hour to create the snapshot (0-23), defaults to midnight (0).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minute": {
                    "description": "Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "snapshots_to_keep": {
                    "description": "The maximum number of snapshots to keep for the weekly schedule.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Weekly schedule policy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Snapshot policy defines the schedule for automatic snapshot creation.\nTo disable automatic snapshot creation you have to remove the whole snapshot_policy block.",
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

func GoogleNetappVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetappVolume), &result)
	return &result
}
