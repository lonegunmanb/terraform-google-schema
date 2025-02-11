package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleColabNotebookExecution = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "Required. The display name of the Notebook Execution.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "execution_timeout": {
        "description": "Max running time of the execution job in seconds (default 86400s / 24 hrs).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_user": {
        "description": "The user email to run the execution as.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gcs_output_uri": {
        "description": "The Cloud Storage location to upload the result to. Format:'gs://bucket-name'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the resource: https://cloud.google.com/colab/docs/locations",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notebook_execution_job_id": {
        "computed": true,
        "description": "User specified ID for the Notebook Execution Job",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notebook_runtime_template_resource_name": {
        "description": "The NotebookRuntimeTemplate to source compute configuration from.",
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
      "service_account": {
        "description": "The service account to run the execution as.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "dataform_repository_source": {
        "block": {
          "attributes": {
            "commit_sha": {
              "description": "The commit SHA to read repository with. If unset, the file will be read at HEAD.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dataform_repository_resource_name": {
              "description": "The resource name of the Dataform Repository.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The Dataform Repository containing the input notebook.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "direct_notebook_source": {
        "block": {
          "attributes": {
            "content": {
              "description": "The base64-encoded contents of the input notebook file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The content of the input notebook in ipynb format.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "gcs_notebook_source": {
        "block": {
          "attributes": {
            "generation": {
              "description": "The version of the Cloud Storage object to read. If unset, the current version of the object is read. See https://cloud.google.com/storage/docs/metadata#generation-number.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "description": "The Cloud Storage uri pointing to the ipynb file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The Cloud Storage uri for the input notebook.",
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

func GoogleColabNotebookExecutionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleColabNotebookExecution), &result)
	return &result
}
