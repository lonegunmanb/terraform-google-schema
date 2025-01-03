package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleHealthcareFhirStore = `{
  "block": {
    "attributes": {
      "complex_data_type_reference_parsing": {
        "computed": true,
        "description": "Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources. Possible values: [\"COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED\", \"DISABLED\", \"ENABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dataset": {
        "description": "Identifies the dataset addressed by this request. Must be in the format\n'projects/{project}/locations/{location}/datasets/{dataset}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_search_handling_strict": {
        "description": "If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.\nIf false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.\nThe handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_referential_integrity": {
        "description": "Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store\ncreation. The default value is false, meaning that the API will enforce referential integrity and fail the\nrequests that will result in inconsistent state in the FHIR store. When this field is set to true, the API\nwill skip referential integrity check. Consequently, operations that rely on references, such as\nPatient.get$everything, will not return all the results if broken references exist.\n\n** Changing this property may recreate the FHIR store (removing all data) **",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_resource_versioning": {
        "description": "Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation\nof FHIR store. If set to false, which is the default behavior, all write operations will cause historical\nversions to be recorded automatically. The historical versions can be fetched through the history APIs, but\ncannot be updated. If set to true, no historical versions will be kept. The server will send back errors for\nattempts to read the historical versions.\n\n** Changing this property may recreate the FHIR store (removing all data) **",
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
      "enable_history_import": {
        "description": "Whether to allow the bulk import API to accept history bundles and directly insert historical resource\nversions into the FHIR store. Importing resource histories creates resource interactions that appear to have\noccurred in the past, which clients may not want to allow. If set to false, history bundles within an import\nwill fail with an error.\n\n** Changing this property may recreate the FHIR store (removing all data) **\n\n** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_update_create": {
        "description": "Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update\noperation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through\nthe Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit\nlogs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient\nidentifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub\nnotifications.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-supplied key-value pairs used to organize FHIR stores.\n\nLabel keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must\nconform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}\n\nLabel values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128\nbytes, and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}\n\nNo more than 64 labels can be associated with a given store.\n\nAn object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The resource name for the FhirStore.\n\n** Changing this property may recreate the FHIR store (removing all data) **",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The fully qualified name of this dataset",
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
      "version": {
        "description": "The FHIR specification version. Possible values: [\"DSTU2\", \"STU3\", \"R4\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "notification_config": {
        "block": {
          "attributes": {
            "pubsub_topic": {
              "description": "The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.\nPubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.\nIt is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message\nwas published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a\nproject. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given\nCloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description": "A nested object resource.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "notification_configs": {
        "block": {
          "attributes": {
            "pubsub_topic": {
              "description": "The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.\nPubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.\nIt is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message\nwas published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a\nproject. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given\nCloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "send_full_resource": {
              "description": "Whether to send full FHIR resource to this Pub/Sub topic for Create and Update operation.\nNote that setting this to true does not guarantee that all resources will be sent in the format of\nfull FHIR resource. When a resource change is too large or during heavy traffic, only the resource name will be\nsent. Clients should always check the \"payloadType\" label from a Pub/Sub message to determine whether\nit needs to fetch the full resource as a separate operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "send_previous_resource_on_delete": {
              "description": "Whether to send full FHIR resource to this Pub/Sub topic for deleting FHIR resource. Note that setting this to\ntrue does not guarantee that all previous resources will be sent in the format of full FHIR resource. When a\nresource change is too large or during heavy traffic, only the resource name will be sent. Clients should always\ncheck the \"payloadType\" label from a Pub/Sub message to determine whether it needs to fetch the full previous\nresource as a separate operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "stream_configs": {
        "block": {
          "attributes": {
            "resource_types": {
              "description": "Supply a FHIR resource type (such as \"Patient\" or \"Observation\"). See\nhttps://www.hl7.org/fhir/valueset-resource-types.html for a list of all FHIR resource types. The server treats\nan empty list as an intent to stream all the supported resource types in this FHIR store.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "bigquery_destination": {
              "block": {
                "attributes": {
                  "dataset_uri": {
                    "description": "BigQuery URI to a dataset, up to 2000 characters long, in the format bq://projectId.bqDatasetId",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "schema_config": {
                    "block": {
                      "attributes": {
                        "recursive_structure_depth": {
                          "description": "The depth for all recursive structures in the output analytics schema. For example, concept in the CodeSystem\nresource is a recursive structure; when the depth is 2, the CodeSystem table will have a column called\nconcept.concept but not concept.concept.concept. If not specified or set to 0, the server will use the default\nvalue 2. The maximum depth allowed is 5.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "schema_type": {
                          "description": "Specifies the output schema type.\n * ANALYTICS: Analytics schema defined by the FHIR community.\n  See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md.\n * ANALYTICS_V2: Analytics V2, similar to schema defined by the FHIR community, with added support for extensions with one or more occurrences and contained resources in stringified JSON.\n * LOSSLESS: A data-driven schema generated from the fields present in the FHIR data being exported, with no additional simplification. Default value: \"ANALYTICS\" Possible values: [\"ANALYTICS\", \"ANALYTICS_V2\", \"LOSSLESS\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "last_updated_partition_config": {
                          "block": {
                            "attributes": {
                              "expiration_ms": {
                                "description": "Number of milliseconds for which to keep the storage for a partition.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "Type of partitioning. Possible values: [\"PARTITION_TYPE_UNSPECIFIED\", \"HOUR\", \"DAY\", \"MONTH\", \"YEAR\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The configuration for exported BigQuery tables to be partitioned by FHIR resource's last updated time column.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The configuration for the exported BigQuery schema.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The destination BigQuery structure that contains both the dataset location and corresponding schema config.\nThe output is organized in one table per resource type. The server reuses the existing tables (if any) that\nare named after the resource types, e.g. \"Patient\", \"Observation\". When there is no existing table for a given\nresource type, the server attempts to create one.\nSee the [streaming config reference](https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.fhirStores#streamconfig) for more details.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A list of streaming configs that configure the destinations of streaming export for every resource mutation in\nthis FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next\nresource mutation is streamed to the new location in addition to the existing ones. When a location is removed\nfrom the list, the server stops streaming to that location. Before adding a new config, you must add the required\nbigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on\nthe order of dozens of seconds) is expected before the results show up in the streaming destination.",
          "description_kind": "plain"
        },
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

func GoogleHealthcareFhirStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleHealthcareFhirStore), &result)
	return &result
}
