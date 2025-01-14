package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeonpremVmwareCluster = `{
  "block": {
    "attributes": {
      "admin_cluster_membership": {
        "description": "The admin cluster this VMware User Cluster belongs to.\nThis is the full resource name of the admin cluster's hub membership.\nIn the future, references to other resource types might be allowed if\nadmin clusters are modeled as their own resources.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "annotations": {
        "description": "Annotations on the VMware User Cluster.\nThis field has the same restrictions as Kubernetes annotations.\nThe total size of all keys and values combined is limited to 256k.\nKey can have 2 segments: prefix (optional) and name (required),\nseparated by a slash (/).\nPrefix must be a DNS subdomain.\nName must be 63 characters or less, begin and end with alphanumerics,\nwith dashes (-), underscores (_), dots (.), and alphanumerics between.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The time at which VMware User Cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "The time at which VMware User Cluster was deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A human readable description of this VMware User Cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_bundled_ingress": {
        "description": "Disable bundled ingress.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "enable_control_plane_v2": {
        "description": "Enable control plane V2. Default to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description": "The DNS name of VMware User Cluster's API server.",
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
        "description": "Fleet configuration for the cluster.",
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
        "description": "The object name of the VMware OnPremUserCluster custom resource on the\nassociated admin cluster. This field is used to support conflicting\nnames when enrolling existing clusters to the API. When used as a part of\ncluster enrollment, this field will differ from the ID in the resource\nname. For new clusters, this field will match the user provided cluster ID\nand be visible in the last component of the resource name. It is not\nmodifiable.\n\nAll users should use this name to access their cluster using gkectl or\nkubectl and should expect to see the local name when viewing admin\ncluster controller logs.",
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
        "description": "The VMware cluster name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "on_prem_version": {
        "description": "The Anthos clusters on the VMware version for your user cluster.",
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
        "description": "If set, there are currently changes in flight to the VMware User Cluster.",
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
        "description": "ResourceStatus representing detailed cluster state.",
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
        "description": "The unique identifier of the VMware User Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which VMware User Cluster was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "validation_check": {
        "computed": true,
        "description": "ValidationCheck represents the result of the preflight check job.",
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
      },
      "vm_tracking_enabled": {
        "computed": true,
        "description": "Enable VM tracking.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "anti_affinity_groups": {
        "block": {
          "attributes": {
            "aag_config_disabled": {
              "description": "Spread nodes across at least three physical hosts (requires at least three\nhosts).\nEnabled by default.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "AAGConfig specifies whether to spread VMware User Cluster nodes across at\nleast three physical hosts in the datacenter.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
                "description": "Users that will be granted the cluster-admin role on the cluster, providing\nfull access to the cluster.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "RBAC policy that will be applied and managed by GKE On-Prem.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "auto_repair_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether auto repair is enabled.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Configuration for auto repairing.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "control_plane_node": {
        "block": {
          "attributes": {
            "cpus": {
              "description": "The number of CPUs for each admin cluster node that serve as control planes\nfor this VMware User Cluster. (default: 4 CPUs)",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory": {
              "description": "The megabytes of memory for each admin cluster node that serves as a\ncontrol plane for this VMware User Cluster (default: 8192 MB memory).",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replicas": {
              "description": "The number of control plane nodes for this VMware User Cluster.\n(default: 1 replica).",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "vsphere_config": {
              "computed": true,
              "description": "Vsphere-specific config.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "datastore": "string",
                    "storage_policy_name": "string"
                  }
                ]
              ]
            }
          },
          "block_types": {
            "auto_resize_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether to enable control plane node auto resizing.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "AutoResizeConfig provides auto resizing configurations.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "VMware User Cluster control plane nodes must have either 1 or 3 replicas.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "dataplane_v2": {
        "block": {
          "attributes": {
            "advanced_networking": {
              "description": "Enable advanced networking which requires dataplane_v2_enabled to be set true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dataplane_v2_enabled": {
              "description": "Enables Dataplane V2.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "windows_dataplane_v2_enabled": {
              "description": "Enable Dataplane V2 for clusters with Windows nodes.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "VmwareDataplaneV2Config specifies configuration for Dataplane V2.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "load_balancer": {
        "block": {
          "block_types": {
            "f5_config": {
              "block": {
                "attributes": {
                  "address": {
                    "description": "The load balancer's IP address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partition": {
                    "description": "he preexisting partition to be used by the load balancer. T\nhis partition is usually created for the admin cluster for example:\n'my-f5-admin-partition'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "snat_pool": {
                    "computed": true,
                    "description": "The pool name. Only necessary, if using SNAT.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Configuration for F5 Big IP typed load balancers.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "manual_lb_config": {
              "block": {
                "attributes": {
                  "control_plane_node_port": {
                    "computed": true,
                    "description": "NodePort for control plane service. The Kubernetes API server in the admin\ncluster is implemented as a Service of type NodePort (ex. 30968).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ingress_http_node_port": {
                    "computed": true,
                    "description": "NodePort for ingress service's http. The ingress service in the admin\ncluster is implemented as a Service of type NodePort (ex. 32527).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ingress_https_node_port": {
                    "computed": true,
                    "description": "NodePort for ingress service's https. The ingress service in the admin\ncluster is implemented as a Service of type NodePort (ex. 30139).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "konnectivity_server_node_port": {
                    "computed": true,
                    "description": "NodePort for konnectivity server service running as a sidecar in each\nkube-apiserver pod (ex. 30564).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Manually configured load balancers.",
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
                          "description": "The addresses that are part of this pool. Each address\nmust be either in the CIDR form (1.2.3.0/24) or range\nform (1.2.3.1-1.2.3.5).",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "avoid_buggy_ips": {
                          "computed": true,
                          "description": "If true, avoid using IPs ending in .0 or .255.\nThis avoids buggy consumer devices mistakenly dropping IPv4 traffic for\nthose special IP addresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "manual_assign": {
                          "computed": true,
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
                  }
                },
                "description": "Configuration for MetalLB typed load balancers.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "vip_config": {
              "block": {
                "attributes": {
                  "control_plane_vip": {
                    "description": "The VIP which you previously set aside for the Kubernetes API of this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ingress_vip": {
                    "description": "The VIP which you previously set aside for ingress traffic into this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The VIPs used by the load balancer.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Load Balancer configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "attributes": {
            "pod_address_cidr_blocks": {
              "description": "All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges.\nOnly a single range is supported. This field cannot be changed after creation.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "service_address_cidr_blocks": {
              "description": "All services in the cluster are assigned an RFC1918 IPv4 address\nfrom these ranges. Only a single range is supported.. This field\ncannot be changed after creation.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "vcenter_network": {
              "computed": true,
              "description": "vcenter_network specifies vCenter network name. Inherited from the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "control_plane_v2_config": {
              "block": {
                "block_types": {
                  "control_plane_ip_block": {
                    "block": {
                      "attributes": {
                        "gateway": {
                          "description": "The network gateway used by the VMware User Cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "netmask": {
                          "description": "The netmask used by the VMware User Cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "ips": {
                          "block": {
                            "attributes": {
                              "hostname": {
                                "computed": true,
                                "description": "Hostname of the machine. VM's name will be used if this field is empty.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ip": {
                                "description": "IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The node's network configurations used by the VMware User Cluster.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Static IP addresses for the control plane nodes.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Configuration for control plane V2 mode.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "dhcp_ip_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "enabled is a flag to mark if DHCP IP allocation is\nused for VMware user clusters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration settings for a DHCP IP configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "host_config": {
              "block": {
                "attributes": {
                  "dns_search_domains": {
                    "description": "DNS search domains.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "dns_servers": {
                    "description": "DNS servers.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ntp_servers": {
                    "description": "NTP servers.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Represents common network settings irrespective of the host's IP address.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "static_ip_config": {
              "block": {
                "block_types": {
                  "ip_blocks": {
                    "block": {
                      "attributes": {
                        "gateway": {
                          "description": "The network gateway used by the VMware User Cluster.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "netmask": {
                          "description": "The netmask used by the VMware User Cluster.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "ips": {
                          "block": {
                            "attributes": {
                              "hostname": {
                                "computed": true,
                                "description": "Hostname of the machine. VM's name will be used if this field is empty.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ip": {
                                "description": "IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The node's network configurations used by the VMware User Cluster.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents the configuration values for static IP allocation to nodes.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Configuration settings for a static IP configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The VMware User Cluster network configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "storage": {
        "block": {
          "attributes": {
            "vsphere_csi_disabled": {
              "description": "Whether or not to deploy vSphere CSI components in the VMware User Cluster.\nEnabled by default.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Storage configuration.",
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
      },
      "upgrade_policy": {
        "block": {
          "attributes": {
            "control_plane_only": {
              "description": "Controls whether the upgrade applies to the control plane only.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Specifies upgrade policy for the cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "vcenter": {
        "block": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The vCenter IP address.",
              "description_kind": "plain",
              "type": "string"
            },
            "ca_cert_data": {
              "computed": true,
              "description": "Contains the vCenter CA certificate public key for SSL verification.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cluster": {
              "computed": true,
              "description": "The name of the vCenter cluster for the user cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "datacenter": {
              "computed": true,
              "description": "The name of the vCenter datacenter for the user cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "datastore": {
              "computed": true,
              "description": "The name of the vCenter datastore for the user cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "folder": {
              "computed": true,
              "description": "The name of the vCenter folder for the user cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_pool": {
              "computed": true,
              "description": "The name of the vCenter resource pool for the user cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_policy_name": {
              "computed": true,
              "description": "The name of the vCenter storage policy for the user cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "VmwareVCenterConfig specifies vCenter config for the user cluster.\nInherited from the admin cluster.",
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

func GoogleGkeonpremVmwareClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeonpremVmwareCluster), &result)
	return &result
}
