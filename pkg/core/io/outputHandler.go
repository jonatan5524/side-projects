package core

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type outputStdout struct {
	writer *os.File
}

func NewOutputStdout() OutputHandler {
	return &outputStdout{os.Stdout}
}

func (output *outputStdout) PrintString(str string) {
	fmt.Fprintln(output.writer, str)
}

func (output *outputStdout) PrintObject(object fmt.Stringer) {
	fmt.Fprintf(output.writer, "%v\n", object)
}

func (output *outputStdout) PrintTable(objects []Tabler) {
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(output.writer)
	tableWriter.SetStyle(table.StyleLight)

	if len(objects) > 0 {
		tableWriter.AppendHeader(objects[0].TableHeader())

		for _, object := range objects {
			tableWriter.AppendRow(object.TableData())
			tableWriter.AppendSeparator()
		}

	}

	tableWriter.Render()
}
