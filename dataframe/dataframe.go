// Package dataframe implements a dataframe.
// This dataframe does not enforce first normal form.
// It is built to handle JSON format files
package dataframe

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Dataframe is the basic struct for this package.
type Dataframe struct {
	Data map[int][]string
}

// ReadCSV takes the name of a csv file as an argument
// and returns a pointer to a dataframe
func ReadCSV(file string) *Dataframe {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file!")
		return nil
	}
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}
	var df Dataframe
	m := make(map[int][]string)
	for i := 0; i < len(data); i++ {
		m[i] = data[i]
	}
	df.Data = m
	return &df
}
