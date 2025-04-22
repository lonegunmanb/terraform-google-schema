package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBatchOperationsJob = `{
  "block": {
    "attributes": {
      "complete_time": {
        "computed": true,
        "description": "The time that the job was completed.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp at which this storage batch operation was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_protection": {
        "description": "If set to 'true', the storage batch operation job will not be deleted and new job will be created.",
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
      "job_id": {
        "description": "The ID of the job.",
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
      "schedule_time": {
        "computed": true,
        "description": "The time that the job was scheduled.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the job.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp at which this storage batch operation was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "bucket_list": {
        "block": {
          "block_types": {
            "buckets": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description": "Bucket name for the objects to be transformed.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "manifest": {
                    "block": {
                      "attributes": {
                        "manifest_location": {
                          "description": "Specifies objects in a manifest file.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "contain the manifest source file that is a CSV file in a Google Cloud Storage bucket.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "prefix_list": {
                    "block": {
                      "attributes": {
                        "included_object_prefixes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Specifies objects matching a prefix set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of buckets and their objects to be transformed.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "List of buckets and their objects to be transformed. Currently, only one bucket configuration is supported. If multiple buckets are specified, an error will be returned",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "delete_object": {
        "block": {
          "attributes": {
            "permanent_object_deletion_enabled": {
              "description": "enable flag to permanently delete object and all object versions if versioning is enabled on bucket.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "allows batch operations to delete objects in bucket",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "put_metadata": {
        "block": {
          "attributes": {
            "cache_control": {
              "description": "Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_disposition": {
              "description": "Content-Disposition of the object data.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_encoding": {
              "description": "Content Encoding of the object data.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_language": {
              "description": "Content-Language of the object data.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_type": {
              "description": "Content-Type of the object data.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "custom_metadata": {
              "description": "User-provided metadata, in key/value pairs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "custom_time": {
              "description": "Updates the objects fixed custom time metadata.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "allows batch operations to update metadata for objects in bucket",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "put_object_hold": {
        "block": {
          "attributes": {
            "event_based_hold": {
              "description": "set/unset to update event based hold for objects.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "temporary_hold": {
              "description": "set/unset to update temporary based hold for objects.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "allows to update temporary hold or eventBased hold for objects in bucket.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rewrite_object": {
        "block": {
          "attributes": {
            "kms_key": {
              "description": "valid kms key",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "allows to update encryption key for objects in bucket.",
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

func GoogleStorageBatchOperationsJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBatchOperationsJob), &result)
	return &result
}
