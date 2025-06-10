package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePubsubTopic = `{
  "block": {
    "attributes": {
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
      "ingestion_data_source_settings": {
        "computed": true,
        "description": "Settings for ingestion from a data source into this topic.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aws_kinesis": [
                "list",
                [
                  "object",
                  {
                    "aws_role_arn": "string",
                    "consumer_arn": "string",
                    "gcp_service_account": "string",
                    "stream_arn": "string"
                  }
                ]
              ],
              "aws_msk": [
                "list",
                [
                  "object",
                  {
                    "aws_role_arn": "string",
                    "cluster_arn": "string",
                    "gcp_service_account": "string",
                    "topic": "string"
                  }
                ]
              ],
              "azure_event_hubs": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "event_hub": "string",
                    "gcp_service_account": "string",
                    "namespace": "string",
                    "resource_group": "string",
                    "subscription_id": "string",
                    "tenant_id": "string"
                  }
                ]
              ],
              "cloud_storage": [
                "list",
                [
                  "object",
                  {
                    "avro_format": [
                      "list",
                      [
                        "object",
                        {}
                      ]
                    ],
                    "bucket": "string",
                    "match_glob": "string",
                    "minimum_object_create_time": "string",
                    "pubsub_avro_format": [
                      "list",
                      [
                        "object",
                        {}
                      ]
                    ],
                    "text_format": [
                      "list",
                      [
                        "object",
                        {
                          "delimiter": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "confluent_cloud": [
                "list",
                [
                  "object",
                  {
                    "bootstrap_server": "string",
                    "cluster_id": "string",
                    "gcp_service_account": "string",
                    "identity_pool_id": "string",
                    "topic": "string"
                  }
                ]
              ],
              "platform_logs_settings": [
                "list",
                [
                  "object",
                  {
                    "severity": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "kms_key_name": {
        "computed": true,
        "description": "The resource name of the Cloud KMS CryptoKey to be used to protect access\nto messages published on this topic. Your project's PubSub service account\n('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nThe expected format is 'projects/*/locations/*/keyRings/*/cryptoKeys/*'",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "A set of key/value label pairs to assign to this Topic.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "message_retention_duration": {
        "computed": true,
        "description": "Indicates the minimum duration to retain a message after it is published\nto the topic. If this field is set, messages published to the topic in\nthe last messageRetentionDuration are always available to subscribers.\nFor instance, it allows any attached subscription to seek to a timestamp\nthat is up to messageRetentionDuration in the past. If this field is not\nset, message retention is controlled by settings on individual subscriptions.\nThe rotation period has the format of a decimal number, followed by the\nletter 's' (seconds). Cannot be more than 31 days or less than 10 minutes.",
        "description_kind": "plain",
        "type": "string"
      },
      "message_storage_policy": {
        "computed": true,
        "description": "Policy constraining the set of Google Cloud Platform regions where\nmessages published to the topic may be stored. If not present, then no\nconstraints are in effect.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allowed_persistence_regions": [
                "set",
                "string"
              ],
              "enforce_in_transit": "bool"
            }
          ]
        ]
      },
      "message_transforms": {
        "computed": true,
        "description": "Transforms to be applied to messages published to the topic. Transforms are applied in the\norder specified.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disabled": "bool",
              "javascript_udf": [
                "list",
                [
                  "object",
                  {
                    "code": "string",
                    "function_name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "description": "Name of the topic.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schema_settings": {
        "computed": true,
        "description": "Settings for validating messages published against a schema.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "encoding": "string",
              "schema": "string"
            }
          ]
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GooglePubsubTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePubsubTopic), &result)
	return &result
}
