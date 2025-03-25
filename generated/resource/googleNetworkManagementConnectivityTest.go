package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkManagementConnectivityTest = `{
  "block": {
    "attributes": {
      "bypass_firewall_checks": {
        "description": "Whether the analysis should skip firewall checking. Default value is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "The user-supplied description of the Connectivity Test.\nMaximum of 512 characters.",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Resource labels to represent user-provided metadata.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
      "protocol": {
        "description": "IP Protocol of the test. When not provided, \"TCP\" is assumed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "related_projects": {
        "description": "Other projects that may be relevant for reachability analysis.\nThis is applicable to scenarios where a test can cross project\nboundaries.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "round_trip": {
        "description": "Whether run analysis for the return path from destination to source.\nDefault value is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "destination": {
        "block": {
          "attributes": {
            "cloud_sql_instance": {
              "description": "A Cloud SQL instance URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "forwarding_rule": {
              "description": "Forwarding rule URI. Forwarding rules are frontends for load balancers,\nPSC endpoints, and Protocol Forwarding.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fqdn": {
              "description": "A DNS endpoint of Google Kubernetes Engine cluster control plane.\nRequires gke_master_cluster to be set, can't be used simultaneoulsly with\nip_address or network. Applicable only to destination endpoint.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gke_master_cluster": {
              "description": "A cluster URI for Google Kubernetes Engine cluster control plane.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance": {
              "description": "A Compute Engine instance URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_address": {
              "description": "The IP address of the endpoint, which can be an external or internal IP.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network": {
              "description": "A VPC network URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The IP protocol port of the endpoint. Only applicable when protocol is\nTCP or UDP.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "project_id": {
              "description": "Project ID where the endpoint is located.\nThe project ID can be derived from the URI if you provide a endpoint or\nnetwork URI.\nThe following are two cases where you may need to provide the project ID:\n1. Only the IP address is specified, and the IP address is within a Google\nCloud project.\n2. When you are using Shared VPC and the IP address that you provide is\nfrom the service project. In this case, the network that the IP address\nresides in is defined in the host project.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redis_cluster": {
              "description": "A Redis Cluster URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redis_instance": {
              "description": "A Redis Instance URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Required. Destination specification of the Connectivity Test.\n\nYou can use a combination of destination IP address, URI of a supported\nendpoint, project ID, or VPC network to identify the destination location.\n\nReachability analysis proceeds even if the destination location is\nambiguous. However, the test result might include endpoints or use a\ndestination that you don't intend to test.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "source": {
        "block": {
          "attributes": {
            "cloud_sql_instance": {
              "description": "A Cloud SQL instance URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gke_master_cluster": {
              "description": "A cluster URI for Google Kubernetes Engine cluster control plane.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance": {
              "description": "A Compute Engine instance URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_address": {
              "description": "The IP address of the endpoint, which can be an external or internal IP.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network": {
              "description": "A VPC network URI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_type": {
              "description": "Type of the network where the endpoint is located. Possible values: [\"GCP_NETWORK\", \"NON_GCP_NETWORK\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The IP protocol port of the endpoint. Only applicable when protocol is\nTCP or UDP.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "project_id": {
              "description": "Project ID where the endpoint is located.\nThe project ID can be derived from the URI if you provide a endpoint or\nnetwork URI.\nThe following are two cases where you may need to provide the project ID:\n1. Only the IP address is specified, and the IP address is within a Google\nCloud project.\n2. When you are using Shared VPC and the IP address that you provide is\nfrom the service project. In this case, the network that the IP address\nresides in is defined in the host project.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "app_engine_version": {
              "block": {
                "attributes": {
                  "uri": {
                    "description": "An App Engine service version name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "An App Engine service version.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cloud_function": {
              "block": {
                "attributes": {
                  "uri": {
                    "description": "A Cloud Function name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A Cloud Function.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cloud_run_revision": {
              "block": {
                "attributes": {
                  "uri": {
                    "description": "A Cloud Run revision URI.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A Cloud Run revision.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Required. Source specification of the Connectivity Test.\n\nYou can use a combination of source IP address, URI of a supported\nendpoint, project ID, or VPC network to identify the source location.\n\nReachability analysis might proceed even if the source location is\nambiguous. However, the test result might include endpoints or use a source\nthat you don't intend to test.",
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

func GoogleNetworkManagementConnectivityTestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkManagementConnectivityTest), &result)
	return &result
}
