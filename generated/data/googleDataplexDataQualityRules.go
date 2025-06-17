package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataplexDataQualityRules = `{
  "block": {
    "attributes": {
      "data_scan_id": {
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
      "location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "column": "string",
              "description": "string",
              "dimension": "string",
              "ignore_null": "bool",
              "name": "string",
              "non_null_expectation": [
                "list",
                [
                  "object",
                  {}
                ]
              ],
              "range_expectation": [
                "list",
                [
                  "object",
                  {
                    "max_value": "string",
                    "min_value": "string",
                    "strict_max_enabled": "bool",
                    "strict_min_enabled": "bool"
                  }
                ]
              ],
              "regex_expectation": [
                "list",
                [
                  "object",
                  {
                    "regex": "string"
                  }
                ]
              ],
              "row_condition_expectation": [
                "list",
                [
                  "object",
                  {
                    "sql_expression": "string"
                  }
                ]
              ],
              "set_expectation": [
                "list",
                [
                  "object",
                  {
                    "values": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "sql_assertion": [
                "list",
                [
                  "object",
                  {
                    "sql_statement": "string"
                  }
                ]
              ],
              "statistic_range_expectation": [
                "list",
                [
                  "object",
                  {
                    "max_value": "string",
                    "min_value": "string",
                    "statistic": "string",
                    "strict_max_enabled": "bool",
                    "strict_min_enabled": "bool"
                  }
                ]
              ],
              "suspended": "bool",
              "table_condition_expectation": [
                "list",
                [
                  "object",
                  {
                    "sql_expression": "string"
                  }
                ]
              ],
              "threshold": "number",
              "uniqueness_expectation": [
                "list",
                [
                  "object",
                  {}
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataplexDataQualityRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataplexDataQualityRules), &result)
	return &result
}
