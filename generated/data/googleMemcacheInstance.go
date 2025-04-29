package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMemcacheInstance = `{
  "block": {
    "attributes": {
      "authorized_network": {
        "computed": true,
        "description": "The full name of the GCE network to connect the instance to.  If not provided,\n'default' will be used.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "discovery_endpoint": {
        "computed": true,
        "description": "Endpoint for Discovery API",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "A user-visible name for the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Resource labels to represent user-provided metadata.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "maintenance_policy": {
        "computed": true,
        "description": "Maintenance policy for an instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "description": "string",
              "update_time": "string",
              "weekly_maintenance_window": [
                "list",
                [
                  "object",
                  {
                    "day": "string",
                    "duration": "string",
                    "start_time": [
                      "list",
                      [
                        "object",
                        {
                          "hours": "number",
                          "minutes": "number",
                          "nanos": "number",
                          "seconds": "number"
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
      "maintenance_schedule": {
        "computed": true,
        "description": "Output only. Published maintenance schedule.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "end_time": "string",
              "schedule_deadline_time": "string",
              "start_time": "string"
            }
          ]
        ]
      },
      "memcache_full_version": {
        "computed": true,
        "description": "The full version of memcached server running on this instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "memcache_nodes": {
        "computed": true,
        "description": "Additional information about the instance state, if available.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host": "string",
              "node_id": "string",
              "port": "number",
              "state": "string",
              "zone": "string"
            }
          ]
        ]
      },
      "memcache_parameters": {
        "computed": true,
        "description": "User-specified parameters for this memcache instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "params": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "memcache_version": {
        "computed": true,
        "description": "The major version of Memcached software. If not provided, latest supported version will be used.\nCurrently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically\ndetermined by our system based on the latest supported minor version. Default value: \"MEMCACHE_1_5\" Possible values: [\"MEMCACHE_1_5\", \"MEMCACHE_1_6_15\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The resource name of the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_config": {
        "computed": true,
        "description": "Configuration for memcache nodes.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cpu_count": "number",
              "memory_size_mb": "number"
            }
          ]
        ]
      },
      "node_count": {
        "computed": true,
        "description": "Number of nodes in the memcache instance.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region of the Memcache instance. If it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_ip_range_id": {
        "computed": true,
        "description": "Contains the name of allocated IP address ranges associated with\nthe private service access connection for example, \"test-default\"\nassociated with IP range 10.0.0.0/29.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "zones": {
        "computed": true,
        "description": "Zones where memcache nodes should be provisioned.  If not\nprovided, all zones will be used.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleMemcacheInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMemcacheInstance), &result)
	return &result
}
