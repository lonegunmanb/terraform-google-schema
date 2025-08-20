package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudTasksQueue = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the queue",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The queue name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "app_engine_routing_override": {
        "block": {
          "attributes": {
            "host": {
              "computed": true,
              "description": "The host that the task is sent to.",
              "description_kind": "plain",
              "type": "string"
            },
            "instance": {
              "description": "App instance.\n\nBy default, the task is sent to an instance which is available when the task is attempted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service": {
              "description": "App service.\n\nBy default, the task is sent to the service which is the default service when the task is attempted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description": "App version.\n\nBy default, the task is sent to the version which is the default version when the task is attempted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Overrides for task-level appEngineRouting. These settings apply only\nto App Engine tasks in this queue",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_target": {
        "block": {
          "attributes": {
            "http_method": {
              "computed": true,
              "description": "The HTTP method to use for the request.\n\nWhen specified, it overrides HttpRequest for the task.\nNote that if the value is set to GET the body of the task will be ignored at execution time. Possible values: [\"HTTP_METHOD_UNSPECIFIED\", \"POST\", \"GET\", \"HEAD\", \"PUT\", \"DELETE\", \"PATCH\", \"OPTIONS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "header_overrides": {
              "block": {
                "block_types": {
                  "header": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description": "The Key of the header.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "The Value of the header.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Header embodying a key and a value.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "HTTP target headers.\n\nThis map contains the header field names and values.\nHeaders will be set when running the CreateTask and/or BufferTask.\n\nThese headers represent a subset of the headers that will be configured for the task's HTTP request.\nSome HTTP request headers will be ignored or replaced.\n\nHeaders which can have multiple values (according to RFC2616) can be specified using comma-separated values.\n\nThe size of the headers must be less than 80KB. Queue-level headers to override headers of all the tasks in the queue.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "oauth_token": {
              "block": {
                "attributes": {
                  "scope": {
                    "computed": true,
                    "description": "OAuth scope to be used for generating OAuth access token.\nIf not specified, \"https://www.googleapis.com/auth/cloud-platform\" will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_email": {
                    "description": "Service account email to be used for generating OAuth token.\nThe service account must be within the same project as the queue.\nThe caller must have iam.serviceAccounts.actAs permission for the service account.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "If specified, an OAuth token is generated and attached as the Authorization header in the HTTP request.\n\nThis type of authorization should generally be used only when calling Google APIs hosted on *.googleapis.com.\nNote that both the service account email and the scope MUST be specified when using the queue-level authorization override.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oidc_token": {
              "block": {
                "attributes": {
                  "audience": {
                    "computed": true,
                    "description": "Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_email": {
                    "description": "Service account email to be used for generating OIDC token.\nThe service account must be within the same project as the queue.\nThe caller must have iam.serviceAccounts.actAs permission for the service account.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "If specified, an OIDC token is generated and attached as an Authorization header in the HTTP request.\n\nThis type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.\nNote that both the service account email and the audience MUST be specified when using the queue-level authorization override.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "uri_override": {
              "block": {
                "attributes": {
                  "host": {
                    "description": "Host override.\n\nWhen specified, replaces the host part of the task URL.\nFor example, if the task URL is \"https://www.google.com\", and host value\nis set to \"example.net\", the overridden URI will be changed to \"https://example.net\".\nHost value cannot be an empty string (INVALID_ARGUMENT).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description": "Port override.\n\nWhen specified, replaces the port part of the task URI.\nFor instance, for a URI http://www.google.com/foo and port=123, the overridden URI becomes http://www.google.com:123/foo.\nNote that the port value must be a positive integer.\nSetting the port to 0 (Zero) clears the URI port.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scheme": {
                    "computed": true,
                    "description": "Scheme override.\n\nWhen specified, the task URI scheme is replaced by the provided value (HTTP or HTTPS). Possible values: [\"HTTP\", \"HTTPS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uri_override_enforce_mode": {
                    "computed": true,
                    "description": "URI Override Enforce Mode\n\nWhen specified, determines the Target UriOverride mode. If not specified, it defaults to ALWAYS. Possible values: [\"ALWAYS\", \"IF_NOT_EXISTS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "path_override": {
                    "block": {
                      "attributes": {
                        "path": {
                          "computed": true,
                          "description": "The URI path (e.g., /users/1234). Default is an empty string.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "URI path.\n\nWhen specified, replaces the existing path of the task URL.\nSetting the path value to an empty string clears the URI path segment.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "query_override": {
                    "block": {
                      "attributes": {
                        "query_params": {
                          "computed": true,
                          "description": "The query parameters (e.g., qparam1=123\u0026qparam2=456). Default is an empty string.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "URI query.\n\nWhen specified, replaces the query part of the task URI. Setting the query value to an empty string clears the URI query segment.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "URI override.\n\nWhen specified, overrides the execution URI for all the tasks in the queue.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Modifies HTTP target for HTTP tasks.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rate_limits": {
        "block": {
          "attributes": {
            "max_burst_size": {
              "computed": true,
              "description": "The max burst size.\n\nMax burst size limits how fast tasks in queue are processed when many tasks are\nin the queue and the rate is high. This field allows the queue to have a high\nrate so processing starts shortly after a task is enqueued, but still limits\nresource usage when many tasks are enqueued in a short period of time.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_concurrent_dispatches": {
              "computed": true,
              "description": "The maximum number of concurrent tasks that Cloud Tasks allows to\nbe dispatched for this queue. After this threshold has been\nreached, Cloud Tasks stops dispatching tasks until the number of\nconcurrent requests decreases.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_dispatches_per_second": {
              "computed": true,
              "description": "The maximum rate at which tasks are dispatched from this queue.\n\nIf unspecified when the queue is created, Cloud Tasks will pick the default.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Rate limits for task dispatches.\n\nThe queue's actual dispatch rate is the result of:\n\n* Number of tasks in the queue\n* User-specified throttling: rateLimits, retryConfig, and the queue's state.\n* System throttling due to 429 (Too Many Requests) or 503 (Service\n  Unavailable) responses from the worker, high error rates, or to\n  smooth sudden large traffic spikes.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retry_config": {
        "block": {
          "attributes": {
            "max_attempts": {
              "computed": true,
              "description": "Number of attempts per task.\n\nCloud Tasks will attempt the task maxAttempts times (that is, if\nthe first attempt fails, then there will be maxAttempts - 1\nretries). Must be \u003e= -1.\n\nIf unspecified when the queue is created, Cloud Tasks will pick\nthe default.\n\n-1 indicates unlimited attempts.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_backoff": {
              "computed": true,
              "description": "A task will be scheduled for retry between minBackoff and\nmaxBackoff duration after it fails, if the queue's RetryConfig\nspecifies that the task should be retried.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_doublings": {
              "computed": true,
              "description": "The time between retries will double maxDoublings times.\n\nA task's retry interval starts at minBackoff, then doubles maxDoublings times,\nthen increases linearly, and finally retries retries at intervals of maxBackoff\nup to maxAttempts times.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_retry_duration": {
              "computed": true,
              "description": "If positive, maxRetryDuration specifies the time limit for\nretrying a failed task, measured from when the task was first\nattempted. Once maxRetryDuration time has passed and the task has\nbeen attempted maxAttempts times, no further attempts will be\nmade and the task will be deleted.\n\nIf zero, then the task age is unlimited.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_backoff": {
              "computed": true,
              "description": "A task will be scheduled for retry between minBackoff and\nmaxBackoff duration after it fails, if the queue's RetryConfig\nspecifies that the task should be retried.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Settings that determine the retry behavior.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "stackdriver_logging_config": {
        "block": {
          "attributes": {
            "sampling_ratio": {
              "description": "Specifies the fraction of operations to write to Stackdriver Logging.\nThis field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the\ndefault and means that no operations are logged.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Configuration options for writing logs to Stackdriver Logging.",
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

func GoogleCloudTasksQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudTasksQueue), &result)
	return &result
}
