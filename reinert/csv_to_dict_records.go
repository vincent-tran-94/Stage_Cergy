package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func csv_to_dict() []LemmatisationStruct {
	csvFile, err := os.Open("./corpus/Lematisation.csv")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvlines := csv.NewReader(csvFile)
	if err != nil {
		fmt.Println(err)
	}

	data, _ := csvlines.ReadAll()
	var rows []LemmatisationStruct

	// Convert to struct
	for i, line := range data {
		if i != 0 {
			rows = append(rows, LemmatisationStruct{line[0], line[1]})
		}
	}
	return rows
}
