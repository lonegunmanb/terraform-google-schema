package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOsConfigV2PolicyOrchestrator = `{
  "block": {
    "attributes": {
      "action": {
        "description": "Required. Action to be done by the orchestrator in\n'projects/{project_id}/zones/{zone_id}' locations defined by the\n'orchestration_scope'. Allowed values:\n- 'UPSERT' - Orchestrator will create or update target resources.\n- 'DELETE' - Orchestrator will delete target resources, if they exist",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Timestamp when the policy orchestrator resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. Freeform text describing the purpose of the resource.",
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
        "description": "Optional. Labels as key value pairs\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Immutable. Identifier. In form of\n* 'organizations/{organization_id}/locations/global/policyOrchestrators/{orchestrator_id}'\n* 'folders/{folder_id}/locations/global/policyOrchestrators/{orchestrator_id}'\n* 'projects/{project_id_or_number}/locations/global/policyOrchestrators/{orchestrator_id}'",
        "description_kind": "plain",
        "type": "string"
      },
      "orchestration_state": {
        "computed": true,
        "description": "Describes the state of the orchestration process.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "current_iteration_state": [
                "list",
                [
                  "object",
                  {
                    "error": [
                      "list",
                      [
                        "object",
                        {
                          "code": "number",
                          "details": [
                            "list",
                            [
                              "object",
                              {
                                "type_url": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "message": "string"
                        }
                      ]
                    ],
                    "failed_actions": "string",
                    "finish_time": "string",
                    "performed_actions": "string",
                    "progress": "number",
                    "rollout_resource": "string",
                    "start_time": "string",
                    "state": "string"
                  }
                ]
              ],
              "previous_iteration_state": [
                "list",
                [
                  "object",
                  {
                    "error": [
                      "list",
                      [
                        "object",
                        {
                          "code": "number",
                          "details": [
                            "list",
                            [
                              "object",
                              {
                                "type_url": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "message": "string"
                        }
                      ]
                    ],
                    "failed_actions": "string",
                    "finish_time": "string",
                    "performed_actions": "string",
                    "progress": "number",
                    "rollout_resource": "string",
                    "start_time": "string",
                    "state": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "policy_orchestrator_id": {
        "description": "Required. The logical identifier of the policy orchestrator, with the following\nrestrictions:\n\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.\n* Must be unique within the parent.",
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
      "reconciling": {
        "computed": true,
        "description": "Output only. Set to true, if the there are ongoing changes being applied by the\norchestrator.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "description": "Optional. State of the orchestrator. Can be updated to change orchestrator behaviour.\nAllowed values:\n- 'ACTIVE' - orchestrator is actively looking for actions to be taken.\n- 'STOPPED' - orchestrator won't make any changes.\n\nNote: There might be more states added in the future. We use string here\ninstead of an enum, to avoid the need of propagating new states to all the\nclient code.",
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
        "description": "Output only. Timestamp when the policy orchestrator resource was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "orchestrated_resource": {
        "block": {
          "attributes": {
            "id": {
              "description": "Optional. ID of the resource to be used while generating set of affected resources.\n\nFor UPSERT action the value is auto-generated during PolicyOrchestrator\ncreation when not set. When the value is set it should following next\nrestrictions:\n\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.\n* Must be unique within the project.\n\nFor DELETE action, ID must be specified explicitly during\nPolicyOrchestrator creation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "os_policy_assignment_v1_payload": {
              "block": {
                "attributes": {
                  "baseline": {
                    "computed": true,
                    "description": "Output only. Indicates that this revision has been successfully rolled out in this zone\nand new VMs will be assigned OS policies from this revision.\n\nFor a given OS policy assignment, there is only one revision with a value\nof 'true' for this field.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "deleted": {
                    "computed": true,
                    "description": "Output only. Indicates that this revision deletes the OS policy assignment.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "description": {
                    "description": "OS policy assignment description.\nLength of the description is limited to 1024 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Resource name.\n\nFormat:\n'projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}'\n\nThis field is ignored when you create an OS policy assignment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reconciling": {
                    "computed": true,
                    "description": "Output only. Indicates that reconciliation is in progress for the revision.\nThis value is 'true' when the 'rollout_state' is one of:\n* IN_PROGRESS\n* CANCELLING",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "revision_create_time": {
                    "computed": true,
                    "description": "Output only. The timestamp that the revision was created.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "revision_id": {
                    "computed": true,
                    "description": "Output only. The assignment revision ID\nA new revision is committed whenever a rollout is triggered for a OS policy\nassignment",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rollout_state": {
                    "computed": true,
                    "description": "Output only. OS policy assignment rollout state\nPossible values:\nROLLOUT_STATE_UNSPECIFIED\nIN_PROGRESS\nCANCELLING\nCANCELLED\nSUCCEEDED",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "uid": {
                    "computed": true,
                    "description": "Output only. Server generated unique id for the OS policy assignment resource.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "block_types": {
                  "instance_filter": {
                    "block": {
                      "attributes": {
                        "all": {
                          "description": "Target all VMs in the project. If true, no other criteria is\npermitted.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "exclusion_labels": {
                          "block": {
                            "attributes": {
                              "labels": {
                                "description": "Labels are identified by key/value pairs in this map.\nA VM should contain all the key/value pairs specified in this\nmap to be selected.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description": "List of label sets used for VM exclusion.\n\nIf the list has more than one label set, the VM is excluded if any\nof the label sets are applicable for the VM.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "inclusion_labels": {
                          "block": {
                            "attributes": {
                              "labels": {
                                "description": "Labels are identified by key/value pairs in this map.\nA VM should contain all the key/value pairs specified in this\nmap to be selected.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description": "List of label sets used for VM inclusion.\n\nIf the list has more than one 'LabelSet', the VM is included if any\nof the label sets are applicable for the VM.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "inventories": {
                          "block": {
                            "attributes": {
                              "os_short_name": {
                                "description": "Required. The OS short name",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "os_version": {
                                "description": "The OS version\n\nPrefix matches are supported if asterisk(*) is provided as the\nlast character. For example, to match all versions with a major\nversion of '7', specify the following value for this field '7.*'\n\nAn empty string matches all OS versions.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "List of inventories to select VMs.\n\nA VM is selected if its inventory data matches at least one of the\nfollowing inventories.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Filters to select target VMs for an assignment.\n\nIf more than one filter criteria is specified below, a VM will be selected\nif and only if it satisfies all of them.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "os_policies": {
                    "block": {
                      "attributes": {
                        "allow_no_resource_group_match": {
                          "description": "This flag determines the OS policy compliance status when none of the\nresource groups within the policy are applicable for a VM. Set this value\nto 'true' if the policy needs to be reported as compliant even if the\npolicy has nothing to validate or enforce.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "description": {
                          "description": "Policy description.\nLength of the description is limited to 1024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "description": "Required. The id of the OS policy with the following restrictions:\n\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.\n* Must be unique within the assignment.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "mode": {
                          "description": "Required. Policy mode\nPossible values:\nMODE_UNSPECIFIED\nVALIDATION\nENFORCEMENT",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "resource_groups": {
                          "block": {
                            "block_types": {
                              "inventory_filters": {
                                "block": {
                                  "attributes": {
                                    "os_short_name": {
                                      "description": "Required. The OS short name",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "os_version": {
                                      "description": "The OS version\n\nPrefix matches are supported if asterisk(*) is provided as the\nlast character. For example, to match all versions with a major\nversion of '7', specify the following value for this field '7.*'\n\nAn empty string matches all OS versions.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "List of inventory filters for the resource group.\n\nThe resources in this resource group are applied to the target VM if it\nsatisfies at least one of the following inventory filters.\n\nFor example, to apply this resource group to VMs running either 'RHEL' or\n'CentOS' operating systems, specify 2 items for the list with following\nvalues:\ninventory_filters[0].os_short_name='rhel' and\ninventory_filters[1].os_short_name='centos'\n\nIf the list is empty, this resource group will be applied to the target\nVM unconditionally.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "resources": {
                                "block": {
                                  "attributes": {
                                    "id": {
                                      "description": "Required. The id of the resource with the following restrictions:\n\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.\n* Must be unique within the OS policy.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "exec": {
                                      "block": {
                                        "block_types": {
                                          "enforce": {
                                            "block": {
                                              "attributes": {
                                                "args": {
                                                  "description": "Optional arguments to pass to the source during execution.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                },
                                                "interpreter": {
                                                  "description": "Required. The script interpreter to use.\nPossible values:\nINTERPRETER_UNSPECIFIED\nNONE\nSHELL\nPOWERSHELL",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "output_file_path": {
                                                  "description": "Only recorded for enforce Exec.\nPath to an output file (that is created by this Exec) whose\ncontent will be recorded in OSPolicyResourceCompliance after a\nsuccessful run. Absence or failure to read this file will result in\nthis ExecResource being non-compliant. Output file size is limited to\n500K bytes.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "script": {
                                                  "description": "An inline script.\nThe size of the script is limited to 32KiB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "file": {
                                                  "block": {
                                                    "attributes": {
                                                      "allow_insecure": {
                                                        "description": "Defaults to false. When false, files are subject to validations\nbased on the file type:\n\nRemote: A checksum must be specified.\nCloud Storage: An object generation number must be specified.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
                                                      },
                                                      "local_path": {
                                                        "description": "A local path within the VM to use.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "gcs": {
                                                        "block": {
                                                          "attributes": {
                                                            "bucket": {
                                                              "description": "Required. Bucket of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            },
                                                            "generation": {
                                                              "description": "Generation number of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "object": {
                                                              "description": "Required. Name of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available as a Cloud Storage Object.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "remote": {
                                                        "block": {
                                                          "attributes": {
                                                            "sha256_checksum": {
                                                              "description": "SHA256 checksum of the remote file.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "uri": {
                                                              "description": "Required. URI from which to fetch the object. It should contain both the\nprotocol and path following the format '{protocol}://{location}'.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available via some URI.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      }
                                                    },
                                                    "description": "A remote or local file.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "A file or script to execute.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "validate": {
                                            "block": {
                                              "attributes": {
                                                "args": {
                                                  "description": "Optional arguments to pass to the source during execution.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                },
                                                "interpreter": {
                                                  "description": "Required. The script interpreter to use.\nPossible values:\nINTERPRETER_UNSPECIFIED\nNONE\nSHELL\nPOWERSHELL",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "output_file_path": {
                                                  "description": "Only recorded for enforce Exec.\nPath to an output file (that is created by this Exec) whose\ncontent will be recorded in OSPolicyResourceCompliance after a\nsuccessful run. Absence or failure to read this file will result in\nthis ExecResource being non-compliant. Output file size is limited to\n500K bytes.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "script": {
                                                  "description": "An inline script.\nThe size of the script is limited to 32KiB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "file": {
                                                  "block": {
                                                    "attributes": {
                                                      "allow_insecure": {
                                                        "description": "Defaults to false. When false, files are subject to validations\nbased on the file type:\n\nRemote: A checksum must be specified.\nCloud Storage: An object generation number must be specified.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
                                                      },
                                                      "local_path": {
                                                        "description": "A local path within the VM to use.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "gcs": {
                                                        "block": {
                                                          "attributes": {
                                                            "bucket": {
                                                              "description": "Required. Bucket of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            },
                                                            "generation": {
                                                              "description": "Generation number of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "object": {
                                                              "description": "Required. Name of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available as a Cloud Storage Object.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "remote": {
                                                        "block": {
                                                          "attributes": {
                                                            "sha256_checksum": {
                                                              "description": "SHA256 checksum of the remote file.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "uri": {
                                                              "description": "Required. URI from which to fetch the object. It should contain both the\nprotocol and path following the format '{protocol}://{location}'.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available via some URI.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      }
                                                    },
                                                    "description": "A remote or local file.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "A file or script to execute.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "min_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "A resource that allows executing scripts on the VM.\n\nThe 'ExecResource' has 2 stages: 'validate' and 'enforce' and both stages\naccept a script as an argument to execute.\n\nWhen the 'ExecResource' is applied by the agent, it first executes the\nscript in the 'validate' stage. The 'validate' stage can signal that the\n'ExecResource' is already in the desired state by returning an exit code\nof '100'. If the 'ExecResource' is not in the desired state, it should\nreturn an exit code of '101'. Any other exit code returned by this stage\nis considered an error.\n\nIf the 'ExecResource' is not in the desired state based on the exit code\nfrom the 'validate' stage, the agent proceeds to execute the script from\nthe 'enforce' stage. If the 'ExecResource' is already in the desired\nstate, the 'enforce' stage will not be run.\nSimilar to 'validate' stage, the 'enforce' stage should return an exit\ncode of '100' to indicate that the resource in now in its desired state.\nAny other exit code is considered an error.\n\nNOTE: An exit code of '100' was chosen over '0' (and '101' vs '1') to\nhave an explicit indicator of 'in desired state', 'not in desired state'\nand errors. Because, for example, Powershell will always return an exit\ncode of '0' unless an 'exit' statement is provided in the script. So, for\nreasons of consistency and being explicit, exit codes '100' and '101'\nwere chosen.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "file": {
                                      "block": {
                                        "attributes": {
                                          "content": {
                                            "description": "A a file with this content.\nThe size of the content is limited to 32KiB.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "path": {
                                            "description": "Required. The absolute path of the file within the VM.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "permissions": {
                                            "description": "Consists of three octal digits which represent, in\norder, the permissions of the owner, group, and other users for the\nfile (similarly to the numeric mode used in the linux chmod\nutility). Each digit represents a three bit number with the 4 bit\ncorresponding to the read permissions, the 2 bit corresponds to the\nwrite bit, and the one bit corresponds to the execute permission.\nDefault behavior is 755.\n\nBelow are some examples of permissions and their associated values:\nread, write, and execute: 7\nread and execute: 5\nread and write: 6\nread only: 4",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "state": {
                                            "description": "Required. Desired state of the file.\nPossible values:\nDESIRED_STATE_UNSPECIFIED\nPRESENT\nABSENT\nCONTENTS_MATCH",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "file": {
                                            "block": {
                                              "attributes": {
                                                "allow_insecure": {
                                                  "description": "Defaults to false. When false, files are subject to validations\nbased on the file type:\n\nRemote: A checksum must be specified.\nCloud Storage: An object generation number must be specified.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                },
                                                "local_path": {
                                                  "description": "A local path within the VM to use.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "gcs": {
                                                  "block": {
                                                    "attributes": {
                                                      "bucket": {
                                                        "description": "Required. Bucket of the Cloud Storage object.",
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "generation": {
                                                        "description": "Generation number of the Cloud Storage object.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "object": {
                                                        "description": "Required. Name of the Cloud Storage object.",
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description": "Specifies a file available as a Cloud Storage Object.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                },
                                                "remote": {
                                                  "block": {
                                                    "attributes": {
                                                      "sha256_checksum": {
                                                        "description": "SHA256 checksum of the remote file.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "uri": {
                                                        "description": "Required. URI from which to fetch the object. It should contain both the\nprotocol and path following the format '{protocol}://{location}'.",
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description": "Specifies a file available via some URI.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "A remote or local file.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "A resource that manages the state of a file.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "pkg": {
                                      "block": {
                                        "attributes": {
                                          "desired_state": {
                                            "description": "Required. The desired state the agent should maintain for this package.\nPossible values:\nDESIRED_STATE_UNSPECIFIED\nINSTALLED\nREMOVED",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "apt": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Required. Package name.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A package managed by APT.\n- install: 'apt-get update \u0026\u0026 apt-get -y install [name]'\n- remove: 'apt-get -y remove [name]'",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "deb": {
                                            "block": {
                                              "attributes": {
                                                "pull_deps": {
                                                  "description": "Whether dependencies should also be installed.\n- install when false: 'dpkg -i package'\n- install when true: 'apt-get update \u0026\u0026 apt-get -y install\npackage.deb'",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                }
                                              },
                                              "block_types": {
                                                "source": {
                                                  "block": {
                                                    "attributes": {
                                                      "allow_insecure": {
                                                        "description": "Defaults to false. When false, files are subject to validations\nbased on the file type:\n\nRemote: A checksum must be specified.\nCloud Storage: An object generation number must be specified.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
                                                      },
                                                      "local_path": {
                                                        "description": "A local path within the VM to use.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "gcs": {
                                                        "block": {
                                                          "attributes": {
                                                            "bucket": {
                                                              "description": "Required. Bucket of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            },
                                                            "generation": {
                                                              "description": "Generation number of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "object": {
                                                              "description": "Required. Name of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available as a Cloud Storage Object.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "remote": {
                                                        "block": {
                                                          "attributes": {
                                                            "sha256_checksum": {
                                                              "description": "SHA256 checksum of the remote file.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "uri": {
                                                              "description": "Required. URI from which to fetch the object. It should contain both the\nprotocol and path following the format '{protocol}://{location}'.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available via some URI.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      }
                                                    },
                                                    "description": "A remote or local file.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "min_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "A deb package file. dpkg packages only support INSTALLED state.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "googet": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Required. Package name.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A package managed by GooGet.\n- install: 'googet -noconfirm install package'\n- remove: 'googet -noconfirm remove package'",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "msi": {
                                            "block": {
                                              "attributes": {
                                                "properties": {
                                                  "description": "Additional properties to use during installation.\nThis should be in the format of Property=Setting.\nAppended to the defaults of 'ACTION=INSTALL\nREBOOT=ReallySuppress'.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "source": {
                                                  "block": {
                                                    "attributes": {
                                                      "allow_insecure": {
                                                        "description": "Defaults to false. When false, files are subject to validations\nbased on the file type:\n\nRemote: A checksum must be specified.\nCloud Storage: An object generation number must be specified.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
                                                      },
                                                      "local_path": {
                                                        "description": "A local path within the VM to use.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "gcs": {
                                                        "block": {
                                                          "attributes": {
                                                            "bucket": {
                                                              "description": "Required. Bucket of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            },
                                                            "generation": {
                                                              "description": "Generation number of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "object": {
                                                              "description": "Required. Name of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available as a Cloud Storage Object.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "remote": {
                                                        "block": {
                                                          "attributes": {
                                                            "sha256_checksum": {
                                                              "description": "SHA256 checksum of the remote file.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "uri": {
                                                              "description": "Required. URI from which to fetch the object. It should contain both the\nprotocol and path following the format '{protocol}://{location}'.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available via some URI.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      }
                                                    },
                                                    "description": "A remote or local file.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "min_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "An MSI package. MSI packages only support INSTALLED state.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "rpm": {
                                            "block": {
                                              "attributes": {
                                                "pull_deps": {
                                                  "description": "Whether dependencies should also be installed.\n- install when false: 'rpm --upgrade --replacepkgs package.rpm'\n- install when true: 'yum -y install package.rpm' or\n'zypper -y install package.rpm'",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                }
                                              },
                                              "block_types": {
                                                "source": {
                                                  "block": {
                                                    "attributes": {
                                                      "allow_insecure": {
                                                        "description": "Defaults to false. When false, files are subject to validations\nbased on the file type:\n\nRemote: A checksum must be specified.\nCloud Storage: An object generation number must be specified.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
                                                      },
                                                      "local_path": {
                                                        "description": "A local path within the VM to use.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "gcs": {
                                                        "block": {
                                                          "attributes": {
                                                            "bucket": {
                                                              "description": "Required. Bucket of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            },
                                                            "generation": {
                                                              "description": "Generation number of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "object": {
                                                              "description": "Required. Name of the Cloud Storage object.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available as a Cloud Storage Object.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "remote": {
                                                        "block": {
                                                          "attributes": {
                                                            "sha256_checksum": {
                                                              "description": "SHA256 checksum of the remote file.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "uri": {
                                                              "description": "Required. URI from which to fetch the object. It should contain both the\nprotocol and path following the format '{protocol}://{location}'.",
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "description": "Specifies a file available via some URI.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      }
                                                    },
                                                    "description": "A remote or local file.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "min_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "An RPM package file. RPM packages only support INSTALLED state.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "yum": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Required. Package name.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A package managed by YUM.\n- install: 'yum -y install package'\n- remove: 'yum -y remove package'",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "zypper": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Required. Package name.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A package managed by Zypper.\n- install: 'zypper -y install package'\n- remove: 'zypper -y rm package'",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "A resource that manages a system package.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "repository": {
                                      "block": {
                                        "block_types": {
                                          "apt": {
                                            "block": {
                                              "attributes": {
                                                "archive_type": {
                                                  "description": "Required. Type of archive files in this repository.\nPossible values:\nARCHIVE_TYPE_UNSPECIFIED\nDEB\nDEB_SRC",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "components": {
                                                  "description": "Required. List of components for this repository. Must contain at least one\nitem.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                },
                                                "distribution": {
                                                  "description": "Required. Distribution of this repository.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "gpg_key": {
                                                  "description": "URI of the key file for this repository. The agent maintains a\nkeyring at '/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg'.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "uri": {
                                                  "description": "Required. URI for this repository.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Represents a single apt package repository. These will be added to\na repo file that will be managed at\n'/etc/apt/sources.list.d/google_osconfig.list'.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "goo": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Required. The name of the repository.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "url": {
                                                  "description": "Required. The url of the repository.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Represents a Goo package repository. These are added to a repo file\nthat is managed at\n'C:/ProgramData/GooGet/repos/google_osconfig.repo'.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "yum": {
                                            "block": {
                                              "attributes": {
                                                "base_url": {
                                                  "description": "Required. The location of the repository directory.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "display_name": {
                                                  "description": "The display name of the repository.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "gpg_keys": {
                                                  "description": "URIs of GPG keys.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                },
                                                "id": {
                                                  "description": "Required. A one word, unique name for this repository. This is  the 'repo\nid' in the yum config file and also the 'display_name' if\n'display_name' is omitted. This id is also used as the unique\nidentifier when checking for resource conflicts.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Represents a single yum package repository. These are added to a\nrepo file that is managed at\n'/etc/yum.repos.d/google_osconfig.repo'.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "zypper": {
                                            "block": {
                                              "attributes": {
                                                "base_url": {
                                                  "description": "Required. The location of the repository directory.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "display_name": {
                                                  "description": "The display name of the repository.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "gpg_keys": {
                                                  "description": "URIs of GPG keys.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                },
                                                "id": {
                                                  "description": "Required. A one word, unique name for this repository. This is the 'repo\nid' in the zypper config file and also the 'display_name' if\n'display_name' is omitted. This id is also used as the unique\nidentifier when checking for GuestPolicy conflicts.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Represents a single zypper package repository. These are added to a\nrepo file that is managed at\n'/etc/zypp/repos.d/google_osconfig.repo'.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "A resource that manages a package repository.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Required. List of resources configured for this resource group.\nThe resources are executed in the exact order specified here.",
                                  "description_kind": "plain"
                                },
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Required. List of resource groups for the policy.\nFor a particular VM, resource groups are evaluated in the order specified\nand the first resource group that is applicable is selected and the rest\nare ignored.\n\nIf none of the resource groups are applicable for a VM, the VM is\nconsidered to be non-compliant w.r.t this policy. This behavior can be\ntoggled by the flag 'allow_no_resource_group_match'",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Required. List of OS policies to be applied to the VMs.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "rollout": {
                    "block": {
                      "attributes": {
                        "min_wait_duration": {
                          "description": "Required. This determines the minimum duration of time to wait after the\nconfiguration changes are applied through the current rollout. A\nVM continues to count towards the 'disruption_budget' at least\nuntil this duration of time has passed after configuration changes are\napplied.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "disruption_budget": {
                          "block": {
                            "attributes": {
                              "fixed": {
                                "description": "Specifies a fixed value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "percent": {
                                "description": "Specifies the relative value defined as a percentage, which will be\nmultiplied by a reference value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Message encapsulating a value that can be either absolute (\"fixed\") or\nrelative (\"percent\") to a value.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Message to configure the rollout at the zonal level for the OS policy\nassignment.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "OS policy assignment is an API resource that is used to\napply a set of OS policies to a dynamically targeted group of Compute Engine\nVM instances.\n\nAn OS policy is used to define the desired state configuration for a\nCompute Engine VM instance through a set of configuration resources that\nprovide capabilities such as installing or removing software packages, or\nexecuting a script.\n\nFor more information about the OS policy resource definitions and examples,\nsee\n[OS policy and OS policy\nassignment](https://cloud.google.com/compute/docs/os-configuration-management/working-with-os-policies).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Represents a resource that is being orchestrated by the policy orchestrator.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "orchestration_scope": {
        "block": {
          "block_types": {
            "selectors": {
              "block": {
                "block_types": {
                  "location_selector": {
                    "block": {
                      "attributes": {
                        "included_locations": {
                          "description": "Optional. Names of the locations in scope.\nFormat: 'us-central1-a'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Selector containing locations in scope.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "resource_hierarchy_selector": {
                    "block": {
                      "attributes": {
                        "included_folders": {
                          "description": "Optional. Names of the folders in scope.\nFormat: 'folders/{folder_id}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "included_projects": {
                          "description": "Optional. Names of the projects in scope.\nFormat: 'projects/{project_number}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Selector containing Cloud Resource Manager resource hierarchy nodes.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Selectors of the orchestration scope. There is a logical AND between each\nselector defined.\n\nWhen there is no explicit 'ResourceHierarchySelector' selector specified,\nthe scope is by default bounded to the parent of the policy orchestrator\nresource.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Defines a set of selectors which drive which resources are in scope of policy\norchestration.",
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

func GoogleOsConfigV2PolicyOrchestratorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOsConfigV2PolicyOrchestrator), &result)
	return &result
}
