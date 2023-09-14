// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["pass"] = model.Subcommand{
		Name:        []string{"pass"},
		Description: `Pass - stores, retrieves, generates, and synchronizes passwords securely`,
		Args: []model.Arg{{
			Name:        "pass-name",
			Description: `The password you want to show`,
			Generator:   nil, // TODO: port over generator
		}},
		Options: []model.Option{{
			Name:        []string{"--clip", "-c"},
			Description: `Copy the password to the clipboard`,
		}, {
			Name:        []string{"--qrcode", "-q"},
			Description: `Display a QRcode of the password`,
		}, {
			Name:        []string{"--help"},
			Description: `Show help for pass`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"init"},
			Description: `Initialize new password storage and use gpg-id for encryption`,
			Args: []model.Arg{{
				Name:        "gpg-id",
				Description: `The gpg-id you want to use to encrypt your password store`,
			}},
			Options: []model.Option{{
				Name:        []string{"--path", "-p"},
				Description: `A specific gpg-id or set of gpg-ids is assigned for that specific sub folder of the password store`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "sub-folder",
				}},
			}},
		}, {
			Name:        []string{"insert"},
			Description: `Insert a new password into the password store called pass-name`,
			Args: []model.Arg{{
				Name:        "pass-name",
				Description: `The password name`,
			}},
			Options: []model.Option{{
				Name:        []string{"--echo", "-e"},
				Description: `Disable keyboard echo when the password is entered and confirm the password by asking for it twice`,
			}, {
				Name:        []string{"--multi-line", "-m"},
				Description: `Lines will be read until EOF or Ctrl+D is reached. Otherwise, only a single line from standard in is read`,
			}, {
				Name:        []string{"--force", "-f"},
				Description: `Don't prompt before overwriting an existing password`,
			}},
		}, {
			Name:        []string{"git"},
			Description: `Password store git functions`,
		}, {
			Name:        []string{"version"},
			Description: `Show version information`,
		}, {
			Name:        []string{"help"},
			Description: `Show usage message`,
		}, {
			Name:        []string{"cp"},
			Description: `Copies the password or directory named old-path to new-path`,
			Args: []model.Arg{{
				Name:        "old-path",
				Description: `The old password name or directory`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "new-path",
				Description: `The new password name or directory`,
			}},
			Options: []model.Option{{
				Name:        []string{"--force", "-f"},
				Description: `Do not interactively prompt before moving`,
			}},
		}, {
			Name:        []string{"mv"},
			Description: `Renames the password or directory named old-path to new-path`,
			Args: []model.Arg{{
				Name:        "old-path",
				Description: `The old password name or directory`,
				Generator:   nil, // TODO: port over generator
			}, {
				Name:        "new-path",
				Description: `The new password name or directory`,
			}},
			Options: []model.Option{{
				Name:        []string{"--force", "-f"},
				Description: `Do not interactively prompt before moving`,
			}},
		}, {
			Name:        []string{"rm"},
			Description: `Remove the password named pass-name from the password store`,
			Args: []model.Arg{{
				Name:        "pass-name",
				Description: `The password name`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--recursive", "-r"},
				Description: `Delete pass-name recursively if it is a directory`,
			}, {
				Name:        []string{"--force", "-f"},
				Description: `Do not interactively prompt before removal`,
			}},
		}, {
			Name:        []string{"generate"},
			Description: `Generate a new password of length pass-length and insert into pass-name`,
			Args: []model.Arg{{
				Name:        "pass-name",
				Description: `The password name`,
			}, {
				Name:        "pass-length",
				Description: `The length of the password`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--no-symbols", "-n"},
				Description: `Do not use any non-alphanumeric characters in the generated password`,
			}, {
				Name:        []string{"--clip", "-c"},
				Description: `Do not print the password but instead copy it to the clipboard`,
			}, {
				Name:        []string{"--in-place", "-i"},
				Description: `Do not interactively prompt, and only replace the first line of the password file with the new generated password, keeping the remainder of the file intact`,
			}, {
				Name:        []string{"--force", "-f"},
				Description: `Overwrite the existing password`,
			}},
		}, {
			Name:        []string{"ls", "list"},
			Description: `List names of passwords inside the tree at subfolder by using the tree`,
			Args: []model.Arg{{
				Name:        "password sub-directory",
				Description: `The password sub directory you want to list`,
				Generator:   nil, // TODO: port over generator
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"find"},
			Description: `List names of passwords inside the tree that match pass-names`,
			Args: []model.Arg{{
				Name:        "pass-name",
				Description: `The password name you want to search for`,
			}},
		}, {
			Name:        []string{"show"},
			Description: `Decrypt and print a password`,
			Args: []model.Arg{{
				Name:        "pass-name",
				Description: `The password you want to show`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--clip", "-c"},
				Description: `Copy the password to the clipboard`,
			}, {
				Name:        []string{"--qrcode", "-q"},
				Description: `Display a QRcode of the password`,
			}},
		}, {
			Name:        []string{"edit"},
			Description: `Insert a new password or edit an existing password using the default text editor specified by the environment`,
			Args: []model.Arg{{
				Name:        "pass-name",
				Description: `The password you want to edit`,
				Generator:   nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"grep"},
			Description: `Searches inside each decrypted password file for search-string. Grep options can be used`,
			Args: []model.Arg{{
				Name:        "pass-name",
				Description: `The password name you want to grep for`,
			}},
		}},
	}
}
