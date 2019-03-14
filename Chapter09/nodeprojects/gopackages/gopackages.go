package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	data := "item11,item12,item13\nitem21,item22,item23\nitem31,item32,item33\n"
	csvReader := csv.NewReader(strings.NewReader(data))
	i := 0
	for {
		row, err := csvReader.Read()
		if err != nil {
			break
		}
		i++
		fmt.Println("Line", i, "of CSV data:", row)
	}
}
