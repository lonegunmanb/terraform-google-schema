package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInterconnectGroup = `{
  "block": {
    "attributes": {
      "configured": {
        "computed": true,
        "description": "The status of the group as configured. This has the same\nstructure as the operational field reported by the OperationalStatus\nmethod, but does not take into account the operational status of each\nresource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "topology_capability": [
                "list",
                [
                  "object",
                  {
                    "intended_capability_blockers": [
                      "list",
                      [
                        "object",
                        {
                          "blocker_type": "string",
                          "documentation_link": "string",
                          "explanation": "string",
                          "facilities": [
                            "list",
                            "string"
                          ],
                          "interconnects": [
                            "list",
                            "string"
                          ],
                          "metros": [
                            "list",
                            "string"
                          ],
                          "zones": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "supported_sla": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when you create the resource.",
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
        "description": "Name of the resource. Provided by the client when the resource is created. The name must be\n1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first\ncharacter must be a lowercase letter, and all following characters must be a dash,\nlowercase letter, or digit, except the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "physical_structure": {
        "computed": true,
        "description": "An analysis of the physical layout of Interconnects in this\ngroup. Every Interconnect in the group is shown once in this structure.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "metros": [
                "list",
                [
                  "object",
                  {
                    "facilities": [
                      "list",
                      [
                        "object",
                        {
                          "facility": "string",
                          "zones": [
                            "list",
                            [
                              "object",
                              {
                                "interconnects": [
                                  "list",
                                  "string"
                                ],
                                "zone": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "metro": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "intent": {
        "block": {
          "attributes": {
            "topology_capability": {
              "description": "The reliability the user intends this group to be capable of, in terms\nof the Interconnect product SLAs. Possible values: [\"PRODUCTION_NON_CRITICAL\", \"PRODUCTION_CRITICAL\", \"NO_SLA\", \"AVAILABILITY_SLA_UNSPECIFIED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The user's intent for this group. This is the only required field besides\nthe name that must be specified on group creation.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "interconnects": {
        "block": {
          "attributes": {
            "interconnect": {
              "description": "The URL of an Interconnect in this group. All Interconnects in the group are unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Interconnects in the InterconnectGroup. Keys are arbitrary user-specified\nstrings. Users are encouraged, but not required, to use their preferred\nformat for resource links as keys.\nNote that there are add-members and remove-members methods in gcloud.\nThe size of this map is limited by an \"Interconnects per group\" quota.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func GoogleComputeInterconnectGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInterconnectGroup), &result)
	return &result
}
