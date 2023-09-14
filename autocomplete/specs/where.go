// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["where"] = model.Subcommand{
		Name:        []string{"where"},
		Description: `For each name, indicate how it should be interpreted`,
		Args: []model.Arg{{
			Name:       "names",
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-w"},
			Description: `For each name, print 'name: word', where 'word' is the kind of command`,
		}, {
			Name:        []string{"-p"},
			Description: `Do a path search for the name, even if it's an alias/function/builtin`,
		}, {
			Name:        []string{"-m"},
			Description: `The arguments are taken as patterns (pattern characters must be quoted)`,
		}, {
			Name:        []string{"-s"},
			Description: `If the pathname contains symlinks, print the symlink-free name as well`,
		}, {
			Name:        []string{"-S"},
			Description: `Print intermediate symlinks and the resolved name`,
		}, {
			Name:        []string{"-x"},
			Description: `Expand tabs when outputting shell function`,
			Args: []model.Arg{{
				Name: "num",
			}},
		}},
	}
}
