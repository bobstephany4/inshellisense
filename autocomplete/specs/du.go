// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["du"] = model.Subcommand{
		Name:        []string{"du"},
		Description: `Display disk usage statistics`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:       "files",
			IsOptional: true,
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-a"},
			Description: `Display an entry for each file in a file hierarchy`,
			ExclusiveOn: []string{"-s", "-d"},
		}, {
			Name:        []string{"-c"},
			Description: `Display a grand total`,
		}, {
			Name:        []string{"-H"},
			Description: `Symbolic links on the command line are followed, symbolic links in file hierarchies are not followed`,
			ExclusiveOn: []string{"-L", "-P"},
		}, {
			Name:        []string{"-h"},
			Description: `"Human-readable" output.  Use unit suffixes: Byte, Kilobyte, Megabyte, Gigabyte, Terabyte and Petabyte`,
			ExclusiveOn: []string{"-k", "-m", "-g"},
		}, {
			Name:        []string{"-g"},
			Description: `Display block counts in 1073741824-byte (1-Gbyte) blocks`,
			ExclusiveOn: []string{"-k", "-m", "-h"},
		}, {
			Name:        []string{"-k"},
			Description: `Display block counts in 1024-byte (1-Kbyte) blocks`,
			ExclusiveOn: []string{"-g", "-m", "-h"},
		}, {
			Name:        []string{"-m"},
			Description: `Display block counts in 1048576-byte (1-Mbyte) blocks`,
			ExclusiveOn: []string{"-g", "-k", "-h"},
		}, {
			Name:        []string{"-I"},
			Description: `Ignore files and directories matching the specified mask`,
			Args: []model.Arg{{
				Name: "mask",
			}},
		}, {
			Name:        []string{"-L"},
			Description: `Symbolic links on the command line and in file hierarchies are followed`,
			ExclusiveOn: []string{"-H", "-P"},
		}, {
			Name:        []string{"-r"},
			Description: `Generate messages about directories that cannot be read, files that cannot be opened, and so on.  This is the default case.  This option exists solely for conformance with X/Open Portability Guide Issue 4 (""XPG4'')`,
		}, {
			Name:        []string{"-P"},
			Description: `No symbolic links are followed.  This is the default`,
			ExclusiveOn: []string{"-H", "-L"},
		}, {
			Name:        []string{"-d"},
			Description: `Display an entry for all files and directories depth directories deep`,
			Args: []model.Arg{{
				Name:        "depth",
				Suggestions: []model.Suggestion{{Name: []string{`0`}}, {Name: []string{`1`}}, {Name: []string{`2`}}},
			}},
			ExclusiveOn: []string{"-a", "-s"},
		}, {
			Name:        []string{"-s"},
			Description: `Display an entry for each specified file.  (Equivalent to -d 0)`,
			ExclusiveOn: []string{"-a", "-d"},
		}, {
			Name:        []string{"-x"},
			Description: `Display an entry for each specified file.  (Equivalent to -d 0)`,
		}},
	}
}
