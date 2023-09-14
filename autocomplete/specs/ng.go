// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["ng"] = model.Subcommand{
		Name:        []string{"ng"},
		Description: `CLI interface for Angular`,
		Options: []model.Option{{
			Name:        []string{"--version"},
			Description: `View your Angular CLI version`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"new"},
			Description: `Create a new Angular app`,
			Args: []model.Arg{{
				Name: "name",
			}},
			Options: []model.Option{{
				Name:        []string{"--create-application"},
				Description: `Create a default application?`,
				Args: []model.Arg{{
					Name:        "project",
					Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
				}},
			}},
		}, {
			Name:        []string{"generate"},
			Description: `Generate new files`,
			Args: []model.Arg{{
				Name: "schematic",
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"application"},
				Description: `Generates a new application`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the new app`,
				}},
				Options: []model.Option{{
					Name: []string{"--style"},
					Args: []model.Arg{{
						Name:        "extension",
						Suggestions: []model.Suggestion{{Name: []string{`css`}}, {Name: []string{`scss`}}, {Name: []string{`sass`}}, {Name: []string{`less`}}, {Name: []string{`styl`}}},
					}},
				}},
			}, {
				Name:        []string{"component"},
				Description: `Generate a new component`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Component name`,
					IsOptional:  true,
				}},
				Options: []model.Option{{
					Name:        []string{"--project"},
					Description: `Project name`,
					Args: []model.Arg{{
						Generator: nil, // TODO: port over generator
					}},
				}, {
					Name:        []string{"--change-detection", "-c"},
					Description: `The change detection strategy to use`,
					Args: []model.Arg{{
						Name:        "strategy",
						Suggestions: []model.Suggestion{{Name: []string{`Default`}}, {Name: []string{`OnPush`}}},
					}},
				}, {
					Name:        []string{"--display-block", "-b"},
					Description: `Add :host block to styles`,
					Args: []model.Arg{{
						Name:        "boolean",
						Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
					}},
				}, {
					Name:        []string{"--flat"},
					Description: `Create at the top level`,
					Args: []model.Arg{{
						Name:        "boolean",
						Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
					}},
				}},
			}, {
				Name:        []string{"library"},
				Description: `Generates a new library`,
				Args: []model.Arg{{
					Name:       "name",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"class"},
				Description: `Generates a class`,
				Args: []model.Arg{{
					Name:       "name",
					IsOptional: true,
				}},
				Options: []model.Option{{
					Name:        []string{"--project"},
					Description: `Project name`,
					Args: []model.Arg{{
						Generator: nil, // TODO: port over generator
					}},
				}},
			}},
		}, {
			Name:        []string{"version"},
			Description: `View your Angular CLI version (update for Angular 14+)`,
		}},
	}
}
