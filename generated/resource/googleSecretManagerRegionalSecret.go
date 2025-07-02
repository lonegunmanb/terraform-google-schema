package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerRegionalSecret = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Custom metadata about the regional secret.\n\nAnnotations are distinct from various forms of labels. Annotations exist to allow\nclient tools to store their own state information without requiring a database.\n\nAnnotation keys must be between 1 and 63 characters long, have a UTF-8 encoding of\nmaximum 128 bytes, begin and end with an alphanumeric character ([a-z0-9A-Z]), and\nmay have dashes (-), underscores (_), dots (.), and alphanumerics in between these\nsymbols.\n\nThe total size of annotation keys and values must be less than 16KiB.\n\nAn object containing a list of \"key\": value pairs. Example:\n{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The time at which the regional secret was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "description": "Whether Terraform will be prevented from destroying the regional secret. Defaults to false.\nWhen the field is set to true in Terraform state, a 'terraform apply'\nor 'terraform destroy' that would delete the federation will fail.",
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
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "expire_time": {
        "computed": true,
        "description": "Timestamp in UTC when the regional secret is scheduled to expire. This is always provided on\noutput, regardless of what was sent on input. A timestamp in RFC3339 UTC \"Zulu\" format, with\nnanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and\n\"2014-10-02T15:01:23.045123456Z\". Only one of 'expire_time' or 'ttl' can be provided.",
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
      "labels": {
        "description": "The labels assigned to this regional secret.\n\nLabel keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,\nand must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}\n\nLabel values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,\nand must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}\n\nNo more than 64 labels can be assigned to a given resource.\n\nAn object containing a list of \"key\": value pairs. Example:\n{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the regional secret. eg us-central1",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the regional secret. Format:\n'projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret_id": {
        "description": "This must be unique within the project.",
        "description_kind": "plain",
        "required": true,
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
      "ttl": {
        "description": "The TTL for the regional secret. A duration in seconds with up to nine fractional digits,\nterminated by 's'. Example: \"3.5s\". Only one of 'ttl' or 'expire_time' can be provided.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_aliases": {
        "description": "Mapping from version alias to version name.\n\nA version alias is a string with a maximum length of 63 characters and can contain\nuppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_')\ncharacters. An alias string must start with a letter and cannot be the string\n'latest' or 'NEW'. No more than 50 aliases can be assigned to a given secret.\n\nAn object containing a list of \"key\": value pairs. Example:\n{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version_destroy_ttl": {
        "description": "Secret Version TTL after destruction request.\nThis is a part of the delayed delete feature on Secret Version.\nFor secret with versionDestroyTtl\u003e0, version destruction doesn't happen immediately\non calling destroy instead the version goes to a disabled state and\nthe actual destruction happens after this TTL expires. It must be atleast 24h.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "customer_managed_encryption": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The resource name of the Cloud KMS CryptoKey used to encrypt secret payloads.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The customer-managed encryption configuration of the regional secret.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rotation": {
        "block": {
          "attributes": {
            "next_rotation_time": {
              "description": "Timestamp in UTC at which the Secret is scheduled to rotate.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rotation_period": {
              "description": "The Duration between rotation notifications. Must be in seconds and at least 3600s (1h)\nand at most 3153600000s (100 years). If rotationPeriod is set, 'next_rotation_time' must\nbe set. 'next_rotation_time' will be advanced by this period when the service\nautomatically sends rotation notifications.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager\nwill send a Pub/Sub notification to the topics configured on the Secret. 'topics' must be\nset to configure rotation.",
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
      "topics": {
        "block": {
          "attributes": {
            "name": {
              "description": "The resource name of the Pub/Sub topic that will be published to, in the following format:\nprojects/*/topics/*. For publication to succeed, the Secret Manager Service\nAgent service account must have pubsub.publisher permissions on the topic.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A list of up to 10 Pub/Sub topics to which messages are published when control plane\noperations are called on the regional secret or its versions.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSecretManagerRegionalSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecretManagerRegionalSecret), &result)
	return &result
}
