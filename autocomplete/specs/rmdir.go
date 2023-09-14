// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["rmdir"] = model.Subcommand{
		Name:        []string{"rmdir"},
		Description: `Remove directories`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFolders},
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-p"},
			Description: `Remove each directory of path`,
		}},
	}
}
