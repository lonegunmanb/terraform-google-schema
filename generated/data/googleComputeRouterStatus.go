package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterStatus = `{
  "block": {
    "attributes": {
      "best_routes": {
        "computed": true,
        "description": "Best routes for this router's network.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "as_paths": [
                "list",
                [
                  "object",
                  {
                    "as_lists": [
                      "list",
                      "number"
                    ],
                    "path_segment_type": "string"
                  }
                ]
              ],
              "creation_timestamp": "string",
              "description": "string",
              "dest_range": "string",
              "name": "string",
              "network": "string",
              "next_hop_gateway": "string",
              "next_hop_hub": "string",
              "next_hop_ilb": "string",
              "next_hop_instance": "string",
              "next_hop_instance_zone": "string",
              "next_hop_inter_region_cost": "string",
              "next_hop_ip": "string",
              "next_hop_med": "string",
              "next_hop_network": "string",
              "next_hop_origin": "string",
              "next_hop_peering": "string",
              "next_hop_vpn_tunnel": "string",
              "params": [
                "list",
                [
                  "object",
                  {
                    "resource_manager_tags": [
                      "map",
                      "string"
                    ]
                  }
                ]
              ],
              "priority": "number",
              "project": "string",
              "route_status": "string",
              "route_type": "string",
              "self_link": "string",
              "tags": [
                "set",
                "string"
              ],
              "warnings": [
                "list",
                [
                  "object",
                  {
                    "code": "string",
                    "data": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "message": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "best_routes_for_router": {
        "computed": true,
        "description": "Best routes learned by this router.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "as_paths": [
                "list",
                [
                  "object",
                  {
                    "as_lists": [
                      "list",
                      "number"
                    ],
                    "path_segment_type": "string"
                  }
                ]
              ],
              "creation_timestamp": "string",
              "description": "string",
              "dest_range": "string",
              "name": "string",
              "network": "string",
              "next_hop_gateway": "string",
              "next_hop_hub": "string",
              "next_hop_ilb": "string",
              "next_hop_instance": "string",
              "next_hop_instance_zone": "string",
              "next_hop_inter_region_cost": "string",
              "next_hop_ip": "string",
              "next_hop_med": "string",
              "next_hop_network": "string",
              "next_hop_origin": "string",
              "next_hop_peering": "string",
              "next_hop_vpn_tunnel": "string",
              "params": [
                "list",
                [
                  "object",
                  {
                    "resource_manager_tags": [
                      "map",
                      "string"
                    ]
                  }
                ]
              ],
              "priority": "number",
              "project": "string",
              "route_status": "string",
              "route_type": "string",
              "self_link": "string",
              "tags": [
                "set",
                "string"
              ],
              "warnings": [
                "list",
                [
                  "object",
                  {
                    "code": "string",
                    "data": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "message": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the router to query.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "URI of the network to which this router belongs.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description": "Project ID of the target router.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region of the target router.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRouterStatusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterStatus), &result)
	return &result
}
