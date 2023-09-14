// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["tmutil"] = model.Subcommand{
		Name:        []string{"tmutil"},
		Description: `Time Machine utility`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"version"},
			Description: `Print the current version`,
		}, {
			Name:        []string{"enable"},
			Description: `Turn on automatic backups (required root + full disk access)`,
		}, {
			Name:        []string{"disable"},
			Description: `Turn off automatic backups (required root + full disk access)`,
		}, {
			Name:        []string{"startbackup"},
			Description: `Begin a backup if one is not already running`,
			Options: []model.Option{{
				Name:        []string{"-a", "--auto"},
				Description: `Run the backup in a mode similar to system-scheduled backups`,
			}, {
				Name:        []string{"-b", "--block"},
				Description: `Wait (block) until the backup is finished before exiting`,
			}, {
				Name:        []string{"-r", "--rotation"},
				Description: `Allow automatic destination rotation during backup`,
			}, {
				Name:        []string{"-d", "--destination"},
				Description: `Perform the backup to the destination corresponding to the specified ID`,
				Args: []model.Arg{{
					Name: "destination id",
				}},
			}},
		}, {
			Name:        []string{"stopbackup"},
			Description: `Cancel a backup currently in progress`,
		}, {
			Name:        []string{"delete"},
			Description: `Deletes backups with the specified timestamp from the backup volume mounted at the given mountpoint`,
			Options: []model.Option{{
				Name:        []string{"-d"},
				Description: `Delete from the given mount point`,
				Args: []model.Arg{{
					Name: "backup mount point",
				}},
			}, {
				Name:        []string{"-t"},
				Description: `Delete backups with this timestamp`,
				Args: []model.Arg{{
					Name: "timestamp",
				}},
			}, {
				Name:        []string{"-p"},
				Description: `A specific path to delete (HFS backup disk)`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths, model.TemplateFolders},
					Name:      "path",
				}},
			}},
		}, {
			Name:        []string{"deleteinprogress"},
			Description: `Delete all in-progress backups for a machine directory`,
			Args: []model.Arg{{
				Name: "machine directory",
			}},
		}, {
			Name:        []string{"restore"},
			Description: `Restore the item "source", which is inside a backup, to the location "destination". Same semantics as "cp"`,
			Args: []model.Arg{{
				Name:       "source",
				IsVariadic: true,
			}, {
				Name: "destination",
			}},
			Options: []model.Option{{
				Name:        []string{"-v"},
				Description: `Be verbose, showing files as they are copied`,
			}},
		}, {
			Name:        []string{"compare"},
			Description: `Perform a backup diff`,
			Args: []model.Arg{{
				Name: "backup or path",
			}, {
				Name:       "path (if not backup)",
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-@"},
				Description: `Compare extended attributes`,
			}, {
				Name:        []string{"-a"},
				Description: `Compare all supported metadata`,
			}, {
				Name:        []string{"-c"},
				Description: `Compare creation times`,
			}, {
				Name:        []string{"-d"},
				Description: `Compare file data forks`,
			}, {
				Name:        []string{"-e"},
				Description: `Compare ACLs`,
			}, {
				Name:        []string{"-f"},
				Description: `Compare file flags`,
			}, {
				Name:        []string{"-g"},
				Description: `Compare GIDs`,
			}, {
				Name:        []string{"-h"},
				Description: `Unknown`,
			}, {
				Name:        []string{"-l"},
				Description: `Unknown`,
			}, {
				Name:        []string{"-m"},
				Description: `Compare file modes`,
			}, {
				Name:        []string{"-n"},
				Description: `No metadata comparison`,
			}, {
				Name:        []string{"-s"},
				Description: `Compare sizes`,
			}, {
				Name:        []string{"-t"},
				Description: `Compare modification times`,
			}, {
				Name:        []string{"-u"},
				Description: `Compare UIDs`,
			}, {
				Name:        []string{"-E"},
				Description: `Don't take exclusions into account when comparing items inside volumes`,
			}, {
				Name:        []string{"-U"},
				Description: `Ignore logical volume identity (volume UUIDs) when directly comparing a local volume or volume store to a volume store`,
			}, {
				Name:        []string{"-X"},
				Description: `Print output in XML property list format`,
			}, {
				Name:        []string{"-D"},
				Description: `Limit traversal depth to depth levels from the beginning of iteration`,
				Args: []model.Arg{{
					Name: "depth",
				}},
			}, {
				Name:        []string{"-I"},
				Description: `Ignore paths with a path component equal to name during iteration`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}},
		}, {
			Name:        []string{"setdestination"},
			Description: `Configure a local HFS+ or APFS volume, AFP share, or SMB share as a backup destination (requires root + Full Disk Access)`,
			Args: []model.Arg{{
				Name: "destination",
			}},
			Options: []model.Option{{
				Name:        []string{"-a"},
				Description: `Add the arg to the list of destinations, instead of replacing the list`,
			}, {
				Name:        []string{"-p"},
				Description: `Enter a password in a non-echoing interactive prompt`,
			}},
		}, {
			Name:        []string{"removedestination"},
			Description: `Remove the destination with the specified unique identifier from the Time Machine configuration (requires root + Full Disk Access)`,
			Args: []model.Arg{{
				Name: "destination id",
			}},
		}, {
			Name:        []string{"destinationinfo"},
			Description: `Print information about destinations currently configured for use with Time Machine`,
			Options: []model.Option{{
				Name:        []string{"-X"},
				Description: `Print output in XML plist format`,
			}},
		}, {
			Name:        []string{"addexclusion"},
			Description: `Configure an exclusion that tells Time Machine not to backup a file, directory, or volume during future backups`,
			Args: []model.Arg{{
				Name:       "item",
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-p"},
				Description: `Configure fixed-path exclusions`,
			}, {
				Name:        []string{"-v"},
				Description: `Configure volume exclusions`,
			}},
		}, {
			Name:        []string{"removeexclusion"},
			Description: `Configure  Time Machine to backup a file, directory, or volume during future backups`,
			Args: []model.Arg{{
				Name:       "item",
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-p"},
				Description: `Configure fixed-path exclusions`,
			}, {
				Name:        []string{"-v"},
				Description: `Configure volume exclusions`,
			}},
		}, {
			Name:        []string{"isexcluded"},
			Description: `Determine if a file, directory, or volume are excluded from Time Machine backups`,
			Args: []model.Arg{{
				Name:       "item",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"inheritbackup"},
			Description: `Claim a machine directory or sparsebundle for use by the current machine`,
			Args: []model.Arg{{
				Name: "machine directory or sparsebundle",
			}},
		}, {
			Name:        []string{"associatedisk"},
			Description: `Bind a volume store directory to the specified local disk, thereby reconfiguring the backup history (requires root + Full Disk Access)`,
			Args: []model.Arg{{
				Name: "mount point",
			}, {
				Name: "volume backup directory",
			}},
			Options: []model.Option{{
				Name:        []string{"-a"},
				Description: `Find all volume stores in the same machine directory that match the identity of the volume backup directory`,
			}},
		}, {
			Name:        []string{"latestbackup"},
			Description: `List this computer's latest completed backup`,
			Options: []model.Option{{
				Name:        []string{"-d"},
				Description: `Specify a destination volume to list backups from`,
			}, {
				Name:        []string{"-m"},
				Description: `Attempt to mount the backups and list their mounted paths`,
			}, {
				Name:        []string{"-t"},
				Description: `Only show the backup timestampt rather than the full name or path`,
			}},
		}, {
			Name:        []string{"listbackups"},
			Description: `List all of this computer's completed backups`,
			Options: []model.Option{{
				Name:        []string{"-d"},
				Description: `Specify a destination volume to list backups from`,
			}, {
				Name:        []string{"-m"},
				Description: `Attempt to mount the backups and list their mounted paths`,
			}, {
				Name:        []string{"-t"},
				Description: `Only show the backup timestampt rather than the full name or path`,
			}},
		}, {
			Name:        []string{"machinedirectory"},
			Description: `Print the path to the current machine directory for this computer`,
		}, {
			Name:        []string{"calculatedrift"},
			Description: `Analyze the backups in an HFS machine directory and determine the amount of change between each`,
			Args: []model.Arg{{
				Name: "machine directory",
			}},
		}, {
			Name:        []string{"uniquesize"},
			Description: `Analyze the specified path in an HFS+ backup or path to an APFS backup and determine its unique size`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "path",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"verifychecksums"},
			Description: `Compute a checksum of data contained within a backup and verify the results against checksum information computed at the time of backup`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "path",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"localsnapshot"},
			Description: `Create new local Time Machine snapshots of all APFS volumes included in the Time Machine backup`,
		}, {
			Name:        []string{"listlocalsnapshots"},
			Description: `List local Time Machine snapshots of the specified volume`,
			Args: []model.Arg{{
				Name: "mount point",
			}},
		}, {
			Name:        []string{"listlocalsnapshotdates"},
			Description: `List the creation dates of all local Time Machine snapshots`,
			Args: []model.Arg{{
				Name:       "mount point",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"deletelocalsnapshots"},
			Description: `Delete all local Time Machine snapshots on all mounted disks for the specified date or on the given disk`,
			Args: []model.Arg{{
				Name:       "mount point or snapshot date",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"thinlocalsnapshots"},
			Description: `Think local Time Machine snapshots for the specified volume`,
			Args: []model.Arg{{
				Name: "mount point",
			}, {
				Name:       "purge amount (bytes)",
				IsOptional: true,
			}, {
				Name:        "urgency",
				Suggestions: []model.Suggestion{{Name: []string{`1`}}, {Name: []string{`2`}}, {Name: []string{`3`}}, {Name: []string{`4`}}},
				IsOptional:  true,
			}},
		}},
	}
}
