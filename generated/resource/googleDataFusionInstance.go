package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataFusionInstance = `{
  "block": {
    "attributes": {
      "api_endpoint": {
        "computed": true,
        "description": "Endpoint on which the REST APIs is accessible.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time the instance was created in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataproc_service_account": {
        "description": "User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "An optional description of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name for an instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_rbac": {
        "description": "Option to enable granular role-based access control.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_stackdriver_logging": {
        "description": "Option to enable Stackdriver Logging.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_stackdriver_monitoring": {
        "description": "Option to enable Stackdriver Monitoring.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gcs_bucket": {
        "computed": true,
        "description": "Cloud Storage bucket generated by Data Fusion in the customer project.",
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
        "description": "The resource labels for instance to use to annotate any related underlying resources,\nsuch as Compute Engine VMs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The ID of the instance or a fully qualified identifier for the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "options": {
        "computed": true,
        "description": "Map of additional options used to configure the behavior of Data Fusion instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "p4_service_account": {
        "computed": true,
        "description": "P4 service account for the customer project.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_instance": {
        "description": "Specifies whether the Data Fusion instance should be private. If set to\ntrue, all Data Fusion nodes will have private IP addresses and will not be\nable to access the public internet.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region of the Data Fusion instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_endpoint": {
        "computed": true,
        "description": "Endpoint on which the Data Fusion UI and REST APIs are accessible.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of this Data Fusion instance.\n- CREATING: Instance is being created\n- RUNNING: Instance is running and ready for requests\n- FAILED: Instance creation failed\n- DELETING: Instance is being deleted\n- UPGRADING: Instance is being upgraded\n- RESTARTING: Instance is being restarted",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "Additional information about the current state of this Data Fusion instance if available.",
        "description_kind": "plain",
        "type": "string"
      },
      "tenant_project_id": {
        "computed": true,
        "description": "The name of the tenant project.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "Represents the type of Data Fusion instance. Each type is configured with\nthe default settings for processing and memory.\n- BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines\nusing point and click UI. However, there are certain limitations, such as fewer number\nof concurrent pipelines, no support for streaming pipelines, etc.\n- ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features\navailable, such as support for streaming pipelines, higher number of concurrent pipelines, etc.\n- DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but\nwith restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration\npipelines at low cost. Possible values: [\"BASIC\", \"ENTERPRISE\", \"DEVELOPER\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time the instance was last updated in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "Current version of the Data Fusion.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description": "Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "accelerators": {
        "block": {
          "attributes": {
            "accelerator_type": {
              "description": "The type of an accelator for a CDF instance. Possible values: [\"CDC\", \"HEALTHCARE\", \"CCAI_INSIGHTS\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state": {
              "description": "The type of an accelator for a CDF instance. Possible values: [\"ENABLED\", \"DISABLED\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "List of accelerators enabled for this CDF instance.\n\nIf accelerators are enabled it is possible a permadiff will be created with the Options field.\nUsers will need to either manually update their state file to include these diffed options, or include the field in a [lifecycle ignore changes block](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes).",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "crypto_key_config": {
        "block": {
          "attributes": {
            "key_reference": {
              "description": "The name of the key which is used to encrypt/decrypt customer data. For key in Cloud KMS, the key should be in the format of projects/*/locations/*/keyRings/*/cryptoKeys/*.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "event_publish_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Option to enable Event Publishing.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "topic": {
              "description": "The resource name of the Pub/Sub topic. Format: projects/{projectId}/topics/{topic_id}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Option to enable and pass metadata for event publishing.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "attributes": {
            "ip_allocation": {
              "description": "The IP range in CIDR notation to use for the managed Data Fusion instance\nnodes. This range must not overlap with any other ranges used in the Data Fusion instance network.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network": {
              "description": "Name of the network in the project with which the tenant project\nwill be peered for executing pipelines. In case of shared VPC where the network resides in another host\nproject the network should specified in the form of projects/{host-project-id}/global/networks/{network}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Network configuration options. These are required when a private Data Fusion instance is to be created.",
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

func GoogleDataFusionInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataFusionInstance), &result)
	return &result
}