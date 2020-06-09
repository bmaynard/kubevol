package core

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
)

type UI struct {
	tableWriter table.Writer
}

func SetupTable(header table.Row, out io.Writer) *UI {
	ui := &UI{table.NewWriter()}

	ui.tableWriter.SetOutputMirror(out)
	ui.tableWriter.AppendHeader(header)

	return ui
}

func (ui *UI) AppendRow(row []table.Row) {
	ui.tableWriter.AppendRows(row)
}

func (ui *UI) Render() {
	ui.tableWriter.Render()
}
