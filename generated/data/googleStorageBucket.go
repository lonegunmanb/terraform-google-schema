package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucket = `{
  "block": {
    "attributes": {
      "autoclass": {
        "computed": true,
        "description": "The bucket's autoclass configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "terminal_storage_class": "string"
            }
          ]
        ]
      },
      "cors": {
        "computed": true,
        "description": "The bucket's Cross-Origin Resource Sharing (CORS) configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "max_age_seconds": "number",
              "method": [
                "list",
                "string"
              ],
              "origin": [
                "list",
                "string"
              ],
              "response_header": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "custom_placement_config": {
        "computed": true,
        "description": "The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_locations": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "default_event_based_hold": {
        "computed": true,
        "description": "Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "Enables each object in the bucket to have its own retention policy, which prevents deletion until stored for a specific length of time.",
        "description_kind": "plain",
        "type": "bool"
      },
      "encryption": {
        "computed": true,
        "description": "The bucket's encryption configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "default_kms_key_name": "string"
            }
          ]
        ]
      },
      "force_destroy": {
        "computed": true,
        "description": "When deleting a bucket, this boolean option will delete all contained objects, or anywhereCaches (if any). If you try to delete a bucket that contains objects or anywhereCaches, Terraform will fail that run, deleting anywhereCaches may take 80 minutes to complete.",
        "description_kind": "plain",
        "type": "bool"
      },
      "hierarchical_namespace": {
        "computed": true,
        "description": "The bucket's HNS configuration, which defines bucket can organize folders in logical file system structure.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_filter": {
        "computed": true,
        "description": "The bucket IP filtering configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_all_service_agent_access": "bool",
              "allow_cross_org_vpcs": "bool",
              "mode": "string",
              "public_network_source": [
                "list",
                [
                  "object",
                  {
                    "allowed_ip_cidr_ranges": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "vpc_network_sources": [
                "list",
                [
                  "object",
                  {
                    "allowed_ip_cidr_ranges": [
                      "list",
                      "string"
                    ],
                    "network": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "labels": {
        "computed": true,
        "description": "A set of key/value label pairs to assign to the bucket.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "lifecycle_rule": {
        "computed": true,
        "description": "The bucket's Lifecycle Rules configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action": [
                "set",
                [
                  "object",
                  {
                    "storage_class": "string",
                    "type": "string"
                  }
                ]
              ],
              "condition": [
                "set",
                [
                  "object",
                  {
                    "age": "number",
                    "created_before": "string",
                    "custom_time_before": "string",
                    "days_since_custom_time": "number",
                    "days_since_noncurrent_time": "number",
                    "matches_prefix": [
                      "list",
                      "string"
                    ],
                    "matches_storage_class": [
                      "list",
                      "string"
                    ],
                    "matches_suffix": [
                      "list",
                      "string"
                    ],
                    "noncurrent_time_before": "string",
                    "num_newer_versions": "number",
                    "send_age_if_zero": "bool",
                    "send_days_since_custom_time_if_zero": "bool",
                    "send_days_since_noncurrent_time_if_zero": "bool",
                    "send_num_newer_versions_if_zero": "bool",
                    "with_state": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description": "The Google Cloud Storage location or region.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging": {
        "computed": true,
        "description": "The bucket's Access \u0026 Storage Logs configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "log_bucket": "string",
              "log_object_prefix": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "The name of the bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
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
        "type": "string"
      },
      "requester_pays": {
        "computed": true,
        "description": "Enables Requester Pays on a storage bucket.",
        "description_kind": "plain",
        "type": "bool"
      },
      "retention_policy": {
        "computed": true,
        "description": "Configuration of the bucket's data retention policy for how long objects in the bucket should be retained.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "is_locked": "bool",
              "retention_period": "number"
            }
          ]
        ]
      },
      "rpo": {
        "computed": true,
        "description": "Specifies the RPO setting of bucket. If set 'ASYNC_TURBO', The Turbo Replication will be enabled for the dual-region bucket. Value 'DEFAULT' will set RPO setting to default. Turbo Replication is only for buckets in dual-regions.See the docs for more details.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "soft_delete_policy": {
        "computed": true,
        "description": "The bucket's soft delete policy, which defines the period of time that soft-deleted objects will be retained, and cannot be permanently deleted. If it is not provided, by default Google Cloud Storage sets this to default soft delete policy",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "effective_time": "string",
              "retention_duration_seconds": "number"
            }
          ]
        ]
      },
      "storage_class": {
        "computed": true,
        "description": "The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.",
        "description_kind": "plain",
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
      "time_created": {
        "computed": true,
        "description": "The creation time of the bucket in RFC 3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "uniform_bucket_level_access": {
        "computed": true,
        "description": "Enables uniform bucket-level access on a bucket.",
        "description_kind": "plain",
        "type": "bool"
      },
      "updated": {
        "computed": true,
        "description": "The time at which the bucket's metadata or IAM policy was last updated, in RFC 3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "url": {
        "computed": true,
        "description": "The base URL of the bucket, in the format gs://\u003cbucket-name\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "versioning": {
        "computed": true,
        "description": "The bucket's Versioning configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      },
      "website": {
        "computed": true,
        "description": "Configuration if the bucket acts as a website.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "main_page_suffix": "string",
              "not_found_page": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBucket), &result)
	return &result
}
