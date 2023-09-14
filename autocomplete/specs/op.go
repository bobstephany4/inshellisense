// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["op"] = model.Subcommand{
		Name:        []string{"op"},
		Description: `Official 1Password CLI`,
		Options: []model.Option{{
			Name:        []string{"--account"},
			Description: `Select the account to execute the command by account shorthand, sign-in address, account ID, or user ID. For a list of available accounts, run 'op account list'. Can be set as the OP_ACCOUNT environment variable`,
			Args: []model.Arg{{
				Name:      "account",
				Generator: nil, // TODO: port over generator
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"--cache"},
			Description:  `Store and use cached information`,
			IsPersistent: true,
		}, {
			Name:        []string{"--config"},
			Description: `Use this configuration directory`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "directory",
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"--debug"},
			Description:  `Output debug logs. Can also be set using $OP_DEBUG environment variable`,
			IsPersistent: true,
		}, {
			Name:        []string{"--encoding"},
			Description: `Use this character encoding type. Default: UTF-8. Supported: SHIFT_JIS, gbk`,
			Args: []model.Arg{{
				Name:        "type",
				Suggestions: []model.Suggestion{{Name: []string{`UTF-8`}}, {Name: []string{`SHIFT_JIS`}}, {Name: []string{`gbk`}}},
			}},
			IsPersistent: true,
		}, {
			Name:        []string{"--format"},
			Description: `Use this output format. Can be 'human-readable' or 'json'. Can be set as the OP_FORMAT environment variable. (default "human-readable")`,
			Args: []model.Arg{{
				Name:        "string",
				Suggestions: []model.Suggestion{{Name: []string{`human-readable`}}, {Name: []string{`json`}}},
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"-h", "--help"},
			Description:  `Get help for op`,
			IsPersistent: true,
		}, {
			Name:         []string{"--iso-timestamps"},
			Description:  `Format timestamps according to ISO 8601 / RFC 3339. Can be set as the OP_ISO_TIMESTAMPS environment variable`,
			IsPersistent: true,
		}, {
			Name:         []string{"--no-color"},
			Description:  `Print output without color`,
			IsPersistent: true,
		}, {
			Name:        []string{"--session"},
			Description: `Authenticate with this session token. 1Password CLI outputs session tokens for successful 'op signin' commands when biometric unlock is disabled`,
			Args: []model.Arg{{
				Name: "token",
			}},
			IsPersistent: true,
		}, {
			Name:        []string{"-v", "--version"},
			Description: `Version for op`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"account"},
			Description: `Manage your locally configured 1Password accounts`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"add"},
				Description: `Add an account to sign in to for the first time`,
				Options: []model.Option{{
					Name:        []string{"--address"},
					Description: `The sign-in address for your account`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--email"},
					Description: `The email address associated with your account`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--raw"},
					Description: `Only return the session token`,
				}, {
					Name:        []string{"--shorthand"},
					Description: `Set the short account name`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get details about your account`,
			}, {
				Name:        []string{"list", "ls"},
				Description: `List users and accounts set up on this device`,
			}, {
				Name:        []string{"forget"},
				Description: `Remove a 1Password account from this device`,
				Args: []model.Arg{{
					Name: "account",
				}},
				Options: []model.Option{{
					Name:        []string{"--all"},
					Description: `Forget all authenticated accounts`,
				}},
			}},
		}, {
			Name:        []string{"connect"},
			Description: `Manage Connect instances and Connect tokens in your 1Password account`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"group"},
				Description: `Manage group access to Secrets Automation`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"grant"},
					Description: `Grant a group access to manage Secrets Automation`,
					Options: []model.Option{{
						Name:        []string{"--all-servers"},
						Description: `Grant access to all current and future servers in the authenticated account`,
					}, {
						Name:        []string{"--group"},
						Description: `The group to receive access`,
						Args: []model.Arg{{
							Name: "group",
						}},
					}},
				}, {
					Name:        []string{"revoke"},
					Description: `Revoke a group's access to manage Secrets Automation`,
					Options: []model.Option{{
						Name:        []string{"--all-servers"},
						Description: `Revoke access to all current and future servers in the authenticated account`,
					}, {
						Name:        []string{"--group"},
						Description: `The group to revoke access from`,
						Args: []model.Arg{{
							Name: "group",
						}},
					}},
				}},
			}, {
				Name:        []string{"server"},
				Description: `Manage Connect servers`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"create"},
					Description: `Set up a Connect server`,
				}, {
					Name:        []string{"get"},
					Description: `Get a Connect server`,
					Args: []model.Arg{{
						Name: "serverName | serverID",
						Suggestions: []model.Suggestion{{
							Name:        []string{`-`},
							Description: `Read from stdin`,
						}},
					}},
				}, {
					Name:        []string{"edit"},
					Description: `Rename a Connect server`,
					Args: []model.Arg{{
						Name: "serverName | serverID",
					}},
				}, {
					Name:        []string{"delete", "remove", "rm"},
					Description: `Remove a Connect server`,
					Args: []model.Arg{{
						Name: "serverName | serverID",
						Suggestions: []model.Suggestion{{
							Name:        []string{`-`},
							Description: `Read from stdin`,
						}},
					}},
				}, {
					Name:        []string{"list", "ls"},
					Description: `List Connect servers`,
				}},
			}, {
				Name:        []string{"token"},
				Description: `Manage Connect tokens`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"create"},
					Description: `Issue a token for a 1Password Connect server`,
					Options: []model.Option{{
						Name:        []string{"--expires-in"},
						Description: `Set how long the Connect token is valid for in (s)econds, (m)inutes, or (h)ours`,
						Args: []model.Arg{{
							Name: "duration",
						}},
					}, {
						Name:        []string{"--server"},
						Description: `Issue a token for this server`,
						Args: []model.Arg{{
							Name: "string",
						}},
					}},
				}, {
					Name:        []string{"edit"},
					Description: `Rename a Connect token`,
					Options: []model.Option{{
						Name:        []string{"--name"},
						Description: `Change the token's name`,
						Args: []model.Arg{{
							Name: "string",
						}},
					}},
				}, {
					Name:        []string{"delete", "remove", "rm"},
					Description: `Revoke a token for a Connect server`,
					Args: []model.Arg{{
						Name: "token",
					}},
				}, {
					Name:        []string{"list", "ls"},
					Description: `Get a list of tokens`,
				}},
			}, {
				Name:        []string{"vault"},
				Description: `Manage connect server vault access`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"grant"},
					Description: `Grant a Connect server access to a vault`,
					Options: []model.Option{{
						Name:        []string{"--server"},
						Description: `The server to be granted access`,
						Args: []model.Arg{{
							Name: "string",
						}},
					}},
				}, {
					Name:        []string{"revoke"},
					Description: `Revoke a Connect server's access to a vault`,
					Options: []model.Option{{
						Name:        []string{"--server"},
						Description: `The server to revoke access from`,
						Args: []model.Arg{{
							Name: "server",
						}},
					}},
				}},
			}},
		}, {
			Name:        []string{"document"},
			Description: `Perform CRUD operations on Document items in your vaults`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a document item`,
				Args: []model.Arg{{
					Name: "file",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--file-name"},
					Description: `Set the file's name`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}, {
					Name:        []string{"--tags"},
					Description: `Set the tags to the specified (comma-separated) values`,
					Args: []model.Arg{{
						Name: "tags",
					}},
				}, {
					Name:        []string{"--title"},
					Description: `Set the document item's title`,
					Args: []model.Arg{{
						Name: "title",
					}},
				}},
			}, {
				Name:        []string{"get"},
				Description: `Download a document`,
				Args: []model.Arg{{
					Name: "itemName | itemID",
				}},
				Options: []model.Option{{
					Name:        []string{"--include-archive"},
					Description: `Include document items in the Archive. Can also be set using OP_INCLUDE_ARCHIVE environment variable`,
				}, {
					Name:        []string{"--output"},
					Description: `Save the document to the file path instead of stdout`,
					Args: []model.Arg{{
						Name: "path",
					}},
				}},
			}, {
				Name:        []string{"edit"},
				Description: `Edit a document item`,
				Args: []model.Arg{{
					Name: "itemName | itemID",
				}, {
					Name: "file",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--file-name"},
					Description: `Set the file's name`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}, {
					Name:        []string{"--tags"},
					Description: `Set the tags to the specified (comma-separated) values. An empty value will remove all tags`,
					Args: []model.Arg{{
						Name: "tags",
					}},
				}, {
					Name:        []string{"--title"},
					Description: `Set the document item's title`,
					Args: []model.Arg{{
						Name: "title",
					}},
				}},
			}, {
				Name:        []string{"delete", "remove", "rm"},
				Description: `Delete or archive a document item`,
				Args: []model.Arg{{
					Name: "itemName | itemID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--archive"},
					Description: `Move the document to the Archive`,
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `Get a list of documents`,
				Options: []model.Option{{
					Name:        []string{"--include-archive"},
					Description: `Include document items in the Archive. Can also be set using OP_INCLUDE_ARCHIVE environment variable`,
				}},
			}},
		}, {
			Name:        []string{"events-api"},
			Description: `Manage Events API integrations in your 1Password account`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Set up an integration with the Events API`,
			}},
		}, {
			Name:        []string{"group"},
			Description: `Perform CRUD operations on the groups of users in your 1Password account`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"user"},
				Description: `Manage users in groups`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"grant"},
					Description: `Grant a user access to a group`,
					Options: []model.Option{{
						Name:        []string{"--group"},
						Description: `Specify the group to grant the user access to`,
						Args: []model.Arg{{
							Name: "string",
						}},
					}, {
						Name:        []string{"--role"},
						Description: `Specify the user's role as a member or manager. Default: member`,
						Args: []model.Arg{{
							Name: "string",
						}},
					}},
				}, {
					Name:        []string{"revoke"},
					Description: `Revoke a user's access to a vault or group`,
					Options: []model.Option{{
						Name:        []string{"--group"},
						Description: `Specify the group to revoke the user's access to`,
						Args: []model.Arg{{
							Name: "string",
						}},
					}},
				}, {
					Name:        []string{"list", "ls"},
					Description: `Retrieve users that belong to a group`,
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create a group`,
				Options: []model.Option{{
					Name:        []string{"--description"},
					Description: `Set the group's description`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get details about a group`,
				Args: []model.Arg{{
					Name: "groupName | groupID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
			}, {
				Name:        []string{"edit"},
				Description: `Edit a group's name or description`,
				Args: []model.Arg{{
					Name: "groupName | groupID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--description"},
					Description: `Change the group's description`,
					Args: []model.Arg{{
						Name: "description",
					}},
				}},
			}, {
				Name:        []string{"delete", "remove", "rm"},
				Description: `Remove a group`,
				Args: []model.Arg{{
					Name: "groupName | groupID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `List groups`,
				Options: []model.Option{{
					Name:        []string{"--user"},
					Description: `List groups that a user belongs to`,
					Args: []model.Arg{{
						Name: "user",
					}},
				}},
			}},
		}, {
			Name:        []string{"item"},
			Description: `Perform CRUD operations on the 1Password items in your vaults`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"template"},
				Description: `Manage templates`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"get"},
					Description: `Get an item template`,
					Args: []model.Arg{{
						Name: "category",
						Suggestions: []model.Suggestion{{
							Name:        []string{`-`},
							Description: `Read from stdin`,
						}},
					}},
					Options: []model.Option{{
						Name:        []string{"-f", "--force"},
						Description: `Do not prompt for confirmation`,
					}},
				}, {
					Name:        []string{"list", "ls"},
					Description: `Get a list of templates`,
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create an item`,
				Args: []model.Arg{{
					Name: "assignment...",
				}},
				Options: []model.Option{{
					Name:        []string{"--category"},
					Description: `Set the item's category`,
					Args: []model.Arg{{
						Name: "category",
					}},
				}, {
					Name:        []string{"--dry-run"},
					Description: `Perform a dry run of the command and output a preview of the resulting item`,
				}, {
					Name:        []string{"--generate-password"},
					Description: `Give the item a randomly generated password`,
					Args: []model.Arg{{
						Name: "recipe",
					}},
				}, {
					Name:        []string{"--tags"},
					Description: `Set the tags to the specified (comma-separated) values`,
					Args: []model.Arg{{
						Name: "tags",
					}},
				}, {
					Name:        []string{"--template"},
					Description: `Specify the filepath to read an item template from`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--title"},
					Description: `Set the item's title`,
					Args: []model.Arg{{
						Name: "title",
					}},
				}, {
					Name:        []string{"--url"},
					Description: `Set the URL associated with the item`,
					Args: []model.Arg{{
						Name: "URL",
					}},
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get an item's details`,
				Args: []model.Arg{{
					Name: "itemName | itemID | shareLink",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--fields"},
					Description: `Only return data from these fields. Use 'label=' to get the field by name or 'type=' to filter fields by type`,
					Args: []model.Arg{{
						Name: "fields",
					}},
				}, {
					Name:        []string{"--include-archive"},
					Description: `Include items in the Archive. Can also be set using OP_INCLUDE_ARCHIVE environment variable`,
				}, {
					Name:        []string{"--otp"},
					Description: `Output the primary one-time password for this item`,
				}, {
					Name:        []string{"--share-link"},
					Description: `Get a shareable link for the item`,
				}},
			}, {
				Name:        []string{"edit"},
				Description: `Edit an item's details`,
				Args: []model.Arg{{
					Name: "itemName | itemID | shareLink",
				}, {
					Name: "assignment ...",
				}},
			}, {
				Name:        []string{"delete", "remove", "rm"},
				Description: `Delete or archive an item`,
				Args: []model.Arg{{
					Name: "itemName | itemID | shareLink",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--archive"},
					Description: `Move the item to the Archive`,
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `List items`,
				Options: []model.Option{{
					Name:        []string{"--categories"},
					Description: `Only list items in these categories (comma-separated)`,
					Args: []model.Arg{{
						Name: "categories",
					}},
				}, {
					Name:        []string{"--include-archive"},
					Description: `Include items in the Archive. Can also be set using OP_INCLUDE_ARCHIVE environment variable`,
				}, {
					Name:        []string{"--long"},
					Description: `Output a more detailed item list`,
				}, {
					Name:        []string{"--tags"},
					Description: `Only list items with these tags (comma-separated)`,
					Args: []model.Arg{{
						Name: "tags",
					}},
				}},
			}, {
				Name:        []string{"share"},
				Description: `Share an item`,
				Args: []model.Arg{{
					Name: "itemName | itemID",
				}},
				Options: []model.Option{{
					Name:        []string{"--emails"},
					Description: `Email addresses to share with`,
					Args: []model.Arg{{
						Name: "strings",
					}},
				}, {
					Name:        []string{"--expiry"},
					Description: `Link expiring after the specified duration in (s)econds, (m)inutes, or (h)ours (default 7h)`,
					Args: []model.Arg{{
						Name: "duration",
					}},
				}, {
					Name:        []string{"--vault"},
					Description: `Look for the item in this vault`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}},
		}, {
			Name:        []string{"user"},
			Description: `Manage users within this 1Password account`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"provision"},
				Description: `Provision a user in the authenticated account`,
				Options: []model.Option{{
					Name:        []string{"--email"},
					Description: `Provide the user's email address`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--language"},
					Description: `Provide the user's account language. (default "en")`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}, {
				Name:        []string{"confirm"},
				Description: `Confirm a user`,
				Args: []model.Arg{{
					Name: "email | name | userID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--all"},
					Description: `Confirm all unconfirmed users`,
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get details about a user`,
				Args: []model.Arg{{
					Name: "email | name | userID | --me",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}, {
						Name:        []string{`--me`},
						Description: `Get the authenticated user's details`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--fingerprint"},
					Description: `Get the user's public key fingerprint`,
				}},
			}, {
				Name:        []string{"edit"},
				Description: `Edit a user's name or Travel Mode status`,
				Args: []model.Arg{{
					Name: "email | name | userID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--name"},
					Description: `Set the user's name`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}, {
				Name:        []string{"suspend"},
				Description: `Suspend a user`,
				Args: []model.Arg{{
					Name: "email | name | userID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--deauthorize-devices-after"},
					Description: `Deauthorize the user's devices after a time (rounded down to seconds)`,
					Args: []model.Arg{{
						Name: "duration",
					}},
				}},
			}, {
				Name:        []string{"reactivate"},
				Description: `Reactivate a suspended user`,
				Args: []model.Arg{{
					Name: "email | name | userID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
			}, {
				Name:        []string{"delete", "remove", "rm"},
				Description: `Remove a user and all their data from the account`,
				Args: []model.Arg{{
					Name: "email | name | userID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `List users`,
				Options: []model.Option{{
					Name:        []string{"--group"},
					Description: `List users who belong to a group`,
					Args: []model.Arg{{
						Name: "group",
					}},
				}},
			}},
		}, {
			Name:        []string{"vault"},
			Description: `Manage permissions and perform CRUD operations on your 1Password vaults`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"group"},
				Description: `Manage group vault access`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"grant"},
					Description: `Grant a group permissions to a vault`,
					Options: []model.Option{{
						Name:        []string{"--group"},
						Description: `The group to receive access`,
						Args: []model.Arg{{
							Name: "group",
						}},
					}, {
						Name:        []string{"--no-input"},
						Description: `Do not prompt for input on interactive terminal`,
						Args: []model.Arg{{
							Name: "input",
						}},
					}, {
						Name:        []string{"--permissions"},
						Description: `The permissions to grant to the group`,
						Args: []model.Arg{{
							Name: "permissions",
						}},
					}},
				}, {
					Name:        []string{"revoke"},
					Description: `Revoke a portion or the entire access of a group to a vault`,
					Options: []model.Option{{
						Name:        []string{"--group"},
						Description: `The group to revoke access from`,
						Args: []model.Arg{{
							Name: "group",
						}},
					}, {
						Name:        []string{"--no-input"},
						Description: `Do not prompt for input on interactive terminal`,
						Args: []model.Arg{{
							Name: "input",
						}},
					}, {
						Name:        []string{"--permissions"},
						Description: `The permissions to revoke from the group`,
						Args: []model.Arg{{
							Name: "permissions",
						}},
					}},
				}, {
					Name:        []string{"list", "ls"},
					Description: `List all the groups that have access to the given vault`,
					Args: []model.Arg{{
						Name: "vault",
						Suggestions: []model.Suggestion{{
							Name:        []string{`-`},
							Description: `Read from stdin`,
						}},
					}},
				}},
			}, {
				Name:        []string{"user"},
				Description: `Manage user vault access`,
				Subcommands: []model.Subcommand{{
					Name:        []string{"grant"},
					Description: `Grant a user access to a vault`,
					Options: []model.Option{{
						Name:        []string{"--no-input"},
						Description: `Do not prompt for input on interactive terminal`,
						Args: []model.Arg{{
							Name: "input",
						}},
					}, {
						Name:        []string{"--permissions"},
						Description: `The permissions to grant to the user`,
						Args: []model.Arg{{
							Name: "permissions",
						}},
					}, {
						Name:        []string{"--user"},
						Description: `The user to receive access`,
						Args: []model.Arg{{
							Name: "user",
						}},
					}},
				}, {
					Name:        []string{"revoke"},
					Description: `Revoke a portion or the entire access of a user to a vault`,
					Options: []model.Option{{
						Name:        []string{"--no-input"},
						Description: `Do not prompt for input on interactive terminal`,
						Args: []model.Arg{{
							Name: "input",
						}},
					}, {
						Name:        []string{"--permissions"},
						Description: `The permissions to revoke from the user`,
						Args: []model.Arg{{
							Name: "permissions",
						}},
					}, {
						Name:        []string{"--user"},
						Description: `The user to revoke access from`,
						Args: []model.Arg{{
							Name: "user",
						}},
					}},
				}, {
					Name:        []string{"list", "ls"},
					Description: `List all users with access to the vault and their permissions`,
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create a new vault`,
				Options: []model.Option{{
					Name:        []string{"--allow-admins-to-manage"},
					Description: `Set whether admins can manage the vault. Default: false`,
					Args: []model.Arg{{
						Name: "true|false",
					}},
				}, {
					Name:        []string{"--description"},
					Description: `Set the group's description`,
					Args: []model.Arg{{
						Name: "description",
					}},
				}},
			}, {
				Name:        []string{"get"},
				Description: `Get details about a vault`,
				Args: []model.Arg{{
					Name: "vaultName | vaultID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
			}, {
				Name:        []string{"edit"},
				Description: `Edit a vault's name, description, icon or Travel Mode status`,
				Args: []model.Arg{{
					Name: "vaultName | vaultID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
				Options: []model.Option{{
					Name:        []string{"--description"},
					Description: `Change the vault's description`,
					Args: []model.Arg{{
						Name: "description",
					}},
				}, {
					Name:        []string{"--icon"},
					Description: `Change the vault's icon`,
					Args: []model.Arg{{
						Name: "icon",
					}},
				}, {
					Name:        []string{"--name"},
					Description: `Change the vault's name`,
					Args: []model.Arg{{
						Name: "name",
					}},
				}, {
					Name:        []string{"--travel-mode"},
					Description: `Turn Travel Mode on or off for the vault. Only vaults with Travel Mode enabled will be accessible while a user has Travel Mode turned on. (default off)`,
					Args: []model.Arg{{
						Name: "on|off",
					}},
				}},
			}, {
				Name:        []string{"delete", "remove", "rm"},
				Description: `Remove a vault`,
				Args: []model.Arg{{
					Name: "vaultName | vaultID",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Read from stdin`,
					}},
				}},
			}, {
				Name:        []string{"list", "ls"},
				Description: `List all vaults in the account`,
				Options: []model.Option{{
					Name:        []string{"--group"},
					Description: `List vaults a group has access to`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}},
			}},
		}, {
			Name:        []string{"completion"},
			Description: `Generate shell completion information`,
		}, {
			Name:        []string{"inject"},
			Description: `Inject secrets into a config file`,
		}, {
			Name:        []string{"read"},
			Description: `Read a secret using the secrets reference syntax`,
			Options: []model.Option{{
				Name:        []string{"--file-mode"},
				Description: `Set filemode for the output file if it does not yet exist. It is ignored without the --out-file flag. (default 0600)`,
				Args: []model.Arg{{
					Name: "filemode",
				}},
			}, {
				Name:        []string{"-f", "--force"},
				Description: `Do not prompt for confirmation`,
			}, {
				Name:        []string{"-n", "--no-newline"},
				Description: `Do not print a new line after the secret`,
			}},
		}, {
			Name:        []string{"run"},
			Description: `Pass secrets as environment variables to a process`,
			Args: []model.Arg{{
				Name:      "command",
				IsCommand: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--env-file"},
				Description: `Enable Dotenv integration with specific Dotenv files to parse. For example: --env-file=.env`,
				Args: []model.Arg{{
					Name: "stringArray",
				}},
			}, {
				Name:        []string{"--no-masking"},
				Description: `Disable masking of secrets on stdout and stderr`,
			}},
		}, {
			Name:        []string{"signin"},
			Description: `Sign in to a 1Password account`,
			Options: []model.Option{{
				Name:        []string{"-f", "--force"},
				Description: `Ignore warnings and print raw output from this command`,
			}},
		}, {
			Name:        []string{"signout"},
			Description: `Sign out of a 1Password account`,
			Options: []model.Option{{
				Name:        []string{"--all"},
				Description: `Sign out of all signed-in accounts`,
			}, {
				Name:        []string{"--forget"},
				Description: `Remove the details for a 1Password account from this device`,
			}},
		}, {
			Name:        []string{"update"},
			Description: `Check for and download updates`,
			Options: []model.Option{{
				Name:        []string{"--directory"},
				Description: `Download the update to this ''path''`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "string",
				}},
			}},
		}, {
			Name:        []string{"whoami"},
			Description: `Get information about a signed-in account`,
		}},
	}
}
