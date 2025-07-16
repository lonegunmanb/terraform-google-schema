package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeonpremVmwareAdminCluster = `{
  "block": {
    "attributes": {
      "annotations": {
        "computed": true,
        "description": "Annotations on the VMware Admin Cluster.\nThis field has the same restrictions as Kubernetes annotations.\nThe total size of all keys and values combined is limited to 256k.\nKey can have 2 segments: prefix (optional) and name (required),\nseparated by a slash (/).\nPrefix must be a DNS subdomain.\nName must be 63 characters or less, begin and end with alphanumerics,\nwith dashes (-), underscores (_), dots (.), and alphanumerics between.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "bootstrap_cluster_membership": {
        "computed": true,
        "description": "The bootstrap cluster this VMware admin cluster belongs to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time the cluster was created, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A human readable description of this VMware admin cluster.",
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
      "enable_advanced_cluster": {
        "computed": true,
        "description": "If set, the advanced cluster feature is enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description": "The DNS name of VMware admin cluster's API server.",
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
      "image_type": {
        "computed": true,
        "description": "The OS image type for the VMware admin cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_name": {
        "computed": true,
        "description": "The object name of the VMwareAdminCluster custom resource on the\nassociated admin cluster. This field is used to support conflicting\nnames when enrolling existing clusters to the API. When used as a part of\ncluster enrollment, this field will differ from the ID in the resource\nname. For new clusters, this field will match the user provided cluster ID\nand be visible in the last component of the resource name. It is not\nmodifiable.\nAll users should use this name to access their cluster using gkectl or\nkubectl and should expect to see the local name when viewing admin\ncluster controller logs.",
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
        "description": "The VMware admin cluster resource name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "on_prem_version": {
        "description": "The Anthos clusters on the VMware version for the admin cluster.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "If set, there are currently changes in flight to the VMware admin cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The lifecycle state of the VMware admin cluster.",
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
        "description": "The unique identifier of the VMware Admin Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time the cluster was last updated, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "addon_node": {
        "block": {
          "block_types": {
            "auto_resize_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether to enable controle plane node auto resizing.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Specifies auto resize config.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The VMware admin cluster addon node configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
          "description": "AAGConfig specifies whether to spread VMware Admin Cluster nodes across at\nleast three physical hosts in the datacenter.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "authorization": {
        "block": {
          "block_types": {
            "viewer_users": {
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
          "description": "The VMware admin cluster authorization configuration.",
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
              "description": "The number of vCPUs for the control-plane node of the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory": {
              "description": "The number of mebibytes of memory for the control-plane node of the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replicas": {
              "description": "The number of control plane nodes for this VMware admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "The VMware admin cluster control plane node configuration.",
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
                  "addons_node_port": {
                    "computed": true,
                    "description": "NodePort for add-ons server in the admin cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
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
                "attributes": {
                  "enabled": {
                    "description": "Metal LB is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Metal LB load balancers.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "vip_config": {
              "block": {
                "attributes": {
                  "addons_vip": {
                    "description": "The VIP to configure the load balancer for add-ons.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "control_plane_vip": {
                    "description": "The VIP which you previously set aside for the Kubernetes\nAPI of this VMware Admin Cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specified the VMware Load Balancer Config",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies the load balancer configuration for VMware admin cluster.",
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
              "description": "vcenter_network specifies vCenter network name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "dhcp_ip_config": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "enabled is a flag to mark if DHCP IP allocation is\nused for VMware admin clusters.",
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
            "ha_control_plane_config": {
              "block": {
                "block_types": {
                  "control_plane_ip_block": {
                    "block": {
                      "attributes": {
                        "gateway": {
                          "description": "The network gateway used by the VMware Admin Cluster.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "netmask": {
                          "description": "The netmask used by the VMware Admin Cluster.",
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
                            "description": "The node's network configurations used by the VMware Admin Cluster.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
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
                "description": "Configuration for HA admin cluster control plane.",
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
                          "description": "The network gateway used by the VMware Admin Cluster.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "netmask": {
                          "description": "The netmask used by the VMware Admin Cluster.",
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
                            "description": "The node's network configurations used by the VMware Admin Cluster.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Represents the configuration values for static IP allocation to nodes.",
                      "description_kind": "plain"
                    },
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
          "description": "The VMware admin cluster network configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "platform_config": {
        "block": {
          "attributes": {
            "bundles": {
              "computed": true,
              "description": "The list of bundles installed in the admin cluster.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "status": [
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
                    ],
                    "version": "string"
                  }
                ]
              ]
            },
            "platform_version": {
              "computed": true,
              "description": "The platform version e.g. 1.13.2.",
              "description_kind": "plain",
              "type": "string"
            },
            "required_platform_version": {
              "description": "The required platform version e.g. 1.13.1.\nIf the current platform version is lower than the target version,\nthe platform version will be updated to the target version.\nIf the target version is not installed in the platform\n(bundle versions), download the target version bundle.",
              "description_kind": "plain",
              "optional": true,
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
            }
          },
          "description": "The VMware platform configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "private_registry_config": {
        "block": {
          "attributes": {
            "address": {
              "description": "The registry address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ca_cert": {
              "description": "The CA certificate public key for private registry.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Configuration for private registry.",
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
      "vcenter": {
        "block": {
          "attributes": {
            "address": {
              "description": "The vCenter IP address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ca_cert_data": {
              "description": "Contains the vCenter CA certificate public key for SSL verification.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cluster": {
              "description": "The name of the vCenter cluster for the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_disk": {
              "description": "The name of the virtual machine disk (VMDK) for the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "datacenter": {
              "description": "The name of the vCenter datacenter for the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "datastore": {
              "description": "The name of the vCenter datastore for the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "folder": {
              "description": "The name of the vCenter folder for the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_pool": {
              "description": "The name of the vCenter resource pool for the admin cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_policy_name": {
              "description": "The name of the vCenter storage policy for the user cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Specifies vCenter config for the admin cluster.",
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

func GoogleGkeonpremVmwareAdminClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeonpremVmwareAdminCluster), &result)
	return &result
}
