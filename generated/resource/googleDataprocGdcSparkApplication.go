package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocGdcSparkApplication = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "The annotations to associate with this application. Annotations may be used to store client information, but are not used by the server. \n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "application_environment": {
        "description": "An ApplicationEnvironment from which to inherit configuration properties.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dependency_images": {
        "description": "List of container image uris for additional file dependencies. Dependent files are sequentially copied from each image. If a file with the same name exists in 2 images then the file from later image is used.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description": "User-provided human-readable name to be used in user interfaces.",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels to associate with this application. Labels may be used for filtering and billing tracking. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the spark application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitoring_endpoint": {
        "computed": true,
        "description": "URL for a monitoring UI for this application (for eventual Spark PHS/UI support) Out of scope for private GA",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the application. Format: projects/{project}/locations/{location}/serviceInstances/{service_instance}/sparkApplications/{application}",
        "description_kind": "plain",
        "type": "string"
      },
      "namespace": {
        "description": "The Kubernetes namespace in which to create the application. This namespace must already exist on the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_uri": {
        "computed": true,
        "description": "An HCFS URI pointing to the location of stdout and stdout of the application Mainly useful for Pantheon and gcloud Not in scope for private GA",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "properties": {
        "description": "application-specific properties.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "reconciling": {
        "computed": true,
        "description": "Whether the application is currently reconciling. True if the current state of the resource does not match the intended state, and the system is working to reconcile them, whether or not the change was user initiated.",
        "description_kind": "plain",
        "type": "bool"
      },
      "serviceinstance": {
        "description": "The id of the service instance to which this spark application belongs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "spark_application_id": {
        "description": "The id of the application",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state.\nPossible values:\n* 'STATE_UNSPECIFIED'\n* 'PENDING'\n* 'RUNNING'\n* 'CANCELLING'\n* 'CANCELLED'\n* 'SUCCEEDED'\n* 'FAILED'",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "A message explaining the current state.",
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
        "description": "System generated unique identifier for this application, formatted as UUID4.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the resource was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "description": "The Dataproc version of this application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "pyspark_application_config": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "The arguments to pass to the driver.  Do not include arguments, such as '--conf', that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_python_file_uri": {
              "description": "The HCFS URI of the main Python file to use as the driver. Must be a .py file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "python_file_uris": {
              "description": "HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Represents the PySparkApplicationConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_application_config": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: '.jar', '.tar', '.tar.gz', '.tgz', and '.zip'.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "The arguments to pass to the driver. Do not include arguments that can be set as application properties, such as '--conf', since a collision can occur that causes an incorrect application submission.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "HCFS URIs of files to be placed in the working directory of each executor.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to add to the classpath of the Spark driver and tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_class": {
              "description": "The name of the driver main class. The jar file that contains the class must be in the classpath or specified in 'jar_file_uris'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "main_jar_file_uri": {
              "description": "The HCFS URI of the jar file that contains the main class.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Represents the SparkApplicationConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_r_application_config": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "The arguments to pass to the driver.  Do not include arguments, such as '--conf', that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_r_file_uri": {
              "description": "The HCFS URI of the main R file to use as the driver. Must be a .R file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Represents the SparkRApplicationConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_sql_application_config": {
        "block": {
          "attributes": {
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to be added to the Spark CLASSPATH.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "query_file_uri": {
              "description": "The HCFS URI of the script that contains SQL queries.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "script_variables": {
              "description": "Mapping of query variable names to values (equivalent to the Spark SQL command: SET 'name=\"value\";').",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "query_list": {
              "block": {
                "attributes": {
                  "queries": {
                    "description": "The queries to run.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Represents a list of queries.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Represents the SparkRApplicationConfig.",
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

func GoogleDataprocGdcSparkApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocGdcSparkApplication), &result)
	return &result
}
