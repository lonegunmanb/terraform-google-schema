package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudfunctionsFunction = `{
  "block": {
    "attributes": {
      "available_memory_mb": {
        "description": "Memory (in MB), available to the function. Default value is 256. Possible values include 128, 256, 512, 1024, etc.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "build_environment_variables": {
        "description": " A set of key/value environment variable pairs available during build time.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "build_worker_pool": {
        "description": "Name of the Cloud Build Custom Worker Pool that should be used to build the function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "Description of the function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "docker_registry": {
        "computed": true,
        "description": "Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "docker_repository": {
        "description": "User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry for storing images built with Cloud Build.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "entry_point": {
        "description": "Name of the function that will be executed when the Google Cloud Function is triggered.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_variables": {
        "description": "A set of key/value environment variable pairs to assign to the function.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "https_trigger_security_level": {
        "computed": true,
        "description": "The security level for the function. Defaults to SECURE_OPTIONAL. Valid only if trigger_http is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "https_trigger_url": {
        "computed": true,
        "description": "URL which triggers function execution. Returned only if trigger_http is used.",
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
      "ingress_settings": {
        "description": "String value that controls what traffic can reach the function. Allowed values are ALLOW_ALL and ALLOW_INTERNAL_ONLY. Changes to this field will recreate the cloud function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "description": "Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "max_instances": {
        "computed": true,
        "description": "The limit on the maximum number of function instances that may coexist at a given time.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_instances": {
        "description": "The limit on the minimum number of function instances that may coexist at a given time.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "A user-defined name of the function. Function names must be unique globally.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "Project of the function. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region of function. If it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "description": "The runtime in which the function is going to run. Eg. \"nodejs12\", \"nodejs14\", \"python37\", \"go111\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_account_email": {
        "computed": true,
        "description": " If provided, the self-provided service account to run the function with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_archive_bucket": {
        "description": "The GCS bucket containing the zip archive which contains the function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_archive_object": {
        "description": "The source archive object (file) in archive bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Describes the current stage of a deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "timeout": {
        "description": "Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "trigger_http": {
        "description": "Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as https_trigger_url. Cannot be used with trigger_bucket and trigger_topic.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpc_connector": {
        "description": "The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is projects/*/locations/*/connectors/*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_connector_egress_settings": {
        "computed": true,
        "description": "The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are ALL_TRAFFIC and PRIVATE_RANGES_ONLY. Defaults to PRIVATE_RANGES_ONLY. If unset, this field preserves the previously set value.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "event_trigger": {
        "block": {
          "attributes": {
            "event_type": {
              "description": "The type of event to observe. For example: \"google.storage.object.finalize\". See the documentation on calling Cloud Functions for a full reference of accepted triggers.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource": {
              "description": "The name or partial URI of the resource from which to observe events. For example, \"myBucket\" or \"projects/my-project/topics/my-topic\"",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "failure_policy": {
              "block": {
                "attributes": {
                  "retry": {
                    "description": "Whether the function should be retried on failure. Defaults to false.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Specifies policy for failed executions",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A source that fires events in response to a condition in another service. Cannot be used with trigger_http.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "secret_environment_variables": {
        "block": {
          "attributes": {
            "key": {
              "description": "Name of the environment variable.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "project_id": {
              "computed": true,
              "description": "Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret": {
              "description": "ID of the secret in secret manager (not the full resource name).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "description": "Version of the secret (version number or the string \"latest\"). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new clones start.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Secret environment variables configuration",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "secret_volumes": {
        "block": {
          "attributes": {
            "mount_path": {
              "description": "The path within the container to mount the secret volume. For example, setting the mount_path as \"/etc/secrets\" would mount the secret value files under the \"/etc/secrets\" directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount paths: \"/etc/secrets\" Restricted mount paths: \"/cloudsql\", \"/dev/log\", \"/pod\", \"/proc\", \"/var/log\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "project_id": {
              "computed": true,
              "description": "Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret": {
              "description": "ID of the secret in secret manager (not the full resource name).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "versions": {
              "block": {
                "attributes": {
                  "path": {
                    "description": "Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mount_path as \"/etc/secrets\" and path as \"/secret_foo\" would mount the secret value file at \"/etc/secrets/secret_foo\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "version": {
                    "description": "Version of the secret (version number or the string \"latest\"). It is preferable to use \"latest\" version with secret volumes as secret value changes are reflected immediately.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "List of secret versions to mount for this secret. If empty, the \"latest\" version of the secret will be made available in a file named after the secret under the mount point.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Secret volumes configuration.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "source_repository": {
        "block": {
          "attributes": {
            "deployed_url": {
              "computed": true,
              "description": "The URL pointing to the hosted repository where the function was defined at the time of deployment.",
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "description": "The URL pointing to the hosted repository where the function is defined.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Represents parameters related to source repository where a function is hosted. Cannot be set alongside source_archive_bucket or source_archive_object.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudfunctionsFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudfunctionsFunction), &result)
	return &result
}
