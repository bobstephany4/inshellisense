// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["splash"] = model.Subcommand{
		Name:        []string{"splash-cli"},
		Description: `Get stunning wallpapers from Unsplash`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Help Message`,
		}, {
			Name:        []string{"--version", "-v"},
			Description: `Prints "splash-cli" version`,
		}, {
			Name:        []string{"--scale"},
			Description: `Scale of the image`,
			Args: []model.Arg{{
				Name:        "scale",
				Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`fill`}}, {Name: []string{`fit`}}, {Name: []string{`stretch`}}, {Name: []string{`center`}}},
			}},
		}, {
			Name:        []string{"--screen"},
			Description: `Set wallpaper on selected screen`,
			Args: []model.Arg{{
				Name:        "screen",
				Suggestions: []model.Suggestion{{Name: []string{`all`}}, {Name: []string{`main`}}},
			}},
		}, {
			Name:        []string{"-s", "--save"},
			Description: `Save photo without setting as wallpaper`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFolders},
				Name:       "path",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--set"},
			Description: `Set wallpaper from local file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "filepath",
			}},
		}, {
			Name:        []string{"-i", "--info"},
			Description: `Show image exif data`,
		}, {
			Name:        []string{"-q", "--quiet"},
			Description: `Hide output`,
		}, {
			Name:        []string{"--rotate"},
			Description: `Rotate image`,
			Args: []model.Arg{{
				Name:        "degrees",
				Suggestions: []model.Suggestion{{Name: []string{`90`}}, {Name: []string{`180`}}, {Name: []string{`270`}}},
			}},
		}, {
			Name:        []string{"--colorspace"},
			Description: `Define image colorspace`,
			Args: []model.Arg{{
				Name:        "colorspace",
				Suggestions: []model.Suggestion{{Name: []string{`srgb`}}, {Name: []string{`rgb`}}, {Name: []string{`cmyk`}}, {Name: []string{`lab`}}, {Name: []string{`b-w`}}},
			}},
		}, {
			Name:        []string{"--flip"},
			Description: `Flip image on the Y axis`,
		}, {
			Name:        []string{"-f", "--featured"},
			Description: `Limit to only featured photos`,
		}, {
			Name:        []string{"--query"},
			Description: `Filter by keywords`,
			Args: []model.Arg{{
				Name: "querystring",
			}},
		}, {
			Name:        []string{"--orientation"},
			Description: `Filter by orientation`,
			Args: []model.Arg{{
				Name:        "orientation",
				Suggestions: []model.Suggestion{{Name: []string{`landscape`}}, {Name: []string{`portrait`}}, {Name: []string{`squarish`}}},
			}},
		}, {
			Name:        []string{"-c", "--curated"},
			Description: `Random Curated photo`,
		}, {
			Name:        []string{"-u", "--user"},
			Description: `Random photo from user`,
			Args: []model.Arg{{
				Name: "username",
			}},
		}, {
			Name:        []string{"--collection"},
			Description: `Random photo from collection`,
			Args: []model.Arg{{
				Name: "collection id",
			}},
		}, {
			Name:        []string{"--id"},
			Description: `Get photo by ID`,
			Args: []model.Arg{{
				Name: "photo_id",
			}},
		}, {
			Name:        []string{"--day"},
			Description: `Photo of the day`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"settings"},
			Description: `Manage settings`,
			Subcommands: []model.Subcommand{{
				Name: []string{"get"},
				Args: []model.Arg{{
					Name:        "config key",
					Description: `Config key`,
					IsOptional:  true,
				}},
			}, {
				Name:        []string{"set"},
				Description: `Setup the configuration`,
			}, {
				Name:        []string{"restore"},
				Description: `Restore default settings`,
			}},
		}, {
			Name:        []string{"aliases"},
			Description: `Manage aliases`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"get"},
				Description: `Get an alias`,
				Args: []model.Arg{{
					Name: "alias",
				}},
			}, {
				Name:        []string{"set"},
				Description: `Set an alias`,
				Args: []model.Arg{{
					Name: "key",
				}, {
					Name: "value",
				}},
			}, {
				Name:        []string{"help"},
				Description: `Show help menu`,
			}},
		}, {
			Name:        []string{"collection"},
			Description: `Manage collections`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"get"},
				Description: `Get a collection`,
				Args: []model.Arg{{
					Name: "collection id",
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Delete a collection`,
				Args: []model.Arg{{
					Name: "collection id",
				}},
			}, {
				Name:        []string{"create"},
				Description: `Create a collection`,
			}, {
				Name:        []string{"photos:add"},
				Description: `Add photos to a collection`,
			}, {
				Name:        []string{"photos:remove"},
				Description: `Remove photos to a collection`,
			}, {
				Name:        []string{"help"},
				Description: `Show help for this command`,
			}},
		}, {
			Name:        []string{"dir"},
			Description: `Manage SplashCLI download directory`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"get"},
				Description: `Get the download directory path`,
			}, {
				Name:        []string{"set"},
				Description: `Set the download directory path`,
				Args: []model.Arg{{
					Name: "path",
				}},
			}, {
				Name:        []string{"clean"},
				Description: `Delete all the downloaded photos`,
			}, {
				Name:        []string{"count"},
				Description: `Count all the downloaded photos`,
			}, {
				Name:        []string{"help"},
				Description: `Show help for this command`,
			}},
		}, {
			Name: []string{"user"},
			Subcommands: []model.Subcommand{{
				Name:        []string{"login"},
				Description: `Login with your Unsplash account`,
			}, {
				Name:        []string{"logout"},
				Description: `Removes all user data`,
			}, {
				Name:        []string{"liked"},
				Description: `List last 10 liked photos`,
			}, {
				Name:        []string{"collections"},
				Description: `List all user's collections`,
			}, {
				Name:        []string{"get"},
				Description: `Get user infos`,
			}, {
				Name:        []string{"edit", "update"},
				Description: `Update user infos`,
			}, {
				Name:        []string{"help"},
				Description: `Help Menu`,
			}},
		}},
	}
}
