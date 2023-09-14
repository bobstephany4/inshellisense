// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["sqlite3"] = model.Subcommand{
		Name:        []string{"sqlite3"},
		Description: `A command line interface for SQLite version 3`,
		Args: []model.Arg{{
			Name:      "FILENAME",
			Generator: nil, // TODO: port over generator
		}},
		Options: []model.Option{{
			Name:        []string{"-append"},
			Description: `Append the database to the end of the file`,
		}, {
			Name:        []string{"-ascii"},
			Description: `Set output mode to 'ascii'`,
		}, {
			Name:        []string{"-bail"},
			Description: `Stop after hitting an error`,
		}, {
			Name:        []string{"-batch"},
			Description: `Force batch I/O`,
		}, {
			Name:        []string{"-column"},
			Description: `Set output mode to 'column'`,
		}, {
			Name:        []string{"-cmd"},
			Description: `Run "COMMAND" before reading stdin`,
			Args: []model.Arg{{
				Name: "COMMAND",
			}},
		}, {
			Name:        []string{"-csv"},
			Description: `Set output mode to 'csv'`,
		}, {
			Name:        []string{"-echo"},
			Description: `Print commands before execution`,
		}, {
			Name:        []string{"-init"},
			Description: `Read/process named file`,
			Args: []model.Arg{{
				Name:      "FILENAME",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-header"},
			Description: `Turn headers on`,
		}, {
			Name:        []string{"-noheader"},
			Description: `Turn headers off`,
		}, {
			Name:        []string{"-help"},
			Description: `Show help message`,
		}, {
			Name:        []string{"-html"},
			Description: `Set output mode to HTML`,
		}, {
			Name:        []string{"-interactive"},
			Description: `Force interactive I/O`,
		}, {
			Name:        []string{"-line"},
			Description: `Set output mode to 'line'`,
		}, {
			Name:        []string{"-list"},
			Description: `Set output mode to 'list'`,
		}, {
			Name:        []string{"-lookaside"},
			Description: `Use N entries of SZ bytes for lookaside memory`,
			Args: []model.Arg{{
				Name: "SIZE",
			}, {
				Name: "N",
			}},
		}, {
			Name:        []string{"-memtrace"},
			Description: `Trace all memory allocations and deallocations`,
		}, {
			Name:        []string{"-mmap"},
			Description: `Default mmap size set to N`,
			Args: []model.Arg{{
				Name: "N",
			}},
		}, {
			Name:        []string{"-newline"},
			Description: `Set output row separator`,
			Args: []model.Arg{{
				Name: "SEP",
			}},
		}, {
			Name:        []string{"-nofollow"},
			Description: `Refuse to open symbolic links to database files`,
		}, {
			Name:        []string{"-nullvalue"},
			Description: `Set text string for NULL values`,
			Args: []model.Arg{{
				Name: "TEXT",
			}},
		}, {
			Name:        []string{"-pagecache"},
			Description: `Use N slots of SZ bytes each for page cache memory`,
			Args: []model.Arg{{
				Name: "SIZE",
			}, {
				Name: "N",
			}},
		}, {
			Name:        []string{"-quote"},
			Description: `Set output mode to 'quote'`,
		}, {
			Name:        []string{"-readonly"},
			Description: `Open the database read-only`,
		}, {
			Name:        []string{"-separator"},
			Description: `Set output column separator`,
			Args: []model.Arg{{
				Name: "SEP",
			}},
		}, {
			Name:        []string{"-stats"},
			Description: `Print memory stats before each finalize`,
		}, {
			Name:        []string{"-version"},
			Description: `Show SQLite version`,
		}, {
			Name:        []string{"-vfs"},
			Description: `Use NAME as the default VFS`,
			Args: []model.Arg{{
				Name: "NAME",
			}},
		}},
	}
}
