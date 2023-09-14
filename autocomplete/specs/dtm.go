// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["dtm"] = model.Subcommand{
		Name:        []string{"dtm"},
		Description: `DevStream is an open-source DevOps toolchain manager`,
		Options: []model.Option{{
			Name:         []string{"--debug"},
			Description:  `Debug level log`,
			IsPersistent: true,
		}, {
			Name:         []string{"--help", "-h"},
			Description:  `Display help`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"apply"},
			Description: `Create or update DevOps tools according to DevStream configuration file`,
			Options: []model.Option{{
				Name:        []string{"--config-file", "-f"},
				Description: `Config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config-file",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--plugin-dir", "-d"},
				Description: `Plugins directory`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "plugin-dir",
				}},
			}, {
				Name:        []string{"--yes", "-y"},
				Description: `Apply directly without confirmation`,
			}},
		}, {
			Name:        []string{"completion"},
			Description: `Generate the autocompletion script for dtm for the specified shell`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"bash"},
				Description: `Generate autocompletion script for bash`,
			}, {
				Name:        []string{"fish"},
				Description: `Generate autocompletion script for fish`,
			}, {
				Name:        []string{"powershell"},
				Description: `Generate autocompletion script for powershell`,
			}, {
				Name:        []string{"zsh"},
				Description: `Generate autocompletion script for zsh`,
			}},
		}, {
			Name:        []string{"delete"},
			Description: `Delete DevOps tools according to DevStream configuration file`,
			Options: []model.Option{{
				Name:        []string{"--config-file", "-f"},
				Description: `Config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config-file",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--force"},
				Description: `Force delete by config`,
			}, {
				Name:        []string{"--plugin-dir", "-d"},
				Description: `Plugins directory`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "plugin-dir",
				}},
			}, {
				Name:        []string{"--yes", "-y"},
				Description: `Delete directly without confirmation`,
			}},
		}, {
			Name:        []string{"destroy"},
			Description: `Destroy DevOps tools deployment according to DevStream configuration file & state file`,
			Options: []model.Option{{
				Name:        []string{"--config-file", "-f"},
				Description: `Config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config-file",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--force"},
				Description: `Force destroy by config`,
			}, {
				Name:        []string{"--plugin-dir", "-d"},
				Description: `Plugins directory`,
				Args: []model.Arg{{
					Name: "plugin-dir",
				}},
			}, {
				Name:        []string{"--yes", "-y"},
				Description: `Destroy directly without confirmation`,
			}},
		}, {
			Name:        []string{"develop"},
			Description: `Develop is used for develop a new plugin`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create-plugin"},
				Description: `Create a new plugin`,
				Options: []model.Option{{
					Name:        []string{"--name", "-n"},
					Description: `Specify name of the plugin to be created`,
					Args: []model.Arg{{
						Name: "name",
					}},
					IsPersistent: true,
				}},
			}, {
				Name:        []string{"validate-plugin"},
				Description: `Validate a plugin`,
				Options: []model.Option{{
					Name:         []string{"--all", "-a"},
					Description:  `Validate all plugins`,
					IsPersistent: true,
				}, {
					Name:        []string{"--name", "-n"},
					Description: `Specify name of the plugin to be validated`,
					Args: []model.Arg{{
						Name: "name",
					}},
					IsPersistent: true,
				}},
			}},
		}, {
			Name:        []string{"init"},
			Description: `Download needed plugins according to the config file`,
			Options: []model.Option{{
				Name:        []string{"--all", "-a"},
				Description: `Download all plugins`,
			}, {
				Name:        []string{"--arch"},
				Description: `Download plugins for specific arch`,
				Args: []model.Arg{{
					Name: "arch",
				}},
			}, {
				Name:        []string{"--config-file", "-f"},
				Description: `Config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config-file",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--download-only"},
				Description: `Download plugins only`,
			}, {
				Name:        []string{"--os"},
				Description: `Download plugins for specific os`,
				Args: []model.Arg{{
					Name: "os",
				}},
			}, {
				Name:        []string{"--plugin-dir", "-d"},
				Description: `Plugins directory`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "plugin-dir",
				}},
			}, {
				Name:        []string{"--plugins", "-p"},
				Description: `The plugins to be downloaded`,
				Args: []model.Arg{{
					Name: "plugins",
				}},
			}},
		}, {
			Name:        []string{"list"},
			Description: `This command only supports listing plugins now`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"plugins"},
				Description: `List all plugins`,
				Options: []model.Option{{
					Name:        []string{"--filter", "-r"},
					Description: `Filter plugin by regex`,
					Args: []model.Arg{{
						Name: "filter",
					}},
					IsPersistent: true,
				}},
			}},
		}, {
			Name:        []string{"show"},
			Description: `Show is used to print plugins' configuration templates or status`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"config"},
				Description: `Show configuration information`,
				Options: []model.Option{{
					Name:        []string{"--plugin", "-p"},
					Description: `Specify name with the plugin`,
					Args: []model.Arg{{
						Name:      "plugin",
						Generator: nil, // TODO: port over generator
					}},
				}, {
					Name:        []string{"--template", "-t"},
					Description: `Print a template config, e.g. quickstart/gitops/`,
					Args: []model.Arg{{
						Name: "template",
					}},
				}},
			}, {
				Name:        []string{"status"},
				Description: `Show status information`,
				Options: []model.Option{{
					Name:        []string{"--all", "-a"},
					Description: `Show all instances of all plugins status`,
				}, {
					Name:        []string{"--config-file", "-f"},
					Description: `Config file`,
					Args: []model.Arg{{
						Name: "config-file",
					}},
				}, {
					Name:        []string{"--id", "-i"},
					Description: `Specify id with the plugin instance`,
					Args: []model.Arg{{
						Name: "id",
					}},
				}, {
					Name:        []string{"--plugin", "-p"},
					Description: `Specify name with the plugin`,
					Args: []model.Arg{{
						Name: "plugin",
					}},
				}, {
					Name:        []string{"--plugin-dir", "-d"},
					Description: `Plugins directory`,
					Args: []model.Arg{{
						Name: "plugin-dir",
					}},
				}},
			}},
		}, {
			Name:        []string{"upgrade"},
			Description: `Upgrade dtm to the latest release version`,
			Options: []model.Option{{
				Name:        []string{"--yes", "-y"},
				Description: `Upgrade directly without confirmation`,
			}},
		}, {
			Name:        []string{"verify"},
			Description: `Verify DevOps tools according to DevStream config file and state`,
			Options: []model.Option{{
				Name:        []string{"--config-file", "-f"},
				Description: `Config file`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config-file",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--plugin-dir", "-d"},
				Description: `Plugins directory`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "plugin-dir",
				}},
			}},
		}, {
			Name:        []string{"version"},
			Description: `Print the version number of DevStream`,
		}, {
			Name:        []string{"help"},
			Description: `Help about any command`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"apply"},
				Description: `Create or update DevOps tools according to DevStream configuration file`,
			}, {
				Name:        []string{"completion"},
				Description: `Generate the autocompletion script for dtm for the specified shell`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"bash"},
					Description: `Generate autocompletion script for bash`,
				}, {
					Name:        []string{"fish"},
					Description: `Generate autocompletion script for fish`,
				}, {
					Name:        []string{"powershell"},
					Description: `Generate autocompletion script for powershell`,
				}, {
					Name:        []string{"zsh"},
					Description: `Generate autocompletion script for zsh`,
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Delete DevOps tools according to DevStream configuration file`,
			}, {
				Name:        []string{"destroy"},
				Description: `Destroy DevOps tools deployment according to DevStream configuration file & state file`,
			}, {
				Name:        []string{"develop"},
				Description: `Develop is used for develop a new plugin`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"create-plugin"},
					Description: `Create a new plugin`,
				}, {
					Name:        []string{"validate-plugin"},
					Description: `Validate a plugin`,
				}},
			}, {
				Name:        []string{"init"},
				Description: `Download needed plugins according to the config file`,
			}, {
				Name:        []string{"list"},
				Description: `This command only supports listing plugins now`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"plugins"},
					Description: `List all plugins`,
				}},
			}, {
				Name:        []string{"show"},
				Description: `Show is used to print plugins' configuration templates or status`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"config"},
					Description: `Show configuration information`,
				}, {
					Name:        []string{"status"},
					Description: `Show status information`,
				}},
			}, {
				Name:        []string{"upgrade"},
				Description: `Upgrade dtm to the latest release version`,
			}, {
				Name:        []string{"verify"},
				Description: `Verify DevOps tools according to DevStream config file and state`,
			}, {
				Name:        []string{"version"},
				Description: `Print the version number of DevStream`,
			}},
		}},
	}
}
