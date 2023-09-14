// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["remix"] = model.Subcommand{
		Name:        []string{"remix"},
		Description: `Remix CLI to start, build and export your application`,
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Output usage information`,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Output the version number`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"build"},
			Description: `Create an optimized production build of your application`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir",
				Description: `Represent the directory of the Remix application`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--sourcemap"},
				Description: `Enables production sourcemap`,
			}},
		}, {
			Name:        []string{"dev"},
			Description: `Start the application in development mode`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir",
				Description: `Represent the directory of the Remix application`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"setup"},
			Description: `Prepare node_modules/remix folder (after installation of packages)`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir",
				Description: `Represent the directory of the Remix application`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"routes"},
			Description: `Generate the route config of the application`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "dir",
				Description: `Represent the directory of the Remix application`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--json"},
				Description: `Print the route config as JSON`,
			}},
		}},
	}
}
