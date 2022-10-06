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
package tiktok

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"time"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/table"
	"portalnesia.com/downloader/utils"
)

func downloadCommand(v *Video, output string) error {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	fmt.Fprintf(writer, "Please wait... we are looking for the video...\n")

	info, err := get_info(v)
	if err != nil {
		utils.Errorf(err)
		return err
	}

	t := utils.PrintTable([]table.Row{
		{"ID", info.Video.ID},
		{"Description", info.Video.Description},
		{"Created", info.Video.CreatedTime.PNformat("full")},
		{"Author", info.Author.Nickname},
	})
	out := t.Render()
	fmt.Fprintln(writer.Bypass(), out)

	resp, err := v.Download() // Change return, from string to *http.Response
	if err != nil {
		utils.Errorf(err)
		return err
	}
	defer resp.Body.Close()
	regex := regexp.MustCompile(`(?m)tiktok\.com\/@([\w.-]+)\/video\/(\d+)/?`)
	regexResult := regex.FindStringSubmatch(v.URL)

	username, videoId := regexResult[1], regexResult[2]
	filename := fmt.Sprintf("%s/%s_%s [Portalnesia].mp4", output, utils.SanitizeFilename(username), utils.SanitizeFilename(videoId))

	pw := utils.GetProgress()
	go pw.Render()
	go downloading(pw, filename, resp.ContentLength, resp.Body)

	time.Sleep(time.Second)
	for pw.IsRenderInProgress() {
		time.Sleep(time.Second)
	}
	color.Green("File saved at %s", filename)
	return nil
}

func downloading(pw progress.Writer, filename string, total int64, stream io.ReadCloser) {
	file, err := os.Create(filename)
	if err != nil {
		utils.Errorf(err)
		return
	}
	defer file.Close()

	tracker := progress.Tracker{
		Message: "Downloading video",
		Total:   total,
		Units:   progress.UnitsBytes,
	}

	pw.AppendTracker(&tracker)
	bytes := make(chan int64)

	go func(done chan int64, fi *os.File, track *progress.Tracker) {
		time.Sleep(time.Second)
		for !tracker.IsDone() {
			select {
			case <-done:
				track.MarkAsDone()
			default:
				info, _ := fi.Stat()
				size := info.Size()
				track.SetValue(size)
			}
			time.Sleep(time.Millisecond * 500)
		}
	}(bytes, file, &tracker)

	n, err := io.Copy(file, stream)
	if err != nil {
		tracker.MarkAsErrored()
		utils.Errorf(err)
		return
	}
	bytes <- n
}
