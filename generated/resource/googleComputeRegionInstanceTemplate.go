package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionInstanceTemplate = `{
  "block": {
    "attributes": {
      "can_ip_forward": {
        "description": "Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "The time at which the instance was created in RFC 3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A brief description of this resource.",
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
      "instance_description": {
        "description": "A description of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_revocation_action_type": {
        "description": "Action to be taken when a customer's encryption key is revoked. Supports \"STOP\" and \"NONE\", with \"NONE\" being the default.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs to assign to instances created from this template,\n\n\t\t\t\t**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\n\t\t\t\tPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "machine_type": {
        "description": "The machine type to create. To create a machine with a custom type (such as extended memory), format the value like custom-VCPUS-MEM_IN_MB like custom-6-20480 for 6 vCPU and 20GB of RAM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metadata": {
        "description": "Metadata key/value pairs to make available from within instances created from this template.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "metadata_fingerprint": {
        "computed": true,
        "description": "The unique fingerprint of the metadata.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata_startup_script": {
        "description": "An alternative to using the startup-script metadata key, mostly to match the compute_instance resource. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_cpu_platform": {
        "description": "Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell or Intel Skylake.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the instance template. If you leave this blank, Terraform will auto-generate a unique name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description": "Creates a unique name beginning with the specified prefix. Conflicts with name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "numeric_id": {
        "computed": true,
        "description": "The ID of the template in numeric format.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region in which the instance template is located. If it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_manager_tags": {
        "description": "A map of resource manager tags.\n\t\t\t\tResource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT \u0026 PATCH) when empty.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "resource_policies": {
        "description": "A list of self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description": "Tags to attach to the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags_fingerprint": {
        "computed": true,
        "description": "The unique fingerprint of the tags.",
        "description_kind": "plain",
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "advanced_machine_features": {
        "block": {
          "attributes": {
            "enable_nested_virtualization": {
              "description": "Whether to enable nested virtualization or not.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_uefi_networking": {
              "description": "Whether to enable UEFI networking or not.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "performance_monitoring_unit": {
              "description": "The PMU is a hardware component within the CPU core that monitors how the processor runs code. Valid values for the level of PMU are \"STANDARD\", \"ENHANCED\", and \"ARCHITECTURAL\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threads_per_core": {
              "description": "The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "turbo_mode": {
              "description": "Turbo frequency mode to use for the instance. Currently supported modes is \"ALL_CORE_MAX\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "visible_core_count": {
              "description": "The number of physical cores to expose to an instance. Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance\\'s nominal CPU count and the underlying platform\\'s SMT width.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Controls for advanced machine-related behavior features.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "confidential_instance_config": {
        "block": {
          "attributes": {
            "confidential_instance_type": {
              "description": "\n\t\t\t\t\t\t\t\tThe confidential computing technology the instance uses.\n\t\t\t\t\t\t\t\tSEV is an AMD feature. TDX is an Intel feature. One of the following\n\t\t\t\t\t\t\t\tvalues is required: SEV, SEV_SNP, TDX. If SEV_SNP, min_cpu_platform =\n\t\t\t\t\t\t\t\t\"AMD Milan\" is currently required.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_confidential_compute": {
              "description": "Defines whether the instance should have confidential compute enabled. Field will be deprecated in a future release.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail to create.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "disk": {
        "block": {
          "attributes": {
            "architecture": {
              "computed": true,
              "description": "The architecture of the image. Allowed values are ARM64 or X86_64.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auto_delete": {
              "description": "Whether or not the disk should be auto-deleted. This defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "boot": {
              "computed": true,
              "description": "Indicates that this is a boot disk.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "device_name": {
              "computed": true,
              "description": "A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance. If not specified, the server chooses a default device name to apply to this disk.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_name": {
              "description": "Name of the disk. When not provided, this defaults to the name of the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_size_gb": {
              "computed": true,
              "description": "The size of the image in gigabytes. If not specified, it will inherit the size of its base image. For SCRATCH disks, the size must be one of 375 or 3000 GB, with a default of 375 GB.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_type": {
              "computed": true,
              "description": "The Google Compute Engine disk type. Such as \"pd-ssd\", \"local-ssd\", \"pd-balanced\" or \"pd-standard\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "guest_os_features": {
              "description": "A list of features to enable on the guest operating system. Applicable only for bootable images.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "interface": {
              "computed": true,
              "description": "Specifies the disk interface to use for attaching this disk.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "description": "A set of key/value label pairs to assign to disks,",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "mode": {
              "computed": true,
              "description": "The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If you are attaching or creating a boot disk, this must read-write mode.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provisioned_iops": {
              "computed": true,
              "description": "Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. For more details, see the [Extreme persistent disk documentation](https://cloud.google.com/compute/docs/disks/extreme-persistent-disk) or the [Hyperdisk documentation](https://cloud.google.com/compute/docs/disks/hyperdisks) depending on the selected disk_type.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "provisioned_throughput": {
              "computed": true,
              "description": "Indicates how much throughput to provision for the disk, in MB/s. This sets the amount of data that can be read or written from the disk per second. Values must greater than or equal to 1. For more details, see the [Hyperdisk documentation](https://cloud.google.com/compute/docs/disks/hyperdisks).",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "resource_manager_tags": {
              "description": "A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT \u0026 PATCH) when empty.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "resource_policies": {
              "description": "A list (short name or id) of resource policies to attach to this disk. Currently a max of 1 resource policy is supported.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "source": {
              "description": "The name (not self_link) of the disk (such as those managed by google_compute_disk) to attach. ~\u003e Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_image": {
              "computed": true,
              "description": "The image from which to initialize this disk. This can be one of: the image's self_link, projects/{project}/global/images/{image}, projects/{project}/global/images/family/{family}, global/images/{image}, global/images/family/{family}, family/{family}, {project}/{family}, {project}/{image}, {family}, or {image}. ~\u003e Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_snapshot": {
              "description": "The source snapshot to create this disk. When creating\na new instance, one of initializeParams.sourceSnapshot,\ninitializeParams.sourceImage, or disks.source is\nrequired except for local SSD.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of Google Compute Engine disk, can be either \"SCRATCH\" or \"PERSISTENT\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "disk_encryption_key": {
              "block": {
                "attributes": {
                  "kms_key_self_link": {
                    "description": "The self link of the encryption key that is stored in Google Cloud KMS.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_service_account": {
                    "description": "The service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Encrypts or decrypts a disk using a customer-supplied encryption key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "source_image_encryption_key": {
              "block": {
                "attributes": {
                  "kms_key_self_link": {
                    "description": "The self link of the encryption key that is stored in\nGoogle Cloud KMS. Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_service_account": {
                    "description": "The service account being used for the encryption\nrequest for the given KMS key. If absent, the Compute\nEngine default service account is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "raw_key": {
                    "description": "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.  Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "rsa_encrypted_key": {
                    "description": "Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.  Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description": "The customer-supplied encryption key of the source\nimage. Required if the source image is protected by a\ncustomer-supplied encryption key.\n\nInstance templates do not store customer-supplied\nencryption keys, so you cannot create disks for\ninstances in a managed instance group if the source\nimages are encrypted with your own keys.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "source_snapshot_encryption_key": {
              "block": {
                "attributes": {
                  "kms_key_self_link": {
                    "description": "The self link of the encryption key that is stored in\nGoogle Cloud KMS. Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_service_account": {
                    "description": "The service account being used for the encryption\nrequest for the given KMS key. If absent, the Compute\nEngine default service account is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "raw_key": {
                    "description": "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource. Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "rsa_encrypted_key": {
                    "description": "Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.  Only one of kms_key_self_link, rsa_encrypted_key and raw_key may be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description": "The customer-supplied encryption key of the source snapshot.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Disks to attach to instances created from this template. This can be specified multiple times for multiple disks.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "guest_accelerator": {
        "block": {
          "attributes": {
            "count": {
              "description": "The number of the guest accelerator cards exposed to this instance.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "type": {
              "description": "The accelerator type resource to expose to this instance. E.g. nvidia-tesla-k80.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "List of the type and count of accelerator cards attached to the instance.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "network_interface": {
        "block": {
          "attributes": {
            "internal_ipv6_prefix_length": {
              "computed": true,
              "description": "The prefix length of the primary internal IPv6 range.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv6_access_type": {
              "computed": true,
              "description": "One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork.",
              "description_kind": "plain",
              "type": "string"
            },
            "ipv6_address": {
              "computed": true,
              "description": "An IPv6 internal network address for this network interface. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the network_interface.",
              "description_kind": "plain",
              "type": "string"
            },
            "network": {
              "computed": true,
              "description": "The name or self_link of the network to attach this interface to. Use network attribute for Legacy or Auto subnetted networks and subnetwork for custom subnetted networks.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_ip": {
              "description": "The private IP address to assign to the instance. If empty, the address will be automatically assigned.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nic_type": {
              "description": "The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET, MRDMA, and IRDMA",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "queue_count": {
              "description": "The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "stack_type": {
              "computed": true,
              "description": "The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnetwork": {
              "computed": true,
              "description": "The name of the subnetwork to attach this interface to. The subnetwork must exist in the same region this instance will be created in. Either network or subnetwork must be provided.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnetwork_project": {
              "computed": true,
              "description": "The ID of the project in which the subnetwork belongs. If it is not provided, the provider project is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "access_config": {
              "block": {
                "attributes": {
                  "nat_ip": {
                    "computed": true,
                    "description": "The IP address that will be 1:1 mapped to the instance's network ip. If not given, one will be generated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network_tier": {
                    "computed": true,
                    "description": "The networking tier used for configuring this instance template. This field can take the following values: PREMIUM, STANDARD, FIXED_STANDARD. If this field is not specified, it is assumed to be PREMIUM.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "public_ptr_domain_name": {
                    "computed": true,
                    "description": "The DNS domain name for the public PTR record.The DNS domain name for the public PTR record.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Access configurations, i.e. IPs via which this instance can be accessed via the Internet. Omit to ensure that the instance is not accessible from the Internet (this means that ssh provisioners will not work unless you are running Terraform can send traffic to the instance's network (e.g. via tunnel or because it is running on another cloud instance on that network). This block can be repeated multiple times.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "alias_ip_range": {
              "block": {
                "attributes": {
                  "ip_cidr_range": {
                    "description": "The IP CIDR range represented by this alias IP range. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. At the time of writing only a netmask (e.g. /24) may be supplied, with a CIDR format resulting in an API error.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "subnetwork_range_name": {
                    "description": "The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. If left unspecified, the primary range of the subnetwork will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "An array of alias IP ranges for this network interface. Can only be specified for network interfaces on subnet-mode networks.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ipv6_access_config": {
              "block": {
                "attributes": {
                  "external_ipv6": {
                    "computed": true,
                    "description": "The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. The field is output only, an IPv6 address from a subnetwork associated with the instance will be allocated dynamically.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "external_ipv6_prefix_length": {
                    "computed": true,
                    "description": "The prefix length of the external IPv6 range.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of this access configuration.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_tier": {
                    "description": "The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "public_ptr_domain_name": {
                    "computed": true,
                    "description": "The domain name to be used when creating DNSv6 records for the external IPv6 ranges.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Networks to attach to instances created from this template. This can be specified multiple times for multiple networks.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "network_performance_config": {
        "block": {
          "attributes": {
            "total_egress_bandwidth_tier": {
              "description": "The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configures network performance settings for the instance. If not specified, the instance will be created with its default network performance configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "reservation_affinity": {
        "block": {
          "attributes": {
            "type": {
              "description": "The type of reservation from which this instance can consume resources.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "specific_reservation": {
              "block": {
                "attributes": {
                  "key": {
                    "description": "Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description": "Corresponds to the label values of a reservation resource.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Specifies the label selector for the reservation to use.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies the reservations that this instance can consume from.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "scheduling": {
        "block": {
          "attributes": {
            "automatic_restart": {
              "description": "Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). This defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "availability_domain": {
              "description": "Specifies the availability domain, which this instance should be scheduled on.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_termination_action": {
              "description": "Specifies the action GCE should take when SPOT VM is preempted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_node_cpus": {
              "description": "Minimum number of cpus for the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "on_host_maintenance": {
              "computed": true,
              "description": "Defines the maintenance behavior for this instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preemptible": {
              "description": "Allows instance to be preempted. This defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "provisioning_model": {
              "computed": true,
              "description": "Whether the instance is spot. If this is set as SPOT.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "termination_time": {
              "description": "Specifies the timestamp, when the instance will be terminated,\nin RFC3339 text format. If specified, the instance termination action\nwill be performed at the termination time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "local_ssd_recovery_timeout": {
              "block": {
                "attributes": {
                  "nanos": {
                    "description": "Span of time that's a fraction of a second at nanosecond\nresolution. Durations less than one second are represented\nwith a 0 seconds field and a positive nanos field. Must\nbe from 0 to 999,999,999 inclusive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "seconds": {
                    "description": "Span of time at a resolution of a second.\nMust be from 0 to 315,576,000,000 inclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Specifies the maximum amount of time a Local Ssd Vm should wait while\n  recovery of the Local Ssd state is attempted. Its value should be in\n  between 0 and 168 hours with hour granularity and the default value being 1\n  hour.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "max_run_duration": {
              "block": {
                "attributes": {
                  "nanos": {
                    "description": "Span of time that's a fraction of a second at nanosecond\nresolution. Durations less than one second are represented\nwith a 0 seconds field and a positive nanos field. Must\nbe from 0 to 999,999,999 inclusive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "seconds": {
                    "description": "Span of time at a resolution of a second.\nMust be from 0 to 315,576,000,000 inclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "The timeout for new network connections to hosts.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "node_affinities": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "on_instance_stop_action": {
              "block": {
                "attributes": {
                  "discard_local_ssd": {
                    "description": "If true, the contents of any attached Local SSD disks will be discarded.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Defines the behaviour for instances with the instance_termination_action.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The scheduling strategy to use.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "service_account": {
        "block": {
          "attributes": {
            "email": {
              "computed": true,
              "description": "The service account e-mail address. If not given, the default Google Compute Engine service account is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scopes": {
              "description": "A list of service scopes. Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the cloud-platform scope.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "Service account to attach to the instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "shielded_instance_config": {
        "block": {
          "attributes": {
            "enable_integrity_monitoring": {
              "description": "Compare the most recent boot measurements to the integrity policy baseline and return a pair of pass/fail results depending on whether they match or not. Defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_secure_boot": {
              "description": "Verify the digital signature of all boot components, and halt the boot process if signature verification fails. Defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_vtpm": {
              "description": "Use a virtualized trusted platform module, which is a specialized computer chip you can use to encrypt objects like keys and certificates. Defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Enable Shielded VM on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Note: shielded_instance_config can only be used with boot images with shielded vm support.",
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleComputeRegionInstanceTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionInstanceTemplate), &result)
	return &result
}
