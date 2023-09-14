// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["vr"] = model.Subcommand{
		Name:        []string{"vr"},
		Description: `The npm-style script runner for Deno`,
		Args: []model.Arg{{
			Name:      "script",
			Generator: nil, // TODO: port over generator
		}},
		Options: []model.Option{{
			Name:         []string{"--help", "-h"},
			Description:  `Show help for Velociraptor`,
			IsPersistent: true,
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Show the version number for Velociraptor`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"run"},
			Description: `Run a script`,
			Args: []model.Arg{{
				Name:      "script",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"export"},
			Description: `Export one or more scripts as standalone executable files`,
			Args: []model.Arg{{
				Name:       "script",
				Generator:  nil, // TODO: port over generator
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"upgrade"},
			Description: `Upgrade Velociraptor to the latest version or to a specific one`,
			Args: []model.Arg{{
				Name:       "version",
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-o", "--out-dir"},
				Description: `The folder where the scripts will be exported`,
				Args: []model.Arg{{
					Name: "dir",
				}},
			}},
		}},
	}
}
