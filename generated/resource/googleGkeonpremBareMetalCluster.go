package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeonpremBareMetalCluster = `{
  "block": {
    "attributes": {
      "admin_cluster_membership": {
        "description": "The Admin Cluster this Bare Metal User Cluster belongs to.\nThis is the full resource name of the Admin Cluster's hub membership.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "annotations": {
        "description": "Annotations on the Bare Metal User Cluster.\nThis field has the same restrictions as Kubernetes annotations.\nThe total size of all keys and values combined is limited to 256k.\nKey can have 2 segments: prefix (optional) and name (required),\nseparated by a slash (/).\nPrefix must be a DNS subdomain.\nName must be 63 characters or less, begin and end with alphanumerics,\nwith dashes (-), underscores (_), dots (.), and alphanumerics between.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "bare_metal_version": {
        "description": "A human readable description of this Bare Metal User Cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time the cluster was created, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "The time the cluster was deleted, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A human readable description of this Bare Metal User Cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "endpoint": {
        "computed": true,
        "description": "The IP address name of Bare Metal User Cluster's API server.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.\nAllows clients to perform consistent read-modify-writes\nthrough optimistic concurrency control.",
        "description_kind": "plain",
        "type": "string"
      },
      "fleet": {
        "computed": true,
        "description": "Fleet related configuration.\nFleets are a Google Cloud concept for logically organizing clusters,\nletting you use and manage multi-cluster capabilities and apply\nconsistent policies across your systems.\nSee [Anthos Fleets](https://cloud.google.com/anthos/multicluster-management/fleets) for\nmore details on Anthos multi-cluster capabilities using Fleets.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "membership": "string"
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
      "local_name": {
        "computed": true,
        "description": "The object name of the Bare Metal Cluster custom resource on the\nassociated admin cluster. This field is used to support conflicting\nnames when enrolling existing clusters to the API. When used as a part of\ncluster enrollment, this field will differ from the ID in the resource\nname. For new clusters, this field will match the user provided cluster ID\nand be visible in the last component of the resource name. It is not\nmodifiable.\nAll users should use this name to access their cluster using gkectl or\nkubectl and should expect to see the local name when viewing admin\ncluster controller logs.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The bare metal cluster name.",
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
      "reconciling": {
        "computed": true,
        "description": "If set, there are currently changes in flight to the Bare Metal User Cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of this cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Specifies detailed cluster status.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "conditions": [
                "list",
                [
                  "object",
                  {
                    "last_transition_time": "string",
                    "message": "string",
                    "reason": "string",
                    "state": "string",
                    "type": "string"
                  }
                ]
              ],
              "error_message": "string"
            }
          ]
        ]
      },
      "uid": {
        "computed": true,
        "description": "The unique identifier of the Bare Metal User Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time the cluster was last updated, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "validation_check": {
        "computed": true,
        "description": "Specifies the security related settings for the Bare Metal User Cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "options": "string",
              "scenario": "string",
              "status": [
                "list",
                [
                  "object",
                  {
                    "result": [
                      "list",
                      [
                        "object",
                        {
                          "category": "string",
                          "description": "string",
                          "details": "string",
                          "options": "string",
                          "reason": "string"
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
    },
    "block_types": {
      "binary_authorization": {
        "block": {
          "attributes": {
            "evaluation_mode": {
              "description": "Mode of operation for binauthz policy evaluation. If unspecified,\ndefaults to DISABLED. Possible values: [\"DISABLED\", \"PROJECT_SINGLETON_POLICY_ENFORCE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Binary Authorization related configurations.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cluster_operations": {
        "block": {
          "attributes": {
            "enable_application_logs": {
              "description": "Whether collection of application logs/metrics should be enabled (in addition to system logs/metrics).",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Specifies the User Cluster's observability infrastructure.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "control_plane": {
        "block": {
          "block_types": {
            "api_server_args": {
              "block": {
                "attributes": {
                  "argument": {
                    "description": "The argument name as it appears on the API Server command line please make sure to remove the leading dashes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The value of the arg as it will be passed to the API Server command line.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Customizes the default API server args. Only a subset of\ncustomized flags are supported. Please refer to the API server\ndocumentation below to know the exact format:\nhttps://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "control_plane_node_pool_config": {
              "block": {
                "block_types": {
                  "node_pool_config": {
                    "block": {
                      "attributes": {
                        "labels": {
                          "computed": true,
                          "description": "The map of Kubernetes labels (key/value pairs) to be applied to\neach node. These will added in addition to any default label(s)\nthat Kubernetes may apply to the node. In case of conflict in\nlabel keys, the applied set may differ depending on the Kubernetes\nversion -- it's best to assume the behavior is undefined and\nconflicts should be avoided. For more information, including usage\nand the valid values, see:\n  - http://kubernetes.io/v1.1/docs/user-guide/labels.html\nAn object containing a list of \"key\": value pairs.\nFor example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "operating_system": {
                          "description": "Specifies the nodes operating system (default: LINUX).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "node_configs": {
                          "block": {
                            "attributes": {
                              "labels": {
                                "description": "The map of Kubernetes labels (key/value pairs) to be applied to\neach node. These will added in addition to any default label(s)\nthat Kubernetes may apply to the node. In case of conflict in\nlabel keys, the applied set may differ depending on the Kubernetes\nversion -- it's best to assume the behavior is undefined and\nconflicts should be avoided. For more information, including usage\nand the valid values, see:\n  - http://kubernetes.io/v1.1/docs/user-guide/labels.html\nAn object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "node_ip": {
                                "description": "The default IPv4 address for SSH access and Kubernetes node.\nExample: 192.168.0.1",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The list of machine addresses in the Bare Metal Node Pool.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "taints": {
                          "block": {
                            "attributes": {
                              "effect": {
                                "description": "Specifies the nodes operating system (default: LINUX). Possible values: [\"EFFECT_UNSPECIFIED\", \"PREFER_NO_SCHEDULE\", \"NO_EXECUTE\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key": {
                                "description": "Key associated with the effect.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "Value associated with the effect.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The initial taints assigned to nodes of this node pool.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The generic configuration for a node pool running the control plane.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Configures the node pool running the control plane. If specified the corresponding NodePool will be created for the cluster's control plane. The NodePool will have the same name and namespace as the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies the control plane configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "load_balancer": {
        "block": {
          "block_types": {
            "bgp_lb_config": {
              "block": {
                "attributes": {
                  "asn": {
                    "description": "BGP autonomous system number (ASN) of the cluster.\nThis field can be updated after cluster creation.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "address_pools": {
                    "block": {
                      "attributes": {
                        "addresses": {
                          "description": "The addresses that are part of this pool. Each address must be either in the CIDR form (1.2.3.0/24) or range form (1.2.3.1-1.2.3.5).",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "avoid_buggy_ips": {
                          "description": "If true, avoid using IPs ending in .0 or .255.\nThis avoids buggy consumer devices mistakenly dropping IPv4 traffic for those special IP addresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "manual_assign": {
                          "description": "If true, prevent IP addresses from being automatically assigned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "pool": {
                          "description": "The name of the address pool.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "AddressPools is a list of non-overlapping IP pools used by load balancer\ntyped services. All addresses must be routable to load balancer nodes.\nIngressVIP must be included in the pools.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "bgp_peer_configs": {
                    "block": {
                      "attributes": {
                        "asn": {
                          "description": "BGP autonomous system number (ASN) for the network that contains the\nexternal peer device.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "control_plane_nodes": {
                          "description": "The IP address of the control plane node that connects to the external\npeer.\nIf you don't specify any control plane nodes, all control plane nodes\ncan connect to the external peer. If you specify one or more IP addresses,\nonly the nodes specified participate in peering sessions.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "ip_address": {
                          "description": "The IP address of the external peer device.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The list of BGP peers that the cluster will connect to.\nAt least one peer must be configured for each control plane node.\nControl plane nodes will connect to these peers to advertise the control\nplane VIP. The Services load balancer also uses these peers by default.\nThis field can be updated after cluster creation.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "load_balancer_node_pool_config": {
                    "block": {
                      "block_types": {
                        "node_pool_config": {
                          "block": {
                            "attributes": {
                              "labels": {
                                "description": "The map of Kubernetes labels (key/value pairs) to be applied to\neach node. These will added in addition to any default label(s)\nthat Kubernetes may apply to the node. In case of conflict in\nlabel keys, the applied set may differ depending on the Kubernetes\nversion -- it's best to assume the behavior is undefined and\nconflicts should be avoided. For more information, including usage\nand the valid values, see:\n  - http://kubernetes.io/v1.1/docs/user-guide/labels.html\nAn object containing a list of \"key\": value pairs.\nFor example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "operating_system": {
                                "description": "Specifies the nodes operating system (default: LINUX).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "kubelet_config": {
                                "block": {
                                  "attributes": {
                                    "registry_burst": {
                                      "description": "The maximum size of bursty pulls, temporarily allows pulls to burst to this\nnumber, while still not exceeding registry_pull_qps.\nThe value must not be a negative number.\nUpdating this field may impact scalability by changing the amount of\ntraffic produced by image pulls.\nDefaults to 10.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "registry_pull_qps": {
                                      "description": "The limit of registry pulls per second.\nSetting this value to 0 means no limit.\nUpdating this field may impact scalability by changing the amount of\ntraffic produced by image pulls.\nDefaults to 5.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "serialize_image_pulls_disabled": {
                                      "description": "Prevents the Kubelet from pulling multiple images at a time.\nWe recommend *not* changing the default value on nodes that run docker\ndaemon with version  \u003c 1.9 or an Another Union File System (Aufs) storage\nbackend. Issue https://github.com/kubernetes/kubernetes/issues/10959 has\nmore details.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "The modifiable kubelet configurations for the baremetal machines.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "node_configs": {
                                "block": {
                                  "attributes": {
                                    "labels": {
                                      "description": "The map of Kubernetes labels (key/value pairs) to be applied to\neach node. These will added in addition to any default label(s)\nthat Kubernetes may apply to the node. In case of conflict in\nlabel keys, the applied set may differ depending on the Kubernetes\nversion -- it's best to assume the behavior is undefined and\nconflicts should be avoided. For more information, including usage\nand the valid values, see:\n  - http://kubernetes.io/v1.1/docs/user-guide/labels.html\nAn object containing a list of \"key\": value pairs.\nFor example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "node_ip": {
                                      "description": "The default IPv4 address for SSH access and Kubernetes node.\nExample: 192.168.0.1",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The list of machine addresses in the Bare Metal Node Pool.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "taints": {
                                "block": {
                                  "attributes": {
                                    "effect": {
                                      "description": "Specifies the nodes operating system (default: LINUX). Possible values: [\"EFFECT_UNSPECIFIED\", \"PREFER_NO_SCHEDULE\", \"NO_EXECUTE\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "description": "Key associated with the effect.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "Value associated with the effect.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The initial taints assigned to nodes of this node pool.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The generic configuration for a node pool running a load balancer.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies the node pool running data plane load balancing. L2 connectivity\nis required among nodes in this pool. If missing, the control plane node\npool is used for data plane load balancing.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Configuration for BGP typed load balancers.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "manual_lb_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether manual load balancing is enabled.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "A nested object resource",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "metal_lb_config": {
              "block": {
                "block_types": {
                  "address_pools": {
                    "block": {
                      "attributes": {
                        "addresses": {
                          "description": "The addresses that are part of this pool. Each address must be either in the CIDR form (1.2.3.0/24) or range form (1.2.3.1-1.2.3.5).",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "avoid_buggy_ips": {
                          "description": "If true, avoid using IPs ending in .0 or .255.\nThis avoids buggy consumer devices mistakenly dropping IPv4 traffic for those special IP addresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "manual_assign": {
                          "description": "If true, prevent IP addresses from being automatically assigned.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "pool": {
                          "description": "The name of the address pool.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "AddressPools is a list of non-overlapping IP pools used by load balancer\ntyped services. All addresses must be routable to load balancer nodes.\nIngressVIP must be included in the pools.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "load_balancer_node_pool_config": {
                    "block": {
                      "block_types": {
                        "node_pool_config": {
                          "block": {
                            "attributes": {
                              "labels": {
                                "computed": true,
                                "description": "The map of Kubernetes labels (key/value pairs) to be applied to\neach node. These will added in addition to any default label(s)\nthat Kubernetes may apply to the node. In case of conflict in\nlabel keys, the applied set may differ depending on the Kubernetes\nversion -- it's best to assume the behavior is undefined and\nconflicts should be avoided. For more information, including usage\nand the valid values, see:\n  - http://kubernetes.io/v1.1/docs/user-guide/labels.html\nAn object containing a list of \"key\": value pairs.\nFor example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "operating_system": {
                                "computed": true,
                                "description": "Specifies the nodes operating system (default: LINUX).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "node_configs": {
                                "block": {
                                  "attributes": {
                                    "labels": {
                                      "description": "The map of Kubernetes labels (key/value pairs) to be applied to\neach node. These will added in addition to any default label(s)\nthat Kubernetes may apply to the node. In case of conflict in\nlabel keys, the applied set may differ depending on the Kubernetes\nversion -- it's best to assume the behavior is undefined and\nconflicts should be avoided. For more information, including usage\nand the valid values, see:\n  - http://kubernetes.io/v1.1/docs/user-guide/labels.html\nAn object containing a list of \"key\": value pairs.\nFor example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "node_ip": {
                                      "description": "The default IPv4 address for SSH access and Kubernetes node.\nExample: 192.168.0.1",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The list of machine addresses in the Bare Metal Node Pool.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "taints": {
                                "block": {
                                  "attributes": {
                                    "effect": {
                                      "description": "Specifies the nodes operating system (default: LINUX). Possible values: [\"EFFECT_UNSPECIFIED\", \"PREFER_NO_SCHEDULE\", \"NO_EXECUTE\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "description": "Key associated with the effect.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "Value associated with the effect.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The initial taints assigned to nodes of this node pool.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The generic configuration for a node pool running a load balancer.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies the load balancer's node pool configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A nested object resource",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "port_config": {
              "block": {
                "attributes": {
                  "control_plane_load_balancer_port": {
                    "description": "The port that control plane hosted load balancers will listen on.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Specifies the load balancer ports.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "vip_config": {
              "block": {
                "attributes": {
                  "control_plane_vip": {
                    "description": "The VIP which you previously set aside for the Kubernetes API of this Bare Metal User Cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ingress_vip": {
                    "description": "The VIP which you previously set aside for ingress traffic into this Bare Metal User Cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specified the Bare Metal Load Balancer Config",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies the load balancer configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_config": {
        "block": {
          "attributes": {
            "maintenance_address_cidr_blocks": {
              "description": "All IPv4 address from these ranges will be placed into maintenance mode.\nNodes in maintenance mode will be cordoned and drained. When both of these\nare true, the \"baremetal.cluster.gke.io/maintenance\" annotation will be set\non the node resource.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Specifies the workload node configurations.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "attributes": {
            "advanced_networking": {
              "description": "Enables the use of advanced Anthos networking features, such as Bundled\nLoad Balancing with BGP or the egress NAT gateway.\nSetting configuration for advanced networking features will automatically\nset this flag.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "island_mode_cidr": {
              "block": {
                "attributes": {
                  "pod_address_cidr_blocks": {
                    "description": "All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges. This field cannot be changed after creation.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "service_address_cidr_blocks": {
                    "description": "All services in the cluster are assigned an RFC1918 IPv4 address from these ranges. This field cannot be changed after creation.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "A nested object resource",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "multiple_network_interfaces_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether to enable multiple network interfaces for your pods.\nWhen set network_config.advanced_networking is automatically\nset to true.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for multiple network interfaces.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sr_iov_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether to install the SR-IOV operator.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for SR-IOV.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Network configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "node_access_config": {
        "block": {
          "attributes": {
            "login_user": {
              "computed": true,
              "description": "LoginUser is the user name used to access node machines.\nIt defaults to \"root\" if not set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Specifies the node access related settings for the bare metal user cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_config": {
        "block": {
          "attributes": {
            "container_runtime": {
              "computed": true,
              "description": "The available runtimes that can be used to run containers in a Bare Metal User Cluster. Possible values: [\"CONTAINER_RUNTIME_UNSPECIFIED\", \"DOCKER\", \"CONTAINERD\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_pods_per_node": {
              "computed": true,
              "description": "The maximum number of pods a node can run. The size of the CIDR range\nassigned to the node will be derived from this parameter.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Specifies the workload node configurations.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "os_environment_config": {
        "block": {
          "attributes": {
            "package_repo_excluded": {
              "description": "Whether the package repo should not be included when initializing\nbare metal machines.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "OS environment related configurations.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "proxy": {
        "block": {
          "attributes": {
            "no_proxy": {
              "description": "A list of IPs, hostnames, and domains that should skip the proxy.\nFor example [\"127.0.0.1\", \"example.com\", \".corp\", \"localhost\"].",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "uri": {
              "description": "Specifies the address of your proxy server.\nFor example: http://domain\nWARNING: Do not provide credentials in the format\nof http://(username:password@)domain these will be rejected by the server.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Specifies the cluster proxy configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "security_config": {
        "block": {
          "block_types": {
            "authorization": {
              "block": {
                "block_types": {
                  "admin_users": {
                    "block": {
                      "attributes": {
                        "username": {
                          "description": "The name of the user, e.g. 'my-gcp-id@gmail.com'.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Users that will be granted the cluster-admin role on the cluster, providing full access to the cluster.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Configures user access to the Bare Metal User cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies the security related settings for the Bare Metal User Cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "storage": {
        "block": {
          "block_types": {
            "lvp_node_mounts_config": {
              "block": {
                "attributes": {
                  "path": {
                    "description": "The host machine path.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "storage_class": {
                    "description": "The StorageClass name that PVs will be created with.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specifies the config for local PersistentVolumes backed\nby mounted node disks. These disks need to be formatted and mounted by the\nuser, which can be done before or after cluster creation.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "lvp_share_config": {
              "block": {
                "attributes": {
                  "shared_path_pv_count": {
                    "description": "The number of subdirectories to create under path.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "lvp_config": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "The host machine path.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "storage_class": {
                          "description": "The StorageClass name that PVs will be created with.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Defines the machine path and storage class for the LVP Share.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies the config for local PersistentVolumes backed by\nsubdirectories in a shared filesystem. These subdirectores are\nautomatically created during cluster creation.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies the cluster storage configuration.",
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
      },
      "upgrade_policy": {
        "block": {
          "attributes": {
            "policy": {
              "description": "Specifies which upgrade policy to use. Possible values: [\"SERIAL\", \"CONCURRENT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The cluster upgrade policy.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleGkeonpremBareMetalClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeonpremBareMetalCluster), &result)
	return &result
}
