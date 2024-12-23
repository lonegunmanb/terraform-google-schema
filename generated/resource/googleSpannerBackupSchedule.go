package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerBackupSchedule = `{
  "block": {
    "attributes": {
      "database": {
        "description": "The database to create the backup schedule on.",
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
      "instance": {
        "description": "The instance to create the database on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A unique identifier for the backup schedule, which cannot be changed after\nthe backup schedule is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].",
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
      "retention_duration": {
        "description": "At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: '3.5s'.\nYou can set this to a value up to 366 days.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "full_backup_spec": {
        "block": {
          "description": "The schedule creates only full backups..",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "incremental_backup_spec": {
        "block": {
          "description": "The schedule creates incremental backup chains.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "block_types": {
            "cron_spec": {
              "block": {
                "attributes": {
                  "text": {
                    "description": "Textual representation of the crontab. User can customize the\nbackup frequency and the backup version time using the cron\nexpression. The version time must be in UTC timzeone.\nThe backup will contain an externally consistent copy of the\ndatabase at the version time. Allowed frequencies are 12 hour, 1 day,\n1 week and 1 month. Examples of valid cron specifications:\n  0 2/12 * * * : every 12 hours at (2, 14) hours past midnight in UTC.\n  0 2,14 * * * : every 12 hours at (2,14) hours past midnight in UTC.\n  0 2 * * *    : once a day at 2 past midnight in UTC.\n  0 2 * * 0    : once a week every Sunday at 2 past midnight in UTC.\n  0 2 8 * *    : once a month on 8th day at 2 past midnight in UTC.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Cron style schedule specification..",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines specifications of the backup schedule.",
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

func GoogleSpannerBackupScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSpannerBackupSchedule), &result)
	return &result
}
