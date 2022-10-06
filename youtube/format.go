package youtube

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	youtube "github.com/kkdai/youtube/v2"
	util "github.com/portalnesia/go-utils"
)

func get_format(formats youtube.FormatList) []table.Row {
	n := 0
	for _, f := range formats {
		if f.ContentLength > 0 {
			formats[n] = f
			n++
		}
	}
	formats = formats[:n]
	rows := []table.Row{}
	for i, f := range formats {
		mime := strings.Split(f.MimeType, ";")
		f.MimeType = mime[0]
		var quality string
		isAudio := strings.Contains(f.MimeType, "audio")
		if isAudio {
			quality = "-"
		} else {
			quality = fmt.Sprintf("%s (%s)", f.Quality, f.QualityLabel)
		}

		label := "Video + Audio"
		if f.AudioChannels > 0 && isAudio {
			label = "Audio only"
		}
		size := ""
		if f.AudioChannels <= 0 {
			size = "*"
		}
		size = fmt.Sprintf("%s%s", util.NumberSize(float64(f.ContentLength), 2), size)
		rows = append(rows, table.Row{i + 1, quality, f.MimeType, label, size})
	}
	return rows
}
