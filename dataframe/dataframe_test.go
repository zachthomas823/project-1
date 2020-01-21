package dataframe

import (
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {
	df := ReadCSV("./test.csv")
	if df.Data[0][0] != "Column 0" {
		t.Error("Error in header, column 0. Should be Column 0, was", df.Data[0][0])
	}
	if df.Data[0][1] != "Column B" {
		t.Error("Error in header, column 1. Should be Column B, was", df.Data[0][1])
	}
	if df.Data[0][2] != "Column 2" {
		t.Error("Error in header, column 2. Should be Column 2, was", df.Data[0][2])
	}
	if df.Data[1][0] != "1" {
		t.Error("Error in row 0, column 0. Should be 1, was", df.Data[1][0])
	}
	if df.Data[1][1] != "B" {
		t.Error("Error in row 0, column 1. Should be B, was", df.Data[1][1])
	}
	if df.Data[1][2] != "3" {
		t.Error("Error in row 0, column 3. Should be 3, was", df.Data[1][2])
	}
	if df.Data[2][0] != "2.0" {
		t.Error("Error in row 1, column 0. Should be 2.0, was", df.Data[2][0])
	}
	if df.Data[2][1] != "b" {
		t.Error("Error in row 1, column 1. Should be b, was", df.Data[2][1])
	}
	if df.Data[2][2] != "1.0" {
		t.Error("Error in row 1, column 2. Should be 1.0, was", df.Data[2][2])
	}
}

func TestSort(t *testing.T) {
	df := ReadCSV("./test.csv")
	df.Sort(2)
	if df.Data[0][0] != "Column 0" {
		t.Error("Error in header, column 0. Should be Column 0, was", df.Data[0][0])
	}
	if df.Data[0][1] != "Column B" {
		t.Error("Error in header, column 1. Should be Column B, was", df.Data[0][1])
	}
	if df.Data[0][2] != "Column 2" {
		t.Error("Error in header, column 2. Should be Column 2, was", df.Data[0][2])
	}
	if df.Data[1][0] != "1" {
		t.Error("Error in row 0, column 0. Should be 1, was", df.Data[1][0])
	}
	if df.Data[1][1] != "B" {
		t.Error("Error in row 0, column 1. Should be B, was", df.Data[1][1])
	}
	if df.Data[1][2] != "3" {
		t.Error("Error in row 0, column 3. Should be 3, was", df.Data[1][2])
	}
	if df.Data[2][0] != "2.0" {
		t.Error("Error in row 1, column 0. Should be 2.0, was", df.Data[2][0])
	}
	if df.Data[2][1] != "b" {
		t.Error("Error in row 1, column 1. Should be b, was", df.Data[2][1])
	}
	if df.Data[2][2] != "1.0" {
		t.Error("Error in row 1, column 2. Should be 1.0, was", df.Data[2][2])
	}
	var dfCompare Dataframe
	dfCompare.Data = df.Data
	df.Sort(1)
	if !reflect.DeepEqual(df.Data, dfCompare.Data) {
		t.Error("Sort on non-numeric column changed dataframe")
	}
	df.Sort(0)
	if df.Data[0][0] != "Column 0" {
		t.Error("Error in header, column 0. Should be Column 0, was", df.Data[0][0])
	}
	if df.Data[0][1] != "Column B" {
		t.Error("Error in header, column 1. Should be Column B, was", df.Data[0][1])
	}
	if df.Data[0][2] != "Column 2" {
		t.Error("Error in header, column 2. Should be Column 2, was", df.Data[0][2])
	}
	if df.Data[1][0] != "2.0" {
		t.Error("Error in row 0, column 0. Should be 2.0, was", df.Data[1][0])
	}
	if df.Data[1][1] != "b" {
		t.Error("Error in row 0, column 1. Should be b, was", df.Data[1][1])
	}
	if df.Data[1][2] != "1.0" {
		t.Error("Error in row 0, column 3. Should be 1.0, was", df.Data[1][2])
	}
	if df.Data[2][0] != "1" {
		t.Error("Error in row 1, column 0. Should be 1, was", df.Data[2][0])
	}
	if df.Data[2][1] != "B" {
		t.Error("Error in row 1, column 1. Should be B, was", df.Data[2][1])
	}
	if df.Data[2][2] != "3" {
		t.Error("Error in row 1, column 2. Should be 3, was", df.Data[2][2])
	}
}

func TestDropCol(t *testing.T) {
	df := ReadCSV("./test.csv")
	df.DropCol(0)
	if !reflect.DeepEqual(df.Data[0], []string{"Column B", "Column 2"}) ||
		!reflect.DeepEqual(df.Data[1], []string{"B", "3"}) ||
		!reflect.DeepEqual(df.Data[2], []string{"b", "1.0"}) {
		t.Error("DropCol(0) did not properly drop the first column")
	}
	df.DropCol(1)
	if !reflect.DeepEqual(df.Data[0], []string{"Column B"}) ||
		!reflect.DeepEqual(df.Data[1], []string{"B"}) ||
		!reflect.DeepEqual(df.Data[2], []string{"b"}) {
		t.Error("DropCol(1) did not properly drop the second column")
	}
	df.DropCol(0)
	df.DropCol(0)
}

func TestDropRow(t *testing.T) {
	df := ReadCSV("./test.csv")
	df.DropRow(0)
	if !reflect.DeepEqual(df.Data[0], []string{"1", "B", "3"}) ||
		!reflect.DeepEqual(df.Data[1], []string{"2.0", "b", "1.0"}) {
		t.Error("DropRow(0) did not properly drop the first row")
	}
}
