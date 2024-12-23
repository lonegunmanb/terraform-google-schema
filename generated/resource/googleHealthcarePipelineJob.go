package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleHealthcarePipelineJob = `{
  "block": {
    "attributes": {
      "dataset": {
        "description": "Healthcare Dataset under which the Pipeline Job is to run",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disable_lineage": {
        "description": "If true, disables writing lineage for the pipeline.",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-supplied key-value pairs used to organize Pipeline Jobs.\nLabel keys must be between 1 and 63 characters long, have a UTF-8 encoding of\nmaximum 128 bytes, and must conform to the following PCRE regular expression:\n[\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}\nLabel values are optional, must be between 1 and 63 characters long, have a\nUTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE\nregular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}\nNo more than 64 labels can be associated with a given pipeline.\nAn object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location where the Pipeline Job is to run",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Specifies the name of the pipeline job. This field is user-assigned.",
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
      }
    },
    "block_types": {
      "backfill_pipeline_job": {
        "block": {
          "attributes": {
            "mapping_pipeline_job": {
              "description": "Specifies the mapping pipeline job to backfill, the name format\nshould follow: projects/{projectId}/locations/{locationId}/datasets/{datasetId}/pipelineJobs/{pipelineJobId}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Specifies the backfill configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mapping_pipeline_job": {
        "block": {
          "attributes": {
            "fhir_store_destination": {
              "description": "If set, the mapping pipeline will write snapshots to this\nFHIR store without assigning stable IDs. You must\ngrant your pipeline project's Cloud Healthcare Service\nAgent serviceaccount healthcare.fhirResources.executeBundle\nand healthcare.fhirResources.create permissions on the\ndestination store. The destination store must set\n[disableReferentialIntegrity][FhirStore.disable_referential_integrity]\nto true. The destination store must use FHIR version R4.\nFormat: project/{projectID}/locations/{locationID}/datasets/{datasetName}/fhirStores/{fhirStoreID}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reconciliation_destination": {
              "description": "If set to true, a mapping pipeline will send output snapshots\nto the reconciliation pipeline in its dataset. A reconciliation\npipeline must exist in this dataset before a mapping pipeline\nwith a reconciliation destination can be created.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "fhir_streaming_source": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "Describes the streaming FHIR data source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fhir_store": {
                    "description": "The path to the FHIR store in the format projects/{projectId}/locations/{locationId}/datasets/{datasetId}/fhirStores/{fhirStoreId}.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A streaming FHIR data source.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "mapping_config": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "Describes the mapping configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "whistle_config_source": {
                    "block": {
                      "attributes": {
                        "import_uri_prefix": {
                          "description": "Directory path where all the Whistle files are located.\nExample: gs://{bucket-id}/{path/to/import-root/dir}",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "uri": {
                          "description": "Main configuration file which has the entrypoint or the root function.\nExample: gs://{bucket-id}/{path/to/import-root/dir}/entrypoint-file-name.wstl.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies the path to the mapping configuration for harmonization pipeline.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The location of the mapping configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies mapping configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "reconciliation_pipeline_job": {
        "block": {
          "attributes": {
            "fhir_store_destination": {
              "description": "The harmonized FHIR store to write harmonized FHIR resources to,\nin the format of: project/{projectID}/locations/{locationID}/datasets/{datasetName}/fhirStores/{id}",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "matching_uri_prefix": {
              "description": "Specifies the top level directory of the matching configs used\nin all mapping pipelines, which extract properties for resources\nto be matched on.\nExample: gs://{bucket-id}/{path/to/matching/configs}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "merge_config": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "Describes the mapping configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "whistle_config_source": {
                    "block": {
                      "attributes": {
                        "import_uri_prefix": {
                          "description": "Directory path where all the Whistle files are located.\nExample: gs://{bucket-id}/{path/to/import-root/dir}",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "uri": {
                          "description": "Main configuration file which has the entrypoint or the root function.\nExample: gs://{bucket-id}/{path/to/import-root/dir}/entrypoint-file-name.wstl.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies the path to the mapping configuration for harmonization pipeline.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies the location of the reconciliation configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies reconciliation configuration.",
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

func GoogleHealthcarePipelineJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleHealthcarePipelineJob), &result)
	return &result
}
