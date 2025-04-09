package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageBucketObject = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the containing bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cache_control": {
        "description": "Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content": {
        "computed": true,
        "description": "Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "content_disposition": {
        "description": "Content-Disposition of the object data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_encoding": {
        "description": "Content-Encoding of the object data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_language": {
        "description": "Content-Language of the object data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description": "Content-Type of the object data. Defaults to \"application/octet-stream\" or \"text/plain; charset=utf-8\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "crc32c": {
        "computed": true,
        "description": "Base 64 CRC32 hash of the uploaded data.",
        "description_kind": "plain",
        "type": "string"
      },
      "detect_md5hash": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_based_hold": {
        "description": "Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "generation": {
        "computed": true,
        "description": "The content generation of this object. Used for object versioning and soft delete.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "computed": true,
        "description": "Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName value, if any.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "md5hash": {
        "computed": true,
        "description": "Base 64 MD5 hash of the uploaded data.",
        "description_kind": "plain",
        "type": "string"
      },
      "md5hexhash": {
        "computed": true,
        "description": "Hex value of md5hash",
        "description_kind": "plain",
        "type": "string"
      },
      "media_link": {
        "computed": true,
        "description": "A url reference to download this object.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "description": "User-provided metadata, in key/value pairs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name of the object. If you're interpolating the name of this object, see output_name instead.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_name": {
        "computed": true,
        "description": "The name of the object. Use this field in interpolations with google_storage_object_acl to recreate google_storage_object_acl resources when your google_storage_bucket_object is recreated.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "A url reference to this object.",
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "description": "A path to the data you want to upload. Must be defined if content is not.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_class": {
        "computed": true,
        "description": "The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "temporary_hold": {
        "description": "Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "customer_encryption": {
        "block": {
          "attributes": {
            "encryption_algorithm": {
              "description": "The encryption algorithm. Default: AES256",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_key": {
              "description": "Base64 encoded customer supplied encryption key.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "Encryption key; encoded using base64.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention": {
        "block": {
          "attributes": {
            "mode": {
              "description": "The object retention mode. Supported values include: \"Unlocked\", \"Locked\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "retain_until_time": {
              "description": "Time in RFC 3339 (e.g. 2030-01-01T02:03:04Z) until which object retention protects this object.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Object level retention configuration.",
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

func GoogleStorageBucketObjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageBucketObject), &result)
	return &result
}
