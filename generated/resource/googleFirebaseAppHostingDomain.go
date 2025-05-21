package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaseAppHostingDomain = `{
  "block": {
    "attributes": {
      "backend": {
        "description": "The ID of the Backend that this Domain is associated with",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time at which the domain was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain_status": {
        "computed": true,
        "description": "The status of a custom domain's linkage to the Backend.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cert_state": "string",
              "host_state": "string",
              "issues": [
                "list",
                [
                  "object",
                  {
                    "code": "number",
                    "details": "string",
                    "message": "string"
                  }
                ]
              ],
              "ownership_state": "string",
              "required_dns_updates": [
                "list",
                [
                  "object",
                  {
                    "check_time": "string",
                    "desired": [
                      "list",
                      [
                        "object",
                        {
                          "check_error": [
                            "list",
                            [
                              "object",
                              {
                                "code": "number",
                                "details": "string",
                                "message": "string"
                              }
                            ]
                          ],
                          "domain_name": "string",
                          "records": [
                            "list",
                            [
                              "object",
                              {
                                "domain_name": "string",
                                "rdata": "string",
                                "relevant_state": [
                                  "list",
                                  "string"
                                ],
                                "required_action": "string",
                                "type": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "discovered": [
                      "list",
                      [
                        "object",
                        {
                          "check_error": [
                            "list",
                            [
                              "object",
                              {
                                "code": "number",
                                "details": "string",
                                "message": "string"
                              }
                            ]
                          ],
                          "domain_name": "string",
                          "records": [
                            "list",
                            [
                              "object",
                              {
                                "domain_name": "string",
                                "rdata": "string",
                                "relevant_state": [
                                  "list",
                                  "string"
                                ],
                                "required_action": "string",
                                "type": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "domain_name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "delete_time": {
        "computed": true,
        "description": "Time at which the domain was deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "description": "Id of the domain to create.\nMust be a valid domain name, such as \"foo.com\"",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Server-computed checksum based on other values; may be sent\non update or delete to ensure operation is done on expected resource.",
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
        "description": "The location of the Backend that this Domain is associated with",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the domain, e.g.\n'projects/{project}/locations/{locationId}/backends/{backendId}/domains/{domainId}'",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "purge_time": {
        "computed": true,
        "description": "Time at which a soft-deleted domain will be purged, rendering in\npermanently deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-assigned, unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time at which the domain was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "serve": {
        "block": {
          "block_types": {
            "redirect": {
              "block": {
                "attributes": {
                  "status": {
                    "description": "The status code to use in a redirect response. Must be a valid HTTP 3XX\nstatus code. Defaults to 302 if not present.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uri": {
                    "description": "The URI of the redirect's intended destination. This URI will be\nprepended to the original request path. URI without a scheme are\nassumed to be HTTPS.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specifies redirect behavior for a domain.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The serving behavior of the domain. If specified, the domain will\nserve content other than its Backend's live content.",
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

func GoogleFirebaseAppHostingDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaseAppHostingDomain), &result)
	return &result
}
