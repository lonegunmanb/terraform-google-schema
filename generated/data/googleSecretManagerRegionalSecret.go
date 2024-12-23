package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerRegionalSecret = `{
  "block": {
    "attributes": {
      "annotations": {
        "computed": true,
        "description": "Custom metadata about the regional secret.\n\nAnnotations are distinct from various forms of labels. Annotations exist to allow\nclient tools to store their own state information without requiring a database.\n\nAnnotation keys must be between 1 and 63 characters long, have a UTF-8 encoding of\nmaximum 128 bytes, begin and end with an alphanumeric character ([a-z0-9A-Z]), and\nmay have dashes (-), underscores (_), dots (.), and alphanumerics in between these\nsymbols.\n\nThe total size of annotation keys and values must be less than 16KiB.\n\nAn object containing a list of \"key\": value pairs. Example:\n{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
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
      "customer_managed_encryption": {
        "computed": true,
        "description": "The customer-managed encryption configuration of the regional secret.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_name": "string"
            }
          ]
        ]
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
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "The labels assigned to this regional secret.\n\nLabel keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,\nand must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}\n\nLabel values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,\nand must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}\n\nNo more than 64 labels can be assigned to a given resource.\n\nAn object containing a list of \"key\": value pairs. Example:\n{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rotation": {
        "computed": true,
        "description": "The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager\nwill send a Pub/Sub notification to the topics configured on the Secret. 'topics' must be\nset to configure rotation.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "next_rotation_time": "string",
              "rotation_period": "string"
            }
          ]
        ]
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
      "topics": {
        "computed": true,
        "description": "A list of up to 10 Pub/Sub topics to which messages are published when control plane\noperations are called on the regional secret or its versions.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string"
            }
          ]
        ]
      },
      "ttl": {
        "computed": true,
        "description": "The TTL for the regional secret. A duration in seconds with up to nine fractional digits,\nterminated by 's'. Example: \"3.5s\". Only one of 'ttl' or 'expire_time' can be provided.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_aliases": {
        "computed": true,
        "description": "Mapping from version alias to version name.\n\nA version alias is a string with a maximum length of 63 characters and can contain\nuppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_')\ncharacters. An alias string must start with a letter and cannot be the string\n'latest' or 'NEW'. No more than 50 aliases can be assigned to a given secret.\n\nAn object containing a list of \"key\": value pairs. Example:\n{ \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "version_destroy_ttl": {
        "computed": true,
        "description": "Secret Version TTL after destruction request.\nThis is a part of the delayed delete feature on Secret Version.\nFor secret with versionDestroyTtl\u003e0, version destruction doesn't happen immediately\non calling destroy instead the version goes to a disabled state and\nthe actual destruction happens after this TTL expires. It must be atleast 24h.",
        "description_kind": "plain",
        "type": "string"
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
