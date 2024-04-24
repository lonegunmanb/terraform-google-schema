package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClouddeployTarget = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field ` + "`" + `effective_annotations` + "`" + ` for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Time at which the ` + "`" + `Target` + "`" + ` was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deploy_parameters": {
        "description": "Optional. The deploy parameters to use for this target.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "description": {
        "description": "Optional. Description of the ` + "`" + `Target` + "`" + `. Max length is 255 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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
        "description": "Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.",
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
        "description": "Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be \u003c= 128 bytes.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field ` + "`" + `effective_labels` + "`" + ` for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the ` + "`" + `Target` + "`" + `. Format is [a-z][a-z0-9\\-]{0,62}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "require_approval": {
        "description": "Optional. Whether or not the ` + "`" + `Target` + "`" + ` requires approval.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "target_id": {
        "computed": true,
        "description": "Output only. Resource id of the ` + "`" + `Target` + "`" + `.",
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
      "uid": {
        "computed": true,
        "description": "Output only. Unique identifier of the ` + "`" + `Target` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Most recent time at which the ` + "`" + `Target` + "`" + ` was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "anthos_cluster": {
        "block": {
          "attributes": {
            "membership": {
              "description": "Membership of the GKE Hub-registered cluster to which to apply the Skaffold configuration. Format is ` + "`" + `projects/{project}/locations/{location}/memberships/{membership_name}` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Information specifying an Anthos Cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "execution_configs": {
        "block": {
          "attributes": {
            "artifact_storage": {
              "computed": true,
              "description": "Optional. Cloud Storage location in which to store execution outputs. This can either be a bucket (\"gs://my-bucket\") or a path within a bucket (\"gs://my-bucket/my-dir\"). If unspecified, a default bucket located in the same region will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execution_timeout": {
              "computed": true,
              "description": "Optional. Execution timeout for a Cloud Build Execution. This must be between 10m and 24h in seconds format. If unspecified, a default timeout of 1h is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_account": {
              "computed": true,
              "description": "Optional. Google service account to use for execution. If unspecified, the project execution service account (-compute@developer.gserviceaccount.com) is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "usages": {
              "description": "Required. Usages when this configuration should be applied.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "worker_pool": {
              "description": "Optional. The resource name of the ` + "`" + `WorkerPool` + "`" + `, with the format ` + "`" + `projects/{project}/locations/{location}/workerPools/{worker_pool}` + "`" + `. If this optional field is unspecified, the default Cloud Build pool will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configurations for all execution that relates to this ` + "`" + `Target` + "`" + `. Each ` + "`" + `ExecutionEnvironmentUsage` + "`" + ` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the ` + "`" + `RENDER` + "`" + ` and ` + "`" + `DEPLOY` + "`" + ` ` + "`" + `ExecutionEnvironmentUsage` + "`" + ` values. When no configurations are specified, execution will use the default specified in ` + "`" + `DefaultPool` + "`" + `.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "gke": {
        "block": {
          "attributes": {
            "cluster": {
              "description": "Information specifying a GKE Cluster. Format is ` + "`" + `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "internal_ip": {
              "description": "Optional. If true, ` + "`" + `cluster` + "`" + ` is accessed using the private IP address of the control plane endpoint. Otherwise, the default IP address of the control plane endpoint is used. The default IP address is the private IP address for clusters with private control-plane endpoints and the public IP address otherwise. Only specify this option when ` + "`" + `cluster` + "`" + ` is a [private GKE cluster](https://cloud.google.com/kubernetes-engine/docs/concepts/private-cluster-concept).",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Information specifying a GKE Cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "multi_target": {
        "block": {
          "attributes": {
            "target_ids": {
              "description": "Required. The target_ids of this multiTarget.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Information specifying a multiTarget.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "run": {
        "block": {
          "attributes": {
            "location": {
              "description": "Required. The location where the Cloud Run Service should be located. Format is ` + "`" + `projects/{project}/locations/{location}` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Information specifying a Cloud Run deployment target.",
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

func GoogleClouddeployTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClouddeployTarget), &result)
	return &result
}
