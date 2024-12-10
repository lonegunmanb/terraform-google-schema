package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePrivatecaCertificate = `{
  "block": {
    "attributes": {
      "certificate_authority": {
        "description": "The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from\na Certificate Authority with resource name 'projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca',\nargument 'pool' should be set to 'projects/my-project/locations/us-central1/caPools/my-pool', argument 'certificate_authority'\nshould be set to 'my-ca'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_description": {
        "computed": true,
        "description": "Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aia_issuing_certificate_urls": [
                "list",
                "string"
              ],
              "authority_key_id": [
                "list",
                [
                  "object",
                  {
                    "key_id": "string"
                  }
                ]
              ],
              "cert_fingerprint": [
                "list",
                [
                  "object",
                  {
                    "sha256_hash": "string"
                  }
                ]
              ],
              "crl_distribution_points": [
                "list",
                "string"
              ],
              "public_key": [
                "list",
                [
                  "object",
                  {
                    "format": "string",
                    "key": "string"
                  }
                ]
              ],
              "subject_description": [
                "list",
                [
                  "object",
                  {
                    "hex_serial_number": "string",
                    "lifetime": "string",
                    "not_after_time": "string",
                    "not_before_time": "string",
                    "subject": [
                      "list",
                      [
                        "object",
                        {
                          "common_name": "string",
                          "country_code": "string",
                          "locality": "string",
                          "organization": "string",
                          "organizational_unit": "string",
                          "postal_code": "string",
                          "province": "string",
                          "street_address": "string"
                        }
                      ]
                    ],
                    "subject_alt_name": [
                      "list",
                      [
                        "object",
                        {
                          "custom_sans": [
                            "list",
                            [
                              "object",
                              {
                                "critical": "bool",
                                "obect_id": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "object_id_path": [
                                        "list",
                                        "number"
                                      ]
                                    }
                                  ]
                                ],
                                "value": "string"
                              }
                            ]
                          ],
                          "dns_names": [
                            "list",
                            "string"
                          ],
                          "email_addresses": [
                            "list",
                            "string"
                          ],
                          "ip_addresses": [
                            "list",
                            "string"
                          ],
                          "uris": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "subject_key_id": [
                "list",
                [
                  "object",
                  {
                    "key_id": "string"
                  }
                ]
              ],
              "x509_description": [
                "list",
                [
                  "object",
                  {
                    "additional_extensions": [
                      "list",
                      [
                        "object",
                        {
                          "critical": "bool",
                          "object_id": [
                            "list",
                            [
                              "object",
                              {
                                "object_id_path": [
                                  "list",
                                  "number"
                                ]
                              }
                            ]
                          ],
                          "value": "string"
                        }
                      ]
                    ],
                    "aia_ocsp_servers": [
                      "list",
                      "string"
                    ],
                    "ca_options": [
                      "list",
                      [
                        "object",
                        {
                          "is_ca": "bool",
                          "max_issuer_path_length": "number"
                        }
                      ]
                    ],
                    "key_usage": [
                      "list",
                      [
                        "object",
                        {
                          "base_key_usage": [
                            "list",
                            [
                              "object",
                              {
                                "cert_sign": "bool",
                                "content_commitment": "bool",
                                "crl_sign": "bool",
                                "data_encipherment": "bool",
                                "decipher_only": "bool",
                                "digital_signature": "bool",
                                "encipher_only": "bool",
                                "key_agreement": "bool",
                                "key_encipherment": "bool"
                              }
                            ]
                          ],
                          "extended_key_usage": [
                            "list",
                            [
                              "object",
                              {
                                "client_auth": "bool",
                                "code_signing": "bool",
                                "email_protection": "bool",
                                "ocsp_signing": "bool",
                                "server_auth": "bool",
                                "time_stamping": "bool"
                              }
                            ]
                          ],
                          "unknown_extended_key_usages": [
                            "list",
                            [
                              "object",
                              {
                                "object_id_path": [
                                  "list",
                                  "number"
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "name_constraints": [
                      "list",
                      [
                        "object",
                        {
                          "critical": "bool",
                          "excluded_dns_names": [
                            "list",
                            "string"
                          ],
                          "excluded_email_addresses": [
                            "list",
                            "string"
                          ],
                          "excluded_ip_ranges": [
                            "list",
                            "string"
                          ],
                          "excluded_uris": [
                            "list",
                            "string"
                          ],
                          "permitted_dns_names": [
                            "list",
                            "string"
                          ],
                          "permitted_email_addresses": [
                            "list",
                            "string"
                          ],
                          "permitted_ip_ranges": [
                            "list",
                            "string"
                          ],
                          "permitted_uris": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "policy_ids": [
                      "list",
                      [
                        "object",
                        {
                          "object_id_path": [
                            "list",
                            "number"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "certificate_template": {
        "description": "The resource name for a CertificateTemplate used to issue this certificate,\nin the format 'projects/*/locations/*/certificateTemplates/*'. If this is specified,\nthe caller must have the necessary permission to use this template. If this is\nomitted, no template will be used. This template must be in the same location\nas the Certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time that this resource was created on the server.\nThis is in RFC3339 text format.",
        "description_kind": "plain",
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
      "issuer_certificate_authority": {
        "computed": true,
        "description": "The resource name of the issuing CertificateAuthority in the format 'projects/*/locations/*/caPools/*/certificateAuthorities/*'.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels with user-defined metadata to apply to this resource.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "lifetime": {
        "description": "The desired lifetime of the CA certificate. Used to create the \"notBeforeTime\" and\n\"notAfterTime\" fields inside an X.509 certificate. A duration in seconds with up to nine\nfractional digits, terminated by 's'. Example: \"3.5s\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Location of the Certificate. A full list of valid locations can be found by\nrunning 'gcloud privateca locations list'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name for this Certificate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pem_certificate": {
        "computed": true,
        "description": "Output only. The pem-encoded, signed X.509 certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "pem_certificate_chain": {
        "computed": true,
        "description": "The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "pem_csr": {
        "description": "Immutable. A pem-encoded X.509 certificate signing request (CSR).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pool": {
        "description": "The name of the CaPool this Certificate belongs to.",
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
      "revocation_details": {
        "computed": true,
        "description": "Output only. Details regarding the revocation of this Certificate. This Certificate is\nconsidered revoked if and only if this field is present.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "revocation_state": "string",
              "revocation_time": "string"
            }
          ]
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
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time at which this CertificateAuthority was updated.\nThis is in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "config": {
        "block": {
          "block_types": {
            "public_key": {
              "block": {
                "attributes": {
                  "format": {
                    "description": "The format of the public key. Currently, only PEM format is supported. Possible values: [\"KEY_TYPE_UNSPECIFIED\", \"PEM\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Required. A public key. When this is specified in a request, the padding and encoding can be any of the options described by the respective 'KeyType' value. When this is generated by the service, it will always be an RFC 5280 SubjectPublicKeyInfo structure containing an algorithm identifier and a key. A base64-encoded string.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A PublicKey describes a public key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "subject_config": {
              "block": {
                "block_types": {
                  "subject": {
                    "block": {
                      "attributes": {
                        "common_name": {
                          "description": "The common name of the distinguished name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "country_code": {
                          "description": "The country code of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "locality": {
                          "description": "The locality or city of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "organization": {
                          "description": "The organization of the subject.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "organizational_unit": {
                          "description": "The organizational unit of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "postal_code": {
                          "description": "The postal code of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "province": {
                          "description": "The province, territory, or regional state of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "street_address": {
                          "description": "The street address of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Contains distinguished name fields such as the location and organization.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "subject_alt_name": {
                    "block": {
                      "attributes": {
                        "dns_names": {
                          "description": "Contains only valid, fully-qualified host names.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "email_addresses": {
                          "description": "Contains only valid RFC 2822 E-mail addresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "ip_addresses": {
                          "description": "Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "uris": {
                          "description": "Contains only valid RFC 3986 URIs.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "The subject alternative name fields.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies some of the values in a certificate that are related to the subject.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "subject_key_id": {
              "block": {
                "attributes": {
                  "key_id": {
                    "description": "The value of the KeyId in lowercase hexadecimal.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "When specified this provides a custom SKI to be used in the certificate. This should only be used to maintain a SKI of an existing CA originally created outside CA service, which was not generated using method (1) described in RFC 5280 section 4.2.1.2..",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "x509_config": {
              "block": {
                "attributes": {
                  "aia_ocsp_servers": {
                    "description": "Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the\n\"Authority Information Access\" extension in the certificate.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "additional_extensions": {
                    "block": {
                      "attributes": {
                        "critical": {
                          "description": "Indicates whether or not this extension is critical (i.e., if the client does not know how to\nhandle this extension, the client should consider this to be an error).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "value": {
                          "description": "The value of this X.509 extension. A base64-encoded string.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "object_id": {
                          "block": {
                            "attributes": {
                              "object_id_path": {
                                "description": "An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "number"
                                ]
                              }
                            },
                            "description": "Describes values that are relevant in a CA certificate.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies an X.509 extension, which may be used in different parts of X.509 objects like certificates, CSRs, and CRLs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "ca_options": {
                    "block": {
                      "attributes": {
                        "is_ca": {
                          "description": "When true, the \"CA\" in Basic Constraints extension will be set to true.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "max_issuer_path_length": {
                          "description": "Refers to the \"path length constraint\" in Basic Constraints extension. For a CA certificate, this value describes the depth of\nsubordinate CA certificates that are allowed. If this value is less than 0, the request will fail.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "non_ca": {
                          "description": "When true, the \"CA\" in Basic Constraints extension will be set to false.\nIf both 'is_ca' and 'non_ca' are unset, the extension will be omitted from the CA certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "zero_max_issuer_path_length": {
                          "description": "When true, the \"path length constraint\" in Basic Constraints extension will be set to 0.\nif both 'max_issuer_path_length' and 'zero_max_issuer_path_length' are unset,\nthe max path length will be omitted from the CA certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Describes values that are relevant in a CA certificate.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "key_usage": {
                    "block": {
                      "block_types": {
                        "base_key_usage": {
                          "block": {
                            "attributes": {
                              "cert_sign": {
                                "description": "The key may be used to sign certificates.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "content_commitment": {
                                "description": "The key may be used for cryptographic commitments. Note that this may also be referred to as \"non-repudiation\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "crl_sign": {
                                "description": "The key may be used sign certificate revocation lists.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "data_encipherment": {
                                "description": "The key may be used to encipher data.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "decipher_only": {
                                "description": "The key may be used to decipher only.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "digital_signature": {
                                "description": "The key may be used for digital signatures.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "encipher_only": {
                                "description": "The key may be used to encipher only.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "key_agreement": {
                                "description": "The key may be used in a key agreement protocol.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "key_encipherment": {
                                "description": "The key may be used to encipher other keys.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Describes high-level ways in which a key may be used.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "extended_key_usage": {
                          "block": {
                            "attributes": {
                              "client_auth": {
                                "description": "Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as \"TLS WWW client authentication\", though regularly used for non-WWW TLS.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "code_signing": {
                                "description": "Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as \"Signing of downloadable executable code client authentication\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "email_protection": {
                                "description": "Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as \"Email protection\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "ocsp_signing": {
                                "description": "Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as \"Signing OCSP responses\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "server_auth": {
                                "description": "Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as \"TLS WWW server authentication\", though regularly used for non-WWW TLS.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "time_stamping": {
                                "description": "Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as \"Binding the hash of an object to a time\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Describes high-level ways in which a key may be used.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "unknown_extended_key_usages": {
                          "block": {
                            "attributes": {
                              "object_id_path": {
                                "description": "An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "number"
                                ]
                              }
                            },
                            "description": "An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Indicates the intended use for keys that correspond to a certificate.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "name_constraints": {
                    "block": {
                      "attributes": {
                        "critical": {
                          "description": "Indicates whether or not the name constraints are marked critical.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "excluded_dns_names": {
                          "description": "Contains excluded DNS names. Any DNS name that can be\nconstructed by simply adding zero or more labels to\nthe left-hand side of the name satisfies the name constraint.\nFor example, 'example.com', 'www.example.com', 'www.sub.example.com'\nwould satisfy 'example.com' while 'example1.com' does not.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "excluded_email_addresses": {
                          "description": "Contains the excluded email addresses. The value can be a particular\nemail address, a hostname to indicate all email addresses on that host or\na domain with a leading period (e.g. '.example.com') to indicate\nall email addresses in that domain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "excluded_ip_ranges": {
                          "description": "Contains the excluded IP ranges. For IPv4 addresses, the ranges\nare expressed using CIDR notation as specified in RFC 4632.\nFor IPv6 addresses, the ranges are expressed in similar encoding as IPv4\naddresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "excluded_uris": {
                          "description": "Contains the excluded URIs that apply to the host part of the name.\nThe value can be a hostname or a domain with a\nleading period (like '.example.com')",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "permitted_dns_names": {
                          "description": "Contains permitted DNS names. Any DNS name that can be\nconstructed by simply adding zero or more labels to\nthe left-hand side of the name satisfies the name constraint.\nFor example, 'example.com', 'www.example.com', 'www.sub.example.com'\nwould satisfy 'example.com' while 'example1.com' does not.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "permitted_email_addresses": {
                          "description": "Contains the permitted email addresses. The value can be a particular\nemail address, a hostname to indicate all email addresses on that host or\na domain with a leading period (e.g. '.example.com') to indicate\nall email addresses in that domain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "permitted_ip_ranges": {
                          "description": "Contains the permitted IP ranges. For IPv4 addresses, the ranges\nare expressed using CIDR notation as specified in RFC 4632.\nFor IPv6 addresses, the ranges are expressed in similar encoding as IPv4\naddresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "permitted_uris": {
                          "description": "Contains the permitted URIs that apply to the host part of the name.\nThe value can be a hostname or a domain with a\nleading period (like '.example.com')",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Describes the X.509 name constraints extension.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "policy_ids": {
                    "block": {
                      "attributes": {
                        "object_id_path": {
                          "description": "An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        }
                      },
                      "description": "Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Describes how some of the technical X.509 fields in a certificate should be populated.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The config used to create a self-signed X.509 certificate or CSR.",
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

func GooglePrivatecaCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePrivatecaCertificate), &result)
	return &result
}
