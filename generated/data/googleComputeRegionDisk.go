package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionDisk = `{
  "block": {
    "attributes": {
      "async_primary_disk": {
        "computed": true,
        "description": "A nested object resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disk": "string"
            }
          ]
        ]
      },
      "create_snapshot_before_destroy": {
        "computed": true,
        "description": "If set to true, a snapshot of the disk will be created before it is destroyed.\nIf your disk is encrypted with customer managed encryption keys these will be reused for the snapshot creation.\nThe name of the snapshot by default will be '{{disk-name}}-YYYYMMDD-HHmm'",
        "description_kind": "plain",
        "type": "bool"
      },
      "create_snapshot_before_destroy_prefix": {
        "computed": true,
        "description": "This will set a custom name prefix for the snapshot that's created when the disk is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "disk_encryption_key": {
        "computed": true,
        "description": "Encrypts the disk using a customer-supplied encryption key.\n\nAfter you encrypt a disk with a customer-supplied key, you must\nprovide the same key if you use the disk later (e.g. to create a disk\nsnapshot or an image, or to attach the disk to a virtual machine).\n\nCustomer-supplied encryption keys do not protect access to metadata of\nthe disk.\n\nIf you do not provide an encryption key when creating the disk, then\nthe disk will be encrypted using an automatically generated key and\nyou do not need to provide a key to use the disk later.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_name": "string",
              "raw_key": "string",
              "rsa_encrypted_key": "string",
              "sha256": "string"
            }
          ]
        ]
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
      "guest_os_features": {
        "computed": true,
        "description": "A list of features to enable on the guest operating system.\nApplicable only for bootable disks.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "type": "string"
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
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.  Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Labels to apply to this disk.  A list of key-\u003evalue pairs.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "last_attach_timestamp": {
        "computed": true,
        "description": "Last attach timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_detach_timestamp": {
        "computed": true,
        "description": "Last detach timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "licenses": {
        "computed": true,
        "description": "Any applicable license URI.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "physical_block_size_bytes": {
        "computed": true,
        "description": "Physical block size of the persistent disk, in bytes. If not present\nin a request, a default value is used. Currently supported sizes\nare 4096 and 16384, other sizes may be added in the future.\nIf an unsupported value is requested, the error message will list\nthe supported values for the caller's project.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "A reference to the region where the disk resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica_zones": {
        "computed": true,
        "description": "URLs of the zones where the disk should be replicated to.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "Size of the persistent disk, specified in GB. You can specify this\nfield when creating a persistent disk using the sourceImage or\nsourceSnapshot parameter, or specify it alone to create an empty\npersistent disk.\n\nIf you specify this field along with sourceImage or sourceSnapshot,\nthe value of sizeGb must not be less than the size of the sourceImage\nor the size of the snapshot.",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot": {
        "computed": true,
        "description": "The source snapshot used to create this disk. You can provide this as\na partial or full URL to the resource. For example, the following are\nvalid values:\n\n* 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot'\n* 'projects/project/global/snapshots/snapshot'\n* 'global/snapshots/snapshot'\n* 'snapshot'",
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk": {
        "computed": true,
        "description": "The source disk used to create this disk. You can provide this as a partial or full URL to the resource.\nFor example, the following are valid values:\n\n* https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/disks/{disk}\n* https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/disks/{disk}\n* projects/{project}/zones/{zone}/disks/{disk}\n* projects/{project}/regions/{region}/disks/{disk}\n* zones/{zone}/disks/{disk}\n* regions/{region}/disks/{disk}",
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk_id": {
        "computed": true,
        "description": "The ID value of the disk used to create this image. This value may\nbe used to determine whether the image was taken from the current\nor a previous instance of a given disk name.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_snapshot_encryption_key": {
        "computed": true,
        "description": "The customer-supplied encryption key of the source snapshot. Required\nif the source snapshot is protected by a customer-supplied encryption\nkey.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "raw_key": "string",
              "sha256": "string"
            }
          ]
        ]
      },
      "source_snapshot_id": {
        "computed": true,
        "description": "The unique ID of the snapshot used to create this disk. This value\nidentifies the exact snapshot that was used to create this persistent\ndisk. For example, if you created the persistent disk from a snapshot\nthat was later deleted and recreated under the same name, the source\nsnapshot ID would identify the exact version of the snapshot that was\nused.",
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
      "type": {
        "computed": true,
        "description": "URL of the disk type resource describing which disk type to use to\ncreate the disk. Provide this when creating the disk.",
        "description_kind": "plain",
        "type": "string"
      },
      "users": {
        "computed": true,
        "description": "Links to the users of the disk (attached instances) in form:\nproject/zones/zone/instances/instance",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRegionDiskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionDisk), &result)
	return &result
}
