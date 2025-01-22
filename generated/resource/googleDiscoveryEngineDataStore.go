package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDiscoveryEngineDataStore = `{
  "block": {
    "attributes": {
      "content_config": {
        "description": "The content config of the data store. Possible values: [\"NO_CONTENT\", \"CONTENT_REQUIRED\", \"PUBLIC_WEBSITE\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_advanced_site_search": {
        "description": "If true, an advanced data store for site search will be created. If the\ndata store is not configured as site search (GENERIC vertical and\nPUBLIC_WEBSITE contentConfig), this flag will be ignored.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description": "Timestamp when the DataStore was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_store_id": {
        "description": "The unique id of the data store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_schema_id": {
        "computed": true,
        "description": "The id of the default Schema associated with this data store.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the data store. This field must be a UTF-8 encoded\nstring with a length limit of 128 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "industry_vertical": {
        "description": "The industry vertical that the data store registers. Possible values: [\"GENERIC\", \"MEDIA\", \"HEALTHCARE_FHIR\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The geographic location where the data store should reside. The value can\nonly be one of \"global\", \"us\" and \"eu\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique full resource name of the data store. Values are of the format\n'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}'.\nThis field must be a UTF-8 encoded string with a length limit of 1024\ncharacters.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_default_schema_creation": {
        "description": "A boolean flag indicating whether to skip the default schema creation for\nthe data store. Only enable this flag if you are certain that the default\nschema is incompatible with your use case.\nIf set to true, you must manually create a schema for the data store\nbefore any documents can be ingested.\nThis flag cannot be specified if 'data_store.starting_schema' is\nspecified.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "solution_types": {
        "description": "The solutions that the data store enrolls. Possible values: [\"SOLUTION_TYPE_RECOMMENDATION\", \"SOLUTION_TYPE_SEARCH\", \"SOLUTION_TYPE_CHAT\", \"SOLUTION_TYPE_GENERATIVE_CHAT\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "advanced_site_search_config": {
        "block": {
          "attributes": {
            "disable_automatic_refresh": {
              "description": "If set true, automatic refresh is disabled for the DataStore.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disable_initial_index": {
              "description": "If set true, initial indexing is disabled for the DataStore.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Configuration data for advance site search.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "document_processing_config": {
        "block": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The full resource name of the Document Processing Config. Format:\n'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}/documentProcessingConfig'.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "chunking_config": {
              "block": {
                "block_types": {
                  "layout_based_chunking_config": {
                    "block": {
                      "attributes": {
                        "chunk_size": {
                          "description": "The token size limit for each chunk.\nSupported values: 100-500 (inclusive). Default value: 500.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "include_ancestor_headings": {
                          "description": "Whether to include appending different levels of headings to chunks from the middle of the document to prevent context loss.\nDefault value: False.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Configuration for the layout based chunking.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Whether chunking mode is enabled.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "default_parsing_config": {
              "block": {
                "block_types": {
                  "digital_parsing_config": {
                    "block": {
                      "description": "Configurations applied to digital parser.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "layout_parsing_config": {
                    "block": {
                      "description": "Configurations applied to layout parser.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ocr_parsing_config": {
                    "block": {
                      "attributes": {
                        "use_native_text": {
                          "description": "If true, will use native text instead of OCR text on pages containing native text.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Configurations applied to OCR parser. Currently it only applies to PDFs.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Configurations for default Document parser. If not specified, this resource\nwill be configured to use a default DigitalParsingConfig, and the default parsing\nconfig will be applied to all file types for Document parsing.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "parsing_config_overrides": {
              "block": {
                "attributes": {
                  "file_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "digital_parsing_config": {
                    "block": {
                      "description": "Configurations applied to digital parser.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "layout_parsing_config": {
                    "block": {
                      "description": "Configurations applied to layout parser.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ocr_parsing_config": {
                    "block": {
                      "attributes": {
                        "use_native_text": {
                          "description": "If true, will use native text instead of OCR text on pages containing native text.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Configurations applied to OCR parser. Currently it only applies to PDFs.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Map from file type to override the default parsing configuration based on the file type. Supported keys:\n  * 'pdf': Override parsing config for PDF files, either digital parsing, ocr parsing or layout parsing is supported.\n  * 'html': Override parsing config for HTML files, only digital parsing and or layout parsing are supported.\n  * 'docx': Override parsing config for DOCX files, only digital parsing and or layout parsing are supported.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "Configuration for Document understanding and enrichment.",
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

func GoogleDiscoveryEngineDataStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDiscoveryEngineDataStore), &result)
	return &result
}
