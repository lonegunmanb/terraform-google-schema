package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaseAppHostingBuild = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Unstructured key value map that may be set by external tools to\nstore and arbitrary metadata. They are not queryable and should be\npreserved when modifying objects.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "backend": {
        "description": "The ID of the Backend that this Build applies to",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "build_id": {
        "description": "The user-specified ID of the build being created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "build_logs_uri": {
        "computed": true,
        "description": "The location of the [Cloud Build\nlogs](https://cloud.google.com/build/docs/view-build-results) for the build\nprocess.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time at which the build was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Human-readable name. 63 character limit.",
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
      "environment": {
        "computed": true,
        "description": "The environment name of the backend when this build was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "error": {
        "computed": true,
        "description": "The 'Status' type defines a logical error model that is suitable for\ndifferent programming environments, including REST APIs and RPC APIs. It is\nused by [gRPC](https://github.com/grpc). Each 'Status' message contains\nthree pieces of data: error code, error message, and error details.\n\nYou can find out more about this error model and how to work with it in the\n[API Design Guide](https://cloud.google.com/apis/design/errors).",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "number",
              "details": [
                "list",
                [
                  "map",
                  "string"
                ]
              ],
              "message": "string"
            }
          ]
        ]
      },
      "error_source": {
        "computed": true,
        "description": "The source of the error for the build, if in a 'FAILED' state.\nPossible values:\nCLOUD_BUILD\nCLOUD_RUN",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Server-computed checksum based on other values; may be sent\non update or delete to ensure operation is done on expected resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image": {
        "computed": true,
        "description": "The Artifact Registry\n[container\nimage](https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.dockerImages)\nURI, used by the Cloud Run\n['revision'](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services.revisions)\nfor this build.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Unstructured key value map that can be used to organize and categorize\nobjects.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the Backend that this Build applies to",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the build.\n\nFormat:\n\n'projects/{project}/locations/{locationId}/backends/{backendId}/builds/{buildId}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the build.\nPossible values:\nBUILDING\nBUILT\nDEPLOYING\nREADY\nFAILED",
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
        "description": "System-assigned, unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time at which the build was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "source": {
        "block": {
          "block_types": {
            "codebase": {
              "block": {
                "attributes": {
                  "author": {
                    "computed": true,
                    "description": "Version control metadata for a user associated with a resolved codebase.\nCurrently assumes a Git user.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "display_name": "string",
                          "email": "string",
                          "image_uri": "string"
                        }
                      ]
                    ]
                  },
                  "branch": {
                    "description": "The branch in the codebase to build from, using the latest commit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "commit": {
                    "description": "The commit in the codebase to build from.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "commit_message": {
                    "computed": true,
                    "description": "The message of a codebase change.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "commit_time": {
                    "computed": true,
                    "description": "The time the change was made.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "display_name": {
                    "computed": true,
                    "description": "The human-friendly name to use for this Codebase when displaying a build.\nWe use the first eight characters of the SHA-1 hash for GitHub.com.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "hash": {
                    "computed": true,
                    "description": "The full SHA-1 hash of a Git commit, if available.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "uri": {
                    "computed": true,
                    "description": "A URI linking to the codebase on an hosting provider's website. May\nnot be valid if the commit has been rebased or force-pushed out of\nexistence in the linked repository.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "A codebase source, representing the state of the codebase\nthat the build will be created at.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "container": {
              "block": {
                "attributes": {
                  "image": {
                    "description": "A URI representing a container for the backend to use.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The URI of an Artifact Registry\n[container\nimage](https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.dockerImages)\nto use as the build source.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The source for the build.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleFirebaseAppHostingBuildSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaseAppHostingBuild), &result)
	return &result
}
