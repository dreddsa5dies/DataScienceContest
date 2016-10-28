package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type infoID struct {
	Name   string
	Gender float64
}

// genderSorter block by sorting []infoID slise noGenders
type genderSorter []infoID

func (a genderSorter) Len() int           { return len(a) }
func (a genderSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a genderSorter) Less(i, j int) bool { return a[i].Gender < a[j].Gender }

func main() {
	var xGender []infoID
	allID := createAllIDSlice()

	// counts % woman
	countWoman := 0.0
	allGenderRead := 0.0
	for i := 0; i < len(allID); i++ {
		if allID[i].Gender == 0 {
			countWoman++
			allGenderRead++
		} else if allID[i].Gender == 1 {
			allGenderRead++
		} else {
			// read noGenders
			xGender = append(xGender, allID[i])
		}
	}
	procentWoman := (countWoman / allGenderRead) * 100

	// range noGenders & comparsion on % WOMAN.
	sort.Sort(genderSorter(xGender))

	// STDOUT finally answer
	fmt.Println("customer_id,gender")
	for i := 0; i < len(xGender); i++ {
		if float64(i) < (float64(len(xGender))/100)*procentWoman {
			fmt.Printf("%v,0\n", xGender[i].Name)
		} else {
			fmt.Printf("%v,1\n", xGender[i].Name)
		}
	}
}

// create []infoID for all ID infoID
func createAllIDSlice() []infoID {
	var (
		allID []infoID
		x     infoID
	)

	//read data files
	csvData := readCsv("$GOPATH/src/github.com/dreddsa5dies/sdsjOlimpic/A/step_2/resultGender.csv")

	// all ID & Gender read
	for o := 0; o < len(csvData); o++ {
		x.Name = csvData[o][0]
		x.Gender, _ = strconv.ParseFloat(csvData[o][1], 64)
		allID = append(allID, x)
	}

	return allID
}

// check error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// read & encoding csv
func readCsv(addr string) [][]string {

	// read file
	dat, err := ioutil.ReadFile(addr)
	check(err)

	in := string(dat)

	// encoding csv
	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	check(err)

	return records
}
