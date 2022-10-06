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
package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

var TableNoBorder = table.Style{
	Name: "NoBorder",
	Box: table.BoxStyle{
		PaddingLeft:  " ",
		PaddingRight: " ",
	},
}

func PrintTable(rows []table.Row) table.Writer {
	t := table.NewWriter()
	t.SetStyle(TableNoBorder)
	t.AppendRows(rows)
	return t
}

func PrintTableBorder(rows []table.Row) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRows(rows)
	return t
}

func Errorf(err error) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s %s\n", red("ERROR: "), err.Error())
	os.Exit(1)
}
