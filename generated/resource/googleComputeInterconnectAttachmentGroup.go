package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInterconnectAttachmentGroup = `{
  "block": {
    "attributes": {
      "configured": {
        "computed": true,
        "description": "The redundancy this group is configured to support. The way a\nuser queries what SLA their Attachment gets is by looking at this field of\nthe Attachment's AttachmentGroup.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_sla": [
                "list",
                [
                  "object",
                  {
                    "effective_sla": "string",
                    "intended_sla_blockers": [
                      "list",
                      [
                        "object",
                        {
                          "attachments": [
                            "list",
                            "string"
                          ],
                          "blocker_type": "string",
                          "documentation_link": "string",
                          "explanation": "string",
                          "metros": [
                            "list",
                            "string"
                          ],
                          "regions": [
                            "list",
                            "string"
                          ],
                          "zones": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
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
      "interconnect_group": {
        "description": "The URL of an InterconnectGroup that groups these Attachments'\nInterconnects. Customers do not need to set this unless directed by\nGoogle Support.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logical_structure": {
        "computed": true,
        "description": "An analysis of the logical layout of Attachments in this\ngroup. Every Attachment in the group is shown once in this structure.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "regions": [
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
                                      "attachment": [
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
                    ],
                    "region": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is created. The name must be\n1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first\ncharacter must be a lowercase letter, and all following characters must be a dash,\nlowercase letter, or digit, except the last character, which cannot be a dash.",
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
      "attachments": {
        "block": {
          "attributes": {
            "attachment": {
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
          "description": "Attachments in the AttachmentGroup. Keys are arbitrary user-specified\nstrings. Users are encouraged, but not required, to use their preferred\nformat for resource links as keys.\nNote that there are add-members and remove-members methods in gcloud.\nThe size of this map is limited by an \"Attachments per group\" quota.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "intent": {
        "block": {
          "attributes": {
            "availability_sla": {
              "description": "Which SLA the user intends this group to support. Possible values: [\"PRODUCTION_NON_CRITICAL\", \"PRODUCTION_CRITICAL\", \"NO_SLA\", \"AVAILABILITY_SLA_UNSPECIFIED\"]",
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

func GoogleComputeInterconnectAttachmentGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInterconnectAttachmentGroup), &result)
	return &result
}
