package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleColabSchedule = `{
  "block": {
    "attributes": {
      "allow_queueing": {
        "description": "Whether new scheduled runs can be queued when max_concurrent_runs limit is reached. If set to true, new runs will be queued instead of skipped. Default to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cron": {
        "description": "Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "desired_state": {
        "description": "Desired state of the Colab Schedule. Set this field to 'ACTIVE' to start/resume the schedule, and 'PAUSED' to pause the schedule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Required. The display name of the Schedule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "end_time": {
        "description": "Timestamp after which no new runs can be scheduled. If specified, the schedule will be completed when either end_time is reached or when scheduled_run_count \u003e= max_run_count. Must be in the RFC 3339 (https://www.ietf.org/rfc/rfc3339.txt) format.",
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
      "location": {
        "description": "The location for the resource: https://cloud.google.com/colab/docs/locations",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_concurrent_run_count": {
        "description": "Maximum number of runs that can be started concurrently for this Schedule. This is the limit for starting the scheduled requests and not the execution of the notebook execution jobs created by the requests.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_run_count": {
        "description": "Maximum run count of the schedule. If specified, The schedule will be completed when either startedRunCount \u003e= maxRunCount or when endTime is reached. If not specified, new runs will keep getting scheduled until this Schedule is paused or deleted. Already scheduled runs will be allowed to complete. Unset if not specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Schedule",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description": "The timestamp after which the first run can be scheduled. Defaults to the schedule creation time. Must be in the RFC 3339 (https://www.ietf.org/rfc/rfc3339.txt) format.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The state of the schedule.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "create_notebook_execution_job_request": {
        "block": {
          "block_types": {
            "notebook_execution_job": {
              "block": {
                "attributes": {
                  "display_name": {
                    "description": "Required. The display name of the Notebook Execution.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "execution_timeout": {
                    "description": "Max running time of the execution job in seconds (default 86400s / 24 hrs). A duration in seconds with up to nine fractional digits, ending with \"s\". Example: \"3.5s\".",
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
                  "notebook_runtime_template_resource_name": {
                    "description": "The NotebookRuntimeTemplate to source compute configuration from.",
                    "description_kind": "plain",
                    "required": true,
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
                          "description": "The Cloud Storage uri pointing to the ipynb file. Format: gs://bucket/notebook_file.ipynb",
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
                  }
                },
                "description": "The NotebookExecutionJob to create.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Request for google_colab_notebook_execution.",
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

func GoogleColabScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleColabSchedule), &result)
	return &result
}
