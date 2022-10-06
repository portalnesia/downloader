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
package youtube

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/portalnesia/go-utils/goment"
	"portalnesia.com/downloader/utils"
)

func info_video(videoID string) {
	var err error
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	fmt.Fprintf(writer, "Please wait... we are looking for the video...\n")

	video, err := client.GetVideo(videoID)
	if err != nil {
		utils.Errorf(err)
		return
	}

	t := utils.PrintTable([]table.Row{
		{"ID", video.ID},
		{"Title", video.Title},
		{"Description", video.Description},
		{"Duration", video.Duration.String()},
		{"Author", video.Author},
		{"Publish Date", goment.Must(video.PublishDate).PNformat("full")},
	})
	out := t.Render()
	fmt.Fprintln(writer.Bypass(), out)

	formats := video.Formats
	rows := get_format(formats)

	t = utils.PrintTableBorder(rows)
	t.AppendHeader(table.Row{"ID", "Quality", "Mime Type", "Label", "Size"})
	t.SetOutputMirror(nil)
	out = t.Render()
	fmt.Fprintln(writer.Newline(), out)
	red := color.New(color.FgRed)
	red.Fprintf(writer.Newline(), "* Size is only for video not including audio\n")
}
