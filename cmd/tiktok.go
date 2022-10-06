/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"portalnesia.com/downloader/tiktok"
)

// tiktokCmd represents the tiktok command
var tiktokCmd = &cobra.Command{
	Use:   "tiktok [url tiktok]",
	Short: "Tiktok video downloader",
	Long: `Download video from tiktok.
With this command, you can download video from tiktok
and save it tou your library`,
	RunE: TiktokCmdRunE,
}

func TiktokCmdRunE(cmd *cobra.Command, args []string) error {
	checkOutput(out, true)
	return tiktok.Command(args[0], out, info)
}

func init() {
	rootCmd.AddCommand(tiktokCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tiktokCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tiktokCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tiktokCmd.Flags().BoolVarP(&info, "info", "i", false, "Show only video information without downloading")
}
