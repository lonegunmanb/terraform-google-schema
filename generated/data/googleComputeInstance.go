package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstance = `{
  "block": {
    "attributes": {
      "advanced_machine_features": {
        "computed": true,
        "description": "Controls for advanced machine-related behavior features.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_nested_virtualization": "bool",
              "enable_uefi_networking": "bool",
              "performance_monitoring_unit": "string",
              "threads_per_core": "number",
              "turbo_mode": "string",
              "visible_core_count": "number"
            }
          ]
        ]
      },
      "allow_stopping_for_update": {
        "computed": true,
        "description": "If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail.",
        "description_kind": "plain",
        "type": "bool"
      },
      "attached_disk": {
        "computed": true,
        "description": "List of disks attached to the instance",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "device_name": "string",
              "disk_encryption_key_raw": "string",
              "disk_encryption_key_sha256": "string",
              "kms_key_self_link": "string",
              "mode": "string",
              "source": "string"
            }
          ]
        ]
      },
      "boot_disk": {
        "computed": true,
        "description": "The boot disk for the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_delete": "bool",
              "device_name": "string",
              "disk_encryption_key_raw": "string",
              "disk_encryption_key_sha256": "string",
              "initialize_params": [
                "list",
                [
                  "object",
                  {
                    "enable_confidential_compute": "bool",
                    "image": "string",
                    "labels": [
                      "map",
                      "string"
                    ],
                    "provisioned_iops": "number",
                    "provisioned_throughput": "number",
                    "resource_manager_tags": [
                      "map",
                      "string"
                    ],
                    "resource_policies": [
                      "list",
                      "string"
                    ],
                    "size": "number",
                    "storage_pool": "string",
                    "type": "string"
                  }
                ]
              ],
              "interface": "string",
              "kms_key_self_link": "string",
              "mode": "string",
              "source": "string"
            }
          ]
        ]
      },
      "can_ip_forward": {
        "computed": true,
        "description": "Whether sending and receiving of packets with non-matching source or destination IPs is allowed.",
        "description_kind": "plain",
        "type": "bool"
      },
      "confidential_instance_config": {
        "computed": true,
        "description": "The Confidential VM config being used by the instance.  on_host_maintenance has to be set to TERMINATE or this will fail to create.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "confidential_instance_type": "string",
              "enable_confidential_compute": "bool"
            }
          ]
        ]
      },
      "cpu_platform": {
        "computed": true,
        "description": "The CPU platform used by this instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "current_status": {
        "computed": true,
        "description": "\n\t\t\t\t\tCurrent status of the instance.\n\t\t\t\t\tThis could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED.\n\t\t\t\t\tFor more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether deletion protection is enabled on this instance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A brief description of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "desired_status": {
        "computed": true,
        "description": "Desired status of the instance. Either \"RUNNING\", \"SUSPENDED\" or \"TERMINATED\".",
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
      "enable_display": {
        "computed": true,
        "description": "Whether the instance has virtual displays enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "guest_accelerator": {
        "computed": true,
        "description": "List of the type and count of accelerator cards attached to the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "count": "number",
              "type": "string"
            }
          ]
        ]
      },
      "hostname": {
        "computed": true,
        "description": "A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.",
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
        "computed": true,
        "description": "The server-assigned unique identifier of this instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_revocation_action_type": {
        "computed": true,
        "description": "Action to be taken when a customer's encryption key is revoked. Supports \"STOP\" and \"NONE\", with \"NONE\" being the default.",
        "description_kind": "plain",
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The unique fingerprint of the labels.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "A set of key/value label pairs assigned to the instance.\n\n\t\t\t\t**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\n\t\t\t\tPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "machine_type": {
        "computed": true,
        "description": "The machine type to create.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description": "Metadata key/value pairs made available within the instance.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "Metadata startup scripts made available within the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "min_cpu_platform": {
        "computed": true,
        "description": "The minimum CPU platform specified for the VM instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the instance. One of name or self_link must be provided.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_interface": {
        "computed": true,
        "description": "The networks attached to the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_config": [
                "list",
                [
                  "object",
                  {
                    "nat_ip": "string",
                    "network_tier": "string",
                    "public_ptr_domain_name": "string"
                  }
                ]
              ],
              "alias_ip_range": [
                "list",
                [
                  "object",
                  {
                    "ip_cidr_range": "string",
                    "subnetwork_range_name": "string"
                  }
                ]
              ],
              "internal_ipv6_prefix_length": "number",
              "ipv6_access_config": [
                "list",
                [
                  "object",
                  {
                    "external_ipv6": "string",
                    "external_ipv6_prefix_length": "string",
                    "name": "string",
                    "network_tier": "string",
                    "public_ptr_domain_name": "string"
                  }
                ]
              ],
              "ipv6_access_type": "string",
              "ipv6_address": "string",
              "name": "string",
              "network": "string",
              "network_attachment": "string",
              "network_ip": "string",
              "nic_type": "string",
              "queue_count": "number",
              "stack_type": "string",
              "subnetwork": "string",
              "subnetwork_project": "string"
            }
          ]
        ]
      },
      "network_performance_config": {
        "computed": true,
        "description": "Configures network performance settings for the instance. If not specified, the instance will be created with its default network performance configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "total_egress_bandwidth_tier": "string"
            }
          ]
        ]
      },
      "params": {
        "computed": true,
        "description": "Stores additional params passed with the request, but not persisted as part of resource payload.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_manager_tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "project": {
        "description": "The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither self_link nor project are provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reservation_affinity": {
        "computed": true,
        "description": "Specifies the reservations that this instance can consume from.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "specific_reservation": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "values": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "type": "string"
            }
          ]
        ]
      },
      "resource_policies": {
        "computed": true,
        "description": "A list of self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "scheduling": {
        "computed": true,
        "description": "The scheduling strategy being used by the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_restart": "bool",
              "availability_domain": "number",
              "instance_termination_action": "string",
              "local_ssd_recovery_timeout": [
                "list",
                [
                  "object",
                  {
                    "nanos": "number",
                    "seconds": "number"
                  }
                ]
              ],
              "max_run_duration": [
                "list",
                [
                  "object",
                  {
                    "nanos": "number",
                    "seconds": "number"
                  }
                ]
              ],
              "min_node_cpus": "number",
              "node_affinities": [
                "set",
                [
                  "object",
                  {
                    "key": "string",
                    "operator": "string",
                    "values": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "on_host_maintenance": "string",
              "on_instance_stop_action": [
                "list",
                [
                  "object",
                  {
                    "discard_local_ssd": "bool"
                  }
                ]
              ],
              "preemptible": "bool",
              "provisioning_model": "string"
            }
          ]
        ]
      },
      "scratch_disk": {
        "computed": true,
        "description": "The scratch disks attached to the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "device_name": "string",
              "interface": "string",
              "size": "number"
            }
          ]
        ]
      },
      "self_link": {
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description": "The service account to attach to the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "email": "string",
              "scopes": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "shielded_instance_config": {
        "computed": true,
        "description": "The shielded vm config being used by the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_integrity_monitoring": "bool",
              "enable_secure_boot": "bool",
              "enable_vtpm": "bool"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description": "The list of tags attached to the instance.",
        "description_kind": "plain",
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
      },
      "zone": {
        "description": "The zone of the instance. If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInstance), &result)
	return &result
}
