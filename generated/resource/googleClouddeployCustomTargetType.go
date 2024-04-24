package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClouddeployCustomTargetType = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Time at which the 'CustomTargetType' was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_target_type_id": {
        "computed": true,
        "description": "Resource id of the 'CustomTargetType'.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the 'CustomTargetType'. Max length is 255 characters.",
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
        "description": "The weak etag of the 'CustomTargetType' resource. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.",
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
        "description": "Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be \u003c= 128 bytes.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the source.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the 'CustomTargetType'.",
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
        "description": "Unique identifier of the 'CustomTargetType'.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time at which the 'CustomTargetType' was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "custom_actions": {
        "block": {
          "attributes": {
            "deploy_action": {
              "description": "The Skaffold custom action responsible for deploy operations.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "render_action": {
              "description": "The Skaffold custom action responsible for render operations. If not provided then Cloud Deploy will perform the render operations via 'skaffold render'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "include_skaffold_modules": {
              "block": {
                "attributes": {
                  "configs": {
                    "description": "The Skaffold Config modules to use from the specified source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "git": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Relative path from the repository root to the Skaffold file.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "ref": {
                          "description": "Git ref the package should be cloned from.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "repo": {
                          "description": "Git repository the package should be cloned from.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Remote git repository containing the Skaffold Config modules.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "google_cloud_storage": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Relative path from the source to the Skaffold file.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "source": {
                          "description": "Cloud Storage source paths to copy recursively. For example, providing 'gs://my-bucket/dir/configs/*' will result in Skaffold copying all files within the 'dir/configs' directory in the bucket 'my-bucket'.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Cloud Storage bucket containing Skaffold Config modules.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of Skaffold modules Cloud Deploy will include in the Skaffold Config as required before performing diagnose.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Configures render and deploy for the 'CustomTargetType' using Skaffold custom actions.",
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

func GoogleClouddeployCustomTargetTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClouddeployCustomTargetType), &result)
	return &result
}
