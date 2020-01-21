// Package dataframe implements a dataframe.
// A dataframe is a two dimensional table.
package dataframe

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Dataframe is the two dimensional table structure for this package
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

// PrettyPrint prints out the dataframe in an easier to read than normal format
func (df *Dataframe) PrettyPrint() {
	sizes := make([]int, len(df.Data[0])+1)
	sizes[0] = 2
	for i := 0; i < len(df.Data); i++ {
		for k := 0; k < len(df.Data[i]); k++ {
			if len(df.Data[i][k]) > sizes[k+1] {
				sizes[k+1] = len(df.Data[i][k])
			}
		}
	}
	for i := 0; i < len(df.Data); i++ {
		fmt.Print(i)
		if i < 10 {
			fmt.Print("  ")
		}
		if i >= 10 && i < 100 {
			fmt.Print(" ")
		}
		fmt.Print("|")
		for k := 0; k < len(df.Data[i]); k++ {
			spaces := sizes[k+1] - len(df.Data[i][k])
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
	}
}

// PrettyString returns a string of the dataframe in an easier to read than normal format
func (df *Dataframe) PrettyString() string {
	sizes := make([]int, len(df.Data[0])+1)
	sizes[0] = 2
	for i := 0; i < len(df.Data); i++ {
		for k := 0; k < len(df.Data[i]); k++ {
			if len(df.Data[i][k]) > sizes[k+1] {
				sizes[k+1] = len(df.Data[i][k])
			}
		}
	}
	var result string
	for i := 0; i < len(df.Data); i++ {
		result = result + strconv.Itoa(i)
		if i < 10 {
			result = result + "  "
		}
		if i >= 10 && i < 100 {
			result = result + " "
		}
		result = result + "|"
		for k := 0; k < len(df.Data[i]); k++ {
			spaces := sizes[k+1] - len(df.Data[i][k])
			result = result + strings.Repeat(" ", spaces/2)
			result = result + df.Data[i][k]
			result = result + strings.Repeat(" ", spaces-spaces/2)
			result = result + "|"
		}
		result = result + "\n"
		for k := range sizes {
			result = result + "+"
			result = result + strings.Repeat("-", sizes[k])
		}
		result = result + "+\n"
	}
	return result
}

// ToJSON will return a string of the dataframe in JSON format.
// The first column of the dataframe will be treated as a key to a map
// which maps to a map. In the secondary map each column name will be
// a key to the value that is held for that row.
func (df *Dataframe) ToJSON() string {
	var result string
	result = "{"
	for i := 1; i < len(df.Data); i++ {
		var entry string
		key := df.Data[i][0]
		entry = entry + "\"" + key + "\":{\n\t"
		for j := 1; j < len(df.Data[0]); j++ {
			entry = entry + "\"" + df.Data[0][j] + "\": "
			entry = entry + "\"" + df.Data[i][j]
			if j == len(df.Data[0])-1 {
				entry = entry + "\"\n"
			} else {
				entry = entry + "\",\n\t"
			}
		}
		if i == len(df.Data)-1 {
			entry = entry + "}\n"
		} else {
			entry = entry + "},\n"
		}
		result += entry
	}
	result = result + "}"
	return result
}
