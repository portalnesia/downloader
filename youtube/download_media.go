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
	"io"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	youtube "github.com/kkdai/youtube/v2"
	"portalnesia.com/downloader/utils"
)

func download_media(pw progress.Writer, video *youtube.Video, format *youtube.Format, filename string) {
	stream, total, err := client.GetStream(video, format)
	if err != nil {
		utils.Errorf(err)
		return
	}

	tipe := "video"
	if strings.Contains(format.MimeType, "audio") {
		tipe = "audio"
	}

	file, err := os.Create(filename)
	if err != nil {
		utils.Errorf(err)
		return
	}
	defer file.Close()

	tracker := progress.Tracker{
		Message: fmt.Sprintf("Downloading %s", tipe),
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
