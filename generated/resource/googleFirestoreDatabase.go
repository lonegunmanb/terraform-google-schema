package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirestoreDatabase = `{
  "block": {
    "attributes": {
      "app_engine_integration_mode": {
        "computed": true,
        "description": "The App Engine integration mode to use for this database. Possible values: [\"ENABLED\", \"DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "concurrency_mode": {
        "computed": true,
        "description": "The concurrency control mode to use for this database. Possible values: [\"OPTIMISTIC\", \"PESSIMISTIC\", \"OPTIMISTIC_WITH_ENTITY_GROUPS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The timestamp at which this database was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_protection_state": {
        "computed": true,
        "description": "State of delete protection for the database. Possible values: [\"DELETE_PROTECTION_STATE_UNSPECIFIED\", \"DELETE_PROTECTION_ENABLED\", \"DELETE_PROTECTION_DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "earliest_version_time": {
        "computed": true,
        "description": "Output only. The earliest timestamp at which older versions of the data can be read from the database. See versionRetentionPeriod above; this field is populated with now - versionRetentionPeriod.\nThis value is continuously updated, and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Output only. This checksum is computed by the server based on the value of other fields,\nand may be sent on update and delete requests to ensure the client has an\nup-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_prefix": {
        "computed": true,
        "description": "Output only. The keyPrefix for this database.\nThis keyPrefix is used, in combination with the project id (\"~\") to construct the application id\nthat is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.\nThis value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).",
        "description_kind": "plain",
        "type": "string"
      },
      "location_id": {
        "description": "The location of the database. Available locations are listed at\nhttps://cloud.google.com/firestore/docs/locations.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The ID to use for the database, which will become the final\ncomponent of the database's resource name. This value should be 4-63\ncharacters. Valid characters are /[a-z][0-9]-/ with first character\na letter and the last a letter or a number. Must not be\nUUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.\n\"(default)\" database id is also valid.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "point_in_time_recovery_enablement": {
        "description": "Whether to enable the PITR feature on this database.\nIf 'POINT_IN_TIME_RECOVERY_ENABLED' is selected, reads are supported on selected versions of the data from within the past 7 days.\nversionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour\nand reads against 1-minute snapshots beyond 1 hour and within 7 days.\nIf 'POINT_IN_TIME_RECOVERY_DISABLED' is selected, reads are supported on any version of the data from within the past 1 hour. Default value: \"POINT_IN_TIME_RECOVERY_DISABLED\" Possible values: [\"POINT_IN_TIME_RECOVERY_ENABLED\", \"POINT_IN_TIME_RECOVERY_DISABLED\"]",
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
      "type": {
        "description": "The type of the database.\nSee https://cloud.google.com/datastore/docs/firestore-or-datastore\nfor information about how to choose. Possible values: [\"FIRESTORE_NATIVE\", \"DATASTORE_MODE\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. The system-generated UUID4 for this Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The timestamp at which this database was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_retention_period": {
        "computed": true,
        "description": "Output only. The period during which past versions of data are retained in the database.\nAny read or query can specify a readTime within this window, and will read the state of the database at that time.\nIf the PITR feature is enabled, the retention period is 7 days. Otherwise, the retention period is 1 hour.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleFirestoreDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirestoreDatabase), &result)
	return &result
}
