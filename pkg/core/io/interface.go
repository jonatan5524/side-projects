package core

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

type OutputHandler interface {
	PrintString(string)
	PrintObject(fmt.Stringer)
	PrintTable([]Tabler)
}

type TerminalObject interface {
	fmt.Stringer
	Tabler
}

type Tabler interface {
	TableHeader() table.Row
	TableData() table.Row
}
