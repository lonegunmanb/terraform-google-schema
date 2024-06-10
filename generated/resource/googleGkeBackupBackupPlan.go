package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeBackupBackupPlan = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "The source cluster from which Backups will be created via this BackupPlan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deactivated": {
        "computed": true,
        "description": "This flag indicates whether this BackupPlan has been deactivated.\nSetting this field to True locks the BackupPlan such that no further updates will be allowed\n(except deletes), including the deactivated field itself. It also prevents any new Backups\nfrom being created via this BackupPlan (including scheduled Backups).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "User specified descriptive string for this BackupPlan.",
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
        "description": "etag is used for optimistic concurrency control as a way to help prevent simultaneous\nupdates of a backup plan from overwriting each other. It is strongly suggested that\nsystems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates\nin order to avoid race conditions: An etag is returned in the response to backupPlans.get,\nand systems are expected to put that etag in the request to backupPlans.patch or\nbackupPlans.delete to ensure that their change will be applied to the same version of the resource.",
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
        "description": "Description: A set of custom labels supplied by the user.\nA list of key-\u003evalue pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The region of the Backup Plan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The full name of the BackupPlan Resource.",
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
      "protected_pod_count": {
        "computed": true,
        "description": "The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.",
        "description_kind": "plain",
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "The State of the BackupPlan.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_reason": {
        "computed": true,
        "description": "Detailed description of why BackupPlan is in its current state.",
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
      "uid": {
        "computed": true,
        "description": "Server generated, unique identifier of UUID format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "backup_config": {
        "block": {
          "attributes": {
            "all_namespaces": {
              "description": "If True, include all namespaced resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_secrets": {
              "computed": true,
              "description": "This flag specifies whether Kubernetes Secret resources should be included\nwhen they fall into the scope of Backups.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_volume_data": {
              "computed": true,
              "description": "This flag specifies whether volume data should be backed up when PVCs are\nincluded in the scope of a Backup.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "permissive_mode": {
              "description": "This flag specifies whether Backups will not fail when\nBackup for GKE detects Kubernetes configuration that is\nnon-standard or requires additional setup to restore.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "encryption_key": {
              "block": {
                "attributes": {
                  "gcp_kms_encryption_key": {
                    "description": "Google Cloud KMS encryption key. Format: projects/*/locations/*/keyRings/*/cryptoKeys/*",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "This defines a customer managed encryption key that will be used to encrypt the \"config\"\nportion (the Kubernetes resources) of Backups created via this plan.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "selected_applications": {
              "block": {
                "block_types": {
                  "namespaced_names": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "The name of a Kubernetes Resource.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "namespace": {
                          "description": "The namespace of a Kubernetes Resource.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of namespaced Kubernetes resources.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of namespaced Kubernetes Resources.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "selected_namespaces": {
              "block": {
                "attributes": {
                  "namespaces": {
                    "description": "A list of Kubernetes Namespaces.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "If set, include just the resources in the listed namespaces.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines the configuration of Backups created via this BackupPlan.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "backup_schedule": {
        "block": {
          "attributes": {
            "cron_schedule": {
              "description": "A standard cron string that defines a repeating schedule for\ncreating Backups via this BackupPlan.\nThis is mutually exclusive with the rpoConfig field since at most one\nschedule can be defined for a BackupPlan.\nIf this is defined, then backupRetainDays must also be defined.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "paused": {
              "computed": true,
              "description": "This flag denotes whether automatic Backup creation is paused for this BackupPlan.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "rpo_config": {
              "block": {
                "attributes": {
                  "target_rpo_minutes": {
                    "description": "Defines the target RPO for the BackupPlan in minutes, which means the target\nmaximum data loss in time that is acceptable for this BackupPlan. This must be\nat least 60, i.e., 1 hour, and at most 86400, i.e., 60 days.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "exclusion_windows": {
                    "block": {
                      "attributes": {
                        "daily": {
                          "description": "The exclusion window occurs every day if set to \"True\".\nSpecifying this field to \"False\" is an error.\nOnly one of singleOccurrenceDate, daily and daysOfWeek may be set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "duration": {
                          "description": "Specifies duration of the window in seconds with up to nine fractional digits,\nterminated by 's'. Example: \"3.5s\". Restrictions for duration based on the\nrecurrence type to allow some time for backup to happen:\n  - single_occurrence_date:  no restriction\n  - daily window: duration \u003c 24 hours\n  - weekly window:\n    - days of week includes all seven days of a week: duration \u003c 24 hours\n    - all other weekly window: duration \u003c 168 hours (i.e., 24 * 7 hours)",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "days_of_week": {
                          "block": {
                            "attributes": {
                              "days_of_week": {
                                "description": "A list of days of week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "The exclusion window occurs on these days of each week in UTC.\nOnly one of singleOccurrenceDate, daily and daysOfWeek may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "single_occurrence_date": {
                          "block": {
                            "attributes": {
                              "day": {
                                "description": "Day of a month.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "month": {
                                "description": "Month of a year.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "year": {
                                "description": "Year of the date.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "No recurrence. The exclusion window occurs only once and on this date in UTC.\nOnly one of singleOccurrenceDate, daily and daysOfWeek may be set.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "start_time": {
                          "block": {
                            "attributes": {
                              "hours": {
                                "description": "Hours of day in 24 hour format.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "minutes": {
                                "description": "Minutes of hour of day.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "nanos": {
                                "description": "Fractions of seconds in nanoseconds.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "seconds": {
                                "description": "Seconds of minutes of the time.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Specifies the start time of the window using time of the day in UTC.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "User specified time windows during which backup can NOT happen for this BackupPlan.\nBackups should start and finish outside of any given exclusion window. Note: backup\njobs will be scheduled to start and finish outside the duration of the window as\nmuch as possible, but running jobs will not get canceled when it runs into the window.\nAll the time and date values in exclusionWindows entry in the API are in UTC. We\nonly allow \u003c=1 recurrence (daily or weekly) exclusion window for a BackupPlan while no\nrestriction on number of single occurrence windows.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Defines the RPO schedule configuration for this BackupPlan. This is mutually\nexclusive with the cronSchedule field since at most one schedule can be defined\nfor a BackupPLan. If this is defined, then backupRetainDays must also be defined.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines a schedule for automatic Backup creation via this BackupPlan.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_policy": {
        "block": {
          "attributes": {
            "backup_delete_lock_days": {
              "computed": true,
              "description": "Minimum age for a Backup created via this BackupPlan (in days).\nMust be an integer value between 0-90 (inclusive).\nA Backup created under this BackupPlan will not be deletable\nuntil it reaches Backup's (create time + backup_delete_lock_days).\nUpdating this field of a BackupPlan does not affect existing Backups.\nBackups created after a successful update will inherit this new value.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "backup_retain_days": {
              "computed": true,
              "description": "The default maximum age of a Backup created via this BackupPlan.\nThis field MUST be an integer value \u003e= 0 and \u003c= 365. If specified,\na Backup created under this BackupPlan will be automatically deleted\nafter its age reaches (createTime + backupRetainDays).\nIf not specified, Backups created under this BackupPlan will NOT be\nsubject to automatic deletion. Updating this field does NOT affect\nexisting Backups under it. Backups created AFTER a successful update\nwill automatically pick up the new value.\nNOTE: backupRetainDays must be \u003e= backupDeleteLockDays.\nIf cronSchedule is defined, then this must be \u003c= 360 * the creation interval.\nIf rpo_config is defined, then this must be\n\u003c= 360 * targetRpoMinutes/(1440minutes/day)",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "locked": {
              "computed": true,
              "description": "This flag denotes whether the retention policy of this BackupPlan is locked.\nIf set to True, no further update is allowed on this policy, including\nthe locked field itself.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "RetentionPolicy governs lifecycle of Backups created under this plan.",
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

func GoogleGkeBackupBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeBackupBackupPlan), &result)
	return &result
}
