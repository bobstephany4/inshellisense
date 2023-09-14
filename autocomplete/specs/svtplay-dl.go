// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["svtplay-dl"] = model.Subcommand{
		Name:        []string{"svtplay-dl"},
		Description: `Download videos from your favourite play services`,
		Args: []model.Arg{{
			Name:        "urls",
			Description: `URLs for the videos you want to download`,
			Generator:   nil, // TODO: port over generator
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for svtplay-dl`,
		}, {
			Name:        []string{"--version"},
			Description: `Show program's version number and exit`,
		}, {
			Name:        []string{"-o", "--output"},
			Description: `Outputs to the given filename or folder`,
		}, {
			Name:        []string{"--subfolder"},
			Description: `Create a subfolder titled as the show, non-series gets in folder movies`,
		}, {
			Name:        []string{"--config"},
			Description: `Specify configuration file`,
			Args: []model.Arg{{
				Name: "configfile",
			}},
		}, {
			Name:        []string{"-f", "--force"},
			Description: `Overwrite if file exists already`,
		}, {
			Name:        []string{"-r", "--resume"},
			Description: `Resume a download (RTMP obsolete)`,
		}, {
			Name:        []string{"-l", "--live"},
			Description: `Enable for live streams (RTMP based ones)`,
		}, {
			Name:        []string{"-c", "--capture_time"},
			Description: `Define capture time in minutes of a live stream`,
			Args: []model.Arg{{
				Name: "capture_time",
			}},
		}, {
			Name:        []string{"-s", "--silent"},
			Description: `Be less verbose`,
		}, {
			Name:        []string{"--silent-semi"},
			Description: `Only show a message when the file is downloaded`,
		}, {
			Name:        []string{"-u", "--username"},
			Description: `Username (if applicable)`,
			Args: []model.Arg{{
				Name: "username",
			}},
		}, {
			Name:        []string{"-p", "--password"},
			Description: `Password (if applicable)`,
			Args: []model.Arg{{
				Name: "password",
			}},
		}, {
			Name:        []string{"-t", "--thumbnail"},
			Description: `Download thumbnail from the site if available`,
		}, {
			Name:        []string{"-g", "--get-url"},
			Description: `Do not download any video, but instead print the URL`,
		}, {
			Name:        []string{"--get-only-episode-url"},
			Description: `Do not get video URLs, only print the episode URL`,
		}, {
			Name:        []string{"--dont-verify-ssl-cert"},
			Description: `Don't attempt to verify SSL certificates`,
		}, {
			Name:        []string{"--http-header"},
			Description: `A header to add to each HTTP request`,
			Args: []model.Arg{{
				Name: "header1=value;header2=value2",
			}},
		}, {
			Name:        []string{"--cookies"},
			Description: `A cookies to add to each HTTP request`,
			Args: []model.Arg{{
				Name: "cookie1=value;cookie2=value2",
			}},
		}, {
			Name:        []string{"--exclude"},
			Description: `Exclude videos with the WORD(s) in the filename. Comma separated`,
			Args: []model.Arg{{
				Name: "WORD1,WORD2,...",
			}},
		}, {
			Name:        []string{"--after-date"},
			Description: `Only videos published on or after this date`,
			Args: []model.Arg{{
				Name: "yyyy-MM-dd",
			}},
		}, {
			Name:        []string{"--proxy"},
			Description: `Use the specified HTTP/HTTPS/SOCKS proxy. To enable experimental SOCKS proxy, specify a proper scheme. For example socks5://127.0.0.1:1080/`,
			Args: []model.Arg{{
				Name: "proxy",
			}},
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Explain what is going on`,
		}, {
			Name:        []string{"--nfo"},
			Description: `Create a NFO file`,
		}, {
			Name:        []string{"--force-nfo"},
			Description: `Download only NFO if used with --nfo`,
		}, {
			Name:        []string{"--only-audio"},
			Description: `Only download audio if audio and video is separated`,
		}, {
			Name:        []string{"--only-video"},
			Description: `Only download video if audio and video is separated`,
		}, {
			Name:        []string{"-q", "--quality"},
			Description: `Choose what format to download based on bitrate/video resolution. It will download the best format by default`,
			Args: []model.Arg{{
				Name: "quality",
			}},
		}, {
			Name:        []string{"-Q", "--flexible-quality"},
			Description: `Allow given quality (as above) to differ by an amount`,
			Args: []model.Arg{{
				Name: "amount",
			}},
		}, {
			Name:        []string{"-P", "--preferred"},
			Description: `Preferred download method (dash, hls, or http)`,
			Args: []model.Arg{{
				Name: "preferred",
			}},
		}, {
			Name:        []string{"--list-quality"},
			Description: `List the quality for a video`,
		}, {
			Name:        []string{"--stream-priority"},
			Description: `If two streams have the same quality, choose the one you prefer`,
			Args: []model.Arg{{
				Name: "dash,hls,http",
			}},
		}, {
			Name:        []string{"--format-preferred"},
			Description: `Choose the format you prefer, use --list-quality to show which one to choose from`,
			Args: []model.Arg{{
				Name: "h264,h264-51",
			}},
		}, {
			Name:        []string{"--audio-language"},
			Description: `Choose the language of the audio (it can override the default one), use --list-quality to show which one to choose from`,
			Args: []model.Arg{{
				Name: "AUDIO_LANGUAGE",
			}},
		}, {
			Name:        []string{"--audio-role"},
			Description: `Choose the role of the audio (it can override the default one), use --list-quality to show which one to choose from`,
			Args: []model.Arg{{
				Name: "AUDIO_ROLE",
			}},
		}, {
			Name:        []string{"-S", "--subtitle"},
			Description: `Download subtitle from the site if available`,
		}, {
			Name:        []string{"-M", "--merge-subtitle"},
			Description: `Merge subtitle with video/audio file with corresponding ISO639-3 language code. This invokes --remux automatically`,
		}, {
			Name:        []string{"--force-subtitle"},
			Description: `Download only subtitle if its used with -S`,
		}, {
			Name:        []string{"--require-subtitle"},
			Description: `Download only if a subtitle is available`,
		}, {
			Name:        []string{"--all-subtitles"},
			Description: `Download all available subtitles for the video`,
		}, {
			Name:        []string{"--raw-subtitles"},
			Description: `Also download the subtitles in their native format`,
		}, {
			Name:        []string{"--convert-subtitle-colors"},
			Description: `Converts the color information in subtitles, to <font color=""> tags`,
		}, {
			Name:        []string{"-A", "--all-episodes"},
			Description: `Try to download all episodes`,
		}, {
			Name:        []string{"--all-last"},
			Description: `Get last NN episodes instead of all episodes`,
			Args: []model.Arg{{
				Name: "NN",
			}},
		}, {
			Name:        []string{"--include-clips"},
			Description: `Include clips from websites when using -A`,
		}, {
			Name:        []string{"--cmore-operatorlist"},
			Description: `Show operatorlist for cmore`,
		}, {
			Name: []string{"--cmore-operator"},
			Args: []model.Arg{{
				Name: "operator",
			}},
		}, {
			Name:        []string{"--no-remux"},
			Description: `Do not automatically remux to mp4`,
		}, {
			Name:        []string{"--no-merge"},
			Description: `Do not automatically merge video, audio and possibly also subtitle(s) together`,
		}, {
			Name:        []string{"--no-postprocess"},
			Description: `Do not postprocess anything`,
		}, {
			Name:        []string{"--keep-original"},
			Description: `Do postprocessing while also keeping original files`,
		}, {
			Name:        []string{"--output-format"},
			Description: `Format you want resulting file in (mkv or mp4), mp4 is default`,
			Args: []model.Arg{{
				Name: "mp4,mkv",
			}},
		}},
	}
}
