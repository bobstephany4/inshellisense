// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["head"] = model.Subcommand{
		Name:        []string{"head"},
		Description: `Output the first part of files`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFilepaths},
			Name:      "file",
		}},
		Options: []model.Option{{
			Name:        []string{"-c", "--bytes"},
			Description: `Print the first [numBytes] bytes of each file`,
			Args: []model.Arg{{
				Name: "numBytes",
			}},
		}, {
			Name:        []string{"-n", "--lines"},
			Description: `Print the first [numLines] lines instead of the first 10`,
			Args: []model.Arg{{
				Name: "numLines",
			}},
		}, {
			Name:        []string{"-q", "--quiet", "--silent"},
			Description: `Never print headers giving file names`,
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Always print headers giving file names`,
		}, {
			Name:        []string{"--help"},
			Description: `Display this help and exit`,
		}, {
			Name:        []string{"--version"},
			Description: `Output version information and exit`,
		}},
	}
}
