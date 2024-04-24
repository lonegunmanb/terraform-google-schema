package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeForwardingRules = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "all_ports": "bool",
              "allow_global_access": "bool",
              "allow_psc_global_access": "bool",
              "backend_service": "string",
              "base_forwarding_rule": "string",
              "creation_timestamp": "string",
              "description": "string",
              "effective_labels": [
                "map",
                "string"
              ],
              "ip_address": "string",
              "ip_protocol": "string",
              "ip_version": "string",
              "is_mirroring_collector": "bool",
              "label_fingerprint": "string",
              "labels": [
                "map",
                "string"
              ],
              "load_balancing_scheme": "string",
              "name": "string",
              "network": "string",
              "network_tier": "string",
              "no_automate_dns_zone": "bool",
              "port_range": "string",
              "ports": [
                "set",
                "string"
              ],
              "project": "string",
              "psc_connection_id": "string",
              "psc_connection_status": "string",
              "recreate_closed_psc": "bool",
              "region": "string",
              "self_link": "string",
              "service_directory_registrations": [
                "list",
                [
                  "object",
                  {
                    "namespace": "string",
                    "service": "string"
                  }
                ]
              ],
              "service_label": "string",
              "service_name": "string",
              "source_ip_ranges": [
                "list",
                "string"
              ],
              "subnetwork": "string",
              "target": "string",
              "terraform_labels": [
                "map",
                "string"
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeForwardingRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeForwardingRules), &result)
	return &result
}
