// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["ts-node"] = model.Subcommand{
		Name:        []string{"ts-node"},
		Description: `Run the TypeScript interpreter for Node.JS`,
		Args: []model.Arg{{
			Name:      "script",
			Generator: nil, // TODO: port over generator
		}},
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for ts-node`,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Print version information of the ts-node module`,
		}, {
			Name:        []string{"-e", "--eval"},
			Description: `Evaluate script`,
			Args: []model.Arg{{
				Name: "script",
			}},
		}, {
			Name:        []string{"-p", "--print"},
			Description: `Evaluate script and print result`,
		}, {
			Name:        []string{"-r", "--require"},
			Description: `Require module before executing`,
			Args: []model.Arg{{
				Name:           "module",
				FilterStrategy: model.FilterStrategyFuzzy,
				Generator:      nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-i", "--interactive"},
			Description: `Always open interactive REPL`,
		}, {
			Name:        []string{"--show-config"},
			Description: `Print resolved Typescript config to the terminal`,
		}, {
			Name:        []string{"--cwd-mode"},
			Description: `Resolve Typescript config based on the current working directory`,
		}, {
			Name:        []string{"-T", "--transpile-only"},
			Description: `Use the Typescript transpile module mode`,
		}, {
			Name:        []string{"-H", "--compiler-host"},
			Description: `Use the Typescript compiler host API`,
		}, {
			Name:        []string{"-I", "--ignore"},
			Description: `Ignore patterns from Typescript compilation`,
			Args: []model.Arg{{
				Name: "pattern",
			}},
		}, {
			Name:        []string{"-P", "--project"},
			Description: `Specify TypeScript project location`,
			Args: []model.Arg{{
				Name:      "project",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-C", "--compiler"},
			Description: `Use a custom compiler`,
			Args: []model.Arg{{
				Name: "compiler",
			}},
		}, {
			Name:        []string{"--transpiler"},
			Description: `Use a custom transpiler`,
			Args: []model.Arg{{
				Name: "transpiler",
			}},
		}, {
			Name:        []string{"-D", "--ignore-diagnostics"},
			Description: `Specify Typescript diagnostic code to ignore`,
			Args: []model.Arg{{
				Name: "code",
			}},
		}, {
			Name:        []string{"-O", "--compiler-options"},
			Description: `JSON object that will be merged with the compiler options`,
			Args: []model.Arg{{
				Name: "options",
			}},
		}, {
			Name:        []string{"--cwd"},
			Description: `Specify working directory`,
			Args: []model.Arg{{
				Name:      "cwd",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--files"},
			Description: `Load files, include and exclude from Typescript config on startup`,
		}, {
			Name:        []string{"--pretty"},
			Description: `Use the pretty formatter for diagnostic errors`,
		}, {
			Name:        []string{"--skip-project"},
			Description: `Skip reading Typescript config`,
		}, {
			Name:        []string{"--scope"},
			Description: `Scope compilation to scope directory specified`,
		}, {
			Name:        []string{"--scope-dir"},
			Description: `Directory for scope parameter`,
			Args: []model.Arg{{
				Name:      "directory",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--skip-ignore"},
			Description: `Skip --ignore checks`,
		}, {
			Name:        []string{"--prefer-ts-exts"},
			Description: `Prefer Typescript files over JavaScript files when importing files`,
		}, {
			Name:        []string{"--log-error"},
			Description: `Pipe Typescript errors to stderr instead of throwing exceptions`,
		}, {
			Name:        []string{"--no-experimental-repl-await"},
			Description: `Disable the top-level await function in REPL`,
		}},
	}
}
