// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["space"] = model.Subcommand{
		Name:        []string{"space"},
		Description: `Deta Space CLI for mananging Deta Space projects`,
		Options: []model.Option{{
			Name:         []string{"--help", "-h"},
			Description:  `Display help`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"completion"},
			Description: `Generate the autocompletion script for the specified shell`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"bash"},
				Description: `Generate the autocompletion script for bash`,
				Options: []model.Option{{
					Name:        []string{"--no-descriptions"},
					Description: `Disable completion descriptions`,
				}},
			}, {
				Name:        []string{"fish"},
				Description: `Generate the autocompletion script for fish`,
				Options: []model.Option{{
					Name:        []string{"--no-descriptions"},
					Description: `Disable completion descriptions`,
				}},
			}, {
				Name:        []string{"powershell"},
				Description: `Generate the autocompletion script for powershell`,
				Options: []model.Option{{
					Name:        []string{"--no-descriptions"},
					Description: `Disable completion descriptions`,
				}},
			}, {
				Name:        []string{"zsh"},
				Description: `Generate the autocompletion script for zsh`,
				Options: []model.Option{{
					Name:        []string{"--no-descriptions"},
					Description: `Disable completion descriptions`,
				}},
			}},
		}, {
			Name:        []string{"link"},
			Description: `Link code to project`,
			Options: []model.Option{{
				Name:        []string{"--dir", "-d"},
				Description: `Src of project to link`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "dir",
				}},
			}, {
				Name:        []string{"--id", "-i"},
				Description: `Project id of project to link`,
				Args: []model.Arg{{
					Name: "id",
				}},
			}},
		}, {
			Name:        []string{"login"},
			Description: `Login to space`,
		}, {
			Name:        []string{"new"},
			Description: `Create new project`,
			Options: []model.Option{{
				Name:        []string{"--blank", "-b"},
				Description: `Create blank project`,
			}, {
				Name:        []string{"--dir", "-d"},
				Description: `Src of project to release`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "dir",
				}},
			}, {
				Name:        []string{"--name", "-n"},
				Description: `Project name`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}},
		}, {
			Name:        []string{"open"},
			Description: `Open current project in browser`,
			Options: []model.Option{{
				Name:        []string{"--dir", "-d"},
				Description: `Src of project to open`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "dir",
				}},
			}, {
				Name:        []string{"--id", "-i"},
				Description: `Project id of project to open`,
				Args: []model.Arg{{
					Name: "id",
				}},
			}},
		}, {
			Name:        []string{"push"},
			Description: `Push code for project`,
			Options: []model.Option{{
				Name:        []string{"--dir", "-d"},
				Description: `Src of project to push`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "dir",
				}},
			}, {
				Name:        []string{"--id", "-i"},
				Description: `Project id of project to push`,
				Args: []model.Arg{{
					Name: "id",
				}},
			}, {
				Name:        []string{"--open", "-o"},
				Description: `Open builder instance/project in browser after push`,
			}, {
				Name:        []string{"--skip-logs"},
				Description: `Skip following logs after push`,
			}, {
				Name:        []string{"--tag", "-t"},
				Description: `Tag to identify this push`,
				Args: []model.Arg{{
					Name: "tag",
				}},
			}},
		}, {
			Name:        []string{"release"},
			Description: `Create release for a project`,
			Options: []model.Option{{
				Name:        []string{"--confirm", "-c"},
				Description: `Release latest revision`,
			}, {
				Name:        []string{"--dir", "-d"},
				Description: `Src of project to release`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "dir",
				}},
			}, {
				Name:        []string{"--id", "-i"},
				Description: `Project id of an existing project`,
				Args: []model.Arg{{
					Name: "id",
				}},
			}, {
				Name:        []string{"--listed", "-l"},
				Description: `Listed on discovery`,
			}, {
				Name:        []string{"--notes", "-n"},
				Description: `Release notes`,
				Args: []model.Arg{{
					Name: "notes",
				}},
			}, {
				Name:        []string{"--rid", "-r"},
				Description: `Revision id for release`,
				Args: []model.Arg{{
					Name: "rid",
				}},
			}, {
				Name:        []string{"--version", "-v"},
				Description: `Version for the release`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}},
		}, {
			Name:        []string{"validate"},
			Description: `Validate spacefile in dir`,
			Options: []model.Option{{
				Name:        []string{"--dir", "-d"},
				Description: `Src of project to validate`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "dir",
				}},
			}},
		}, {
			Name:        []string{"version"},
			Description: `Space CLI version`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"upgrade"},
				Description: `Upgrade Space CLI version`,
				Options: []model.Option{{
					Name:        []string{"--version", "-v"},
					Description: `Version number`,
					Args: []model.Arg{{
						Name: "version",
					}},
				}},
			}},
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"completion"},
				Description: `Generate the autocompletion script for the specified shell`,
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
				Name:        []string{"link"},
				Description: `Link code to project`,
			}, {
				Name:        []string{"login"},
				Description: `Login to space`,
			}, {
				Name:        []string{"new"},
				Description: `Create new project`,
			}, {
				Name:        []string{"open"},
				Description: `Open current project in browser`,
			}, {
				Name:        []string{"push"},
				Description: `Push code for project`,
			}, {
				Name:        []string{"release"},
				Description: `Create release for a project`,
			}, {
				Name:        []string{"validate"},
				Description: `Validate spacefile in dir`,
			}, {
				Name:        []string{"version"},
				Description: `Space CLI version`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"upgrade"},
					Description: `Upgrade Space CLI version`,
				}},
			}},
		}},
	}
}
