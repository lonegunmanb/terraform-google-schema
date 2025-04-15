package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeImage = `{
  "block": {
    "attributes": {
      "archive_size_bytes": {
        "computed": true,
        "description": "Size of the image tar.gz archive stored in Google Cloud Storage (in\nbytes).",
        "description_kind": "plain",
        "type": "number"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_size_gb": {
        "computed": true,
        "description": "Size of the image when restored onto a persistent disk (in GB).",
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
      "family": {
        "description": "The name of the image family to which this image belongs. You can\ncreate disks by specifying an image family instead of a specific\nimage name. The image family always returns its latest image that is\nnot deprecated. The name of the image family must comply with\nRFC1035.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource. Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this Image.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "licenses": {
        "computed": true,
        "description": "Any applicable license URI.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the\nlast character, which cannot be a dash.",
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk": {
        "description": "The source disk to create this image based on.\nYou must provide either this property or the\nrawDisk.source property but not both to create an image.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_image": {
        "description": "URL of the source image used to create this image. In order to create an image, you must provide the full or partial\nURL of one of the following:\n\n* The selfLink URL\n* This property\n* The rawDisk.source URL\n* The sourceDisk URL",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_snapshot": {
        "description": "URL of the source snapshot used to create this image.\n\nIn order to create an image, you must provide the full or partial URL of one of the following:\n\n* The selfLink URL\n* This property\n* The sourceImage URL\n* The rawDisk.source URL\n* The sourceDisk URL",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_locations": {
        "computed": true,
        "description": "Cloud Storage bucket storage location of the image\n(regional or multi-regional).\nReference link: https://cloud.google.com/compute/docs/reference/rest/v1/images",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "guest_os_features": {
        "block": {
          "attributes": {
            "type": {
              "description": "The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: [\"MULTI_IP_SUBNET\", \"SECURE_BOOT\", \"SEV_CAPABLE\", \"UEFI_COMPATIBLE\", \"VIRTIO_SCSI_MULTIQUEUE\", \"WINDOWS\", \"GVNIC\", \"IDPF\", \"SEV_LIVE_MIGRATABLE\", \"SEV_SNP_CAPABLE\", \"SUSPEND_RESUME_COMPATIBLE\", \"TDX_CAPABLE\", \"SEV_LIVE_MIGRATABLE_V2\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A list of features to enable on the guest operating system.\nApplicable only for bootable images.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "image_encryption_key": {
        "block": {
          "attributes": {
            "kms_key_self_link": {
              "description": "The self link of the encryption key that is stored in Google Cloud\nKMS.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_service_account": {
              "description": "The service account being used for the encryption request for the\ngiven KMS key. If absent, the Compute Engine default service\naccount is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "rsa_encrypted_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "Encrypts the image using a customer-supplied encryption key.\n\nAfter you encrypt an image with a customer-supplied key, you must\nprovide the same key if you use the image later (e.g. to create a\ndisk from the image)",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "raw_disk": {
        "block": {
          "attributes": {
            "container_type": {
              "description": "The format used to encode and transmit the block device, which\nshould be TAR. This is just a container and transmission format\nand not a runtime format. Provided by the client when the disk\nimage is created. Default value: \"TAR\" Possible values: [\"TAR\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sha1": {
              "description": "An optional SHA1 checksum of the disk image before unpackaging.\nThis is provided by the client when the disk image is created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source": {
              "description": "The full Google Cloud Storage URL where disk storage is stored\nYou must provide either this property or the sourceDisk property\nbut not both.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The parameters of the raw disk image.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "shielded_instance_initial_state": {
        "block": {
          "block_types": {
            "dbs": {
              "block": {
                "attributes": {
                  "content": {
                    "description": "The raw content in the secure keys file.\n\nA base64-encoded string.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "file_type": {
                    "description": "The file type of source file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The Key Database (db).",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "dbxs": {
              "block": {
                "attributes": {
                  "content": {
                    "description": "The raw content in the secure keys file.\n\nA base64-encoded string.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "file_type": {
                    "description": "The file type of source file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The forbidden key database (dbx).",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "keks": {
              "block": {
                "attributes": {
                  "content": {
                    "description": "The raw content in the secure keys file.\n\nA base64-encoded string.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "file_type": {
                    "description": "The file type of source file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The Key Exchange Key (KEK).",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "pk": {
              "block": {
                "attributes": {
                  "content": {
                    "description": "The raw content in the secure keys file.\n\nA base64-encoded string.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "file_type": {
                    "description": "The file type of source file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The Platform Key (PK).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Set the secure boot keys of shielded instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_disk_encryption_key": {
        "block": {
          "attributes": {
            "kms_key_self_link": {
              "description": "The self link of the encryption key used to decrypt this resource. Also called KmsKeyName\nin the cloud console. Your project's Compute Engine System service account\n('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nSee https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_service_account": {
              "description": "The service account being used for the encryption request for the\ngiven KMS key. If absent, the Compute Engine default service\naccount is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "rsa_encrypted_key": {
              "description": "Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit\ncustomer-supplied encryption key to either encrypt or decrypt\nthis resource. You can provide either the rawKey or the rsaEncryptedKey.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "The customer-supplied encryption key of the source disk. Required if\nthe source disk is protected by a customer-supplied encryption key.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_image_encryption_key": {
        "block": {
          "attributes": {
            "kms_key_self_link": {
              "description": "The self link of the encryption key used to decrypt this resource. Also called KmsKeyName\nin the cloud console. Your project's Compute Engine System service account\n('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nSee https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_service_account": {
              "description": "The service account being used for the encryption request for the\ngiven KMS key. If absent, the Compute Engine default service\naccount is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "rsa_encrypted_key": {
              "description": "Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit\ncustomer-supplied encryption key to either encrypt or decrypt\nthis resource. You can provide either the rawKey or the rsaEncryptedKey.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "The customer-supplied encryption key of the source image. Required if\nthe source image is protected by a customer-supplied encryption key.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_snapshot_encryption_key": {
        "block": {
          "attributes": {
            "kms_key_self_link": {
              "description": "The self link of the encryption key used to decrypt this resource. Also called KmsKeyName\nin the cloud console. Your project's Compute Engine System service account\n('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nSee https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_service_account": {
              "description": "The service account being used for the encryption request for the\ngiven KMS key. If absent, the Compute Engine default service\naccount is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "rsa_encrypted_key": {
              "description": "Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit\ncustomer-supplied encryption key to either encrypt or decrypt\nthis resource. You can provide either the rawKey or the rsaEncryptedKey.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "The customer-supplied encryption key of the source snapshot. Required if\nthe source snapshot is protected by a customer-supplied encryption key.",
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

func GoogleComputeImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeImage), &result)
	return &result
}
