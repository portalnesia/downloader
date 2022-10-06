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
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"portalnesia.com/downloader/utils"
)

var (
	out  string
	info = false
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "downloader",
	Short: "Command line utility to download media",
	Long: `Downloader is CLI library to download media.
For now, it only supports video or audio from YouTube and videos from Tiktok.`,
	Version: "v1.0.0-beta.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, _ := os.UserHomeDir()
	dir := "Videos"
	if runtime.GOOS == "windows" {
		dir = "Video"
	}
	out = fmt.Sprintf("%s/%s", home, dir)
	if err := checkOutput(out, false); err != nil {
		if err = os.Mkdir(out, os.ModePerm); err != nil {
			utils.Errorf(err)
		}
	}
	info = false
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.downloader.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
