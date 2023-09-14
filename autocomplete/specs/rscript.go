// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["rscript"] = model.Subcommand{
		Name:        []string{"Rscript"},
		Description: `Scripting Front-End for R`,
		Args: []model.Arg{{
			Name:        "file",
			Description: `R script to run`,
			Generator:   nil, // TODO: port over generator
		}, {
			Name:        "args",
			Description: `Arguments to pass to the script`,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-e"},
			Description: `R expression to run`,
			Args: []model.Arg{{
				Name:        "expression",
				Description: `R expression`,
			}},
		}, {
			Name:        []string{"--help"},
			Description: `Print usage and exit`,
		}, {
			Name:        []string{"--version"},
			Description: `Print version and exit`,
		}, {
			Name:        []string{"--verbose"},
			Description: `Print information on progress`,
		}, {
			Name:        []string{"--no-echo"},
			Description: `Run as quietly as possible`,
		}, {
			Name:        []string{"--no-restore"},
			Description: `Don't restore anything`,
		}, {
			Name:        []string{"--save"},
			Description: `Do save workspace at the end of the session`,
		}, {
			Name:        []string{"--no-environ"},
			Description: `Don't read the site and user environment files`,
		}, {
			Name:        []string{"--no-site-file"},
			Description: `Don't read the site-wide Rprofile`,
		}, {
			Name:        []string{"--no-init-file"},
			Description: `Don't read the user R profile`,
		}, {
			Name:        []string{"--restore"},
			Description: `Do restore previously saved objects at startup`,
		}, {
			Name:        []string{"--vanilla"},
			Description: `Combine --no-save, --no-restore, --no-site-file --no-init-file and --no-environ`,
		}, {
			Name:        []string{"--default-packages"},
			Description: `Comma separated list of default packages`,
			Args: []model.Arg{{
				Name:        "packages",
				Description: `A comma-separated set of package names, or 'NULL'`,
			}},
		}},
	}
}
