// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["echo"] = model.Subcommand{
		Name:        []string{"echo"},
		Description: `Write arguments to the standard output`,
		Args: []model.Arg{{
			Name:       "string",
			Generator:  nil, // TODO: port over generator
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-n"},
			Description: `Do not print the trailing newline character`,
		}, {
			Name:        []string{"-e"},
			Description: `Interpret escape sequences`,
		}, {
			Name:        []string{"-E"},
			Description: `Disable escape sequences`,
		}},
	}
}
