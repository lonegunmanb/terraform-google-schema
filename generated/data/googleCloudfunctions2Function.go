package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudfunctions2Function = `{
  "block": {
    "attributes": {
      "build_config": {
        "computed": true,
        "description": "Describes the Build step of the function that builds a container\nfrom the given source.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_update_policy": [
                "list",
                [
                  "object",
                  {}
                ]
              ],
              "build": "string",
              "docker_repository": "string",
              "entry_point": "string",
              "environment_variables": [
                "map",
                "string"
              ],
              "on_deploy_update_policy": [
                "list",
                [
                  "object",
                  {
                    "runtime_version": "string"
                  }
                ]
              ],
              "runtime": "string",
              "service_account": "string",
              "source": [
                "list",
                [
                  "object",
                  {
                    "repo_source": [
                      "list",
                      [
                        "object",
                        {
                          "branch_name": "string",
                          "commit_sha": "string",
                          "dir": "string",
                          "invert_regex": "bool",
                          "project_id": "string",
                          "repo_name": "string",
                          "tag_name": "string"
                        }
                      ]
                    ],
                    "storage_source": [
                      "list",
                      [
                        "object",
                        {
                          "bucket": "string",
                          "generation": "number",
                          "object": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "worker_pool": "string"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "User-provided description of a function.",
        "description_kind": "plain",
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
      "environment": {
        "computed": true,
        "description": "The environment the function is hosted on.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_trigger": {
        "computed": true,
        "description": "An Eventarc trigger managed by Google Cloud Functions that fires events in\nresponse to a condition in another service.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "event_filters": [
                "set",
                [
                  "object",
                  {
                    "attribute": "string",
                    "operator": "string",
                    "value": "string"
                  }
                ]
              ],
              "event_type": "string",
              "pubsub_topic": "string",
              "retry_policy": "string",
              "service_account_email": "string",
              "trigger": "string",
              "trigger_region": "string"
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
      "kms_key_name": {
        "computed": true,
        "description": "Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.\nIt must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "A set of key/value label pairs associated with this Cloud Function.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of this cloud function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A user-defined name of the function. Function names must\nbe unique globally and match pattern 'projects/*/locations/*/functions/*'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_config": {
        "computed": true,
        "description": "Describes the Service being deployed.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "all_traffic_on_latest_revision": "bool",
              "available_cpu": "string",
              "available_memory": "string",
              "binary_authorization_policy": "string",
              "environment_variables": [
                "map",
                "string"
              ],
              "gcf_uri": "string",
              "ingress_settings": "string",
              "max_instance_count": "number",
              "max_instance_request_concurrency": "number",
              "min_instance_count": "number",
              "secret_environment_variables": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "project_id": "string",
                    "secret": "string",
                    "version": "string"
                  }
                ]
              ],
              "secret_volumes": [
                "list",
                [
                  "object",
                  {
                    "mount_path": "string",
                    "project_id": "string",
                    "secret": "string",
                    "versions": [
                      "list",
                      [
                        "object",
                        {
                          "path": "string",
                          "version": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "service": "string",
              "service_account_email": "string",
              "timeout_seconds": "number",
              "uri": "string",
              "vpc_connector": "string",
              "vpc_connector_egress_settings": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "Describes the current state of the function.",
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
      "update_time": {
        "computed": true,
        "description": "The last update timestamp of a Cloud Function.",
        "description_kind": "plain",
        "type": "string"
      },
      "url": {
        "computed": true,
        "description": "Output only. The deployed url for the function.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudfunctions2FunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudfunctions2Function), &result)
	return &result
}
