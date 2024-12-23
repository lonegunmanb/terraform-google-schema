package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityAuthzPolicy = `{
  "block": {
    "attributes": {
      "action": {
        "description": "When the action is CUSTOM, customProvider must be specified.\nWhen the action is ALLOW, only requests matching the policy will be allowed.\nWhen the action is DENY, only requests matching the policy will be denied.\n\nWhen a request arrives, the policies are evaluated in the following order:\n1. If there is a CUSTOM policy that matches the request, the CUSTOM policy is evaluated using the custom authorization providers and the request is denied if the provider rejects the request.\n2. If there are any DENY policies that match the request, the request is denied.\n3. If there are no ALLOW policies for the resource or if any of the ALLOW policies match the request, the request is allowed.\n4. Else the request is denied by default if none of the configured AuthzPolicies with ALLOW action match the request. Possible values: [\"ALLOW\", \"DENY\", \"CUSTOM\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A human-readable description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "description": "Set of labels associated with the AuthzExtension resource.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Identifier. Name of the AuthzPolicy resource.",
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
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the resource was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "custom_provider": {
        "block": {
          "block_types": {
            "authz_extension": {
              "block": {
                "attributes": {
                  "resources": {
                    "description": "A list of references to authorization extensions that will be invoked for requests matching this policy. Limited to 1 custom provider.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Delegate authorization decision to user authored Service Extension. Only one of cloudIap or authzExtension can be specified.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cloud_iap": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Enable Cloud IAP at the AuthzPolicy level.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Delegates authorization decisions to Cloud IAP. Applicable only for managed load balancers. Enabling Cloud IAP at the AuthzPolicy level is not compatible with Cloud IAP settings in the BackendService. Enabling IAP in both places will result in request failure. Ensure that IAP is enabled in either the AuthzPolicy or the BackendService but not in both places.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Required if the action is CUSTOM. Allows delegating authorization decisions to Cloud IAP or to Service Extensions. One of cloudIap or authzExtension must be specified.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_rules": {
        "block": {
          "attributes": {
            "when": {
              "description": "CEL expression that describes the conditions to be satisfied for the action. The result of the CEL expression is ANDed with the from and to. Refer to the CEL language reference for a list of available attributes.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "from": {
              "block": {
                "block_types": {
                  "not_sources": {
                    "block": {
                      "block_types": {
                        "principals": {
                          "block": {
                            "attributes": {
                              "contains": {
                                "description": "The input string must have the substring specified here. Note: empty contains match is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc.def",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "exact": {
                                "description": "The input string must match exactly the string specified here.\nExamples:\n* abc only matches the value abc.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ignore_case": {
                                "description": "If true, indicates the exact/prefix/suffix/contains matching should be case insensitive. For example, the matcher data will match both input string Data and data if set to true.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "prefix": {
                                "description": "The input string must have the prefix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value abc.xyz",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "suffix": {
                                "description": "The input string must have the suffix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A list of identities derived from the client's certificate. This field will not match on a request unless mutual TLS is enabled for the Forwarding rule or Gateway. Each identity is a string whose value is matched against the URI SAN, or DNS SAN or the subject field in the client's certificate. The match can be exact, prefix, suffix or a substring match. One of exact, prefix, suffix or contains must be specified.\nLimited to 5 principals.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "resources": {
                          "block": {
                            "block_types": {
                              "iam_service_account": {
                                "block": {
                                  "attributes": {
                                    "contains": {
                                      "description": "The input string must have the substring specified here. Note: empty contains match is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc.def",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "exact": {
                                      "description": "The input string must match exactly the string specified here.\nExamples:\n* abc only matches the value abc.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "ignore_case": {
                                      "description": "If true, indicates the exact/prefix/suffix/contains matching should be case insensitive. For example, the matcher data will match both input string Data and data if set to true.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "prefix": {
                                      "description": "The input string must have the prefix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value abc.xyz",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suffix": {
                                      "description": "The input string must have the suffix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "An IAM service account to match against the source service account of the VM sending the request.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "tag_value_id_set": {
                                "block": {
                                  "attributes": {
                                    "ids": {
                                      "description": "A list of resource tag value permanent IDs to match against the resource manager tags value associated with the source VM of a request. The match follows AND semantics which means all the ids must match.\nLimited to 5 matches.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "A list of resource tag value permanent IDs to match against the resource manager tags value associated with the source VM of a request.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of resources to match against the resource of the source VM of a request.\nLimited to 5 resources.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Describes the properties of a request's sources. At least one of sources or notSources must be specified. Limited to 5 sources. A match occurs when ANY source (in sources or notSources) matches the request. Within a single source, the match follows AND semantics across fields and OR semantics within a single field, i.e. a match occurs when ANY principal matches AND ANY ipBlocks match.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "sources": {
                    "block": {
                      "block_types": {
                        "principals": {
                          "block": {
                            "attributes": {
                              "contains": {
                                "description": "The input string must have the substring specified here. Note: empty contains match is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc.def",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "exact": {
                                "description": "The input string must match exactly the string specified here.\nExamples:\n* abc only matches the value abc.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ignore_case": {
                                "description": "If true, indicates the exact/prefix/suffix/contains matching should be case insensitive. For example, the matcher data will match both input string Data and data if set to true.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "prefix": {
                                "description": "The input string must have the prefix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value abc.xyz",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "suffix": {
                                "description": "The input string must have the suffix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A list of identities derived from the client's certificate. This field will not match on a request unless mutual TLS is enabled for the Forwarding rule or Gateway. Each identity is a string whose value is matched against the URI SAN, or DNS SAN or the subject field in the client's certificate. The match can be exact, prefix, suffix or a substring match. One of exact, prefix, suffix or contains must be specified.\nLimited to 5 principals.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "resources": {
                          "block": {
                            "block_types": {
                              "iam_service_account": {
                                "block": {
                                  "attributes": {
                                    "contains": {
                                      "description": "The input string must have the substring specified here. Note: empty contains match is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc.def",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "exact": {
                                      "description": "The input string must match exactly the string specified here.\nExamples:\n* abc only matches the value abc.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "ignore_case": {
                                      "description": "If true, indicates the exact/prefix/suffix/contains matching should be case insensitive. For example, the matcher data will match both input string Data and data if set to true.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "prefix": {
                                      "description": "The input string must have the prefix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value abc.xyz",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suffix": {
                                      "description": "The input string must have the suffix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "An IAM service account to match against the source service account of the VM sending the request.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "tag_value_id_set": {
                                "block": {
                                  "attributes": {
                                    "ids": {
                                      "description": "A list of resource tag value permanent IDs to match against the resource manager tags value associated with the source VM of a request. The match follows AND semantics which means all the ids must match.\nLimited to 5 matches.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "A list of resource tag value permanent IDs to match against the resource manager tags value associated with the source VM of a request.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of resources to match against the resource of the source VM of a request.\nLimited to 5 resources.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Describes the properties of a request's sources. At least one of sources or notSources must be specified. Limited to 5 sources. A match occurs when ANY source (in sources or notSources) matches the request. Within a single source, the match follows AND semantics across fields and OR semantics within a single field, i.e. a match occurs when ANY principal matches AND ANY ipBlocks match.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Describes properties of one or more sources of a request.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "to": {
              "block": {
                "block_types": {
                  "operations": {
                    "block": {
                      "attributes": {
                        "methods": {
                          "description": "A list of HTTP methods to match against. Each entry must be a valid HTTP method name (GET, PUT, POST, HEAD, PATCH, DELETE, OPTIONS). It only allows exact match and is always case sensitive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "header_set": {
                          "block": {
                            "block_types": {
                              "headers": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Specifies the name of the header in the request.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "value": {
                                      "block": {
                                        "attributes": {
                                          "contains": {
                                            "description": "The input string must have the substring specified here. Note: empty contains match is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc.def",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "exact": {
                                            "description": "The input string must match exactly the string specified here.\nExamples:\n* abc only matches the value abc.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "ignore_case": {
                                            "description": "If true, indicates the exact/prefix/suffix/contains matching should be case insensitive. For example, the matcher data will match both input string Data and data if set to true.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "prefix": {
                                            "description": "The input string must have the prefix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value abc.xyz",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "suffix": {
                                            "description": "The input string must have the suffix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Specifies how the header match will be performed.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A list of headers to match against in http header. The match can be one of exact, prefix, suffix, or contains (substring match). The match follows AND semantics which means all the headers must match. Matches are always case sensitive unless the ignoreCase is set. Limited to 5 matches.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of headers to match against in http header.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "hosts": {
                          "block": {
                            "attributes": {
                              "contains": {
                                "description": "The input string must have the substring specified here. Note: empty contains match is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc.def",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "exact": {
                                "description": "The input string must match exactly the string specified here.\nExamples:\n* abc only matches the value abc.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ignore_case": {
                                "description": "If true, indicates the exact/prefix/suffix/contains matching should be case insensitive. For example, the matcher data will match both input string Data and data if set to true.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "prefix": {
                                "description": "The input string must have the prefix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value abc.xyz",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "suffix": {
                                "description": "The input string must have the suffix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A list of HTTP Hosts to match against. The match can be one of exact, prefix, suffix, or contains (substring match). Matches are always case sensitive unless the ignoreCase is set.\nLimited to 5 matches.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "paths": {
                          "block": {
                            "attributes": {
                              "contains": {
                                "description": "The input string must have the substring specified here. Note: empty contains match is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc.def",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "exact": {
                                "description": "The input string must match exactly the string specified here.\nExamples:\n* abc only matches the value abc.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ignore_case": {
                                "description": "If true, indicates the exact/prefix/suffix/contains matching should be case insensitive. For example, the matcher data will match both input string Data and data if set to true.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "prefix": {
                                "description": "The input string must have the prefix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value abc.xyz",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "suffix": {
                                "description": "The input string must have the suffix specified here. Note: empty prefix is not allowed, please use regex instead.\nExamples:\n* abc matches the value xyz.abc",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A list of paths to match against. The match can be one of exact, prefix, suffix, or contains (substring match). Matches are always case sensitive unless the ignoreCase is set.\nLimited to 5 matches.\nNote that this path match includes the query parameters. For gRPC services, this should be a fully-qualified name of the form /package.service/method.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Describes properties of one or more targets of a request. At least one of operations or notOperations must be specified. Limited to 5 operations. A match occurs when ANY operation (in operations or notOperations) matches. Within an operation, the match follows AND semantics across fields and OR semantics within a field, i.e. a match occurs when ANY path matches AND ANY header matches and ANY method matches.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Describes properties of one or more targets of a request",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A list of authorization HTTP rules to match against the incoming request.A policy match occurs when at least one HTTP rule matches the request or when no HTTP rules are specified in the policy. At least one HTTP Rule is required for Allow or Deny Action.\nLimited to 5 rules.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "target": {
        "block": {
          "attributes": {
            "load_balancing_scheme": {
              "description": "All gateways and forwarding rules referenced by this policy and extensions must share the same load balancing scheme.\nFor more information, refer to [Backend services overview](https://cloud.google.com/load-balancing/docs/backend-service). Possible values: [\"INTERNAL_MANAGED\", \"EXTERNAL_MANAGED\", \"INTERNAL_SELF_MANAGED\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resources": {
              "description": "A list of references to the Forwarding Rules on which this policy will be applied.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Specifies the set of resources to which this policy should be applied to.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleNetworkSecurityAuthzPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityAuthzPolicy), &result)
	return &result
}
