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
	"encoding/json"
	"fmt"

	"github.com/gosuri/uilive"
	"github.com/jedib0t/go-pretty/v6/table"
	"portalnesia.com/downloader/utils"
)

func get_info(v *Video) (*Data, error) {
	err := v.FetchInfo() // Change #sigi-persisted-data to #SIGI_STATE
	if err != nil {
		return nil, err
	}
	infoStr, err := v.GetInfo()
	if err != nil {
		return nil, err
	}
	var info Data
	err = json.Unmarshal([]byte(infoStr), &info)
	if err != nil {
		return nil, err
	}
	info.Format()

	return &info, nil
}

func infoCommand(v *Video) error {
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
		{"VIDEO", ""},
		{"ID", info.Video.ID},
		{"Description", info.Video.Description},
		{"URL", info.Video.URL},
		{"Likes", info.Video.Likes},
		{"Shares", info.Video.Shares},
		{"Comments", info.Video.Comment},
		{"Played", info.Video.Played},
		{"Created", info.Video.CreatedTime.PNformat("full")},
		{"", ""},
		{"AUTHOR", ""},
		{"ID", info.Author.ID},
		{"Nickname", info.Author.Nickname},
		{"URL", info.Author.Url},
	})
	out := t.Render()
	fmt.Fprintln(writer.Bypass(), out)

	return nil
}
