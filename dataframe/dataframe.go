// Package dataframe implements a dataframe.
// A dataframe is a two dimensional table
package dataframe

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Dataframe is the basic struct for this package.
type Dataframe struct {
	Data map[int][]string
}

// ReadCSV takes the name of a csv file as an argument
// and returns a pointer to a dataframe.
// The first line will be interpreted as the header
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

// Sort is a method for a Dataframe that takes a column number as a parameter
// The column must be made of numeric characters in string form
// Uses insertion sort algorithm
// Changes are made in place
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

// DropCol removes the column specified by index from the dataframe
// Changes are made in place
func (df *Dataframe) DropCol(column int) {
	m := make(map[int][]string)
	for i := 0; i < len(df.Data); i++ {
		row1 := df.Data[i][0:column]
		row2 := df.Data[i][column+1 : len(df.Data[i])]
		newRow := append(row1, row2...)
		m[i] = newRow
	}
	df.Data = m
}

// DropRow removes the row specified by index from the dataframe
// Changes are made in place
func (df *Dataframe) DropRow(row int) {
	if row >= len(df.Data) {
		fmt.Println("Index out of range")
		return
	}
	for i := row; i < len(df.Data)-1; i++ {
		df.Data[i] = df.Data[i+1]
	}
	delete(df.Data, len(df.Data)-1)
}

func (df *Dataframe) PrettyPrint() {
	sizes := make([]int, len(df.Data[0]))
	for i := 0; i < len(df.Data); i++ {
		for k := 0; k < len(df.Data[i]); k++ {
			if len(df.Data[i][k]) > sizes[k] {
				sizes[k] = len(df.Data[i][k])
			}
		}
	}
	// fullLine := len(df.Data[0]) + 1
	// for i := 0; i < len(sizes); i++ {
	// 	fullLine += sizes[i]
	// }
	for i := 0; i < len(df.Data); i++ {
		fmt.Print("|")
		for k := 0; k < len(df.Data[i]); k++ {
			spaces := sizes[k] - len(df.Data[i][k])
			fmt.Print(strings.Repeat(" ", spaces/2))
			fmt.Print(df.Data[i][k])
			fmt.Print(strings.Repeat(" ", spaces-spaces/2))
			fmt.Print("|")
		}
		fmt.Println()
		for k := range sizes {
			fmt.Print("+")
			fmt.Print(strings.Repeat("-", sizes[k]))
		}
		fmt.Println("+")
		// fmt.Println(strings.Repeat("-", fullLine))
	}
}
