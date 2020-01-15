// Package dataframe implements a dataframe.
// This dataframe does not enforce first normal form.
// It is built to handle JSON format files
package dataframe

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

func (df *Dataframe) Sort(column int) {
	sortedM := make(map[int][]string)
	sortedIdx := make([]int, len(df.Data))
	sortedIdx[1] = 1
	sortedValue := make([]float64, len(df.Data))
	sortedValue[1], _ = strconv.ParseFloat(df.Data[1][column], 64)
	for i := 2; i < len(df.Data); i++ {
		placed := false
		stat, err := strconv.ParseFloat(df.Data[i][column], 64)
		if err != nil {
			fmt.Println("non numeric value in column", column)
			return
		}
		for j := 1; j < len(sortedIdx); j++ {
			if sortedValue[j] <= stat {
				for k := len(sortedIdx) - 1; k > j; k-- {
					sortedIdx[k] = sortedIdx[k-1]
					sortedValue[k] = sortedValue[k-1]
				}
				sortedIdx[j] = i
				sortedValue[j] = stat
				placed = true
				break
			}
		}
		if !placed {
			sortedIdx[i] = i
			sortedValue[i] = stat
		}
	}
	sortedM[0] = df.Data[0]
	for i := 1; i < len(sortedIdx); i++ {
		sortedM[i] = df.Data[sortedIdx[i]]
	}
	df.Data = sortedM
}
