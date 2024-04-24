package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingMetric = `{
  "block": {
    "attributes": {
      "bucket_name": {
        "description": "The resource name of the Log Bucket that owns the Log Metric. Only Log Buckets in projects\nare supported. The bucket has to be in the same project as the metric.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "A description of this metric, which is used in documentation. The maximum length of the\ndescription is 8000 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description": "If set to True, then this metric is disabled and it does not generate any points.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "filter": {
        "description": "An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which\nis used to match log entries.",
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
      "label_extractors": {
        "description": "A map from a label key string to an extractor expression which is used to extract data from a log\nentry field and assign as the label value. Each label key specified in the LabelDescriptor must\nhave an associated extractor expression in this map. The syntax of the extractor expression is\nthe same as for the valueExtractor field.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The client-assigned metric identifier. Examples - \"error_count\", \"nginx/requests\".\nMetric identifiers are limited to 100 characters and can include only the following\ncharacters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash\ncharacter (/) denotes a hierarchy of name pieces, and it cannot be the first character\nof the name.",
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
      "value_extractor": {
        "description": "A valueExtractor is required when using a distribution logs-based metric to extract the values to\nrecord from a log entry. Two functions are supported for value extraction - EXTRACT(field) or\nREGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which\nthe value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax\n(https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified\nlog entry field. The value of the field is converted to a string before applying the regex. It is an\nerror to specify a regex that does not include exactly one capture group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "bucket_options": {
        "block": {
          "block_types": {
            "explicit_buckets": {
              "block": {
                "attributes": {
                  "bounds": {
                    "description": "The values must be monotonically increasing.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "description": "Specifies a set of buckets with arbitrary widths.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "exponential_buckets": {
              "block": {
                "attributes": {
                  "growth_factor": {
                    "description": "Must be greater than 1.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "num_finite_buckets": {
                    "description": "Must be greater than 0.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "scale": {
                    "description": "Must be greater than 0.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Specifies an exponential sequence of buckets that have a width that is proportional to the value of\nthe lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "linear_buckets": {
              "block": {
                "attributes": {
                  "num_finite_buckets": {
                    "description": "Must be greater than 0.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "offset": {
                    "description": "Lower bound of the first bucket.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "width": {
                    "description": "Must be greater than 0.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).\nEach bucket represents a constant absolute uncertainty on the specific value in the bucket.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it\ndescribes the bucket boundaries used to create a histogram of the extracted values.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "metric_descriptor": {
        "block": {
          "attributes": {
            "display_name": {
              "description": "A concise name for the metric, which can be displayed in user interfaces. Use sentence case\nwithout an ending period, for example \"Request count\". This field is optional but it is\nrecommended to be set for any metrics associated with user-visible concepts, such as Quota.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_kind": {
              "description": "Whether the metric records instantaneous values, changes to a value, etc.\nSome combinations of metricKind and valueType might not be supported.\nFor counter metrics, set this to DELTA. Possible values: [\"DELTA\", \"GAUGE\", \"CUMULATIVE\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "unit": {
              "description": "The unit in which the metric value is reported. It is only applicable if the valueType is\n'INT64', 'DOUBLE', or 'DISTRIBUTION'. The supported units are a subset of\n[The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value_type": {
              "description": "Whether the measurement is an integer, a floating-point number, etc.\nSome combinations of metricKind and valueType might not be supported.\nFor counter metrics, set this to INT64. Possible values: [\"BOOL\", \"INT64\", \"DOUBLE\", \"STRING\", \"DISTRIBUTION\", \"MONEY\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "labels": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "A human-readable description for the label.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "The label key.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value_type": {
                    "description": "The type of data that can be assigned to the label. Default value: \"STRING\" Possible values: [\"BOOL\", \"INT64\", \"STRING\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The set of labels that can be used to describe a specific instance of this metric type. For\nexample, the appengine.googleapis.com/http/server/response_latencies metric type has a label\nfor the HTTP response code, response_code, so you can look at latencies for successful responses\nor just for responses that failed.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "The optional metric descriptor associated with the logs-based metric.\nIf unspecified, it uses a default metric descriptor with a DELTA metric kind,\nINT64 value type, with no labels and a unit of \"1\". Such a metric counts the\nnumber of log entries matching the filter expression.",
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

func GoogleLoggingMetricSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingMetric), &result)
	return &result
}
