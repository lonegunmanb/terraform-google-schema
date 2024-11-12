package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDiscoveryEngineTargetSite = `{
  "block": {
    "attributes": {
      "data_store_id": {
        "description": "The unique id of the data store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "exact_match": {
        "description": "If set to false, a uri_pattern is generated to include all pages whose\naddress contains the provided_uri_pattern. If set to true, an uri_pattern\nis generated to try to be an exact match of the provided_uri_pattern or\njust the specific page if the provided_uri_pattern is a specific one.\nprovided_uri_pattern is always normalized to generate the URI pattern to\nbe used by the search engine.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "failure_reason": {
        "computed": true,
        "description": "Site search indexing failure reasons.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "quota_failure": [
                "list",
                [
                  "object",
                  {
                    "total_required_quota": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "generated_uri_pattern": {
        "computed": true,
        "description": "This is system-generated based on the 'provided_uri_pattern'.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "indexing_status": {
        "computed": true,
        "description": "The indexing status.",
        "description_kind": "plain",
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
        "description": "The unique full resource name of the target site. Values are of the format\n'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}/siteSearchEngine/targetSites/{target_site_id}'.\nThis field must be a UTF-8 encoded string with a length limit of 1024\ncharacters.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provided_uri_pattern": {
        "description": "The user provided URI pattern from which the 'generated_uri_pattern' is\ngenerated.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "root_domain_uri": {
        "computed": true,
        "description": "Root domain of the 'provided_uri_pattern'.",
        "description_kind": "plain",
        "type": "string"
      },
      "site_verification_info": {
        "computed": true,
        "description": "Site ownership and validity verification status.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "site_verification_state": "string",
              "verify_time": "string"
            }
          ]
        ]
      },
      "target_site_id": {
        "computed": true,
        "description": "The unique id of the target site.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "The possible target site types. Possible values: [\"INCLUDE\", \"EXCLUDE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The target site's last updated time.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleDiscoveryEngineTargetSiteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDiscoveryEngineTargetSite), &result)
	return &result
}
