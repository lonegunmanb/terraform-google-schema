package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDatabaseMigrationServiceMigrationJob = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC 'Zulu' format, accurate to nanoseconds. Example: '2014-10-02T15:01:23.045123456Z'.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination": {
        "description": "The name of the destination connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{destinationConnectionProfile}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "The migration job display name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dump_path": {
        "description": "The path to the dump file in Google Cloud Storage,\nin the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).\nThis field and the \"dump_flags\" field are mutually exclusive.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dump_type": {
        "description": "The type of the data dump. Supported for MySQL to CloudSQL for MySQL\nmigrations only. Possible values: [\"LOGICAL\", \"PHYSICAL\"]",
        "description_kind": "plain",
        "optional": true,
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
      "error": {
        "computed": true,
        "description": "Output only. The error details in case of state FAILED.",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the migration job should reside.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "migration_job_id": {
        "description": "The ID of the migration job.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of this migration job resource in the form of projects/{project}/locations/{location}/migrationJobs/{migrationJob}.",
        "description_kind": "plain",
        "type": "string"
      },
      "phase": {
        "computed": true,
        "description": "The current migration job phase.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "description": "The name of the source connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{sourceConnectionProfile}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current migration job state.",
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
      "type": {
        "description": "The type of the migration job. Possible values: [\"ONE_TIME\", \"CONTINUOUS\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "dump_flags": {
        "block": {
          "block_types": {
            "dump_flags": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The name of the flag",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The vale of the flag",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A list of dump flags",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The initial dump flags.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "performance_config": {
        "block": {
          "attributes": {
            "dump_parallel_level": {
              "description": "Initial dump parallelism level. Possible values: [\"MIN\", \"OPTIMAL\", \"MAX\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Data dump parallelism settings used by the migration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "reverse_ssh_connectivity": {
        "block": {
          "attributes": {
            "vm": {
              "description": "The name of the virtual machine (Compute Engine) used as the bastion server\nfor the SSH tunnel.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vm_ip": {
              "description": "The IP of the virtual machine (Compute Engine) used as the bastion server\nfor the SSH tunnel.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vm_port": {
              "description": "The forwarding port of the virtual machine (Compute Engine) used as the\nbastion server for the SSH tunnel.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "vpc": {
              "description": "The name of the VPC to peer with the Cloud SQL private network.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The details of the VPC network that the source database is located in.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "static_ip_connectivity": {
        "block": {
          "description": "If set to an empty object ('{}'), the source database will allow incoming\nconnections from the public IP of the destination database.\nYou can retrieve the public IP of the Cloud SQL instance from the\nCloud SQL console or using Cloud SQL APIs.",
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
      },
      "vpc_peering_connectivity": {
        "block": {
          "attributes": {
            "vpc": {
              "description": "The name of the VPC network to peer with the Cloud SQL private network.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The details of the VPC network that the source database is located in.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDatabaseMigrationServiceMigrationJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDatabaseMigrationServiceMigrationJob), &result)
	return &result
}
