// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["hx"] = model.Subcommand{
		Name:        []string{"hx"},
		Description: `A post-modern text editor`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:       "files",
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Show help`,
		}, {
			Name:        []string{"--tutor"},
			Description: `Open the tutorial`,
		}, {
			Name:        []string{"--health"},
			Description: `Check for errors in editor setup`,
			Args: []model.Arg{{
				Name:       "language",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"-v"},
			Description: `Increases logging verbosity`,
		}, {
			Name:        []string{"-g", "--grammar"},
			Description: `Fetch or build tree-sitter grammars`,
			Args: []model.Arg{{
				Name: "action",
				Suggestions: []model.Suggestion{{
					Name: []string{`fetch`},
				}, {
					Name: []string{`build`},
				}},
			}},
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Print version information`,
		}},
	}
}
