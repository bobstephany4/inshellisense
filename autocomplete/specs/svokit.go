// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["svokit"] = model.Subcommand{
		Name:        []string{"svokit"},
		Description: `Runs built svokit project`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for svokit`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"setup"},
			Description: `Creates svokit config (experimental)`,
		}, {
			Name:        []string{"run"},
			Description: `Runs build svokit project`,
		}},
	}
}
