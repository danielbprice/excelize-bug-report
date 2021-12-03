package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	if len(os.Args) != 2 {
		panic("need file argument")
	}

	f, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	sheets := f.GetSheetList()
	rows, err := f.GetRows(sheets[0])
	fmt.Printf("Rows: %#v", rows)
}
