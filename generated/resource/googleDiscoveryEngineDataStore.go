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
        "description": "The industry vertical that the data store registers. Possible values: [\"GENERIC\", \"MEDIA\"]",
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
      "solution_types": {
        "description": "The solutions that the data store enrolls. Possible values: [\"SOLUTION_TYPE_RECOMMENDATION\", \"SOLUTION_TYPE_SEARCH\", \"SOLUTION_TYPE_CHAT\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
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
