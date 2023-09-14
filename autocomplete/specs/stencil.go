// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["stencil"] = model.Subcommand{
		Name:        []string{"stencil"},
		Description: `CLI to build Stencil projects and generate components`,
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Display the help output explaining the different flags`,
		}, {
			Name:        []string{"--version"},
			Description: `Prints the current Stencil version`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"build"},
			Description: `Build components for development or production`,
			Options: []model.Option{{
				Name:        []string{"--ci"},
				Description: `Run a build using recommended settings for a Continuous Integration (CI) environment`,
			}, {
				Name:        []string{"--config", "-c"},
				Description: `Set stencil config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config file",
				}},
			}, {
				Name:        []string{"--debug"},
				Description: `Set the log level to debug`,
			}, {
				Name:        []string{"--dev"},
				Description: `Development build`,
			}, {
				Name:        []string{"--docs-readme"},
				Description: `Generate component readme.md docs`,
			}, {
				Name:        []string{"--es5"},
				Description: `Creates an ES5 compatible build`,
			}, {
				Name:        []string{"--log"},
				Description: `Write stencil-build.log file`,
			}, {
				Name:        []string{"--prerender"},
				Description: `Prerender the application`,
			}, {
				Name:        []string{"--prod"},
				Description: `Runs a production build`,
			}, {
				Name:        []string{"--max-workers"},
				Description: `Max number of workers the compiler should use`,
				Args: []model.Arg{{
					Name:        "workers",
					Description: `Number of workers`,
				}},
			}, {
				Name:        []string{"--next"},
				Description: `Opt-in to test the 'next' Stencil compiler features`,
			}, {
				Name:        []string{"--no-cache"},
				Description: `Disables using the cache`,
			}, {
				Name:        []string{"--no-open"},
				Description: `Do not automatically open the browser window`,
			}, {
				Name:        []string{"--port", "--p"},
				Description: `Port for the Integrated Dev Server`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--serve"},
				Description: `Start the dev-server`,
			}, {
				Name:        []string{"--stats"},
				Description: `Write stencil-stats.json file`,
			}, {
				Name:        []string{"--verbose"},
				Description: `Logs additional information about each step of the build`,
			}, {
				Name:        []string{"--watch"},
				Description: `Rebuild when files update`,
			}},
		}, {
			Name:        []string{"test"},
			Description: `Run unit and end-to-end tests`,
			Options: []model.Option{{
				Name:        []string{"--spec"},
				Description: `Run unit tests with Jest`,
			}, {
				Name:        []string{"--e2e"},
				Description: `Run e2e tests with Puppeteer`,
			}, {
				Name:        []string{"--no-build"},
				Description: `Skips build process before running the tests`,
			}},
		}, {
			Name:        []string{"generate", "g"},
			Description: `Bootstrap components`,
			Args: []model.Arg{{
				Name:        "name",
				Description: `Name of new component (must contain dash (-))`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"telemetry"},
			Description: `Opt in or out of telemetry`,
			Args: []model.Arg{{
				Name: "state",
				Suggestions: []model.Suggestion{{
					Name:        []string{`off`},
					Description: `Disable sharing anonymous usage data`,
				}, {
					Name:        []string{`on`},
					Description: `Enable sharing anonymous usage data`,
				}},
			}},
		}},
	}
}
