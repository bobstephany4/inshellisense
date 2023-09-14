// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["asr"] = model.Subcommand{
		Name:        []string{"asr"},
		Description: `Asr efficiently copies disk images onto volumes, either directly or via a multicast network stream`,
		Options: []model.Option{{
			Name:        []string{"--buffers"},
			Description: `One of the options that control how asr uses memory`,
			Args: []model.Arg{{
				Name:        "num",
				Description: `Specifies the num buffers should be used`,
			}},
		}, {
			Name:        []string{"--buffersize"},
			Description: `One of the options that control how asr uses memory`,
			Args: []model.Arg{{
				Name:        "size",
				Description: `Specifies the size of each buffer`,
			}},
		}, {
			Name:        []string{"--csumbuffers"},
			Description: `One of the options that control how asr uses memory`,
			Args: []model.Arg{{
				Name:        "num",
				Description: `Specifies that num buffers should be used for checksumming operations`,
			}},
		}, {
			Name:        []string{"--csumbuffersize"},
			Description: `One of the options that control how asr uses memory`,
			Args: []model.Arg{{
				Name:        "size",
				Description: `Specifies the size of each buffer used for checksumming`,
			}},
		}, {
			Name:        []string{"--verbose"},
			Description: `Enables verbose progress and error messages`,
		}, {
			Name:        []string{"--debug"},
			Description: `Enables other progress and error messages`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"help"},
			Description: `Displays asr usage information`,
		}, {
			Name:        []string{"version"},
			Description: `Displays asr version`,
		}, {
			Name:        []string{"restore"},
			Description: `Restores a disk image or volume to another volume`,
			Options: []model.Option{{
				Name:        []string{"--source"},
				Description: `Can be a disk image, /dev entry, or volume mountpoint`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "source",
					Description: `Disk image, /dev entry, or volume mountpoint`,
				}},
			}, {
				Name:        []string{"--target"},
				Description: `Can be a /dev entry, or volume mountpoint`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "target",
					Description: `/dev entry, or volume mountpoint`,
				}},
			}, {
				Name:        []string{"--file"},
				Description: `When performing a multicast restore, --file can be specified instead of --target`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "source",
					Description: `Disk image, /dev entry, or volume mountpoint`,
				}},
			}, {
				Name:        []string{"--erase"},
				Description: `Erases target and is required. --erase must always be used, as file copies are no longer supported by asr`,
			}, {
				Name:        []string{"--format"},
				Description: `Specifies the destination filesystem format`,
				Args: []model.Arg{{
					Name:        "HFS+ | HFSX",
					Description: `Specify the destination filesystem format`,
					IsOptional:  true,
				}},
			}, {
				Name:        []string{"--noprompt"},
				Description: `Suppresses the prompt which usually occurs before target`,
			}, {
				Name:        []string{"--timeout"},
				Description: `Specifies num seconds that a multicast client should wait when no payload data has been received over a multicast stream before exiting`,
				Args: []model.Arg{{
					Name:        "num",
					Description: `Number of seconds that a multicast client should wait`,
					Suggestions: []model.Suggestion{{Name: []string{`0`}}},
					IsOptional:  true,
				}},
			}, {
				Name:        []string{"--puppetstrings"},
				Description: `Provide progress output that is easy for another program to parse`,
			}, {
				Name:        []string{"--noverify"},
				Description: `Skips the verification steps normally taken to ensure that a volume has been properly restored`,
			}, {
				Name:        []string{"--allowfragmentedcatalog"},
				Description: `Allows restores to proceed even if the source's catalog file is fragmented`,
			}, {
				Name:        []string{"--SHA256"},
				Description: `Forces the restore to use the SHA-256 hash in the image during verification`,
			}, {
				Name:        []string{"--sourcevolumename"},
				Description: `Tells asr which volume in the source container to invert when doing an APFS restore`,
			}, {
				Name:        []string{"--sourcevolumeUUID"},
				Description: `Tells asr which volume in the source container to invert when doing an APFS restore`,
			}, {
				Name:        []string{"--useReplication"},
				Description: `Forces asr to use replication for restoring APFS volumes`,
			}, {
				Name:        []string{"--useInverter"},
				Description: `Forces asr to use the inverter for restoring APFS volumes`,
			}, {
				Name:        []string{"--toSnapshot"},
				Description: `Specifies the snapshot on the source APFS volume to restore to the target APFS volume`,
			}, {
				Name:        []string{"--fromSnapshot"},
				Description: `Names a snapshot on the source APFS volume to use in combination with --toSnapshot to specify a snapshot delta to restore to the target APFS volume`,
			}},
		}, {
			Name:        []string{"restoreexact"},
			Description: `Same as restore verb, except that for an HFS Plus volume, the target partition is resized to match the size of the source partition/ volume`,
			Options: []model.Option{{
				Name:        []string{"--source"},
				Description: `Can be a disk image, /dev entry, or volume mountpoint`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "source",
					Description: `Disk image, /dev entry, or volume mountpoint`,
				}},
			}, {
				Name:        []string{"--target"},
				Description: `Can be a /dev entry, or volume mountpoint`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "target",
					Description: `/dev entry, or volume mountpoint`,
				}},
			}, {
				Name:        []string{"--file"},
				Description: `When performing a multicast restore, --file can be specified instead of --target`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "source",
					Description: `Disk image, /dev entry, or volume mountpoint`,
				}},
			}, {
				Name:        []string{"--erase"},
				Description: `Erases target and is required. --erase must always be used, as file copies are no longer supported by asr`,
			}, {
				Name:        []string{"--format"},
				Description: `Specifies the destination filesystem format`,
				Args: []model.Arg{{
					Name:        "HFS+ | HFSX",
					Description: `Specify the destination filesystem format`,
					IsOptional:  true,
				}},
			}, {
				Name:        []string{"--noprompt"},
				Description: `Suppresses the prompt which usually occurs before target`,
			}, {
				Name:        []string{"--timeout"},
				Description: `Specifies num seconds that a multicast client should wait when no payload data has been received over a multicast stream before exiting`,
				Args: []model.Arg{{
					Name:        "num",
					Description: `Number of seconds that a multicast client should wait`,
					Suggestions: []model.Suggestion{{Name: []string{`0`}}},
					IsOptional:  true,
				}},
			}, {
				Name:        []string{"--puppetstrings"},
				Description: `Provide progress output that is easy for another program to parse`,
			}, {
				Name:        []string{"--noverify"},
				Description: `Skips the verification steps normally taken to ensure that a volume has been properly restored`,
			}, {
				Name:        []string{"--allowfragmentedcatalog"},
				Description: `Allows restores to proceed even if the source's catalog file is fragmented`,
			}, {
				Name:        []string{"--SHA256"},
				Description: `Forces the restore to use the SHA-256 hash in the image during verification`,
			}, {
				Name:        []string{"--sourcevolumename"},
				Description: `Tells asr which volume in the source container to invert when doing an APFS restore`,
			}, {
				Name:        []string{"--sourcevolumeUUID"},
				Description: `Tells asr which volume in the source container to invert when doing an APFS restore`,
			}, {
				Name:        []string{"--useReplication"},
				Description: `Forces asr to use replication for restoring APFS volumes`,
			}, {
				Name:        []string{"--useInverter"},
				Description: `Forces asr to use the inverter for restoring APFS volumes`,
			}, {
				Name:        []string{"--toSnapshot"},
				Description: `Specifies the snapshot on the source APFS volume to restore to the target APFS volume`,
			}, {
				Name:        []string{"--fromSnapshot"},
				Description: `Names a snapshot on the source APFS volume to use in combination with --toSnapshot to specify a snapshot delta to restore to the target APFS volume`,
			}},
		}, {
			Name:        []string{"server"},
			Description: `Multicasts source over the network`,
			Options: []model.Option{{
				Name:        []string{"--source"},
				Description: `Source has to be a UDIF disk image`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "source",
					Description: `UDIF disk image local/remote path`,
				}},
			}, {
				Name:        []string{"--interface"},
				Description: `The network interface to be used for multicasting`,
				Args: []model.Arg{{
					Name:        "interface",
					Description: `The network interface to be used for multicasting`,
				}},
			}, {
				Name:        []string{"--config"},
				Description: `Server requires a configuration file to be passed`,
				Args: []model.Arg{{
					Name:        "configuration",
					Description: `Configuration file in standard property list format`,
					Generator:   nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"imagescan"},
			Description: `Calculate checksums of the data in the provided image and store them in the image`,
		}, {
			Name:        []string{"info"},
			Description: `Report the image metadata which was placed in the image by a previous use of the imagescan verb`,
			Options: []model.Option{{
				Name:        []string{"--plist"},
				Description: `Writes its output as an XML-formatted plist`,
			}},
		}},
	}
}
