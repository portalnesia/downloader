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
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	youtube "github.com/kkdai/youtube/v2"
	"portalnesia.com/downloader/utils"
)

func download_video_audio(pw progress.Writer, v *youtube.Video, videoFormat *youtube.Format, filename string) {
	if checkFfmpeg, err := exec.Command("ffmpeg", "-version").CombinedOutput(); err != nil {
		err = errors.New(fmt.Sprint(err) + ": " + string(checkFfmpeg))
		time.Sleep(time.Second)
		utils.Errorf(err)
		return
	}

	audioFormat, err := getAudioFormats(v)
	if err != nil {
		utils.Errorf(err)
		return
	}
	outputDir := filepath.Dir(filename)

	// Create temporary video file
	videoFile, err := os.CreateTemp(outputDir, "youtube_*.m4v")
	if err != nil {
		utils.Errorf(err)
		return
	}
	defer func() {
		videoFile.Close()
		os.Remove(videoFile.Name())
	}()

	// Create temporary audio file
	audioFile, err := os.CreateTemp(outputDir, "youtube_*.m4a")
	if err != nil {
		utils.Errorf(err)
		return
	}
	defer func() {
		audioFile.Close()
		os.Remove(audioFile.Name())
	}()

	// Download video
	doneVideoCh, doneAudioCh := make(chan bool), make(chan bool)
	doneVideo, doneAudio := false, false

	tracker := progress.Tracker{
		Message: "Combining video and audio",
		Total:   100,
		Units:   progress.UnitsDefault,
	}
	pw.AppendTracker(&tracker)

	go downloading(doneVideoCh, pw, videoFile, v, videoFormat, "Downloading video")
	go downloading(doneAudioCh, pw, audioFile, v, audioFormat, "Downloading audio")

	for !doneVideo || !doneAudio {
		select {
		case <-doneAudioCh:
			doneAudio = true
			tracker.Increment(40)
		case <-doneVideoCh:
			doneVideo = true
			tracker.Increment(40)
		}
		time.Sleep(time.Second)
		/*
			perc := ((percVid + percAud) / 2) - 10
			perc = math.Max(perc, 0)
			tracker.SetValue(int64(perc))
		*/
	}

	ffmpegVersionCmd := exec.Command("ffmpeg", "-y",
		"-i", videoFile.Name(),
		"-i", audioFile.Name(),
		"-c", "copy", // Just copy without re-encoding
		"-shortest", // Finish encoding when the shortest input stream ends
		filename,
		"-loglevel", "warning",
	)

	output, err := ffmpegVersionCmd.CombinedOutput()
	if err != nil {
		err = errors.New(fmt.Sprint(err) + ": " + string(output))
		videoFile.Close()
		audioFile.Close()
		os.Remove(videoFile.Name())
		os.Remove(audioFile.Name())
		time.Sleep(time.Second)
		utils.Errorf(err)
		return
	}
	videoFile.Close()
	audioFile.Close()
	os.Remove(videoFile.Name())
	os.Remove(audioFile.Name())
	tracker.SetValue(100)
	time.Sleep(time.Second)
	tracker.MarkAsDone()
}

func downloading(dn chan bool, pw progress.Writer, file *os.File, v *youtube.Video, format *youtube.Format, message string) {

	stream, total, err := client.GetStream(v, format)
	if err != nil {
		utils.Errorf(err)
		return
	}

	tracker := progress.Tracker{
		Message: message,
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
				dn <- true
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

func getAudioFormats(v *youtube.Video) (*youtube.Format, error) {
	var audioFormat *youtube.Format
	var audioFormats youtube.FormatList

	formats := v.Formats

	audioFormats = formats.Type("audio")

	if len(audioFormats) > 0 {
		audioFormats.Sort()
		audioFormat = &audioFormats[0]
	}

	if audioFormat == nil {
		return nil, fmt.Errorf("no audio format found after filtering")
	}

	return audioFormat, nil
}

// go run downloader.go youtube https://www.youtube.com/watch?v=6jnuwGWtpH8
// go run downloader.go youtube https://www.youtube.com/watch?v=JZ-FfGx50ro
