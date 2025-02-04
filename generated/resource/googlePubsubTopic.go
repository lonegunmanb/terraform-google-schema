package resource

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
      "kms_key_name": {
        "description": "The resource name of the Cloud KMS CryptoKey to be used to protect access\nto messages published on this topic. Your project's PubSub service account\n('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nThe expected format is 'projects/*/locations/*/keyRings/*/cryptoKeys/*'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs to assign to this Topic.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "message_retention_duration": {
        "description": "Indicates the minimum duration to retain a message after it is published\nto the topic. If this field is set, messages published to the topic in\nthe last messageRetentionDuration are always available to subscribers.\nFor instance, it allows any attached subscription to seek to a timestamp\nthat is up to messageRetentionDuration in the past. If this field is not\nset, message retention is controlled by settings on individual subscriptions.\nThe rotation period has the format of a decimal number, followed by the\nletter 's' (seconds). Cannot be more than 31 days or less than 10 minutes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the topic.",
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
      "ingestion_data_source_settings": {
        "block": {
          "block_types": {
            "aws_kinesis": {
              "block": {
                "attributes": {
                  "aws_role_arn": {
                    "description": "AWS role ARN to be used for Federated Identity authentication with\nKinesis. Check the Pub/Sub docs for how to set up this role and the\nrequired permissions that need to be attached to it.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "consumer_arn": {
                    "description": "The Kinesis consumer ARN to used for ingestion in\nEnhanced Fan-Out mode. The consumer must be already\ncreated and ready to be used.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "gcp_service_account": {
                    "description": "The GCP service account to be used for Federated Identity authentication\nwith Kinesis (via a 'AssumeRoleWithWebIdentity' call for the provided\nrole). The 'awsRoleArn' must be set up with 'accounts.google.com:sub'\nequals to this service account number.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "stream_arn": {
                    "description": "The Kinesis stream ARN to ingest data from.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Settings for ingestion from Amazon Kinesis Data Streams.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "aws_msk": {
              "block": {
                "attributes": {
                  "aws_role_arn": {
                    "description": "AWS role ARN to be used for Federated Identity authentication with\nMSK. Check the Pub/Sub docs for how to set up this role and the\nrequired permissions that need to be attached to it.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cluster_arn": {
                    "description": "ARN that uniquely identifies the MSK cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "gcp_service_account": {
                    "description": "The GCP service account to be used for Federated Identity authentication\nwith MSK (via a 'AssumeRoleWithWebIdentity' call for the provided\nrole). The 'awsRoleArn' must be set up with 'accounts.google.com:sub'\nequals to this service account number.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "topic": {
                    "description": "The name of the MSK topic that Pub/Sub will import from.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Settings for ingestion from Amazon Managed Streaming for Apache Kafka.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "azure_event_hubs": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The Azure event hub client ID to use for ingestion.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "event_hub": {
                    "description": "The Azure event hub to ingest data from.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gcp_service_account": {
                    "description": "The GCP service account to be used for Federated Identity authentication\nwith Azure (via a 'AssumeRoleWithWebIdentity' call for the provided\nrole).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description": "The Azure event hub namespace to ingest data from.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource_group": {
                    "description": "The name of the resource group within an Azure subscription.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subscription_id": {
                    "description": "The Azure event hub subscription ID to use for ingestion.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tenant_id": {
                    "description": "The Azure event hub tenant ID to use for ingestion.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Settings for ingestion from Azure Event Hubs.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cloud_storage": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description": "Cloud Storage bucket. The bucket name must be without any\nprefix like \"gs://\". See the bucket naming requirements:\nhttps://cloud.google.com/storage/docs/buckets#naming.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "match_glob": {
                    "description": "Glob pattern used to match objects that will be ingested. If unset, all\nobjects will be ingested. See the supported patterns:\nhttps://cloud.google.com/storage/docs/json_api/v1/objects/list#list-objects-and-prefixes-using-glob",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "minimum_object_create_time": {
                    "description": "The timestamp set in RFC3339 text format. If set, only objects with a\nlarger or equal timestamp will be ingested. Unset by default, meaning\nall objects will be ingested.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "avro_format": {
                    "block": {
                      "description": "Configuration for reading Cloud Storage data in Avro binary format. The\nbytes of each object will be set to the 'data' field of a Pub/Sub message.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "pubsub_avro_format": {
                    "block": {
                      "description": "Configuration for reading Cloud Storage data written via Cloud Storage\nsubscriptions(See https://cloud.google.com/pubsub/docs/cloudstorage). The\ndata and attributes fields of the originally exported Pub/Sub message\nwill be restored when publishing.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "text_format": {
                    "block": {
                      "attributes": {
                        "delimiter": {
                          "description": "The delimiter to use when using the 'text' format. Each line of text as\nspecified by the delimiter will be set to the 'data' field of a Pub/Sub\nmessage. When unset, '\\n' is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Configuration for reading Cloud Storage data in text format. Each line of\ntext as specified by the delimiter will be set to the 'data' field of a\nPub/Sub message.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Settings for ingestion from Cloud Storage.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "confluent_cloud": {
              "block": {
                "attributes": {
                  "bootstrap_server": {
                    "description": "The Confluent Cloud bootstrap server. The format is url:port.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cluster_id": {
                    "description": "The Confluent Cloud cluster ID.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gcp_service_account": {
                    "description": "The GCP service account to be used for Federated Identity authentication\nwith Confluent Cloud.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "identity_pool_id": {
                    "description": "Identity pool ID to be used for Federated Identity authentication with Confluent Cloud.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "topic": {
                    "description": "Name of the Confluent Cloud topic that Pub/Sub will import from.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Settings for ingestion from Confluent Cloud.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "platform_logs_settings": {
              "block": {
                "attributes": {
                  "severity": {
                    "description": "The minimum severity level of Platform Logs that will be written. If unspecified,\nno Platform Logs will be written. Default value: \"SEVERITY_UNSPECIFIED\" Possible values: [\"SEVERITY_UNSPECIFIED\", \"DISABLED\", \"DEBUG\", \"INFO\", \"WARNING\", \"ERROR\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Settings for Platform Logs regarding ingestion to Pub/Sub. If unset,\nno Platform Logs will be generated.'",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Settings for ingestion from a data source into this topic.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "message_storage_policy": {
        "block": {
          "attributes": {
            "allowed_persistence_regions": {
              "description": "A list of IDs of GCP regions where messages that are published to\nthe topic may be persisted in storage. Messages published by\npublishers running in non-allowed GCP regions (or running outside\nof GCP altogether) will be routed for storage in one of the\nallowed regions. An empty list means that no regions are allowed,\nand is not a valid configuration.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "enforce_in_transit": {
              "description": "If true, 'allowedPersistenceRegions' is also used to enforce in-transit\nguarantees for messages. That is, Pub/Sub will fail topics.publish\noperations on this topic and subscribe operations on any subscription\nattached to this topic in any region that is not in 'allowedPersistenceRegions'.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Policy constraining the set of Google Cloud Platform regions where\nmessages published to the topic may be stored. If not present, then no\nconstraints are in effect.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "schema_settings": {
        "block": {
          "attributes": {
            "encoding": {
              "description": "The encoding of messages validated against schema. Default value: \"ENCODING_UNSPECIFIED\" Possible values: [\"ENCODING_UNSPECIFIED\", \"JSON\", \"BINARY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schema": {
              "description": "The name of the schema that messages published should be\nvalidated against. Format is projects/{project}/schemas/{schema}.\nThe value of this field will be _deleted-schema_\nif the schema has been deleted.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Settings for validating messages published against a schema.",
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

func GooglePubsubTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePubsubTopic), &result)
	return &result
}
