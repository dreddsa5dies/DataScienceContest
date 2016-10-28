package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type infoID struct {
	Name         string
	MccCodeCount map[string]int
	Gender       float64
}

func main() {
	var (
		manX   infoID
		womanX infoID
	)

	// create reference Man & Woman
	manX.Name = "MAN"
	womanX.Name = "WOMAN"
	manX = mccCodeCounts(manX)
	womanX = mccCodeCounts(womanX)
	manX.Gender = 1
	womanX.Gender = 0

	allID := createAllIDSlice()

	// calculation of means values for MCC_CODE reference Man & Woman
	countReferenseMCCMan := make(map[string][]int, 200)
	countReferenseMCCWoman := make(map[string][]int, 200)
	for p := 0; p < len(allID); p++ {
		switch {
		// Woman
		case allID[p].Gender == 0:
			for key, counts := range allID[p].MccCodeCount {
				if counts != 0 {
					countReferenseMCCWoman[key] = append(countReferenseMCCWoman[key], counts)
				}
			}
		//Man
		case allID[p].Gender == 1:
			for key, counts := range allID[p].MccCodeCount {
				if counts != 0 {
					countReferenseMCCMan[key] = append(countReferenseMCCMan[key], counts)
				}
			}
		}
	}

	// counts mean value for MCC code on reference Genders
	referenseMCCMan := make(map[string]int, 200)
	for key, counts := range countReferenseMCCMan {
		b := 0
		count := 0
		for i := 0; i < len(counts); i++ {
			b = b + counts[i]
			count++
		}
		referenseMCCMan[key] = b / count
	}

	referenseMCCWoman := make(map[string]int, 200)
	for key, counts := range countReferenseMCCWoman {
		b := 0
		count := 0
		for i := 0; i < len(counts); i++ {
			b = b + counts[i]
			count++
		}
		referenseMCCWoman[key] = b / count
	}

	// write means value MCC_CODE for reference man & woman
	manX.MccCodeCount = referenseMCCMan
	womanX.MccCodeCount = referenseMCCWoman

	// calculation the probability of the gender
	for p := 0; p < len(allID); p++ {
		// no Gender
		if allID[p].Gender == 2 {
			arrCountsMCCRef := []float64{}
			countMCCRefForArr := 0.0
			// logic calculate
			for key, _ := range allID[p].MccCodeCount {
				// allID[p].MccCodeCount[key] - mcc code for ID
				// manX.MccCodeCount[key] reference mcc code for MAN
				// womanX.MccCodeCount[key] reference mcc code for WoMAN
				x := allID[p].MccCodeCount[key]
				y := manX.MccCodeCount[key]
				z := womanX.MccCodeCount[key]
				if y > z {
					switch {
					case x <= y && x >= z:
						b := (y + z) / 2
						if x >= b {
							countMCCRefForArr = 1.0
						} else {
							countMCCRefForArr = 0.0
						}
					case x > y:
						countMCCRefForArr = 1.0
					case x < z:
						countMCCRefForArr = 0.0
					}
				} else if y < z {
					switch {
					case x <= z && x >= y:
						b := (y + z) / 2
						if x >= b {
							countMCCRefForArr = 0.0
						} else {
							countMCCRefForArr = 1.0
						}
					case x > z:
						countMCCRefForArr = 0.0
					case x < y:
						countMCCRefForArr = 1.0
					}
				}
				arrCountsMCCRef = append(arrCountsMCCRef, countMCCRefForArr)
			}
			a := 0.0
			for j := 0; j < len(arrCountsMCCRef); j++ {
				a = a + arrCountsMCCRef[j]
			}
			allID[p].Gender = a / float64(len(arrCountsMCCRef))
		}
	}

	// create stdout for input csv
	fmt.Println("customer_id,gender")
	for p := 0; p < len(allID); p++ {
		fmt.Printf("%v,%v\n", allID[p].Name, allID[p].Gender)
	}
	fmt.Println()
}

func mccCodeCounts(x infoID) infoID {
	mccCode := readCsv("mccTest.csv")

	x.MccCodeCount = make(map[string]int, 200)
	for i := 1; i < len(mccCode); i++ {
		// read mccCode
		b := mccCode[i][0]
		// add all mccCode & count is 0
		x.MccCodeCount[b] = 0
	}
	return x
}

// create []infoID for all ID infoID
func createAllIDSlice() []infoID {
	var (
		allID   []infoID
		x       infoID
		namesID []string
	)

	//read data files
	csvData := readCsv("transactions.csv")
	getGenderID := readCsv("customers_gender_train.csv")

	// search unduplicated ID
	for i := 1; i < len(csvData)-1; i++ {
		if csvData[i][0] != csvData[i+1][0] {
			namesID = append(namesID, csvData[i][0])
		}
	}

	// all ID read & counts mcc
	for o := 0; o < len(namesID); o++ {
		x.Name = namesID[o]
		x.Gender = 2 // no men or woman

		x = mccCodeCounts(x)

		// count mccCode exploit on ID
		for i := 1; i < len(csvData); i++ {
			if x.Name == csvData[i][0] {
				x.MccCodeCount[csvData[i][2]]++
			}
		}

		// gender check for ID
		for i := 1; i < len(getGenderID); i++ {
			switch x.Name == getGenderID[i][0] {
			case true:
				b, _ := strconv.Atoi(getGenderID[i][1])
				x.Gender = float64(b)
			}
		}

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
