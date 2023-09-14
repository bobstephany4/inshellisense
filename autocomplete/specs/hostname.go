// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["hostname"] = model.Subcommand{
		Name:        []string{"hostname"},
		Description: `Set or print name of current host system`,
		Args: []model.Arg{{
			Name:        "hostname",
			Description: `The hostname to use for this machine`,
		}},
		Options: []model.Option{{
			Name:        []string{"-f"},
			Description: `Include domain information in the printed name`,
		}, {
			Name:        []string{"-s"},
			Description: `Trim off any domain information from the printed name`,
		}, {
			Name:        []string{"-d"},
			Description: `Only print domain information`,
		}},
	}
}
