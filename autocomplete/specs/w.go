// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["w"] = model.Subcommand{
		Name:        []string{"w"},
		Description: `Display who is logged in and what they are doing`,
		Args: []model.Arg{{
			Name:       "user",
			IsOptional: true,
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-h"},
			Description: `Suppress the heading`,
		}, {
			Name:        []string{"-i"},
			Description: `Output is sorted by idle time`,
		}},
	}
}
