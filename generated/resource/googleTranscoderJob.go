package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleTranscoderJob = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time the job was created.",
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
      "end_time": {
        "computed": true,
        "description": "The time the transcoding finished.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels associated with this job. You can use these to organize and group your jobs.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the transcoding job resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the job.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description": "The time the transcoding started.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the job.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_id": {
        "computed": true,
        "description": "Specify the templateId to use for populating Job.config.\nThe default is preset/web-hd, which is the only supported preset.",
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
      }
    },
    "block_types": {
      "config": {
        "block": {
          "block_types": {
            "ad_breaks": {
              "block": {
                "attributes": {
                  "start_time_offset": {
                    "computed": true,
                    "description": "Start time in seconds for the ad break, relative to the output file timeline",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Ad break.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "edit_list": {
              "block": {
                "attributes": {
                  "inputs": {
                    "computed": true,
                    "description": "List of values identifying files that should be used in this atom.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "key": {
                    "computed": true,
                    "description": "A unique key for this atom.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start_time_offset": {
                    "computed": true,
                    "description": "Start time in seconds for the atom, relative to the input file timeline. The default is '0s'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "List of input assets stored in Cloud Storage.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "elementary_streams": {
              "block": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "A unique key for this atom.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "audio_stream": {
                    "block": {
                      "attributes": {
                        "bitrate_bps": {
                          "description": "Audio bitrate in bits per second.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "channel_count": {
                          "computed": true,
                          "description": "Number of audio channels. The default is '2'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "channel_layout": {
                          "computed": true,
                          "description": "A list of channel names specifying layout of the audio channels. The default is [\"fl\", \"fr\"].",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "codec": {
                          "computed": true,
                          "description": "The codec for this audio stream. The default is 'aac'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sample_rate_hertz": {
                          "computed": true,
                          "description": "The audio sample rate in Hertz. The default is '48000'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Encoding of an audio stream.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "video_stream": {
                    "block": {
                      "block_types": {
                        "h264": {
                          "block": {
                            "attributes": {
                              "bitrate_bps": {
                                "description": "The video bitrate in bits per second.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "crf_level": {
                                "computed": true,
                                "description": "Target CRF level. The default is '21'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "entropy_coder": {
                                "computed": true,
                                "description": "The entropy coder to use. The default is 'cabac'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "frame_rate": {
                                "description": "The target video frame rate in frames per second (FPS).",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "gop_duration": {
                                "computed": true,
                                "description": "Select the GOP size based on the specified duration. The default is '3s'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "height_pixels": {
                                "computed": true,
                                "description": "The height of the video in pixels.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "pixel_format": {
                                "computed": true,
                                "description": "Pixel format to use. The default is 'yuv420p'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "preset": {
                                "computed": true,
                                "description": "Enforces the specified codec preset. The default is 'veryfast'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "profile": {
                                "computed": true,
                                "description": "Enforces the specified codec profile.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "rate_control_mode": {
                                "computed": true,
                                "description": "Specify the mode. The default is 'vbr'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "vbv_fullness_bits": {
                                "computed": true,
                                "description": "Initial fullness of the Video Buffering Verifier (VBV) buffer in bits.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "vbv_size_bits": {
                                "computed": true,
                                "description": "Size of the Video Buffering Verifier (VBV) buffer in bits.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "width_pixels": {
                                "computed": true,
                                "description": "The width of the video in pixels.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "hlg": {
                                "block": {
                                  "description": "HLG color format setting for H264.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "sdr": {
                                "block": {
                                  "description": "SDR color format setting for H264.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "H264 codec settings",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Encoding of a video stream.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of input assets stored in Cloud Storage.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "encryptions": {
              "block": {
                "attributes": {
                  "id": {
                    "description": "Identifier for this set of encryption options.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "aes128": {
                    "block": {
                      "description": "Configuration for AES-128 encryption.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "drm_systems": {
                    "block": {
                      "block_types": {
                        "clearkey": {
                          "block": {
                            "description": "Clearkey configuration.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "fairplay": {
                          "block": {
                            "description": "Fairplay configuration.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "playready": {
                          "block": {
                            "description": "Playready configuration.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "widevine": {
                          "block": {
                            "description": "Widevine configuration.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "DRM system(s) to use; at least one must be specified. If a DRM system is omitted, it is considered disabled.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "mpeg_cenc": {
                    "block": {
                      "attributes": {
                        "scheme": {
                          "description": "Specify the encryption scheme.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Configuration for MPEG Common Encryption (MPEG-CENC).",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "sample_aes": {
                    "block": {
                      "description": "Configuration for SAMPLE-AES encryption.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "secret_manager_key_source": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The name of the Secret Version containing the encryption key in the following format: projects/{project}/secrets/{secret_id}/versions/{version_number}.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Configuration for secrets stored in Google Secret Manager.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of encryption configurations for the content.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "inputs": {
              "block": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "A unique key for this input. Must be specified when using advanced mapping and edit lists.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uri": {
                    "computed": true,
                    "description": "URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, gs://bucket/inputs/file.mp4).\nIf empty, the value is populated from Job.input_uri.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "List of input assets stored in Cloud Storage.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "manifests": {
              "block": {
                "attributes": {
                  "file_name": {
                    "computed": true,
                    "description": "The name of the generated file. The default is 'manifest'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mux_streams": {
                    "computed": true,
                    "description": "List of user supplied MuxStream.key values that should appear in this manifest.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of the manifest. Possible values: [\"MANIFEST_TYPE_UNSPECIFIED\", \"HLS\", \"DASH\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Manifest configuration.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "mux_streams": {
              "block": {
                "attributes": {
                  "container": {
                    "computed": true,
                    "description": "The container format. The default is 'mp4'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "elementary_streams": {
                    "computed": true,
                    "description": "List of ElementaryStream.key values multiplexed in this stream.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "encryption_id": {
                    "computed": true,
                    "description": "Identifier of the encryption configuration to use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_name": {
                    "computed": true,
                    "description": "The name of the generated file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "computed": true,
                    "description": "A unique key for this multiplexed stream.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "segment_settings": {
                    "block": {
                      "attributes": {
                        "segment_duration": {
                          "computed": true,
                          "description": "Duration of the segments in seconds. The default is '6.0s'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Segment settings for ts, fmp4 and vtt.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Multiplexing settings for output stream.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "output": {
              "block": {
                "attributes": {
                  "uri": {
                    "computed": true,
                    "description": "URI for the output file(s). For example, gs://my-bucket/outputs/.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Location of output file(s) in a Cloud Storage bucket.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "overlays": {
              "block": {
                "block_types": {
                  "animations": {
                    "block": {
                      "block_types": {
                        "animation_fade": {
                          "block": {
                            "attributes": {
                              "end_time_offset": {
                                "computed": true,
                                "description": "The time to end the fade animation, in seconds.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fade_type": {
                                "description": "Required. Type of fade animation: 'FADE_IN' or 'FADE_OUT'.\nThe possible values are:\n\n* 'FADE_TYPE_UNSPECIFIED': The fade type is not specified.\n\n* 'FADE_IN': Fade the overlay object into view.\n\n* 'FADE_OUT': Fade the overlay object out of view. Possible values: [\"FADE_TYPE_UNSPECIFIED\", \"FADE_IN\", \"FADE_OUT\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "start_time_offset": {
                                "computed": true,
                                "description": "The time to start the fade animation, in seconds.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "xy": {
                                "block": {
                                  "attributes": {
                                    "x": {
                                      "computed": true,
                                      "description": "Normalized x coordinate.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "y": {
                                      "computed": true,
                                      "description": "Normalized y coordinate.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "Normalized coordinates based on output video resolution.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Display overlay object with fade animation.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List of animations. The list should be chronological, without any time overlap.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "image": {
                    "block": {
                      "attributes": {
                        "uri": {
                          "description": "URI of the image in Cloud Storage. For example, gs://bucket/inputs/image.png.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Image overlay.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of overlays on the output video, in descending Z-order.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "pubsub_destination": {
              "block": {
                "attributes": {
                  "topic": {
                    "description": "The name of the Pub/Sub topic to publish job completion notification to. For example: projects/{project}/topics/{topic}.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Pub/Sub destination.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The configuration for this template.",
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

func GoogleTranscoderJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleTranscoderJob), &result)
	return &result
}
