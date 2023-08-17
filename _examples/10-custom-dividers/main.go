package main

import (
	"os"

	"github.com/khulnasoft-lab/table"
)

func main() {

	t := table.New(os.Stdout)
	t.SetDividers(table.Dividers{
		ALL: "@",
		NES: "@",
		NSW: "@",
		NEW: "@",
		ESW: "@",
		NE:  "@",
		NW:  "@",
		SW:  "@",
		ES:  "@",
		EW:  "~",
		NS:  "!",
	})

	t.SetHeaders("ID", "Fruit", "Stock")

	t.AddRow("1", "Apple", "14")
	t.AddRow("2", "Banana", "88,041")
	t.AddRow("3", "Cherry", "342")
	t.AddRow("4", "Dragonfruit", "1")

	t.Render()
}
