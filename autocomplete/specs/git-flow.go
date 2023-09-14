// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["git-flow"] = model.Subcommand{
		Name:        []string{"git-flow"},
		Description: `A collection of Git extensions to provide high-level repository operations for Vincent Driessen's branching model`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"init"},
			Description: `Initialize a new git repo with support for the branching model`,
			Options: []model.Option{{
				Name:        []string{"-d"},
				Description: `Use default branch naming conventions`,
			}, {
				Name:        []string{"-f"},
				Description: `Force setting of gitflow branches, even if already configured`,
			}},
		}, {
			Name:        []string{"feature"},
			Description: `List all feature branches`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"start"},
				Description: `Create a new feature branch`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The name of the new feature branch`,
				}},
			}, {
				Name:        []string{"finish"},
				Description: `Merge a feature branch into develop`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The name of the feature branch to finish`,
					Generator:   nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"publish"},
				Description: `Push a feature branch to the remote repository`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The name of the feature branch to publish`,
				}},
			}, {
				Name:        []string{"pull"},
				Description: `Pull a feature branch from the remote repository`,
				Args: []model.Arg{{
					Name:        "origin",
					Description: `The name of the remote feature branch`,
				}, {
					Name:        "name",
					Description: `The name of the local feature branch`,
				}},
			}},
		}, {
			Name:        []string{"release"},
			Description: `List all release branches`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"start"},
				Description: `Create a new release branch`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The name of the new release branch`,
				}},
			}, {
				Name:        []string{"finish"},
				Description: `Merge a release branch into develop`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The name of the release branch to finish`,
					Generator:   nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"hotfix"},
			Description: `List all hotfix branches`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"start"},
				Description: `Create a new hotfix branch`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The name of the new hotfix branch`,
				}},
			}, {
				Name:        []string{"finish"},
				Description: `Merge a hotfix branch into develop`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The name of the hotfix branch to finish`,
					Generator:   nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"support"},
			Description: `List all support branches`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"start"},
				Description: `Create a new support branch`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `The name of the new support branch`,
				}},
			}},
		}},
	}
}
