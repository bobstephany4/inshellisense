// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["ipatool"] = model.Subcommand{
		Name:        []string{"ipatool"},
		Description: `IPATool is a command line tool that allows you to search for iOS apps on the App Store and download a copy of the app package, known as an ipa file`,
		Options: []model.Option{{
			Name:         []string{"-h", "--help"},
			Description:  `Show help for ipatool`,
			IsPersistent: true,
		}, {
			Name:         []string{"--non-interactive"},
			Description:  `Run in non-interactive session`,
			IsPersistent: true,
		}, {
			Name:         []string{"--verbose"},
			Description:  `Enables verbose logs`,
			IsPersistent: true,
		}, {
			Name:        []string{"--format"},
			Description: `Output format`,
			Args: []model.Arg{{
				Name:        "format",
				Suggestions: []model.Suggestion{{Name: []string{`text`}}, {Name: []string{`json`}}},
			}},
			IsPersistent: true,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Show version for ipatool`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"auth"},
			Description: `Authenticate with the App Store`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"info"},
				Description: `Show current account info`,
			}, {
				Name:        []string{"login"},
				Description: `Login to the App Store`,
				Options: []model.Option{{
					Name:        []string{"--auth-code"},
					Description: `2FA code for the Apple ID`,
					Args: []model.Arg{{
						Name: "2FA code",
					}},
				}, {
					Name:        []string{"-e", "--email"},
					Description: `Apple ID email address`,
				}, {
					Name:        []string{"-p", "--password"},
					Description: `Apple ID password`,
				}},
			}, {
				Name:        []string{"revoke"},
				Description: `Revoke your App Store credentials`,
			}},
		}, {
			Name:        []string{"completion"},
			Description: `Generate shell completion script`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"bash"},
				Description: `Generate the autocompletion script for bash`,
			}, {
				Name:        []string{"fish"},
				Description: `Generate the autocompletion script for fish`,
			}, {
				Name:        []string{"powershell"},
				Description: `Generate the autocompletion script for powershell`,
			}, {
				Name:        []string{"zsh"},
				Description: `Generate the autocompletion script for zsh`,
			}},
		}, {
			Name:        []string{"download"},
			Description: `Download (encrypted) iOS app packages from the App Store`,
			Options: []model.Option{{
				Name:        []string{"-b", "--bundle-identifier"},
				Description: `Bundle identifier of the app`,
				Args: []model.Arg{{
					Name:           "identifier",
					FilterStrategy: model.FilterStrategyFuzzy,
					Generator:      nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"-o", "--output"},
				Description: `The destination path of the downloaded app package`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "output path",
				}},
			}, {
				Name:        []string{"--purchase"},
				Description: `Obtain a license for the app if needed`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Display help for command`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateHelp},
				Name:       "command",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"purchase"},
			Description: `Obtain a license for the app from the App Store`,
			Options: []model.Option{{
				Name:        []string{"-b", "--bundle-identifier"},
				Description: `Bundle identifier of the app`,
				Args: []model.Arg{{
					Name:           "identifier",
					FilterStrategy: model.FilterStrategyFuzzy,
					Generator:      nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"search"},
			Description: `Search for iOS apps available on the App Store`,
			Args: []model.Arg{{
				Name: "query",
			}},
			Options: []model.Option{{
				Name:        []string{"-l", "--limit"},
				Description: `Limit the number of results`,
				Args: []model.Arg{{
					Name: "limit",
				}},
			}},
		}},
	}
}
