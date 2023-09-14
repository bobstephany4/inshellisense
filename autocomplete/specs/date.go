// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["date"] = model.Subcommand{
		Name:        []string{"date"},
		Description: `Display or set date and time`,
		Args: []model.Arg{{
			Name:        "new_time OR output_fmt",
			Description: `New_time: [[[mm]dd]HH]MM[[cc]yy][.ss], output_fmt: '+' followed by user-defined format string`,
			IsOptional:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-d"},
			Description: `Set the kernel's value for daylight saving time`,
			Args: []model.Arg{{
				Name: "dst",
			}},
		}, {
			Name:        []string{"-f"},
			Description: `Use specified format for input instead of the default [[[mm]dd]HH]MM[[cc]yy][.ss] format`,
			Args: []model.Arg{{
				Name:        "input_fmt",
				Description: `The format with which to parse the new date value`,
			}, {
				Name:        "new_date",
				Description: `The new date to set`,
			}},
		}, {
			Name:        []string{"-j"},
			Description: `Don't try to set the date`,
		}, {
			Name:        []string{"-n"},
			Description: `Only set time on the current machine, instead of all machines in the local group`,
		}, {
			Name:        []string{"-R"},
			Description: `Use RFC 2822 date and time output format`,
		}, {
			Name:        []string{"-r"},
			Description: `Print the date and time represented by the specified number of seconds since the Epoch`,
			Args: []model.Arg{{
				Name:        "seconds",
				Description: `Number of seconds since the Epoch (00:00:00 UTC, January 1, 1970)`,
			}},
		}, {
			Name:        []string{"-t"},
			Description: `Set the system's value for minutes west of GMT`,
			Args: []model.Arg{{
				Name: "minutes_west",
			}},
		}, {
			Name:        []string{"-u"},
			Description: `Display or set the date in UTC (Coordinated Universal) time`,
		}, {
			Name:        []string{"-v"},
			Description: `Adjust and print (but don't set) the second, minute, hour, month day, week day, month, or year according to val`,
			Args: []model.Arg{{
				Name:        "val",
				Description: `[+|-]val[ymwdHMS]`,
			}},
		}},
	}
}
