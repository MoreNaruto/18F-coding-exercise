package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Plan struct {
	Planid     string
	State      string
	Metalfield string
	Rate       float64
	Ratearea   int64
}

type SLCSP struct {
	Zipcode string
	Rate    float64
}

type Zip struct {
	Zipcode     string
	State       string
	Countrycode string
	Name        string
	Ratearea    int64
}

func CreatePlans() []Plan {
	reader := ReadCSV("plans.csv")
	var plans []Plan

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		rate, _ := strconv.ParseFloat(strings.TrimSpace(line[3]), 64)
		rateArea, _ := strconv.ParseInt(strings.TrimSpace(line[4]), 10, 64)

		plans = append(plans, Plan{
			Planid:     strings.TrimSpace(line[0]),
			State:      strings.TrimSpace(line[1]),
			Metalfield: strings.TrimSpace(line[2]),
			Rate:       rate,
			Ratearea:   rateArea,
		})
	}

	return plans
}

func CreateSLCSPs() []SLCSP {
	reader := ReadCSV("slcsp.csv")
	var slcsps []SLCSP

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		rate, _ := strconv.ParseFloat(strings.TrimSpace(line[1]), 64)

		slcsps = append(slcsps, SLCSP{
			Zipcode: strings.TrimSpace(line[0]),
			Rate:    rate,
		})
	}

	return slcsps
}

func CreateZip() []Zip {
	reader := ReadCSV("zips.csv")
	var zips []Zip

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		rateArea, _ := strconv.ParseInt(strings.TrimSpace(line[4]), 10, 64)

		zips = append(zips, Zip{
			Zipcode: strings.TrimSpace(line[0]),
			State:   strings.TrimSpace(line[1]),
			Countrycode:   strings.TrimSpace(line[2]),
			Name: strings.TrimSpace(line[3]),
			Ratearea: rateArea,
		})
	}

	return zips
}

func ReadCSV(fileName string) *csv.Reader {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Read()
	return reader
}
