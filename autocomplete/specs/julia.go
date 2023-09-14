// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["julia"] = model.Subcommand{
		Name:        []string{"julia"},
		Description: `The Julia Programming Language`,
		Args: []model.Arg{{
			Name:       "julia script",
			Generator:  nil, // TODO: port over generator
			IsOptional: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-v", "--version"},
			Description: `Display version information`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Print help message for julia (--help-hidden for more)`,
		}, {
			Name:        []string{"--help-hidden"},
			Description: `Uncommon options not shown by "-h"`,
		}, {
			Name:        []string{"--project"},
			Description: `Set given directory as the home project/environment`,
			Args: []model.Arg{{
				Name:        "project folder",
				Description: `Julia project/environment`,
				Suggestions: []model.Suggestion{{
					Name:        []string{`@.`},
					Description: `Search through parent directories until a Project.toml or JuliaProject.toml file is found`,
				}},
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
		}, {
			Name:        []string{"-J", "--sysimage"},
			Description: `Start up with the given system image file`,
			Args: []model.Arg{{
				Name:      "system image",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-H", "--home"},
			Description: `Set location of "julia" executable`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
			}},
		}, {
			Name:        []string{"--startup-file"},
			Description: `Load "~/.julia/config/startup.jl"`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"--handle-signals"},
			Description: `Enable or disable Julia's default signal handlers`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"--sysimage-native-code"},
			Description: `Use native code from system image if available`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"--compiled-modules"},
			Description: `Enable or disable incremental precompilation of modules`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"-e", "--eval"},
			Description: `Evaluate given expr`,
			Args: []model.Arg{{
				Name: "expr",
			}},
		}, {
			Name:        []string{"-E", "--print"},
			Description: `Evaluate given expr and display the result`,
			Args: []model.Arg{{
				Name: "expr",
			}},
		}, {
			Name:        []string{"-L", "--load"},
			Description: `Load given file immediately on all processors`,
			Args: []model.Arg{{
				Name:      "julia script",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-t", "--threads"},
			Description: `Enable N threads; "auto" sets N to the number of local CPU threads`,
			Args: []model.Arg{{
				Description: `Number of threads`,
				Suggestions: []model.Suggestion{{
					Name: []string{`auto`},
				}},
			}},
		}, {
			Name:        []string{"-p", "--procs"},
			Description: `Integer value N launches N additional local worker processes "auto" launches as many workers as the number of local CPU threads`,
			Args: []model.Arg{{
				Description: `Number of additional local worker processes`,
				Suggestions: []model.Suggestion{{
					Name: []string{`auto`},
				}},
			}},
		}, {
			Name:        []string{"--machine-file"},
			Description: `Run processes on hosts listed in given file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}, {
			Name:        []string{"-i"},
			Description: `Interactive mode; REPL runs and isinteractive() is true`,
		}, {
			Name:        []string{"-q", "--quiet"},
			Description: `Quiet startup: no banner, suppress REPL warnings`,
		}, {
			Name:        []string{"--banner"},
			Description: `Enable or disable startup banner`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}, {
					Name: []string{`auto`},
				}},
			}},
		}, {
			Name:        []string{"--color"},
			Description: `Enable or disable color text`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}, {
					Name: []string{`auto`},
				}},
			}},
		}, {
			Name:        []string{"--history-file"},
			Description: `Load or save history`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"--depwarn"},
			Description: `Enable or disable syntax and method deprecation warnings ("error" turns warnings into errors)`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}, {
					Name: []string{`error`},
				}},
			}},
		}, {
			Name:        []string{"--warn-overwrite"},
			Description: `Enable or disable method overwrite warnings`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"--warn-scope"},
			Description: `Enable or disable warning for ambiguous top-level scope`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"-C", "--cpu-target"},
			Description: `Limit usage of CPU features up to <target>; set to "help" to see the available options`,
			Args:        []model.Arg{{}},
		}, {
			Name:        []string{"-O", "--optimize"},
			Description: `Set the optimization level (default level is 2 if unspecified or 3 if used without a level)`,
			Args: []model.Arg{{
				Name:        "level",
				Description: `Level of optimization`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"-g"},
			Description: `Enable / Set the level of debug info generation`,
			Args: []model.Arg{{
				Name:        "level",
				Description: `Level of debug info generation`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"--inline"},
			Description: `Control whether inlining is permitted, including overriding @inline declarations`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"--check-bounds"},
			Description: `Emit bounds checks always, never, or respect @inbounds declarations`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}, {
					Name: []string{`auto`},
				}},
			}},
		}, {
			Name:        []string{"--polly"},
			Description: `Enable or disable the polyhedral optimizer Polly (overrides @polly declaration)`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"--math-mode"},
			Description: `Disallow or enable unsafe floating point optimizations (overrides @fastmath declaration)`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`ieee`},
				}, {
					Name: []string{`fast`},
				}},
			}},
		}, {
			Name:        []string{"--code-coverage"},
			Description: `Count executions of source lines (omitting setting is equivalent to "user")`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`none`},
				}, {
					Name: []string{`user`},
				}, {
					Name: []string{`all`},
				}},
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--track-allocation"},
			Description: `Count bytes allocated by each source line (omitting setting is equivalent to "user")`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`none`},
				}, {
					Name: []string{`user`},
				}, {
					Name: []string{`all`},
				}},
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--bug-report"},
			Description: `Launch a bug report session. It can be used to start a REPL, run a script, or evaluate  expressions. It first tries to use BugReporting.jl installed in current environment and fallbacks to the latest compatible BugReporting.jl if not. For more information, see --bug-report=help`,
			Args: []model.Arg{{
				Name: "KIND",
			}},
		}, {
			Name:        []string{"--compile"},
			Description: `Enable or disable JIT compiler, or request exhaustive or minimal compilation`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}, {
					Name: []string{`all`},
				}, {
					Name: []string{`min`},
				}},
			}},
		}, {
			Name:        []string{"--output-o"},
			Description: `Generate an object file (including system image data)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "name",
			}},
		}, {
			Name:        []string{"--output-ji"},
			Description: `Generate a system image data file (.ji)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "name",
			}},
		}, {
			Name:        []string{"--output-unopt-bc"},
			Description: `Generate unoptimized LLVM bitcode (.bc)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "name",
			}},
		}, {
			Name:        []string{"--output-jit-bc"},
			Description: `Dump all IR generated by the frontend (not including system image)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "name",
			}},
		}, {
			Name:        []string{"--output-bc"},
			Description: `Generate LLVM bitcode (.bc)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "name",
			}},
		}, {
			Name:        []string{"--output-asm"},
			Description: `Generate an assembly file (.s)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "name",
			}},
		}, {
			Name:        []string{"--output-incremental"},
			Description: `Generate an incremental output file (rather than complete)`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name: []string{`yes`},
				}, {
					Name: []string{`no`},
				}},
			}},
		}, {
			Name:        []string{"--image-codegen"},
			Description: `Force generate code in imaging mode`,
		}},
	}
}
