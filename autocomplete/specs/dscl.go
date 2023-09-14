// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["dscl"] = model.Subcommand{
		Name:        []string{"dscl"},
		Description: `Directory Service command line utility`,
		Args: []model.Arg{{
			Name:        "datasource",
			Suggestions: []model.Suggestion{{Name: []string{`.`}}, {Name: []string{`..`}}, {Name: []string{`localhost`}}, {Name: []string{`localonly`}}},
			IsOptional:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-p"},
			Description: `Prompt for password`,
		}, {
			Name:        []string{"-u"},
			Description: `Authenticate as user`,
			Args: []model.Arg{{
				Name:        "user",
				Description: `User to authenticate as`,
			}},
		}, {
			Name:        []string{"-P"},
			Description: `Authenticate with password`,
			Args: []model.Arg{{
				Name:        "password",
				Description: `Password to authenticate with`,
			}},
		}, {
			Name:        []string{"-f"},
			Description: `Targeted local node database file path`,
			Args: []model.Arg{{
				Name:        "file",
				Description: `File path`,
			}},
		}, {
			Name:        []string{"-raw"},
			Description: `Don't strip off prefix from DirectoryService API constants`,
		}, {
			Name:        []string{"-plist"},
			Description: `Print out record(s) or attribute(s) in XML plist format`,
		}, {
			Name:        []string{"-url"},
			Description: `Print record attribute values in URL-style encoding`,
		}, {
			Name:        []string{"-q"},
			Description: `Quiet - no interactive prompt`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"read", "-read"},
			Description: `Prints a directory`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}, {
				Name:        "keys",
				Description: `Directory Service record attribute type`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"readall", "-readall"},
			Description: `Prints all the records of a given type`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}, {
				Name:        "keys",
				Description: `Directory Service record attribute type`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"readpl", "-readpl"},
			Description: `Prints the contents of plist_path`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "plist path",
				Description: `Path to the plist file`,
			}},
		}, {
			Name:        []string{"readpli", "-readpli"},
			Description: `Prints the contents of plist_path for the plist at value_index of the key`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "value index",
				Description: `Value index of the key`,
			}, {
				Name:        "plist path",
				Description: `Path to the plist file`,
			}},
		}, {
			Name:        []string{"list", "-list", "ls", "-ls"},
			Description: `Lists the subdirectories of the given directory`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"search", "-search"},
			Description: `Searches for records that match a pattern`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "val",
				Description: `Value of the key`,
			}},
		}, {
			Name:        []string{"create", "-create", "mk", "-mk"},
			Description: `Creates a new record`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
				IsOptional:  true,
			}, {
				Name:        "vals",
				Description: `Value of the key`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"createpl", "-createpl"},
			Description: `Creates a string, or array of strings at plist_path`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "plist path",
				Description: `Path to the plist file`,
			}, {
				Name:        "val",
				Description: `Value of the key`,
			}, {
				Name:        "vals",
				Description: `Value of the key`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"createpli", "-createpli"},
			Description: `Creates a string, or array of strings at plist_path for the plist at value_index of the key`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "value index",
				Description: `Value index of the key`,
			}, {
				Name:        "plist path",
				Description: `Path to the plist file`,
			}, {
				Name:        "val",
				Description: `Value of the key`,
			}, {
				Name:        "vals",
				Description: `Value of the key`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"append", "-append"},
			Description: `Appends one or more values to a property in a given record`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "vals",
				Description: `Value of the key`,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"merge", "-merge"},
			Description: `Appends one or more values to a property in a given directory if the property does not already have those values`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "vals",
				Description: `Value of the key`,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"delete", "-delete", "rm", "-rm"},
			Description: `Delete a directory, property, or value`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
				IsOptional:  true,
			}, {
				Name:        "vals",
				Description: `Value of the key`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"deletepl", "-deletepl"},
			Description: `Deletes a value in a plist`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "plist path",
				Description: `Path to the plist file`,
			}, {
				Name:        "vals",
				Description: `Value of the key`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"deletepli", "-deletepli"},
			Description: `Deletes a value for the plist at value_index of the key`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "value index",
				Description: `Value index of the key`,
			}, {
				Name:        "plist path",
				Description: `Path to the plist file`,
			}, {
				Name:        "vals",
				Description: `Value of the key`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"change", "-change"},
			Description: `Replaces the given old value in the list of values of the given key with the new value in the specified record`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "old val",
				Description: `Old value of the key`,
			}, {
				Name:        "new val",
				Description: `New value of the key`,
			}},
		}, {
			Name:        []string{"changei", "-changei"},
			Description: `Replaces the value at the given index in the list of values of the given key with the new value in the specified record`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "key",
				Description: `Directory Service record attribute type`,
			}, {
				Name:        "value index",
				Description: `Value index of the key`,
			}, {
				Name:        "new val",
				Description: `New value of the key`,
			}},
		}, {
			Name:        []string{"diff", "-diff"},
			Description: `Compares the data from path1 and path2 looking at the specified keys (or all if no keys are specified)`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "keys",
				Description: `Directory Service record attribute type`,
				IsOptional:  true,
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"passwd", "-passwd"},
			Description: `Changes the password of a user`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to the record`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "new password",
				Description: `New password of the user`,
				IsOptional:  true,
			}},
		}},
	}
}
