// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["coda"] = model.Subcommand{
		Name:        []string{"coda"},
		Description: `Coda Local development CLI tool. It comes bundled with the Pack SDK and makes it easy to build and manage Packs from the CLI`,
		Options: []model.Option{{
			Name:         []string{"--version"},
			Description:  `Show version number`,
			IsPersistent: true,
		}, {
			Name:         []string{"--help"},
			Description:  `Show help`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"init"},
			Description: `Initialize an empty project with the recommended settings and dependencies`,
		}, {
			Name:        []string{"execute"},
			Description: `Execute the formula and print the output to the terminal`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "formula",
				Description: `Formula name to execute`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "params",
				Description: `Arguments to pass to the formula`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--dynamicUrl"},
				Description: `To run a sync for a dynamic sync table, use this parameter to specify which URL to sync from`,
				Args: []model.Arg{{
					Name:        "url",
					Description: `The URL to sync from`,
				}},
			}},
		}, {
			Name:        []string{"auth"},
			Description: `Set up authentication in your development environment so that you can execute Pack formulas with authentication applied to them`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"register"},
			Description: `Create a new API token or register an existing one to be used with future commands`,
			Args: []model.Arg{{
				Name:        "apiToken",
				Description: `The API token to register`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"create"},
			Description: `Create a new Pack on Coda’s servers and get assigned a Pack ID. Run this command just once for each Pack you create`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--name"},
				Description: `Specify a name for the Pack. You can always set or update the name in the Pack management UI later`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The desired Pack name`,
				}},
			}, {
				Name:        []string{"--description"},
				Description: `Specify a description for the Pack. You can always set or update the description in the Pack management UI later`,
				Args: []model.Arg{{
					Name:        "description",
					Description: `The Pack description`,
				}},
			}},
		}, {
			Name:        []string{"upload"},
			Description: `Use this command to upload a new version of your Pack based on your latest code`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"release"},
			Description: `Release a Pack version and make it live for your users`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "version",
				Description: `The release version. Must always be greater than that of any of your previous releases`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"clone"},
			Description: `Migrate the development of a Pack from the Pack Studio to a new local CLI Pack development`,
			Args: []model.Arg{{
				Name:        "urlOrPackId",
				Description: `The URL or ID of the Pack. E.g. https://coda.io/p/123456 or 123456`,
			}},
		}, {
			Name:        []string{"link"},
			Description: `Link the development of a Pack from the Pack studio to an existing local CLI Pack development`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "urlOrPackId",
				Description: `The URL or ID of the Pack. E.g. https://coda.io/p/123456 or 123456`,
			}},
		}, {
			Name:        []string{"whoami"},
			Description: `Looks up information about the API token that is registered in this environment`,
			Args: []model.Arg{{
				Name:        "apiToken",
				Description: `The API token to look up`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"build"},
			Description: `Generate a bundle for your Pack`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"validate"},
			Description: `Validate your Pack definition`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"setOption"},
			Description: `Set a persistent build option for the Pack. This will be used for all builds of the Pack`,
			Args: []model.Arg{{
				Name:        "path/to/pack.ts",
				Description: `The path to the pack.ts file. E.g. src/pack.ts`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "option",
				Description: `Currently the only supported option is 'timerStrategy'`,
				Suggestions: []model.Suggestion{{Name: []string{`timerStrategy`}}},
			}, {
				Name:        "value",
				Description: `Value to set for the option`,
				Suggestions: []model.Suggestion{{Name: []string{`none`}}, {Name: []string{`error`}}, {Name: []string{`fake`}}},
			}},
		}},
	}
}
