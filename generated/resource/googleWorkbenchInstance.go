package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleWorkbenchInstance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ.\nThe milliseconds portion (\".SSS\") is optional.",
        "description_kind": "plain",
        "type": "string"
      },
      "creator": {
        "computed": true,
        "description": "Output only. Email address of entity that sent original CreateInstance request.",
        "description_kind": "plain",
        "type": "string"
      },
      "desired_state": {
        "description": "Desired state of the Workbench Instance. Set this field to 'ACTIVE' to start the Instance, and 'STOPPED' to stop the Instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_proxy_access": {
        "description": "Optional. If true, the workbench instance will not register with the proxy.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "enable_third_party_identity": {
        "description": "Flag that specifies that a notebook can be accessed with third party\nidentity provider.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "health_info": {
        "computed": true,
        "description": "'Output only. Additional information about instance health. Example:\nhealthInfo\": { \"docker_proxy_agent_status\": \"1\", \"docker_status\": \"1\", \"jupyterlab_api_status\":\n\"-1\", \"jupyterlab_status\": \"-1\", \"updated\": \"2020-10-18 09:40:03.573409\" }'",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {}
          ]
        ]
      },
      "health_state": {
        "computed": true,
        "description": "Output only. Instance health_state.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description": "Required. User-defined unique ID of this instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_owners": {
        "description": "'Optional. Input only. The owner of this instance after creation. Format:\n'alias@example.com' Currently supports one owner only. If not specified, all of\nthe service account users of your VM instance''s service account can use the instance.\nIf specified, sets the access mode to 'Single user'. For more details, see\nhttps://cloud.google.com/vertex-ai/docs/workbench/instances/manage-access-jupyterlab'",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "labels": {
        "description": "Optional. Labels to apply to this instance. These can be later modified\nby the UpdateInstance method.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Part of 'parent'. See documentation of 'projectsId'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of this workbench instance. Format: 'projects/{project_id}/locations/{location}/instances/{instance_id}'",
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
      "proxy_uri": {
        "computed": true,
        "description": "Output only. The proxy endpoint that is used to access the Jupyter notebook.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The state of this instance.",
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
      "update_time": {
        "computed": true,
        "description": "An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ.\nThe milliseconds portion (\".SSS\") is optional.",
        "description_kind": "plain",
        "type": "string"
      },
      "upgrade_history": {
        "computed": true,
        "description": "Output only. The upgrade history of this instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action": "string",
              "container_image": "string",
              "create_time": "string",
              "framework": "string",
              "snapshot": "string",
              "state": "string",
              "target_version": "string",
              "version": "string",
              "vm_image": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "gce_setup": {
        "block": {
          "attributes": {
            "disable_public_ip": {
              "computed": true,
              "description": "Optional. If true, no external IP will be assigned to this VM instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_ip_forwarding": {
              "description": "Optional. Flag to enable ip forwarding or not, default false/off.\nhttps://cloud.google.com/vpc/docs/using-routes#canipforward",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "machine_type": {
              "computed": true,
              "description": "Optional. The machine type of the VM instance. https://cloud.google.com/compute/docs/machine-resource",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description": "Optional. Custom metadata to apply to this instance.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "tags": {
              "computed": true,
              "description": "Optional. The Compute Engine tags to add to instance (see [Tagging\ninstances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "accelerator_configs": {
              "block": {
                "attributes": {
                  "core_count": {
                    "description": "Optional. Count of cores of this accelerator.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "Optional. Type of this accelerator. Possible values: [\"NVIDIA_TESLA_P100\", \"NVIDIA_TESLA_V100\", \"NVIDIA_TESLA_P4\", \"NVIDIA_TESLA_T4\", \"NVIDIA_TESLA_A100\", \"NVIDIA_A100_80GB\", \"NVIDIA_L4\", \"NVIDIA_TESLA_T4_VWS\", \"NVIDIA_TESLA_P100_VWS\", \"NVIDIA_TESLA_P4_VWS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The hardware accelerators used on this instance. If you use accelerators, make sure that your configuration has\n[enough vCPUs and memory to support the 'machine_type' you have selected](https://cloud.google.com/compute/docs/gpus/#gpus-list).\nCurrently supports only one accelerator configuration.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "boot_disk": {
              "block": {
                "attributes": {
                  "disk_encryption": {
                    "computed": true,
                    "description": "Optional. Input only. Disk encryption method used on the boot and\ndata disks, defaults to GMEK. Possible values: [\"GMEK\", \"CMEK\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_size_gb": {
                    "computed": true,
                    "description": "Optional. The size of the boot disk in GB attached to this instance,\nup to a maximum of 64000 GB (64 TB). If not specified, this defaults to the\nrecommended value of 150GB.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_type": {
                    "computed": true,
                    "description": "Optional. Indicates the type of the disk. Possible values: [\"PD_STANDARD\", \"PD_SSD\", \"PD_BALANCED\", \"PD_EXTREME\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key": {
                    "description": "'Optional. The KMS key used to encrypt the disks, only\napplicable if disk_encryption is CMEK. Format: 'projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}'\nLearn more about using your own encryption keys.'",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The definition of a boot disk.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "container_image": {
              "block": {
                "attributes": {
                  "repository": {
                    "description": "The path to the container image repository.\nFor example: gcr.io/{project_id}/{imageName}",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tag": {
                    "description": "The tag of the container image. If not specified, this defaults to the latest tag.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Use a container image to start the workbench instance.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "data_disks": {
              "block": {
                "attributes": {
                  "disk_encryption": {
                    "computed": true,
                    "description": "Optional. Input only. Disk encryption method used on the boot\nand data disks, defaults to GMEK. Possible values: [\"GMEK\", \"CMEK\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_size_gb": {
                    "computed": true,
                    "description": "Optional. The size of the disk in GB attached to this VM instance,\nup to a maximum of 64000 GB (64 TB). If not specified, this defaults to\n100.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_type": {
                    "description": "Optional. Input only. Indicates the type of the disk. Possible values: [\"PD_STANDARD\", \"PD_SSD\", \"PD_BALANCED\", \"PD_EXTREME\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key": {
                    "description": "'Optional. The KMS key used to encrypt the disks,\nonly applicable if disk_encryption is CMEK. Format: 'projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}'\nLearn more about using your own encryption keys.'",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Data disks attached to the VM instance. Currently supports only one data disk.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "network_interfaces": {
              "block": {
                "attributes": {
                  "network": {
                    "computed": true,
                    "description": "Optional. The name of the VPC that this VM instance is in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "nic_type": {
                    "description": "Optional. The type of vNIC to be used on this interface. This\nmay be gVNIC or VirtioNet. Possible values: [\"VIRTIO_NET\", \"GVNIC\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnet": {
                    "computed": true,
                    "description": "Optional. The name of the subnet that this VM instance is in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "access_configs": {
                    "block": {
                      "attributes": {
                        "external_ip": {
                          "description": "An external IP address associated with this instance. Specify an unused\nstatic external IP address available to the project or leave this field\nundefined to use an IP from a shared ephemeral IP address pool. If you\nspecify a static external IP address, it must live in the same region as\nthe zone of the instance.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Optional. An array of configurations for this interface. Currently, only one access\nconfig, ONE_TO_ONE_NAT, is supported. If no accessConfigs specified, the\ninstance will have an external internet access through an ephemeral\nexternal IP address.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The network interfaces for the VM. Supports only one interface.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "service_accounts": {
              "block": {
                "attributes": {
                  "email": {
                    "computed": true,
                    "description": "Optional. Email address of the service account.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scopes": {
                    "computed": true,
                    "description": "Output only. The list of scopes to be made available for this\nservice account. Set by the CLH to https://www.googleapis.com/auth/cloud-platform",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The service account that serves as an identity for the VM instance. Currently supports only one service account.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "shielded_instance_config": {
              "block": {
                "attributes": {
                  "enable_integrity_monitoring": {
                    "description": "Optional. Defines whether the VM instance has integrity monitoring\nenabled. Enables monitoring and attestation of the boot integrity of the VM\ninstance. The attestation is performed against the integrity policy baseline.\nThis baseline is initially derived from the implicitly trusted boot image\nwhen the VM instance is created. Enabled by default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_secure_boot": {
                    "description": "Optional. Defines whether the VM instance has Secure Boot enabled.\nSecure Boot helps ensure that the system only runs authentic software by verifying\nthe digital signature of all boot components, and halting the boot process\nif signature verification fails. Disabled by default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_vtpm": {
                    "description": "Optional. Defines whether the VM instance has the vTPM enabled.\nEnabled by default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "A set of Shielded Instance options. See [Images using supported Shielded\nVM features](https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).\nNot all combinations are valid.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "vm_image": {
              "block": {
                "attributes": {
                  "family": {
                    "description": "Optional. Use this VM image family to find the image; the newest\nimage in this family will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Optional. Use VM image name to find the image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project": {
                    "description": "The name of the Google Cloud project that this VM image belongs to.\nFormat: {project_id}",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Definition of a custom Compute Engine virtual machine image for starting\na workbench instance with the environment installed directly on the VM.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The definition of how to configure a VM instance outside of Resources and Identity.",
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

func GoogleWorkbenchInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleWorkbenchInstance), &result)
	return &result
}
