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
	"errors"
	"fmt"
	"strings"
	"time"

	"portalnesia.com/downloader/utils"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/jedib0t/go-pretty/v6/table"
	util "github.com/portalnesia/go-utils"
)

func select_video(videoID, output string) {
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

	fmt.Fprintf(writer.Bypass(), "ID: %s\nTitle: %s\n", video.ID, video.Title)

	formats := video.Formats
	rows := get_format(formats)
	t := utils.PrintTableBorder(rows)
	t.AppendHeader(table.Row{"ID", "Quality", "Mime Type", "Label", "Size"})
	t.Render()
	red := color.New(color.FgRed)
	red.Fprintf(writer.Newline(), "* Size is only for video not including audio\n")
	fmt.Fprintf(writer.Newline(), "\nID of quality? ")

	var q int
	fmt.Scan(&q)

	if q == 0 || q > len(formats) {
		err = errors.New("invalid quality ID")
		utils.Errorf(err)
		return
	}
	format := formats[q-1]
	filename := fmt.Sprintf("%s/%s [Portalnesia]", output, utils.SanitizeFilename(video.Title))
	isAudio := strings.Contains(format.MimeType, "audio")
	if isAudio {
		filename += ".mp3"
	} else {
		filename += utils.YTPickIdealFileExtension(format.MimeType)
	}

	fmt.Fprintf(writer, "Selected: %s, %s, %s\n", format.QualityLabel, format.MimeType, util.NumberSize(float64(format.ContentLength), 2))

	pw := utils.GetProgress()

	// call Render() in async mode; yes we don't have any trackers at the moment
	go pw.Render()

	if format.AudioChannels <= 0 {
		go download_video_audio(pw, video, &format, filename)
	} else {
		go download_media(pw, video, &format, filename)
	}
	time.Sleep(time.Second)
	for pw.IsRenderInProgress() {
		time.Sleep(time.Millisecond * 100)
	}

	time.Sleep(time.Second)
	color.Green("File saved at %s", filename)
}
