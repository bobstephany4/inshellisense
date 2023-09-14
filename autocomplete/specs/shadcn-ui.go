// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["shadcn-ui"] = model.Subcommand{
		Name:        []string{"shadcn-ui"},
		Description: `Shadcn UI CLI`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"add"},
			Description: `Add a component to your project`,
			Args: []model.Arg{{
				Name:        "components",
				Description: `The components to add`,
				Generator:   nil, // TODO: port over generator
				IsVariadic:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-y", "--yes"},
				Description: `Skip confirmation prompt`,
			}, {
				Name:        []string{"-o", "--overwrite"},
				Description: `Overwrite existing files`,
			}, {
				Name:        []string{"-c", "--cwd"},
				Description: `The working directory. defaults to the current directory`,
				Args: []model.Arg{{
					Name: "cwd",
				}},
			}, {
				Name:        []string{"-p", "--path"},
				Description: `The path to add the component to`,
				Args: []model.Arg{{
					Name: "path",
				}},
			}},
		}, {
			Name:        []string{"diff"},
			Description: `Check for updates against the registry`,
			Args: []model.Arg{{
				Name:        "component",
				Description: `The component name`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"-y", "--yes"},
				Description: `Skip confirmation prompt`,
			}, {
				Name:        []string{"-c", "--cwd"},
				Description: `The working directory. defaults to the current directory`,
				Args: []model.Arg{{
					Name: "cwd",
				}},
			}},
		}, {
			Name:        []string{"init"},
			Description: `Initialize your project and install dependencies`,
			Options: []model.Option{{
				Name:        []string{"-y", "--yes"},
				Description: `Skip confirmation prompt`,
			}, {
				Name:        []string{"-c", "--cwd"},
				Description: `The working directory. defaults to the current directory`,
				Args: []model.Arg{{
					Name: "cwd",
				}},
			}},
		}},
	}
}
