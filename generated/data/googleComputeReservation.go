package data

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
      "description": {
        "computed": true,
        "description": "An optional description of this resource.",
        "description_kind": "plain",
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "share_settings": {
        "computed": true,
        "description": "The share setting for reservations.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "project_map": [
                "set",
                [
                  "object",
                  {
                    "id": "string",
                    "project_id": "string"
                  }
                ]
              ],
              "share_type": "string"
            }
          ]
        ]
      },
      "specific_reservation": {
        "computed": true,
        "description": "Reservation for instances with specific machine shapes.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "count": "number",
              "in_use_count": "number",
              "instance_properties": [
                "list",
                [
                  "object",
                  {
                    "guest_accelerators": [
                      "list",
                      [
                        "object",
                        {
                          "accelerator_count": "number",
                          "accelerator_type": "string"
                        }
                      ]
                    ],
                    "local_ssds": [
                      "list",
                      [
                        "object",
                        {
                          "disk_size_gb": "number",
                          "interface": "string"
                        }
                      ]
                    ],
                    "machine_type": "string",
                    "min_cpu_platform": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "specific_reservation_required": {
        "computed": true,
        "description": "When set to true, only VMs that target this reservation by name can\nconsume this reservation. Otherwise, it can be consumed by VMs with\naffinity for any reservation. Defaults to false.",
        "description_kind": "plain",
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeReservation), &result)
	return &result
}
