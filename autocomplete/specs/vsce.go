// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["vsce"] = model.Subcommand{
		Name:        []string{"vsce"},
		Description: `The Visual Studio Code Extension Manager`,
		Options: []model.Option{{
			Name:         []string{"-h", "--help"},
			Description:  `Display help for command`,
			IsPersistent: true,
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Output the version number`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"ls"},
			Description: `Lists all the files that will be published`,
			Options: []model.Option{{
				Name:        []string{"--yarn"},
				Description: `Use yarn instead of npm (default inferred from presence of yarn.lock or .yarnrc)`,
			}, {
				Name:        []string{"--no-yarn"},
				Description: `Use npm instead of yarn (default inferred from lack of yarn.lock or .yarnrc)`,
			}, {
				Name:        []string{"--packagedDependencies"},
				Description: `Select packages that should be published only (includes dependencies)`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFilepaths},
					Name:       "paths",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--ignoreFile"},
				Description: `Indicate alternative .vscodeignore`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "path",
				}},
			}, {
				Name:        []string{"--no-dependencies"},
				Description: `Disable dependency detection via npm or yarn`,
			}},
		}, {
			Name:        []string{"package"},
			Description: `Packages an extension`,
			Options: []model.Option{{
				Name:        []string{"-o", "--out"},
				Description: `Output .vsix extension file to <path> location (defaults to <name>-<version>.vsix)`,
				Args: []model.Arg{{
					Name:      "path",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"-t", "--target"},
				Description: `Target architecture`,
				Args: []model.Arg{{
					Name:        "target",
					Suggestions: []model.Suggestion{{Name: []string{`win32-x64`}}, {Name: []string{`win32-ia32`}}, {Name: []string{`win32-arm64`}}, {Name: []string{`linux-x64`}}, {Name: []string{`linux-arm64`}}, {Name: []string{`linux-armhf`}}, {Name: []string{`alpine-x64`}}, {Name: []string{`alpine-arm64`}}, {Name: []string{`darwin-x64`}}, {Name: []string{`darwin-arm64`}}},
					IsVariadic:  true,
				}},
			}, {
				Name:        []string{"-m", "--message"},
				Description: `Commit message used when calling "npm version"`,
				Args: []model.Arg{{
					Name: "commit message",
				}},
			}, {
				Name:        []string{"--no-git-tag-version"},
				Description: `Do not create a version commit and tag when calling "npm version". Valid only when [version] is provided`,
			}, {
				Name:        []string{"--no-update-package-json"},
				Description: `Do not update "package.json". Valid only when [version] is provided`,
			}, {
				Name:        []string{"--githubBranch"},
				Description: `The GitHub branch used to infer relative links in README.md. Can be overridden by --baseContentUrl and --baseImagesUrl`,
				Args: []model.Arg{{
					Name:      "branch",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--gitlabBranch"},
				Description: `The GitLab branch used to infer relative links in README.md. Can be overridden by --baseContentUrl and --baseImagesUrl`,
				Args: []model.Arg{{
					Name:      "branch",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--no-rewrite-relative-links"},
				Description: `Skip rewriting relative links`,
			}, {
				Name:        []string{"--baseContentUrl"},
				Description: `Prepend all relative links in README.md with this url`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"--baseImagesUrl"},
				Description: `Prepend all relative image links in README.md with this url`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"--yarn"},
				Description: `Use yarn instead of npm (default inferred from presence of yarn.lock or .yarnrc)`,
			}, {
				Name:        []string{"--no-yarn"},
				Description: `Use npm instead of yarn (default inferred from lack of yarn.lock or .yarnrc)`,
			}, {
				Name:        []string{"--ignoreFile"},
				Description: `Indicate alternative .vscodeignore`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "path",
				}},
			}, {
				Name:        []string{"--no-gitHubIssueLinking"},
				Description: `Disable automatic expansion of GitHub-style issue syntax into links`,
			}, {
				Name:        []string{"--no-gitLabIssueLinking"},
				Description: `Disable automatic expansion of GitLab-style issue syntax into links`,
			}, {
				Name:        []string{"--no-dependencies"},
				Description: `Disable dependency detection via npm or yarn`,
			}, {
				Name:        []string{"--pre-release"},
				Description: `Mark this package as a pre-release`,
			}},
		}, {
			Name:        []string{"publish"},
			Description: `Publishes an extension`,
			Options: []model.Option{{
				Name:        []string{"-p", "--pat"},
				Description: `Personal Access Token (defaults to VSCE_PAT environment variable)`,
				Args: []model.Arg{{
					Name: "token",
				}},
			}, {
				Name:        []string{"-t", "--target"},
				Description: `Target architecture`,
				Args: []model.Arg{{
					Name:        "target",
					Suggestions: []model.Suggestion{{Name: []string{`win32-x64`}}, {Name: []string{`win32-ia32`}}, {Name: []string{`win32-arm64`}}, {Name: []string{`linux-x64`}}, {Name: []string{`linux-arm64`}}, {Name: []string{`linux-armhf`}}, {Name: []string{`alpine-x64`}}, {Name: []string{`alpine-arm64`}}, {Name: []string{`darwin-x64`}}, {Name: []string{`darwin-arm64`}}},
					IsVariadic:  true,
				}},
			}, {
				Name:        []string{"-m", "--message"},
				Description: `Commit message used when calling "npm version"`,
				Args: []model.Arg{{
					Name: "commit message",
				}},
			}, {
				Name:        []string{"--no-git-tag-version"},
				Description: `Do not create a version commit and tag when calling "npm version". Valid only when [version] is provided`,
			}, {
				Name:        []string{"--no-update-package-json"},
				Description: `Do not update "package.json". Valid only when [version] is provided`,
			}, {
				Name:        []string{"-i", "--packagePath"},
				Description: `Publish the provided VSIX packages`,
				Args: []model.Arg{{
					Name:       "paths",
					Generator:  nil, // TODO: port over generator
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--githubBranch"},
				Description: `The GitHub branch used to infer relative links in README.md. Can be overridden by --baseContentUrl and --baseImagesUrl`,
				Args: []model.Arg{{
					Name:      "branch",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--gitlabBranch"},
				Description: `The GitLab branch used to infer relative links in README.md. Can be overridden by --baseContentUrl and --baseImagesUrl`,
				Args: []model.Arg{{
					Name:      "branch",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--baseContentUrl"},
				Description: `Prepend all relative links in README.md with this url`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"--baseImagesUrl"},
				Description: `Prepend all relative image links in README.md with this url`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"--yarn"},
				Description: `Use yarn instead of npm (default inferred from presence of yarn.lock or .yarnrc)`,
			}, {
				Name:        []string{"--no-yarn"},
				Description: `Use npm instead of yarn (default inferred from lack of yarn.lock or .yarnrc)`,
			}, {
				Name: []string{"--noVerify"},
			}, {
				Name:        []string{"--ignoreFile"},
				Description: `Indicate alternative .vscodeignore`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "path",
				}},
			}, {
				Name:        []string{"--no-dependencies"},
				Description: `Disable dependency detection via npm or yarn`,
			}, {
				Name:        []string{"--pre-release"},
				Description: `Mark this package as a pre-release`,
			}},
		}, {
			Name:        []string{"unpublish"},
			Description: `Unpublishes an extension. Example extension id: microsoft.csharp`,
			Options: []model.Option{{
				Name:        []string{"-p", "--pat"},
				Description: `Personal Access Token`,
				Args: []model.Arg{{
					Name: "token",
				}},
			}, {
				Name:        []string{"-f", "--force"},
				Description: `Forces Unpublished Extension`,
			}},
		}, {
			Name:        []string{"ls-publishers"},
			Description: `List all known publishers`,
		}, {
			Name:        []string{"delete-publisher"},
			Description: `Deletes a publisher`,
			Args: []model.Arg{{
				Name: "publisher",
			}},
		}, {
			Name:        []string{"login"},
			Description: `Add a publisher to the known publishers list`,
			Args: []model.Arg{{
				Name: "publisher",
			}},
		}, {
			Name:        []string{"logout"},
			Description: `Remove a publisher from the known publishers list`,
			Args: []model.Arg{{
				Name: "publisher",
			}},
		}, {
			Name:        []string{"verify-pat"},
			Description: `Verify if the Personal Access Token has publish rights for the publisher`,
			Args: []model.Arg{{
				Name: "publisher",
			}},
			Options: []model.Option{{
				Name:        []string{"-p", "--pat"},
				Description: `Personal Access Token (defaults to VSCE_PAT environment variable)`,
				Args: []model.Arg{{
					Name: "token",
				}},
			}},
		}, {
			Name:        []string{"show"},
			Description: `Show extension metadata`,
			Args: []model.Arg{{
				Name: "extensionid",
			}},
			Options: []model.Option{{
				Name:        []string{"--json"},
				Description: `Output data in json format (default: false)`,
			}},
		}, {
			Name:        []string{"search"},
			Description: `Search extension gallery`,
			Args: []model.Arg{{
				Name: "text",
			}},
			Options: []model.Option{{
				Name:        []string{"--json"},
				Description: `Output data in json format (default: false)`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Display help for command`,
			Args: []model.Arg{{
				Name:        "command",
				Suggestions: []model.Suggestion{{Name: []string{`ls`}}, {Name: []string{`package`}}, {Name: []string{`publish`}}, {Name: []string{`unpublish`}}, {Name: []string{`ls-publishers`}}, {Name: []string{`delete-publishers`}}, {Name: []string{`login`}}, {Name: []string{`logout`}}, {Name: []string{`verify-pat`}}, {Name: []string{`show`}}, {Name: []string{`search`}}},
			}},
		}},
	}
}
