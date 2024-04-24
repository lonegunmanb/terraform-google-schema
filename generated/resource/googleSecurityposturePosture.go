package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecurityposturePosture = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the Posture was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the posture.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "For Resource freshness validation (https://google.aip.dev/154)",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Location of the resource, eg: global.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the posture.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of the resource, an organization. Format should be 'organizations/{organization_id}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "posture_id": {
        "description": "Id of the posture. It is an immutable field.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "If set, there are currently changes in flight to the posture.",
        "description_kind": "plain",
        "type": "bool"
      },
      "revision_id": {
        "computed": true,
        "description": "Revision_id of the posture.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "description": "State of the posture. Update to state field should not be triggered along with\nwith other field updates. Possible values: [\"DEPRECATED\", \"DRAFT\", \"ACTIVE\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the Posture was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "policy_sets": {
        "block": {
          "attributes": {
            "description": {
              "description": "Description of the policy set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy_set_id": {
              "description": "ID of the policy set.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "policies": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "Description of the policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "policy_id": {
                    "description": "ID of the policy.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "compliance_standards": {
                    "block": {
                      "attributes": {
                        "control": {
                          "description": "Mapping of security controls for the policy.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "standard": {
                          "description": "Mapping of compliance standards for the policy.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Mapping for policy to security standards and controls.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "constraint": {
                    "block": {
                      "block_types": {
                        "org_policy_constraint": {
                          "block": {
                            "attributes": {
                              "canned_constraint_id": {
                                "description": "Organization policy canned constraint Id",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "policy_rules": {
                                "block": {
                                  "attributes": {
                                    "allow_all": {
                                      "description": "Setting this to true means that all values are allowed. This field can be set only in policies for list constraints.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "deny_all": {
                                      "description": "Setting this to true means that all values are denied. This field can be set only in policies for list constraints.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "enforce": {
                                      "description": "If 'true', then the policy is enforced. If 'false', then any configuration is acceptable.\nThis field can be set only in policies for boolean constraints.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "condition": {
                                      "block": {
                                        "attributes": {
                                          "description": {
                                            "description": "Description of the expression",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "expression": {
                                            "description": "Textual representation of an expression in Common Expression Language syntax.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "location": {
                                            "description": "String indicating the location of the expression for error reporting, e.g. a file name and a position in the file",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description": "Title for the expression, i.e. a short string describing its purpose.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.\nThis page details the objects and attributes that are used to the build the CEL expressions for\ncustom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "values": {
                                      "block": {
                                        "attributes": {
                                          "allowed_values": {
                                            "description": "List of values allowed at this resource.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "denied_values": {
                                            "description": "List of values denied at this resource.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "List of values to be used for this policy rule. This field can be set only in policies for list constraints.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Definition of policy rules",
                                  "description_kind": "plain"
                                },
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Organization policy canned constraint definition.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "org_policy_constraint_custom": {
                          "block": {
                            "block_types": {
                              "custom_constraint": {
                                "block": {
                                  "attributes": {
                                    "action_type": {
                                      "description": "The action to take if the condition is met. Possible values: [\"ALLOW\", \"DENY\"]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "condition": {
                                      "description": "A CEL condition that refers to a supported service resource, for example 'resource.management.autoUpgrade == false'. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "description": {
                                      "description": "A human-friendly description of the constraint to display as an error message when the policy is violated.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "display_name": {
                                      "description": "A human-friendly name for the constraint.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "method_types": {
                                      "description": "A list of RESTful methods for which to enforce the constraint. Can be 'CREATE', 'UPDATE', or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "name": {
                                      "description": "Immutable. The name of the custom constraint. This is unique within the organization.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "resource_types": {
                                      "description": "Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, 'container.googleapis.com/NodePool'.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "Organization policy custom constraint definition.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "policy_rules": {
                                "block": {
                                  "attributes": {
                                    "allow_all": {
                                      "description": "Setting this to true means that all values are allowed. This field can be set only in policies for list constraints.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "deny_all": {
                                      "description": "Setting this to true means that all values are denied. This field can be set only in policies for list constraints.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "enforce": {
                                      "description": "If 'true', then the policy is enforced. If 'false', then any configuration is acceptable.\nThis field can be set only in policies for boolean constraints.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "condition": {
                                      "block": {
                                        "attributes": {
                                          "description": {
                                            "description": "Description of the expression",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "expression": {
                                            "description": "Textual representation of an expression in Common Expression Language syntax.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "location": {
                                            "description": "String indicating the location of the expression for error reporting, e.g. a file name and a position in the file",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description": "Title for the expression, i.e. a short string describing its purpose.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.\nThis page details the objects and attributes that are used to the build the CEL expressions for\ncustom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "values": {
                                      "block": {
                                        "attributes": {
                                          "allowed_values": {
                                            "description": "List of values allowed at this resource.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "denied_values": {
                                            "description": "List of values denied at this resource.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "List of values to be used for this policy rule. This field can be set only in policies for list constraints.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Definition of policy rules",
                                  "description_kind": "plain"
                                },
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Organization policy custom constraint policy definition.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "security_health_analytics_custom_module": {
                          "block": {
                            "attributes": {
                              "display_name": {
                                "description": "The display name of the Security Health Analytics custom module. This\ndisplay name becomes the finding category for all findings that are\nreturned by this custom module.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description": "A server generated id of custom module.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "module_enablement_state": {
                                "description": "The state of enablement for the module at its level of the resource hierarchy. Possible values: [\"ENABLEMENT_STATE_UNSPECIFIED\", \"ENABLED\", \"DISABLED\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "config": {
                                "block": {
                                  "attributes": {
                                    "description": {
                                      "description": "Text that describes the vulnerability or misconfiguration that the custom\nmodule detects.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "recommendation": {
                                      "description": "An explanation of the recommended steps that security teams can take to\nresolve the detected issue",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "severity": {
                                      "description": "The severity to assign to findings generated by the module. Possible values: [\"SEVERITY_UNSPECIFIED\", \"CRITICAL\", \"HIGH\", \"MEDIUM\", \"LOW\"]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "custom_output": {
                                      "block": {
                                        "block_types": {
                                          "properties": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Name of the property for the custom output.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "value_expression": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
                                                        "description": "Description of the expression",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "expression": {
                                                        "description": "Textual representation of an expression in Common Expression Language syntax.",
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "location": {
                                                        "description": "String indicating the location of the expression for error reporting, e.g. a file name and a position in the file",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "description": "Title for the expression, i.e. a short string describing its purpose.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description": "The CEL expression for the custom output. A resource property can be\nspecified to return the value of the property or a text string enclosed\nin quotation marks.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "A list of custom output properties to add to the finding.",
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Custom output properties. A set of optional name-value pairs that define custom source properties to\nreturn with each finding that is generated by the custom module. The custom\nsource properties that are defined here are included in the finding JSON\nunder 'sourceProperties'.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "predicate": {
                                      "block": {
                                        "attributes": {
                                          "description": {
                                            "description": "Description of the expression",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "expression": {
                                            "description": "Textual representation of an expression in Common Expression Language syntax.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "location": {
                                            "description": "String indicating the location of the expression for error reporting, e.g. a file name and a position in the file",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description": "Title for the expression, i.e. a short string describing its purpose.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The CEL expression to evaluate to produce findings.When the expression\nevaluates to true against a resource, a finding is generated.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "resource_selector": {
                                      "block": {
                                        "attributes": {
                                          "resource_types": {
                                            "description": "The resource types to run the detector on.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "The resource types that the custom module operates on. Each custom module\ncan specify up to 5 resource types.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Custom module details.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Definition of Security Health Analytics Custom Module.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "security_health_analytics_module": {
                          "block": {
                            "attributes": {
                              "module_enablement_state": {
                                "description": "The state of enablement for the module at its level of the resource hierarchy. Possible values: [\"ENABLEMENT_STATE_UNSPECIFIED\", \"ENABLED\", \"DISABLED\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "module_name": {
                                "description": "The name of the module eg: BIGQUERY_TABLE_CMEK_DISABLED.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Security Health Analytics built-in detector definition.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Policy constraint definition.It can have the definition of one of following constraints: orgPolicyConstraint orgPolicyConstraintCustom securityHealthAnalyticsModule securityHealthAnalyticsCustomModule",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of security policy",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "List of policy sets for the posture.",
          "description_kind": "plain"
        },
        "min_items": 1,
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

func GoogleSecurityposturePostureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecurityposturePosture), &result)
	return &result
}
