// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["tmuxinator"] = model.Subcommand{
		Name:        []string{"tmuxinator"},
		Description: `Create and manage tmux sessions easily`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"commands"},
			Description: `Lists commands available in tmuxinator`,
		}, {
			Name:        []string{"completions"},
			Description: `Used for shell completion`,
		}, {
			Name:        []string{"copy"},
			Description: `Copy an existing project to a new project and open it in your editor`,
			Args: []model.Arg{{
				Name:        "source",
				Description: `The source project name`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "target",
				Description: `The target project name`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"debug"},
			Description: `Output the shell commands that are generated by tmuxinator`,
			Options: []model.Option{{
				Name:        []string{"-a", "--attach"},
				Description: `Attach to tmux session after creation`,
				Args: []model.Arg{{
					Name:        "attach-session",
					Description: `Attach to tmux session`,
				}},
			}, {
				Name:        []string{"-n", "--name"},
				Description: `Give the session a different name`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Give the session a different name`,
				}},
			}, {
				Name:        []string{"-p", "--project-config"},
				Description: `Path to project config file`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFolders, model.TemplateFilepaths},
					Name:        "project-config",
					Description: `Path to project config file`,
				}},
			}},
		}, {
			Name:        []string{"delete"},
			Description: `Deletes given project`,
			Args: []model.Arg{{
				Name:        "project",
				Description: `The project name`,
				Generator:   nil, // TODO: port over generator
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"doctor"},
			Description: `Look for problems in your configuration`,
		}, {
			Name:        []string{"help"},
			Description: `Describe available commands or one specific command`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateHelp},
				Name:        "command",
				Description: `The tmuxinator commands`,
			}},
		}, {
			Name:        []string{"implode"},
			Description: `Deletes all tmuxinator projects`,
		}, {
			Name:        []string{"list"},
			Description: `Lists all tmuxinator projects`,
			Options: []model.Option{{
				Name:        []string{"-n", "--newline"},
				Description: `Force output to be one entry per line`,
			}},
		}, {
			Name:        []string{"local"},
			Description: `Start a tmux session using ./.tmuxinator.y[a]ml`,
			Options: []model.Option{{
				Name:        []string{"--suppress-tmux-version-warning"},
				Description: `Don't show a warning for unsupported tmux versions`,
			}},
		}, {
			Name:        []string{"new"},
			Description: `Create a new project file and open it in your editor`,
			Args: []model.Arg{{
				Name:        "project",
				Description: `The project name`,
			}, {
				Name:        "tmux-session",
				Description: `The tmux session name`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-l"},
				Description: `Create local project file at ./.tmuxinator.y[a]ml`,
			}},
		}, {
			Name:        []string{"edit"},
			Description: `Open a project file it in your editor`,
			Args: []model.Arg{{
				Name:        "project",
				Description: `The project name`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"-l"},
				Description: `Open local project file at ./.tmuxinator.y[a]ml`,
			}},
		}, {
			Name:        []string{"open"},
			Description: `Open a project file it in your editor`,
			Args: []model.Arg{{
				Name:        "project",
				Description: `The project name`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"-l"},
				Description: `Open local project file at ./.tmuxinator.y[a]ml`,
			}},
		}, {
			Name:        []string{"start"},
			Description: `Start a tmux session using a project's name (with an optional [ALIAS] for project reuse) or a path to a project config file (via the -p flag)`,
			Args: []model.Arg{{
				Name:        "project",
				Description: `The project name`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"stop"},
			Description: `Stop a tmux session using a project's tmuxinator config`,
			Args: []model.Arg{{
				Name:        "project",
				Description: `The project name`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"version"},
			Description: `Display installed tmuxinator version`,
		}},
	}
}
