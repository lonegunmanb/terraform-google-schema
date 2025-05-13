package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryRoutine = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time when this routine was created, in milliseconds since the\nepoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "data_governance_type": {
        "description": "If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask Possible values: [\"DATA_MASKING\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dataset_id": {
        "description": "The ID of the dataset containing this routine",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "definition_body": {
        "description": "The body of the routine. For functions, this is the expression in the AS clause.\nIf language=SQL, it is the substring inside (but excluding) the parentheses.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the routine if defined.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "determinism_level": {
        "description": "The determinism level of the JavaScript UDF if defined. Possible values: [\"DETERMINISM_LEVEL_UNSPECIFIED\", \"DETERMINISTIC\", \"NOT_DETERMINISTIC\"]",
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
      "imported_libraries": {
        "description": "Optional. If language = \"JAVASCRIPT\", this field stores the path of the\nimported JAVASCRIPT libraries.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "language": {
        "description": "The language of the routine. Possible values: [\"SQL\", \"JAVASCRIPT\", \"PYTHON\", \"JAVA\", \"SCALA\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time when this routine was modified, in milliseconds since the\nepoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "return_table_type": {
        "description": "Optional. Can be set only if routineType = \"TABLE_VALUED_FUNCTION\".\n\nIf absent, the return table type is inferred from definitionBody at query time in each query\nthat references this routine. If present, then the columns in the evaluated table result will\nbe cast to match the column types specificed in return table type, at query time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "return_type": {
        "description": "A JSON schema for the return type. Optional if language = \"SQL\"; required otherwise.\nIf absent, the return type is inferred from definitionBody at query time in each query\nthat references this routine. If present, then the evaluated result will be cast to\nthe specified returned type at query time. ~\u003e**NOTE**: Because this field expects a JSON\nstring, any changes to the string will create a diff, even if the JSON itself hasn't\nchanged. If the API returns a different value for the same schema, e.g. it switche\nd the order of values or replaced STRUCT field type with RECORD field type, we currently\ncannot suppress the recurring diff this causes. As a workaround, we recommend using\nthe schema as returned by the API.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routine_id": {
        "description": "The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routine_type": {
        "description": "The type of routine. Possible values: [\"SCALAR_FUNCTION\", \"PROCEDURE\", \"TABLE_VALUED_FUNCTION\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_mode": {
        "description": "Optional. The security mode of the routine, if defined. If not defined, the security mode is automatically determined from the routine's configuration. Possible values: [\"DEFINER\", \"INVOKER\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "arguments": {
        "block": {
          "attributes": {
            "argument_kind": {
              "description": "Defaults to FIXED_TYPE. Default value: \"FIXED_TYPE\" Possible values: [\"FIXED_TYPE\", \"ANY_TYPE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_type": {
              "description": "A JSON schema for the data type. Required unless argumentKind = ANY_TYPE.\n~\u003e**NOTE**: Because this field expects a JSON string, any changes to the string\nwill create a diff, even if the JSON itself hasn't changed. If the API returns\na different value for the same schema, e.g. it switched the order of values\nor replaced STRUCT field type with RECORD field type, we currently cannot\nsuppress the recurring diff this causes. As a workaround, we recommend using\nthe schema as returned by the API.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description": "Specifies whether the argument is input or output. Can be set for procedures only. Possible values: [\"IN\", \"OUT\", \"INOUT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "The name of this argument. Can be absent for function return argument.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Input/output argument of a function or a stored procedure.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "remote_function_options": {
        "block": {
          "attributes": {
            "connection": {
              "description": "Fully qualified name of the user-provided connection object which holds\nthe authentication information to send requests to the remote service.\nFormat: \"projects/{projectId}/locations/{locationId}/connections/{connectionId}\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "endpoint": {
              "description": "Endpoint of the user-provided remote service, e.g.\n'https://us-east1-my_gcf_project.cloudfunctions.net/remote_add'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_batching_rows": {
              "description": "Max number of rows in each batch sent to the remote service. If absent or if 0,\nBigQuery dynamically decides the number of rows in a batch.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_defined_context": {
              "computed": true,
              "description": "User-defined context as a set of key/value pairs, which will be sent as function\ninvocation context together with batched arguments in the requests to the remote\nservice. The total number of bytes of keys and values must be less than 8KB.\n\nAn object containing a list of \"key\": value pairs. Example:\n'{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }'.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description": "Remote function specific options.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_options": {
        "block": {
          "attributes": {
            "archive_uris": {
              "computed": true,
              "description": "Archive files to be extracted into the working directory of each executor. For more information about Apache Spark, see Apache Spark.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "connection": {
              "description": "Fully qualified name of the user-provided Spark connection object.\nFormat: \"projects/{projectId}/locations/{locationId}/connections/{connectionId}\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_image": {
              "description": "Custom container image for the runtime environment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "file_uris": {
              "computed": true,
              "description": "Files to be placed in the working directory of each executor. For more information about Apache Spark, see Apache Spark.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "jar_uris": {
              "computed": true,
              "description": "JARs to include on the driver and executor CLASSPATH. For more information about Apache Spark, see Apache Spark.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_class": {
              "description": "The fully qualified name of a class in jarUris, for example, com.example.wordcount.\nExactly one of mainClass and main_jar_uri field should be set for Java/Scala language type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "main_file_uri": {
              "description": "The main file/jar URI of the Spark application.\nExactly one of the definitionBody field and the mainFileUri field must be set for Python.\nExactly one of mainClass and mainFileUri field should be set for Java/Scala language type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "properties": {
              "computed": true,
              "description": "Configuration properties as a set of key/value pairs, which will be passed on to the Spark application.\nFor more information, see Apache Spark and the procedure option list.\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "py_file_uris": {
              "computed": true,
              "description": "Python files to be placed on the PYTHONPATH for PySpark application. Supported file types: .py, .egg, and .zip. For more information about Apache Spark, see Apache Spark.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "runtime_version": {
              "description": "Runtime version. If not specified, the default runtime version is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Optional. If language is one of \"PYTHON\", \"JAVA\", \"SCALA\", this field stores the options for spark stored procedure.",
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

func GoogleBigqueryRoutineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryRoutine), &result)
	return &result
}
