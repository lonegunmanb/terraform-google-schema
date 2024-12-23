package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetappVolumeReplication = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Create time of the active directory. A timestamp in RFC3339 UTC \"Zulu\" format. Examples: \"2023-06-22T09:13:01.617Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_destination_volume": {
        "description": "A destination volume is created as part of replication creation. The destination volume will not became\nunder Terraform management unless you import it manually. If you delete the replication, this volume\nwill remain.\nSetting this parameter to true will delete the *current* destination volume when destroying the\nreplication. If you reversed the replication direction, this will be your former source volume!\nFor production use, it is recommended to keep this parameter false to avoid accidental volume\ndeletion. Handle with care. Default is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "An description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_volume": {
        "computed": true,
        "description": "Full resource name of destination volume with format: 'projects/{{project}}/locations/{{location}}/volumes/{{volumeId}}'",
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
      "force_stopping": {
        "description": "Only replications with mirror_state=MIRRORED can be stopped. A replication in mirror_state=TRANSFERRING\ncurrently receives an update and stopping the update might be undesirable. Set this parameter to true\nto stop anyway. All data transferred to the destination will be discarded and content of destination\nvolume will remain at the state of the last successful update. Default is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "healthy": {
        "computed": true,
        "description": "Condition of the relationship. Can be one of the following:\n  - true: The replication relationship is healthy. It has not missed the most recent scheduled transfer.\n  - false: The replication relationship is not healthy. It has missed the most recent scheduled transfer.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels as key value pairs. Example: '{ \"owner\": \"Bob\", \"department\": \"finance\", \"purpose\": \"testing\" }'\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Name of region for this resource. The resource needs to be created in the region of the destination volume.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mirror_state": {
        "computed": true,
        "description": "Indicates the state of the mirror between source and destination volumes. Depending on the amount of data\nin your source volume, PREPARING phase can take hours or days. mirrorState = MIRRORED indicates your baseline\ntransfer ended and destination volume became accessible read-only. TRANSFERRING means a MIRRORED volume\ncurrently receives an update. Updated every 5 minutes.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the replication. Needs to be unique per location.",
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
      "replication_enabled": {
        "description": "Set to false to stop/break the mirror. Stopping the mirror makes the destination volume read-write\nand act independently from the source volume.\nSet to true to enable/resume the mirror. WARNING: Resuming a mirror overwrites any changes\ndone to the destination volume with the content of the source volume.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "replication_schedule": {
        "description": "Specifies the replication interval. Possible values: [\"EVERY_10_MINUTES\", \"HOURLY\", \"DAILY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role": {
        "computed": true,
        "description": "Reverting a replication can swap source and destination volume roles. This field indicates if the 'location' hosts\nthe source or destination volume. For resume and revert and resume operations it is critical to understand\nwhich volume is the source volume, since it will overwrite changes done to the destination volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_volume": {
        "computed": true,
        "description": "Full resource name of source volume with format: 'projects/{{project}}/locations/{{location}}/volumes/{{volumeId}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Indicates the state of replication resource. State of the mirror itself is indicated in mirrorState.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_details": {
        "computed": true,
        "description": "State details of the replication resource.",
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
      "transfer_stats": {
        "computed": true,
        "description": "Replication transfer statistics. All statistics are updated every 5 minutes.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "lag_duration": "string",
              "last_transfer_bytes": "string",
              "last_transfer_duration": "string",
              "last_transfer_end_time": "string",
              "last_transfer_error": "string",
              "total_transfer_duration": "string",
              "transfer_bytes": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "volume_name": {
        "description": "The name of the existing source volume.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "wait_for_mirror": {
        "description": "Replication resource state is independent of mirror_state. With enough data, it can take many hours\nfor mirror_state to reach MIRRORED. If you want Terraform to wait for the mirror to finish on\ncreate/stop/resume operations, set this parameter to true. Default is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "destination_volume_parameters": {
        "block": {
          "attributes": {
            "description": {
              "description": "Description for the destination volume.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_name": {
              "computed": true,
              "description": "Share name for destination volume. If not specified, name of source volume's share name will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_pool": {
              "description": "Name of an existing storage pool for the destination volume with format: 'projects/{{project}}/locations/{{location}}/storagePools/{{poolId}}'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "volume_id": {
              "computed": true,
              "description": "Name for the destination volume to be created. If not specified, the name of the source volume will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Destination volume parameters.",
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

func GoogleNetappVolumeReplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetappVolumeReplication), &result)
	return &result
}
