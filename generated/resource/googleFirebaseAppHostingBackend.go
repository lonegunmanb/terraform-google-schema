package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaseAppHostingBackend = `{
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
      "app_id": {
        "description": "The [ID of a Web\nApp](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)\nassociated with the backend.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backend_id": {
        "description": "Id of the backend. Also used as the service ID for Cloud Run, and as part\nof the default domain name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time at which the backend was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Time at which the backend was deleted.",
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
        "description": "The environment name of the backend, used to load environment variables\nfrom environment specific configuration.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The canonical IDs of a Google Cloud location such as \"us-east1\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_resources": {
        "computed": true,
        "description": "A list of the resources managed by this backend.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "run_service": [
                "list",
                [
                  "object",
                  {
                    "service": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the backend.\n\nFormat:\n\n'projects/{project}/locations/{locationId}/backends/{backendId}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account": {
        "description": "The name of the service account used for Cloud Build and Cloud Run.\nShould have the role roles/firebaseapphosting.computeRunner\nor equivalent permissions.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "serving_locality": {
        "description": "Immutable. Specifies how App Hosting will serve the content for this backend. It will\neither be contained to a single region (REGIONAL_STRICT) or allowed to use\nApp Hosting's global-replicated serving infrastructure (GLOBAL_ACCESS). Possible values: [\"REGIONAL_STRICT\", \"GLOBAL_ACCESS\"]",
        "description_kind": "plain",
        "required": true,
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
        "description": "Time at which the backend was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "uri": {
        "computed": true,
        "description": "The primary URI to communicate with the backend.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "codebase": {
        "block": {
          "attributes": {
            "repository": {
              "description": "The resource name for the Developer Connect\n['gitRepositoryLink'](https://cloud.google.com/developer-connect/docs/api/reference/rest/v1/projects.locations.connections.gitRepositoryLinks)\nconnected to this backend, in the format:\n\nprojects/{project}/locations/{location}/connections/{connection}/gitRepositoryLinks/{repositoryLink}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "root_directory": {
              "description": "If 'repository' is provided, the directory relative to the root of the\nrepository to use as the root for the deployed web app.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The connection to an external source repository to watch for event-driven\nupdates to the backend.",
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

func GoogleFirebaseAppHostingBackendSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaseAppHostingBackend), &result)
	return &result
}
