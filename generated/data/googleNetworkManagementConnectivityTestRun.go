package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkManagementConnectivityTestRun = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Unique name for the connectivity test.",
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
      "reachability_details": {
        "computed": true,
        "description": "Connectivity test reachability details.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "result": "string",
              "traces": [
                "list",
                [
                  "object",
                  {
                    "endpoint_info": [
                      "list",
                      [
                        "object",
                        {
                          "destination_ip": "string",
                          "destination_network_uri": "string",
                          "destination_port": "number",
                          "protocol": "string",
                          "source_agent_uri": "string",
                          "source_ip": "string",
                          "source_network_uri": "string",
                          "source_port": "number"
                        }
                      ]
                    ],
                    "forward_trace_id": "number",
                    "steps": [
                      "list",
                      [
                        "object",
                        {
                          "causes_drop": "bool",
                          "description": "string",
                          "project_id": "string",
                          "state": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "verify_time": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleNetworkManagementConnectivityTestRunSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkManagementConnectivityTestRun), &result)
	return &result
}
