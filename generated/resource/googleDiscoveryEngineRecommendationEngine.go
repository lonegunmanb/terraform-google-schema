package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDiscoveryEngineRecommendationEngine = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Timestamp the Engine was created at.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_store_ids": {
        "description": "The data stores associated with this engine. For SOLUTION_TYPE_RECOMMENDATION type of engines, they can only associate with at most one data store.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description": "Required. The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_id": {
        "description": "Unique ID to use for Recommendation Engine.",
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
        "description": "The industry vertical that the engine registers. The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to GENERIC. Vertical on Engine has to match vertical of the DataStore liniked to the engine. Default value: \"GENERIC\" Possible values: [\"GENERIC\", \"MEDIA\"]",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The unique full resource name of the recommendation engine. Values are of the format\n'projects/{project}/locations/{location}/collections/{collection}/engines/{engine_id}'.\nThis field must be a UTF-8 encoded string with a length limit of 1024 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Timestamp the Engine was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "common_config": {
        "block": {
          "attributes": {
            "company_name": {
              "description": "The name of the company, business or entity that is associated with the engine. Setting this may help improve LLM related features.cd",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Common config spec that specifies the metadata of the engine.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "media_recommendation_engine_config": {
        "block": {
          "attributes": {
            "optimization_objective": {
              "description": "The optimization objective. e.g., 'cvr'.\nThis field together with MediaRecommendationEngineConfig.type describes\nengine metadata to use to control engine training and serving.\nCurrently supported values: 'ctr', 'cvr'.\nIf not specified, we choose default based on engine type. Default depends on type of recommendation:\n'recommended-for-you' =\u003e 'ctr'\n'others-you-may-like' =\u003e 'ctr'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "training_state": {
              "description": "The training state that the engine is in (e.g. 'TRAINING' or 'PAUSED').\nSince part of the cost of running the service\nis frequency of training - this can be used to determine when to train\nengine in order to control cost. If not specified: the default value for\n'CreateEngine' method is 'TRAINING'. The default value for\n'UpdateEngine' method is to keep the state the same as before. Possible values: [\"PAUSED\", \"TRAINING\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "The type of engine. e.g., 'recommended-for-you'.\nThis field together with MediaRecommendationEngineConfig.optimizationObjective describes\nengine metadata to use to control engine training and serving.\nCurrently supported values: 'recommended-for-you', 'others-you-may-like',\n'more-like-this', 'most-popular-items'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "engine_features_config": {
              "block": {
                "block_types": {
                  "most_popular_config": {
                    "block": {
                      "attributes": {
                        "time_window_days": {
                          "description": "The time window of which the engine is queried at training and\nprediction time. Positive integers only. The value translates to the\nlast X days of events. Currently required for the 'most-popular-items'\nengine.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Feature configurations that are required for creating a Most Popular engine.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "recommended_for_you_config": {
                    "block": {
                      "attributes": {
                        "context_event_type": {
                          "description": "The type of event with which the engine is queried at prediction time.\nIf set to 'generic', only 'view-item', 'media-play',and\n'media-complete' will be used as 'context-event' in engine training. If\nset to 'view-home-page', 'view-home-page' will also be used as\n'context-events' in addition to 'view-item', 'media-play', and\n'media-complete'. Currently supported for the 'recommended-for-you'\nengine. Currently supported values: 'view-home-page', 'generic'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Additional feature configurations for creating a 'recommended-for-you' engine.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "More feature configs of the selected engine type.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "optimization_objective_config": {
              "block": {
                "attributes": {
                  "target_field": {
                    "description": "The name of the field to target. Currently supported values: 'watch-percentage', 'watch-time'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_field_value_float": {
                    "description": "The threshold to be applied to the target (e.g., 0.5).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Name and value of the custom threshold for cvr optimization_objective.\nFor target_field 'watch-time', target_field_value must be an integer\nvalue indicating the media progress time in seconds between (0, 86400]\n(excludes 0, includes 86400) (e.g., 90).\nFor target_field 'watch-percentage', the target_field_value must be a\nvalid float value between (0, 1.0] (excludes 0, includes 1.0) (e.g., 0.5).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configurations for a Media Recommendation Engine. Only applicable on the data stores\nwith SOLUTION_TYPE_RECOMMENDATION solution type and MEDIA industry vertical.",
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

func GoogleDiscoveryEngineRecommendationEngineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDiscoveryEngineRecommendationEngine), &result)
	return &result
}
