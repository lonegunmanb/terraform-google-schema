package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClouddomainsRegistration = `{
  "block": {
    "attributes": {
      "contact_notices": {
        "description": "The list of contact notices that the caller acknowledges. Possible value is PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Time at which the automation was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description": "Required. The domain name. Unicode domain names must be expressed in Punycode format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_notices": {
        "description": "The list of domain notices that you acknowledge. Possible value is HSTS_PRELOADED",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "expire_time": {
        "computed": true,
        "description": "Output only. Time at which the automation was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "issues": {
        "computed": true,
        "description": "Output only. The set of issues with the Registration that require attention.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "labels": {
        "description": "Set of labels associated with the Registration.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. Name of the Registration resource, in the format projects/*/locations/*/registrations/\u003cdomain_name\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "register_failure_reason": {
        "computed": true,
        "description": "Output only. The reason the domain registration failed. Only set for domains in REGISTRATION_FAILED state.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The current state of the Registration.",
        "description_kind": "plain",
        "type": "string"
      },
      "supported_privacy": {
        "computed": true,
        "description": "Output only. Set of options for the contactSettings.privacy field that this Registration supports.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
      "contact_settings": {
        "block": {
          "attributes": {
            "privacy": {
              "description": "Required. Privacy setting for the contacts associated with the Registration.\nValues are PUBLIC_CONTACT_DATA, PRIVATE_CONTACT_DATA, and REDACTED_CONTACT_DATA",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "admin_contact": {
              "block": {
                "attributes": {
                  "email": {
                    "description": "Required. Email address of the contact.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "fax_number": {
                    "description": "Fax number of the contact in international format. For example, \"+1-800-555-0123\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "phone_number": {
                    "description": "Required. Phone number of the contact in international format. For example, \"+1-800-555-0123\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "postal_address": {
                    "block": {
                      "attributes": {
                        "address_lines": {
                          "description": "Unstructured address lines describing the lower levels of an address.\nBecause values in addressLines do not have type information and may sometimes contain multiple values in a single\nfield (e.g. \"Austin, TX\"), it is important that the line order is clear. The order of address lines should be\n\"envelope order\" for the country/region of the address. In places where this can vary (e.g. Japan), address_language\nis used to make it explicit (e.g. \"ja\" for large-to-small ordering and \"ja-Latn\" or \"en\" for small-to-large). This way,\nthe most specific line of an address can be selected based on the language.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "administrative_area": {
                          "description": "Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state,\na province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community\n(e.g. \"Barcelona\" and not \"Catalonia\"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland\nthis should be left unpopulated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "locality": {
                          "description": "Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world\nwhere localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "organization": {
                          "description": "The name of the organization at the address.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "postal_code": {
                          "description": "Postal code of the address. Not all countries use or require postal codes to be present, but where they are used,\nthey may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "recipients": {
                          "description": "The recipient at the address. This field may, under certain circumstances, contain multiline information. For example,\nit might contain \"care of\" information.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "region_code": {
                          "description": "Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to\nensure the value is correct. See https://cldr.unicode.org/ and\nhttps://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: \"CH\" for Switzerland.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Required. Postal address of the contact.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Caution: Anyone with access to this email address, phone number, and/or postal address can take control of the domain.\n\nWarning: For new Registrations, the registrant receives an email confirmation that they must complete within 15 days to\navoid domain suspension.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "registrant_contact": {
              "block": {
                "attributes": {
                  "email": {
                    "description": "Required. Email address of the contact.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "fax_number": {
                    "description": "Fax number of the contact in international format. For example, \"+1-800-555-0123\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "phone_number": {
                    "description": "Required. Phone number of the contact in international format. For example, \"+1-800-555-0123\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "postal_address": {
                    "block": {
                      "attributes": {
                        "address_lines": {
                          "description": "Unstructured address lines describing the lower levels of an address.\nBecause values in addressLines do not have type information and may sometimes contain multiple values in a single\nfield (e.g. \"Austin, TX\"), it is important that the line order is clear. The order of address lines should be\n\"envelope order\" for the country/region of the address. In places where this can vary (e.g. Japan), address_language\nis used to make it explicit (e.g. \"ja\" for large-to-small ordering and \"ja-Latn\" or \"en\" for small-to-large). This way,\nthe most specific line of an address can be selected based on the language.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "administrative_area": {
                          "description": "Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state,\na province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community\n(e.g. \"Barcelona\" and not \"Catalonia\"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland\nthis should be left unpopulated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "locality": {
                          "description": "Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world\nwhere localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "organization": {
                          "description": "The name of the organization at the address.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "postal_code": {
                          "description": "Postal code of the address. Not all countries use or require postal codes to be present, but where they are used,\nthey may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "recipients": {
                          "description": "The recipient at the address. This field may, under certain circumstances, contain multiline information. For example,\nit might contain \"care of\" information.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "region_code": {
                          "description": "Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to\nensure the value is correct. See https://cldr.unicode.org/ and\nhttps://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: \"CH\" for Switzerland.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Required. Postal address of the contact.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Caution: Anyone with access to this email address, phone number, and/or postal address can take control of the domain.\n\nWarning: For new Registrations, the registrant receives an email confirmation that they must complete within 15 days to\navoid domain suspension.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "technical_contact": {
              "block": {
                "attributes": {
                  "email": {
                    "description": "Required. Email address of the contact.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "fax_number": {
                    "description": "Fax number of the contact in international format. For example, \"+1-800-555-0123\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "phone_number": {
                    "description": "Required. Phone number of the contact in international format. For example, \"+1-800-555-0123\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "postal_address": {
                    "block": {
                      "attributes": {
                        "address_lines": {
                          "description": "Unstructured address lines describing the lower levels of an address.\nBecause values in addressLines do not have type information and may sometimes contain multiple values in a single\nfield (e.g. \"Austin, TX\"), it is important that the line order is clear. The order of address lines should be\n\"envelope order\" for the country/region of the address. In places where this can vary (e.g. Japan), address_language\nis used to make it explicit (e.g. \"ja\" for large-to-small ordering and \"ja-Latn\" or \"en\" for small-to-large). This way,\nthe most specific line of an address can be selected based on the language.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "administrative_area": {
                          "description": "Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state,\na province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community\n(e.g. \"Barcelona\" and not \"Catalonia\"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland\nthis should be left unpopulated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "locality": {
                          "description": "Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world\nwhere localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "organization": {
                          "description": "The name of the organization at the address.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "postal_code": {
                          "description": "Postal code of the address. Not all countries use or require postal codes to be present, but where they are used,\nthey may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "recipients": {
                          "description": "The recipient at the address. This field may, under certain circumstances, contain multiline information. For example,\nit might contain \"care of\" information.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "region_code": {
                          "description": "Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to\nensure the value is correct. See https://cldr.unicode.org/ and\nhttps://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: \"CH\" for Switzerland.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Required. Postal address of the contact.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Caution: Anyone with access to this email address, phone number, and/or postal address can take control of the domain.\n\nWarning: For new Registrations, the registrant receives an email confirmation that they must complete within 15 days to\navoid domain suspension.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Required. Settings for contact information linked to the Registration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "dns_settings": {
        "block": {
          "block_types": {
            "custom_dns": {
              "block": {
                "attributes": {
                  "name_servers": {
                    "description": "Required. A list of name servers that store the DNS zone for this domain. Each name server is a domain\nname, with Unicode domain names expressed in Punycode format.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "ds_records": {
                    "block": {
                      "attributes": {
                        "algorithm": {
                          "description": "The algorithm used to generate the referenced DNSKEY.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "digest": {
                          "description": "The digest generated from the referenced DNSKEY.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "digest_type": {
                          "description": "The hash function used to generate the digest of the referenced DNSKEY.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_tag": {
                          "description": "The key tag of the record. Must be set in range 0 -- 65535.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "The list of DS records for this domain, which are used to enable DNSSEC. The domain's DNS provider can provide\nthe values to set here. If this field is empty, DNSSEC is disabled.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Configuration for an arbitrary DNS provider.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "glue_records": {
              "block": {
                "attributes": {
                  "host_name": {
                    "description": "Required. Domain name of the host in Punycode format.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ipv4_addresses": {
                    "description": "List of IPv4 addresses corresponding to this host in the standard decimal format (e.g. 198.51.100.1).\nAt least one of ipv4_address and ipv6_address must be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ipv6_addresses": {
                    "description": "List of IPv4 addresses corresponding to this host in the standard decimal format (e.g. 198.51.100.1).\nAt least one of ipv4_address and ipv6_address must be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The list of glue records for this Registration. Commonly empty.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Settings controlling the DNS configuration of the Registration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "management_settings": {
        "block": {
          "attributes": {
            "preferred_renewal_method": {
              "computed": true,
              "description": "The desired renewal method for this Registration. The actual renewalMethod is automatically updated to reflect this choice.\nIf unset or equal to RENEWAL_METHOD_UNSPECIFIED, the actual renewalMethod is treated as if it were set to AUTOMATIC_RENEWAL.\nYou cannot use RENEWAL_DISABLED during resource creation, and you can update the renewal status only when the Registration\nresource has state ACTIVE or SUSPENDED.\n\nWhen preferredRenewalMethod is set to AUTOMATIC_RENEWAL, the actual renewalMethod can be set to RENEWAL_DISABLED in case of\nproblems with the billing account or reported domain abuse. In such cases, check the issues field on the Registration. After\nthe problem is resolved, the renewalMethod is automatically updated to preferredRenewalMethod in a few hours.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "renewal_method": {
              "computed": true,
              "description": "Output only. The actual renewal method for this Registration. When preferredRenewalMethod is set to AUTOMATIC_RENEWAL,\nthe actual renewalMethod can be equal to RENEWAL_DISABLED—for example, when there are problems with the billing account\nor reported domain abuse. In such cases, check the issues field on the Registration. After the problem is resolved, the\nrenewalMethod is automatically updated to preferredRenewalMethod in a few hours.",
              "description_kind": "plain",
              "type": "string"
            },
            "transfer_lock_state": {
              "computed": true,
              "description": "Controls whether the domain can be transferred to another registrar. Values are UNLOCKED or LOCKED.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Settings for management of the Registration, including renewal, billing, and transfer",
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
      },
      "yearly_price": {
        "block": {
          "attributes": {
            "currency_code": {
              "description": "The three-letter currency code defined in ISO 4217.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "units": {
              "description": "The whole units of the amount. For example if currencyCode is \"USD\", then 1 unit is one US dollar.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Required. Yearly price to register or renew the domain. The value that should be put here can be obtained from\nregistrations.retrieveRegisterParameters or registrations.searchDomains calls.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleClouddomainsRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClouddomainsRegistration), &result)
	return &result
}
