/*
Copyright © 2022 Putu Aditya <aditya@portalnesia.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"portalnesia.com/downloader/youtube"
)

// youtubeCmd represents the youtube command
var youtubeCmd = &cobra.Command{
	Use:   "youtube [url or id youtube video]",
	Short: "Youtube media downloader",
	Long: `Download video or audio from youtube.
With this command, you can download video (with audio) or audio only from youtube
and save it tou your library.

- If you want to download audio only, select "audio only" from format selection.
- For combining video and audio file, please install ffmpeg to your machine.
  Example for linux: sudo apt install ffmpeg`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkOutput(out, true)
		youtube.Command(args[0], out, info)
	},
}

func init() {
	rootCmd.AddCommand(youtubeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// youtubeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	youtubeCmd.Flags().BoolVarP(&info, "info", "i", false, "Show only video information without downloading")
	youtubeCmd.Flags().StringVarP(&out, "output", "o", out, "Directory where to save downloaded video.")
}
