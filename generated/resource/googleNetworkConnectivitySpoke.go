package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkConnectivitySpoke = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The time the spoke was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of the spoke.",
        "description_kind": "plain",
        "optional": true,
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
      "group": {
        "computed": true,
        "description": "The name of the group that this spoke is associated with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hub": {
        "description": "Immutable. The URI of the hub that this spoke is attached to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Immutable. The name of the spoke. Spoke names must be unique.",
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
      "reasons": {
        "computed": true,
        "description": "The reasons for the current state in the lifecycle",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string",
              "message": "string",
              "user_details": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "Output only. The current lifecycle state of this spoke.",
        "description_kind": "plain",
        "type": "string"
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
      "unique_id": {
        "computed": true,
        "description": "Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time the spoke was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "linked_interconnect_attachments": {
        "block": {
          "attributes": {
            "include_import_ranges": {
              "description": "IP ranges allowed to be included during import from hub (does not control transit connectivity).\nThe only allowed value for now is \"ALL_IPV4_RANGES\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "site_to_site_data_transfer": {
              "description": "A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "uris": {
              "description": "The URIs of linked interconnect attachment resources",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "linked_producer_vpc_network": {
        "block": {
          "attributes": {
            "exclude_export_ranges": {
              "description": "IP ranges encompassing the subnets to be excluded from peering.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "include_export_ranges": {
              "description": "IP ranges allowed to be included from peering.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "network": {
              "description": "The URI of the Service Consumer VPC that the Producer VPC is peered with.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "peering": {
              "description": "The name of the VPC peering between the Service Consumer VPC and the Producer VPC (defined in the Tenant project) which is added to the NCC hub. This peering must be in ACTIVE state.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "producer_network": {
              "computed": true,
              "description": "The URI of the Producer VPC.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Producer VPC network that is associated with the spoke.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "linked_router_appliance_instances": {
        "block": {
          "attributes": {
            "include_import_ranges": {
              "description": "IP ranges allowed to be included during import from hub (does not control transit connectivity).\nThe only allowed value for now is \"ALL_IPV4_RANGES\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "site_to_site_data_transfer": {
              "description": "A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "instances": {
              "block": {
                "attributes": {
                  "ip_address": {
                    "description": "The IP address on the VM to use for peering.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "virtual_machine": {
                    "description": "The URI of the virtual machine resource",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The list of router appliance instances",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The URIs of linked Router appliance resources",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "linked_vpc_network": {
        "block": {
          "attributes": {
            "exclude_export_ranges": {
              "description": "IP ranges encompassing the subnets to be excluded from peering.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "include_export_ranges": {
              "description": "IP ranges allowed to be included from peering.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "uri": {
              "description": "The URI of the VPC network resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "VPC network that is associated with the spoke.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "linked_vpn_tunnels": {
        "block": {
          "attributes": {
            "include_import_ranges": {
              "description": "IP ranges allowed to be included during import from hub (does not control transit connectivity).\nThe only allowed value for now is \"ALL_IPV4_RANGES\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "site_to_site_data_transfer": {
              "description": "A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "uris": {
              "description": "The URIs of linked VPN tunnel resources.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "The URIs of linked VPN tunnel resources",
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

func GoogleNetworkConnectivitySpokeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkConnectivitySpoke), &result)
	return &result
}
