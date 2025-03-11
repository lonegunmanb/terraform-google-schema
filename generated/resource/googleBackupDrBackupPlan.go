package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBackupDrBackupPlan = `{
  "block": {
    "attributes": {
      "backup_plan_id": {
        "description": "The ID of the backup plan",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backup_vault": {
        "description": "Backup vault where the backups gets stored using this Backup plan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backup_vault_service_account": {
        "computed": true,
        "description": "The Google Cloud Platform Service Account to be used by the BackupVault for taking backups.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "When the 'BackupPlan' was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The description allows for additional details about 'BackupPlan' and its use cases to be provided.",
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
      "location": {
        "description": "The location for the backup plan",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of backup plan resource created",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "description": "The resource type to which the 'BackupPlan' will be applied. Examples include, \"compute.googleapis.com/Instance\" and \"storage.googleapis.com/Bucket\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "When the 'BackupPlan' was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "backup_rules": {
        "block": {
          "attributes": {
            "backup_retention_days": {
              "description": "Configures the duration for which backup data will be kept. The value should be greater than or equal to minimum enforced retention of the backup vault.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "rule_id": {
              "description": "The unique ID of this 'BackupRule'. The 'rule_id' is unique per 'BackupPlan'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "standard_schedule": {
              "block": {
                "attributes": {
                  "days_of_month": {
                    "description": "Specifies days of months like 1, 5, or 14 on which jobs will run.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "days_of_week": {
                    "description": "Specifies days of week like MONDAY or TUESDAY, on which jobs will run. This is required for 'recurrence_type', 'WEEKLY' and is not applicable otherwise. Possible values: [\"DAY_OF_WEEK_UNSPECIFIED\", \"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "hourly_frequency": {
                    "description": "Specifies frequency for hourly backups. An hourly frequency of 2 means jobs will run every 2 hours from start time till end time defined.\nThis is required for 'recurrence_type', 'HOURLY' and is not applicable otherwise.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "months": {
                    "description": "Specifies values of months Possible values: [\"MONTH_UNSPECIFIED\", \"JANUARY\", \"FEBRUARY\", \"MARCH\", \"APRIL\", \"MAY\", \"JUNE\", \"JULY\", \"AUGUST\", \"SEPTEMBER\", \"OCTOBER\", \"NOVEMBER\", \"DECEMBER\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "recurrence_type": {
                    "description": "RecurrenceType enumerates the applicable periodicity for the schedule. Possible values: [\"HOURLY\", \"DAILY\", \"WEEKLY\", \"MONTHLY\", \"YEARLY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "time_zone": {
                    "description": "The time zone to be used when interpreting the schedule.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "backup_window": {
                    "block": {
                      "attributes": {
                        "end_hour_of_day": {
                          "description": "The hour of the day (1-24) when the window ends, for example, if the value of end hour of the day is 10, that means the backup window end time is 10:00.\nThe end hour of the day should be greater than the start",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "start_hour_of_day": {
                          "description": "The hour of the day (0-23) when the window starts, for example, if the value of the start hour of the day is 6, that means the backup window starts at 6:00.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "A BackupWindow defines the window of the day during which backup jobs will run. Jobs are queued at the beginning of the window and will be marked as\n'NOT_RUN' if they do not start by the end of the window.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "week_day_of_month": {
                    "block": {
                      "attributes": {
                        "day_of_week": {
                          "description": "Specifies the day of the week. Possible values: [\"DAY_OF_WEEK_UNSPECIFIED\", \"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "week_of_month": {
                          "description": "WeekOfMonth enumerates possible weeks in the month, e.g. the first, third, or last week of the month. Possible values: [\"WEEK_OF_MONTH_UNSPECIFIED\", \"FIRST\", \"SECOND\", \"THIRD\", \"FOURTH\", \"LAST\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies a week day of the month like FIRST SUNDAY or LAST MONDAY, on which jobs will run.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "StandardSchedule defines a schedule that runs within the confines of a defined window of days.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The backup rules for this 'BackupPlan'. There must be at least one 'BackupRule' message.",
          "description_kind": "plain"
        },
        "min_items": 1,
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

func GoogleBackupDrBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrBackupPlan), &result)
	return &result
}
