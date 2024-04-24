package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEdgecontainerCluster = `{
  "block": {
    "attributes": {
      "cluster_ca_certificate": {
        "computed": true,
        "description": "The PEM-encoded public certificate of the cluster's CA.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "control_plane_version": {
        "computed": true,
        "description": "The control plane release version.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time the cluster was created, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_max_pods_per_node": {
        "computed": true,
        "description": "The default maximum number of pods per node used if a maximum value is not\nspecified explicitly for a node pool in this cluster. If unspecified, the\nKubernetes default value will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "endpoint": {
        "computed": true,
        "description": "The IP address of the Kubernetes API server.",
        "description_kind": "plain",
        "type": "string"
      },
      "external_load_balancer_ipv4_address_pools": {
        "computed": true,
        "description": "Address pools for cluster data plane external load balancing.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
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
        "description": "User-defined labels for the edgecloud cluster.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maintenance_events": {
        "computed": true,
        "description": "All the maintenance events scheduled for the cluster, including the ones\nongoing, planned for the future and done in the past (up to 90 days).",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "end_time": "string",
              "operation": "string",
              "schedule": "string",
              "start_time": "string",
              "state": "string",
              "target_version": "string",
              "type": "string",
              "update_time": "string",
              "uuid": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "The GDCE cluster name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_version": {
        "computed": true,
        "description": "The lowest release version among all worker nodes. This field can be empty\nif the cluster does not have any worker nodes.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number of the Kubernetes API server.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_channel": {
        "computed": true,
        "description": "The release channel a cluster is subscribed to. Possible values: [\"RELEASE_CHANNEL_UNSPECIFIED\", \"NONE\", \"REGULAR\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Indicates the status of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_version": {
        "computed": true,
        "description": "The target cluster version. For example: \"1.5.0\".",
        "description_kind": "plain",
        "optional": true,
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
      "update_time": {
        "computed": true,
        "description": "The time the cluster was last updated, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "authorization": {
        "block": {
          "block_types": {
            "admin_users": {
              "block": {
                "attributes": {
                  "username": {
                    "description": "An active Google username.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "User that will be granted the cluster-admin role on the cluster, providing\nfull access to the cluster. Currently, this is a singular field, but will\nbe expanded to allow multiple admins in the future.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "RBAC policy that will be applied and managed by GEC.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "control_plane": {
        "block": {
          "block_types": {
            "local": {
              "block": {
                "attributes": {
                  "machine_filter": {
                    "description": "Only machines matching this filter will be allowed to host control\nplane nodes. The filtering language accepts strings like \"name=\u003cname\u003e\",\nand is documented here: [AIP-160](https://google.aip.dev/160).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_count": {
                    "computed": true,
                    "description": "The number of nodes to serve as replicas of the Control Plane.\nOnly 1 and 3 are supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "node_location": {
                    "computed": true,
                    "description": "Name of the Google Distributed Cloud Edge zones where this node pool\nwill be created. For example: 'us-central1-edge-customer-a'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shared_deployment_policy": {
                    "computed": true,
                    "description": "Policy configuration about how user applications are deployed. Possible values: [\"SHARED_DEPLOYMENT_POLICY_UNSPECIFIED\", \"ALLOWED\", \"DISALLOWED\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Local control plane configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "remote": {
              "block": {
                "attributes": {
                  "node_location": {
                    "computed": true,
                    "description": "Name of the Google Distributed Cloud Edge zones where this node pool\nwill be created. For example: 'us-central1-edge-customer-a'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Remote control plane configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The configuration of the cluster control plane.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "control_plane_encryption": {
        "block": {
          "attributes": {
            "kms_key": {
              "computed": true,
              "description": "The Cloud KMS CryptoKey e.g.\nprojects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}\nto use for protecting control plane disks. If not specified, a\nGoogle-managed key will be used instead.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_active_version": {
              "computed": true,
              "description": "The Cloud KMS CryptoKeyVersion currently in use for protecting control\nplane disks. Only applicable if kms_key is set.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_state": {
              "computed": true,
              "description": "Availability of the Cloud KMS CryptoKey. If not 'KEY_AVAILABLE', then\nnodes may go offline as they cannot access their local data. This can be\ncaused by a lack of permissions to use the key, or if the key is disabled\nor deleted.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_status": {
              "computed": true,
              "description": "Error status returned by Cloud KMS when using this key. This field may be\npopulated only if 'kms_key_state' is not 'KMS_KEY_STATE_KEY_AVAILABLE'.\nIf populated, this field contains the error status reported by Cloud KMS.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "code": "number",
                    "message": "string"
                  }
                ]
              ]
            }
          },
          "description": "Remote control plane disk encryption options. This field is only used when\nenabling CMEK support.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "fleet": {
        "block": {
          "attributes": {
            "membership": {
              "computed": true,
              "description": "The name of the managed Hub Membership resource associated to this cluster.\nMembership names are formatted as\n'projects/\u003cproject-number\u003e/locations/global/membership/\u003ccluster-id\u003e'.",
              "description_kind": "plain",
              "type": "string"
            },
            "project": {
              "description": "The name of the Fleet host project where this cluster will be registered.\nProject names are formatted as\n'projects/\u003cproject-number\u003e'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Fleet related configuration.\nFleets are a Google Cloud concept for logically organizing clusters,\nletting you use and manage multi-cluster capabilities and apply\nconsistent policies across your systems.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_policy": {
        "block": {
          "block_types": {
            "window": {
              "block": {
                "block_types": {
                  "recurring_window": {
                    "block": {
                      "attributes": {
                        "recurrence": {
                          "computed": true,
                          "description": "An RRULE (https://tools.ietf.org/html/rfc5545#section-3.8.5.3) for how\nthis window recurs. They go on for the span of time between the start and\nend time.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "window": {
                          "block": {
                            "attributes": {
                              "end_time": {
                                "computed": true,
                                "description": "The time that the window ends. The end time must take place after the\nstart time.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start_time": {
                                "computed": true,
                                "description": "The time that the window first starts.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Represents an arbitrary window of time.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents an arbitrary window of time that recurs.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies the maintenance window in which maintenance may be performed.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Cluster-wide maintenance policy configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "networking": {
        "block": {
          "attributes": {
            "cluster_ipv4_cidr_blocks": {
              "description": "All pods in the cluster are assigned an RFC1918 IPv4 address from these\nblocks. Only a single block is supported. This field cannot be changed\nafter creation.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cluster_ipv6_cidr_blocks": {
              "description": "If specified, dual stack mode is enabled and all pods in the cluster are\nassigned an IPv6 address from these blocks alongside from an IPv4\naddress. Only a single block is supported. This field cannot be changed\nafter creation.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "network_type": {
              "computed": true,
              "description": "IP addressing type of this cluster i.e. SINGLESTACK_V4 vs DUALSTACK_V4_V6.",
              "description_kind": "plain",
              "type": "string"
            },
            "services_ipv4_cidr_blocks": {
              "description": "All services in the cluster are assigned an RFC1918 IPv4 address from these\nblocks. Only a single block is supported. This field cannot be changed\nafter creation.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "services_ipv6_cidr_blocks": {
              "description": "If specified, dual stack mode is enabled and all services in the cluster are\nassigned an IPv6 address from these blocks alongside from an IPv4\naddress. Only a single block is supported. This field cannot be changed\nafter creation.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Fleet related configuration.\nFleets are a Google Cloud concept for logically organizing clusters,\nletting you use and manage multi-cluster capabilities and apply\nconsistent policies across your systems.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "system_addons_config": {
        "block": {
          "block_types": {
            "ingress": {
              "block": {
                "attributes": {
                  "disabled": {
                    "computed": true,
                    "description": "Whether Ingress is disabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ipv4_vip": {
                    "computed": true,
                    "description": "Ingress VIP.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Config for the Ingress add-on which allows customers to create an Ingress\nobject to manage external access to the servers in a cluster. The add-on\nconsists of istiod and istio-ingress.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Config that customers are allowed to define for GDCE system add-ons.",
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

func GoogleEdgecontainerClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEdgecontainerCluster), &result)
	return &result
}
