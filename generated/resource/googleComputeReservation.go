package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeReservation = `{
  "block": {
    "attributes": {
      "commitment": {
        "computed": true,
        "description": "Full or partial URL to a parent commitment. This field displays for\nreservations that are tied to a commitment.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_at_time": {
        "computed": true,
        "description": "Absolute time in future when the reservation will be auto-deleted by Compute Engine. Timestamp is represented in RFC3339 text format.\nCannot be used with delete_after_duration.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "specific_reservation_required": {
        "description": "When set to true, only VMs that target this reservation by name can\nconsume this reservation. Otherwise, it can be consumed by VMs with\naffinity for any reservation. Defaults to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description": "The status of the reservation.",
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "description": "The zone where the reservation is made.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "delete_after_duration": {
        "block": {
          "attributes": {
            "nanos": {
              "description": "Number of nanoseconds for the auto-delete duration.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "seconds": {
              "description": "Number of seconds for the auto-delete duration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Duration after which the reservation will be auto-deleted by Compute Engine. Cannot be used with delete_at_time.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "reservation_sharing_policy": {
        "block": {
          "attributes": {
            "service_share_type": {
              "computed": true,
              "description": "Sharing config for all Google Cloud services. Possible values: [\"ALLOW_ALL\", \"DISALLOW_ALL\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Sharing policy for reservations with Google Cloud managed services.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "share_settings": {
        "block": {
          "attributes": {
            "share_type": {
              "computed": true,
              "description": "Type of sharing for this shared-reservation Possible values: [\"LOCAL\", \"SPECIFIC_PROJECTS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "project_map": {
              "block": {
                "attributes": {
                  "id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The project id/number, should be same as the key of this project config in the project map.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A map of project number and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "The share setting for reservations.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "specific_reservation": {
        "block": {
          "attributes": {
            "count": {
              "description": "The number of resources that are allocated.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "in_use_count": {
              "computed": true,
              "description": "How many instances are in use.",
              "description_kind": "plain",
              "type": "number"
            },
            "source_instance_template": {
              "description": "Specifies the instance template to create the reservation. If you use this field, you must exclude the\ninstanceProperties field.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "instance_properties": {
              "block": {
                "attributes": {
                  "machine_type": {
                    "description": "The name of the machine type to reserve.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "min_cpu_platform": {
                    "computed": true,
                    "description": "The minimum CPU platform for the reservation. For example,\n'\"Intel Skylake\"'. See\nthe CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)\nfor information on available CPU platforms.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "guest_accelerators": {
                    "block": {
                      "attributes": {
                        "accelerator_count": {
                          "description": "The number of the guest accelerator cards exposed to\nthis instance.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "accelerator_type": {
                          "description": "The full or partial URL of the accelerator type to\nattach to this instance. For example:\n'projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100'\n\nIf you are creating an instance template, specify only the accelerator name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Guest accelerator type and count.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "local_ssds": {
                    "block": {
                      "attributes": {
                        "disk_size_gb": {
                          "description": "The size of the disk in base-2 GB.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "interface": {
                          "description": "The disk interface to use for attaching this disk. Default value: \"SCSI\" Possible values: [\"SCSI\", \"NVME\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The amount of local ssd to reserve with each instance. This\nreserves disks of type 'local-ssd'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The instance properties for the reservation.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Reservation for instances with specific machine shapes.",
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

func GoogleComputeReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeReservation), &result)
	return &result
}
