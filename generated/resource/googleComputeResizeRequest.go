package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeResizeRequest = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "The creation timestamp for this resize request in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resize-request.",
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
      "instance_group_manager": {
        "description": "The reference of the instance group manager this ResizeRequest is a part of.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.",
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
      "resize_by": {
        "description": "The number of instances to be created by this resize request. The group's target size will be increased by this number.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "Current state of the request.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of the request.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "error": [
                "list",
                [
                  "object",
                  {
                    "errors": [
                      "list",
                      [
                        "object",
                        {
                          "code": "string",
                          "error_details": [
                            "list",
                            [
                              "object",
                              {
                                "error_info": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "domain": "string",
                                      "metadatas": [
                                        "map",
                                        "string"
                                      ],
                                      "reason": "string"
                                    }
                                  ]
                                ],
                                "help": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "links": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "description": "string",
                                            "url": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "localized_message": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "locale": "string",
                                      "message": "string"
                                    }
                                  ]
                                ],
                                "quota_info": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "dimensions": [
                                        "map",
                                        "string"
                                      ],
                                      "future_limit": "number",
                                      "limit": "number",
                                      "limit_name": "string",
                                      "metric_name": "string",
                                      "rollout_status": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "location": "string",
                          "message": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "last_attempt": [
                "list",
                [
                  "object",
                  {
                    "error": [
                      "list",
                      [
                        "object",
                        {
                          "errors": [
                            "list",
                            [
                              "object",
                              {
                                "code": "string",
                                "error_details": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "error_info": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "domain": "string",
                                            "metadatas": [
                                              "map",
                                              "string"
                                            ],
                                            "reason": "string"
                                          }
                                        ]
                                      ],
                                      "help": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "links": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "description": "string",
                                                  "url": "string"
                                                }
                                              ]
                                            ]
                                          }
                                        ]
                                      ],
                                      "localized_message": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "locale": "string",
                                            "message": "string"
                                          }
                                        ]
                                      ],
                                      "quota_info": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "dimensions": [
                                              "map",
                                              "string"
                                            ],
                                            "future_limit": "number",
                                            "limit": "number",
                                            "limit_name": "string",
                                            "metric_name": "string",
                                            "rollout_status": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "location": "string",
                                "message": "string"
                              }
                            ]
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
      "zone": {
        "description": "The reference of the compute zone scoping this request.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "requested_run_duration": {
        "block": {
          "attributes": {
            "nanos": {
              "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are represented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "seconds": {
              "description": "Span of time at a resolution of a second. Must be from 600 to 604800 inclusive. Note: minimum and maximum allowed range for requestedRunDuration is 10 minutes (600 seconds) and 7 days(604800 seconds) correspondingly.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Requested run duration for instances that will be created by this request. At the end of the run duration instance will be deleted.",
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

func GoogleComputeResizeRequestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeResizeRequest), &result)
	return &result
}
