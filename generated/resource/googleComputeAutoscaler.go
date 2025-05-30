package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeAutoscaler = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
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
      "name": {
        "description": "Name of the resource. The name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target": {
        "description": "URL of the managed instance group that this autoscaler will scale.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description": "URL of the zone where the instance group resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "autoscaling_policy": {
        "block": {
          "attributes": {
            "cooldown_period": {
              "description": "The number of seconds that the autoscaler should wait before it\nstarts collecting information from a new instance. This prevents\nthe autoscaler from collecting information when the instance is\ninitializing, during which the collected usage would not be\nreliable. The default time autoscaler waits is 60 seconds.\n\nVirtual machine initialization times might vary because of\nnumerous factors. We recommend that you test how long an\ninstance may take to initialize. To do this, create an instance\nand time the startup process.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_replicas": {
              "description": "The maximum number of instances that the autoscaler can scale up\nto. This is required when creating or updating an autoscaler. The\nmaximum number of replicas should not be lower than minimal number\nof replicas.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_replicas": {
              "description": "The minimum number of replicas that the autoscaler can scale down\nto. This cannot be less than 0. If not provided, autoscaler will\nchoose a default value depending on maximum number of instances\nallowed.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "mode": {
              "description": "Defines operating mode for this policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "cpu_utilization": {
              "block": {
                "attributes": {
                  "predictive_method": {
                    "description": "Indicates whether predictive autoscaling based on CPU metric is enabled. Valid values are:\n\n- NONE (default). No predictive method is used. The autoscaler scales the group to meet current demand based on real-time metrics.\n\n- OPTIMIZE_AVAILABILITY. Predictive autoscaling improves availability by monitoring daily and weekly load patterns and scaling out ahead of anticipated demand.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target": {
                    "description": "The target CPU utilization that the autoscaler should maintain.\nMust be a float value in the range (0, 1]. If not specified, the\ndefault is 0.6.\n\nIf the CPU level is below the target utilization, the autoscaler\nscales down the number of instances until it reaches the minimum\nnumber of instances you specified or until the average CPU of\nyour instances reaches the target utilization.\n\nIf the average CPU is above the target utilization, the autoscaler\nscales up until it reaches the maximum number of instances you\nspecified or until the average utilization reaches the target\nutilization.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Defines the CPU utilization policy that allows the autoscaler to\nscale based on the average CPU utilization of a managed instance\ngroup.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "load_balancing_utilization": {
              "block": {
                "attributes": {
                  "target": {
                    "description": "Fraction of backend capacity utilization (set in HTTP(s) load\nbalancing configuration) that autoscaler should maintain. Must\nbe a positive float value. If not defined, the default is 0.8.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Configuration parameters of autoscaling based on a load balancer.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "metric": {
              "block": {
                "attributes": {
                  "filter": {
                    "description": "A filter string to be used as the filter string for\na Stackdriver Monitoring TimeSeries.list API call.\nThis filter is used to select a specific TimeSeries for\nthe purpose of autoscaling and to determine whether the metric\nis exporting per-instance or per-group data.\n\nYou can only use the AND operator for joining selectors.\nYou can only use direct equality comparison operator (=) without\nany functions for each selector.\nYou can specify the metric in both the filter string and in the\nmetric field. However, if specified in both places, the metric must\nbe identical.\n\nThe monitored resource type determines what kind of values are\nexpected for the metric. If it is a gce_instance, the autoscaler\nexpects the metric to include a separate TimeSeries for each\ninstance in a group. In such a case, you cannot filter on resource\nlabels.\n\nIf the resource type is any other value, the autoscaler expects\nthis metric to contain values that apply to the entire autoscaled\ninstance group and resource label filtering can be performed to\npoint autoscaler at the correct TimeSeries to scale upon.\nThis is called a per-group metric for the purpose of autoscaling.\n\nIf not specified, the type defaults to gce_instance.\n\nYou should provide a filter that is selective enough to pick just\none TimeSeries for the autoscaled group or for each of the instances\n(if you are using gce_instance resource type). If multiple\nTimeSeries are returned upon the query execution, the autoscaler\nwill sum their respective values to obtain its scaling value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The identifier (type) of the Stackdriver Monitoring metric.\nThe metric cannot have negative values.\n\nThe metric must have a value type of INT64 or DOUBLE.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "single_instance_assignment": {
                    "description": "If scaling is based on a per-group metric value that represents the\ntotal amount of work to be done or resource usage, set this value to\nan amount assigned for a single instance of the scaled group.\nThe autoscaler will keep the number of instances proportional to the\nvalue of this metric, the metric itself should not change value due\nto group resizing.\n\nFor example, a good metric to use with the target is\n'pubsub.googleapis.com/subscription/num_undelivered_messages'\nor a custom metric exporting the total number of requests coming to\nyour instances.\n\nA bad example would be a metric exporting an average or median\nlatency, since this value can't include a chunk assignable to a\nsingle instance, it could be better used with utilization_target\ninstead.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target": {
                    "description": "The target value of the metric that autoscaler should\nmaintain. This must be a positive value. A utilization\nmetric scales number of virtual machines handling requests\nto increase or decrease proportionally to the metric.\n\nFor example, a good metric to use as a utilizationTarget is\nwww.googleapis.com/compute/instance/network/received_bytes_count.\nThe autoscaler will work to keep this value constant for each\nof the instances.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "description": "Defines how target utilization value is expressed for a\nStackdriver Monitoring metric. Possible values: [\"GAUGE\", \"DELTA_PER_SECOND\", \"DELTA_PER_MINUTE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Configuration parameters of autoscaling based on a custom metric.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "scale_in_control": {
              "block": {
                "attributes": {
                  "time_window_sec": {
                    "description": "How long back autoscaling should look when computing recommendations\nto include directives regarding slower scale down, as described above.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "max_scaled_in_replicas": {
                    "block": {
                      "attributes": {
                        "fixed": {
                          "description": "Specifies a fixed number of VM instances. This must be a positive\ninteger.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "percent": {
                          "description": "Specifies a percentage of instances between 0 to 100%, inclusive.\nFor example, specify 80 for 80%.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "A nested object resource.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Defines scale in controls to reduce the risk of response latency\nand outages due to abrupt scale-in events",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "scaling_schedules": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "A description of a scaling schedule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disabled": {
                    "description": "A boolean value that specifies if a scaling schedule can influence autoscaler recommendations. If set to true, then a scaling schedule has no effect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "duration_sec": {
                    "description": "The duration of time intervals (in seconds) for which this scaling schedule will be running. The minimum allowed value is 300.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "min_required_replicas": {
                    "description": "Minimum number of VM instances that autoscaler will recommend in time intervals starting according to schedule.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "schedule": {
                    "description": "The start timestamps of time intervals when this scaling schedule should provide a scaling signal. This field uses the extended cron format (with an optional year field).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "time_zone": {
                    "description": "The time zone to be used when interpreting the schedule. The value of this field must be a time zone name from the tz database: http://en.wikipedia.org/wiki/Tz_database.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Scaling schedules defined for an autoscaler. Multiple schedules can be set on an autoscaler and they can overlap.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "The configuration parameters for the autoscaling algorithm. You can\ndefine one or more of the policies for an autoscaler: cpuUtilization,\ncustomMetricUtilizations, and loadBalancingUtilization.\n\nIf none of these are specified, the default will be to autoscale based\non cpuUtilization to 0.6 or 60%.",
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

func GoogleComputeAutoscalerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeAutoscaler), &result)
	return &result
}
