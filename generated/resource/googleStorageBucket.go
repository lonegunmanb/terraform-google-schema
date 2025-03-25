package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucket = `{
  "block": {
    "attributes": {
      "default_event_based_hold": {
        "description": "Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.",
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
      "enable_object_retention": {
        "description": "Enables each object in the bucket to have its own retention policy, which prevents deletion until stored for a specific length of time.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_destroy": {
        "description": "When deleting a bucket, this boolean option will delete all contained objects. If you try to delete a bucket that contains objects, Terraform will fail that run.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs to assign to the bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The Google Cloud Storage location or region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the bucket.",
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
      "project_number": {
        "computed": true,
        "description": "The project number of the project in which the resource belongs.",
        "description_kind": "plain",
        "type": "number"
      },
      "public_access_prevention": {
        "computed": true,
        "description": "Prevents public access to a bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requester_pays": {
        "description": "Enables Requester Pays on a storage bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "rpo": {
        "computed": true,
        "description": "Specifies the RPO setting of bucket. If set 'ASYNC_TURBO', The Turbo Replication will be enabled for the dual-region bucket. Value 'DEFAULT' will set RPO setting to default. Turbo Replication is only for buckets in dual-regions.See the docs for more details.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_class": {
        "description": "The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.",
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
      },
      "uniform_bucket_level_access": {
        "computed": true,
        "description": "Enables uniform bucket-level access on a bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "url": {
        "computed": true,
        "description": "The base URL of the bucket, in the format gs://\u003cbucket-name\u003e.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "autoclass": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "While set to true, autoclass automatically transitions objects in your bucket to appropriate storage classes based on each object's access pattern.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "terminal_storage_class": {
              "computed": true,
              "description": "The storage class that objects in the bucket eventually transition to if they are not read for a certain length of time. Supported values include: NEARLINE, ARCHIVE.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The bucket's autoclass configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cors": {
        "block": {
          "attributes": {
            "max_age_seconds": {
              "description": "The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "method": {
              "description": "The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: \"*\" is permitted in the list of methods, and means \"any method\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "origin": {
              "description": "The list of Origins eligible to receive CORS response headers. Note: \"*\" is permitted in the list of origins, and means \"any Origin\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "response_header": {
              "description": "The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "The bucket's Cross-Origin Resource Sharing (CORS) configuration.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "custom_placement_config": {
        "block": {
          "attributes": {
            "data_locations": {
              "description": "The list of individual regions that comprise a dual-region bucket. See the docs for a list of acceptable regions. Note: If any of the data_locations changes, it will recreate the bucket.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "encryption": {
        "block": {
          "attributes": {
            "default_kms_key_name": {
              "description": "A Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified. You must pay attention to whether the crypto key is available in the location that this bucket is created in. See the docs for more details.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The bucket's encryption configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "hierarchical_namespace": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Set this field true to organize bucket with logical file system structure.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "The bucket's HNS configuration, which defines bucket can organize folders in logical file system structure.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "lifecycle_rule": {
        "block": {
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "storage_class": {
                    "description": "The target Storage Class of objects affected by this Lifecycle Rule. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "The type of the action of this Lifecycle Rule. Supported values include: Delete, SetStorageClass and AbortIncompleteMultipartUpload.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The Lifecycle Rule's action configuration. A single block of this type is supported.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "set"
            },
            "condition": {
              "block": {
                "attributes": {
                  "age": {
                    "description": "Minimum age of an object in days to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "created_before": {
                    "description": "Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_time_before": {
                    "description": "Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days_since_custom_time": {
                    "description": "Number of days elapsed since the user-specified timestamp set on an object.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "days_since_noncurrent_time": {
                    "description": "Number of days elapsed since the noncurrent timestamp of an object. This\n\t\t\t\t\t\t\t\t\t\tcondition is relevant only for versioned objects.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "matches_prefix": {
                    "description": "One or more matching name prefixes to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "matches_storage_class": {
                    "description": "Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "matches_suffix": {
                    "description": "One or more matching name suffixes to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "noncurrent_time_before": {
                    "description": "Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "num_newer_versions": {
                    "description": "Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "send_age_if_zero": {
                    "description": "While set true, age value will be sent in the request even for zero value of the field. This field is only useful for setting 0 value to the age field. It can be used alone or together with age.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "send_days_since_custom_time_if_zero": {
                    "description": "While set true, days_since_custom_time value will be sent in the request even for zero value of the field. This field is only useful for setting 0 value to the days_since_custom_time field. It can be used alone or together with days_since_custom_time.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "send_days_since_noncurrent_time_if_zero": {
                    "description": "While set true, days_since_noncurrent_time value will be sent in the request even for zero value of the field. This field is only useful for setting 0 value to the days_since_noncurrent_time field. It can be used alone or together with days_since_noncurrent_time.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "send_num_newer_versions_if_zero": {
                    "description": "While set true, num_newer_versions value will be sent in the request even for zero value of the field. This field is only useful for setting 0 value to the num_newer_versions field. It can be used alone or together with num_newer_versions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "with_state": {
                    "computed": true,
                    "description": "Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: \"LIVE\", \"ARCHIVED\", \"ANY\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The Lifecycle Rule's condition configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description": "The bucket's Lifecycle Rules configuration.",
          "description_kind": "plain"
        },
        "max_items": 100,
        "nesting_mode": "list"
      },
      "logging": {
        "block": {
          "attributes": {
            "log_bucket": {
              "description": "The bucket that will receive log objects.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_object_prefix": {
              "computed": true,
              "description": "The object prefix for log objects. If it's not provided, by default Google Cloud Storage sets this to this bucket's name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The bucket's Access \u0026 Storage Logs configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_policy": {
        "block": {
          "attributes": {
            "is_locked": {
              "description": "If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retention_period": {
              "description": "The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 3,155,760,000 seconds.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Configuration of the bucket's data retention policy for how long objects in the bucket should be retained.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "soft_delete_policy": {
        "block": {
          "attributes": {
            "effective_time": {
              "computed": true,
              "description": "Server-determined value that indicates the time from which the policy, or one with a greater retention, was effective. This value is in RFC 3339 format.",
              "description_kind": "plain",
              "type": "string"
            },
            "retention_duration_seconds": {
              "description": "The duration in seconds that soft-deleted objects in the bucket will be retained and cannot be permanently deleted. Default value is 604800.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "The bucket's soft delete policy, which defines the period of time that soft-deleted objects will be retained, and cannot be permanently deleted. If it is not provided, by default Google Cloud Storage sets this to default soft delete policy",
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
            "read": {
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
      },
      "versioning": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "While set to true, versioning is fully enabled for this bucket.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "The bucket's Versioning configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "website": {
        "block": {
          "attributes": {
            "main_page_suffix": {
              "description": "Behaves as the bucket's directory index where missing objects are treated as potential directories.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "not_found_page": {
              "description": "The custom object to return when a requested resource is not found.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration if the bucket acts as a website.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 3
}`

func GoogleStorageBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBucket), &result)
	return &result
}
