package formatter

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func createTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	return t
}

func setBorderOptions(t table.Writer, drawBorder, separateRows, separateColumns bool) {
	t.Style().Options.DrawBorder = drawBorder
	t.Style().Options.SeparateRows = separateRows
	t.Style().Options.SeparateColumns = separateColumns
}
