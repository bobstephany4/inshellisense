// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["ruff"] = model.Subcommand{
		Name:        []string{"ruff"},
		Description: `Ruff: An extremely fast Python linter`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths},
			Name:       "files",
			IsOptional: true,
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"--config"},
			Description: `Path to the "pyproject.toml" or "ruff.toml" file to use for configuration`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "config",
				IsOptional: true,
			}},
			ExclusiveOn: []string{"--isolated"},
		}, {
			Name:        []string{"--select"},
			Description: `Comma-separated list of rule codes to enable (or ALL, to enable all rules)`,
			Args: []model.Arg{{
				Name:       "select",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--extend-select"},
			Description: `Like --select, but adds additional rule codes on top of the selected ones`,
			Args: []model.Arg{{
				Name:       "extend_select",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--ignore"},
			Description: `Comma-separated list of rule codes to disable`,
			Args: []model.Arg{{
				Name:       "ignore",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--extend-ignore"},
			Description: `Like --ignore, but adds additional rule codes on top of the ignored ones`,
			Args: []model.Arg{{
				Name:       "extend_ignore",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--exclude"},
			Description: `List of paths, used to omit files and/or directories from analysis`,
			Args: []model.Arg{{
				Name:       "exclude",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--extend-exclude"},
			Description: `Like --exclude, but adds additional files and directories on top of those already excluded`,
			Args: []model.Arg{{
				Name:       "extend_exclude",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--fixable"},
			Description: `List of rule codes to treat as eligible for autofix. Only applicable when autofix itself is enabled (e.g., via "--fix")`,
			Args: []model.Arg{{
				Name:       "fixable",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--unfixable"},
			Description: `List of rule codes to treat as ineligible for autofix. Only applicable when autofix itself is enabled (e.g., via "--fix")`,
			Args: []model.Arg{{
				Name:       "unfixable",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--per-file-ignores"},
			Description: `List of mappings from file pattern to code to exclude`,
			Args: []model.Arg{{
				Name:       "per_file_ignores",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--format"},
			Description: `Output serialization format for violations`,
			Args: []model.Arg{{
				Name:        "format",
				Suggestions: []model.Suggestion{{Name: []string{`text`}}, {Name: []string{`json`}}, {Name: []string{`junit`}}, {Name: []string{`grouped`}}, {Name: []string{`github`}}, {Name: []string{`gitlab`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"--stdin-filename"},
			Description: `The name of the file when passing it through stdin`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "stdin_filename",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--cache-dir"},
			Description: `Path to the cache directory`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "cache_dir",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--dummy-variable-rgx"},
			Description: `Regular expression matching the name of dummy variables`,
			Args: []model.Arg{{
				Name:       "dummy_variable_rgx",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--target-version"},
			Description: `The minimum Python version that should be supported`,
			Args: []model.Arg{{
				Name:       "target_version",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--line-length"},
			Description: `Set the line-length for length-associated rules and automatic formatting`,
			Args: []model.Arg{{
				Name:       "line_length",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--explain"},
			Description: `Explain a rule`,
			Args: []model.Arg{{
				Name:       "explain",
				IsOptional: true,
			}},
			ExclusiveOn: []string{"--add-noqa", "--clean", "--generate-shell-completion", "--show-files", "--show-settings", "--stdin-filename", "-w", "--watch"},
		}, {
			Name:        []string{"--generate-shell-completion"},
			Description: `Generate shell completion`,
			Args: []model.Arg{{
				Name:        "generate_shell_completion",
				Suggestions: []model.Suggestion{{Name: []string{`bash`}}, {Name: []string{`elvish`}}, {Name: []string{`fig`}}, {Name: []string{`fish`}}, {Name: []string{`powershell`}}, {Name: []string{`zsh`}}},
				IsOptional:  true,
			}},
			ExclusiveOn: []string{"--add-noqa", "--clean", "--explain", "--show-files", "--show-settings", "--stdin-filename", "-w", "--watch"},
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Enable verbose logging`,
		}, {
			Name:        []string{"-q", "--quiet"},
			Description: `Print lint violations, but nothing else`,
		}, {
			Name:        []string{"-s", "--silent"},
			Description: `Disable all logging (but still exit with status code "1" upon detecting lint violations)`,
		}, {
			Name:        []string{"-e", "--exit-zero"},
			Description: `Exit with status code "0", even upon detecting lint violations`,
		}, {
			Name:        []string{"-w", "--watch"},
			Description: `Run in watch mode by re-running whenever files change`,
		}, {
			Name:        []string{"--fix"},
			Description: `Attempt to automatically fix lint violations`,
		}, {
			Name: []string{"--no-fix"},
		}, {
			Name:        []string{"--fix-only"},
			Description: `Fix any fixable lint violations, but don't report on leftover violations. Implies "--fix"`,
		}, {
			Name: []string{"--no-fix-only"},
		}, {
			Name:        []string{"--diff"},
			Description: `Avoid writing any fixed files back; instead, output a diff for each changed file to stdout`,
		}, {
			Name:        []string{"-n", "--no-cache"},
			Description: `Disable cache reads`,
		}, {
			Name:        []string{"--isolated"},
			Description: `Ignore all configuration files`,
			ExclusiveOn: []string{"--config"},
		}, {
			Name:        []string{"--show-source"},
			Description: `Show violations with source code`,
		}, {
			Name: []string{"--no-show-source"},
		}, {
			Name:        []string{"--respect-gitignore"},
			Description: `Respect file exclusions via ".gitignore" and other standard ignore files`,
		}, {
			Name: []string{"--no-respect-gitignore"},
		}, {
			Name:        []string{"--force-exclude"},
			Description: `Enforce exclusions, even for paths passed to Ruff directly on the command-line`,
		}, {
			Name: []string{"--no-force-exclude"},
		}, {
			Name:        []string{"--update-check"},
			Description: `Enable or disable automatic update checks`,
		}, {
			Name: []string{"--no-update-check"},
		}, {
			Name:        []string{"--add-noqa"},
			Description: `Enable automatic additions of "noqa" directives to failing lines`,
			ExclusiveOn: []string{"--clean", "--explain", "--generate-shell-completion", "--show-files", "--show-settings", "--stdin-filename", "-w", "--watch"},
		}, {
			Name:        []string{"--clean"},
			Description: `Clear any caches in the current directory or any subdirectories`,
			ExclusiveOn: []string{"--add-noqa", "--explain", "--generate-shell-completion", "--show-files", "--show-settings", "--stdin-filename", "-w", "--watch"},
		}, {
			Name:        []string{"--show-files"},
			Description: `See the files Ruff will be run against with the current settings`,
			ExclusiveOn: []string{"--add-noqa", "--clean", "--explain", "--generate-shell-completion", "--show-settings", "--stdin-filename", "-w", "--watch"},
		}, {
			Name:        []string{"--show-settings"},
			Description: `See the settings Ruff will use to lint a given Python file`,
			ExclusiveOn: []string{"--add-noqa", "--clean", "--explain", "--generate-shell-completion", "--show-files", "--stdin-filename", "-w", "--watch"},
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Print help`,
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Print version`,
		}},
	}
}
